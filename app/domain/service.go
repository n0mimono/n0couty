package domain

type Service struct {
	Registory Registory
	Scraper   Scraper
}

func (service *Service) ScrapeUserChips(
	char int,
	canContinue func(int, bool) bool,
	onError func(error),
	onNext func(int, []*UserChip, bool),
) {
	page := 1
	hasNext := true

	for canContinue(page, hasNext) {
		out := service.Scraper.UserChip(InUserChip{
			Char: char,
			Page: page,
		})

		if out.Error != nil {
			onError(out.Error)
		}

		for _, chip := range out.Chips {
			_, err := service.Registory.UserChip().Create(chip)
			if err != nil {
				onError(err)
			}
		}
		onNext(page, out.Chips, out.HasNext)

		hasNext = out.HasNext
		page++
	}
}

func (service *Service) ScrapeUserPage(
	qiitaID string,
	onError func(error),
) {
	var err error

	// fetch
	out := service.Scraper.UserPage(InUserPage{
		QiitaID: qiitaID,
	})
	if err := out.Error; err != nil {
		onError(err)
	}

	// user
	user, exist := service.Registory.User().GetByQiitaID(qiitaID)
	if exist {
		out.User.ID = user.ID
		user, err = service.Registory.User().Update(user, out.User)
	} else {
		user, err = service.Registory.User().Create(out.User)
	}
	if err != nil {
		onError(err)
	}

	// proc for ban user
	if out.User.Ban {
		return
	}

	// link, follow tag, followee
	if err = service.Registory.UserSocialLink().UpdateAllByUser(user, out.SocialLinks); err != nil {
		onError(err)
	}
	if err = service.Registory.UserFollowTag().UpdateAllByUser(user, out.FollowTags); err != nil {
		onError(err)
	}
	if err = service.Registory.UserFollowee().UpdateAllByUser(user, out.Followees); err != nil {
		onError(err)
	}

	// stat
	stat, exist := service.Registory.UserStat().GetByUser(user)
	out.Stat.UserID = user.ID
	if exist {
		out.Stat.ID = stat.ID
		_, err = service.Registory.UserStat().Update(stat, out.Stat)
	} else {
		_, err = service.Registory.UserStat().Create(out.Stat)
	}
	if err != nil {
		onError(err)
	}

	// language stat
	if err = service.Registory.UserLanguageStat().UpdateAllByUser(user, out.LangStats); err != nil {
		onError(err)
	}

	// item processing
	procItems := func(pageItems []OutUserPageItem) []*UserItem {
		items := []*UserItem{}
		for _, pageItem := range pageItems {
			i := pageItem.Item
			i.UserID = user.ID
			ts := pageItem.ItemTags

			// item
			item, exist := service.Registory.UserItem().GetByArticleID(i.ArticleID)
			if exist {
				i.ID = item.ID
				item, err = service.Registory.UserItem().Update(item, i)
			} else {
				item, err = service.Registory.UserItem().Create(i)
			}
			if err != nil {
				onError(err)
			}

			// tags
			if err = service.Registory.UserItemTag().UpdateAllByItem(item, ts); err != nil {
				onError(err)
			}

			items = append(items, item)
		}
		return items
	}

	// popular items
	items := procItems(out.PopularItems)
	if err = service.Registory.UserPopularItem().UpdateAllByUser(user, items); err != nil {
		onError(err)
	}

	// recent items
	items = procItems(out.RecentItems)
	if err = service.Registory.UserRecentItem().UpdateAllByUser(user, items); err != nil {
		onError(err)
	}

}

func (service *Service) InitCrawlScores(
	onError func(error),
) {
	var id uint = 1

	service.Registory.UserCrawlScore().DropAndCreate()

	ch := make(chan bool, 20)
	for true {
		chip, exist := service.Registory.UserChip().Get(&UserChip{ID: id})
		if !exist {
			break
		}
		id++

		ch <- true
		go func() {
			score := calculateScore(chip)
			_, err := service.Registory.UserCrawlScore().Create(&UserCrawlScore{
				ChipID:    chip.ID,
				Score:     score,
				Checked:   false,
				Completed: false,
			})
			if err != nil {
				onError(err)
			}
			<-ch
		}()
	}
}

func (service *Service) ScrapeUserPageAll(
	canContinue func() bool,
	onCount func(int, int),
	onLoad func(string, string, int),
	onScrape func(),
	onError func(error),
	onComplete func(),
) {
	for canContinue() {
		cur, size := service.Registory.UserCrawlScore().GetCounts()
		onCount(cur, size)

		score, exist := service.Registory.UserCrawlScore().GetCandidate()
		if !exist {
			break
		}
		chip, _ := service.Registory.UserChip().Get(&UserChip{ID: score.ChipID})
		onLoad(chip.QiitaID, chip.Description, score.Score)

		onScrape()
		service.ScrapeUserPage(chip.QiitaID, func(err error) {
			onError(err)
		})
		service.Registory.UserCrawlScore().Complete(score)
	}
	onComplete()
}

func (service *Service) GetUserSummaries(page int, perPage int, onlyStarred bool) ([]*UserSummary, int, int) {
	toOffset := func(page int, perPage int) int {
		return (page - 1) * perPage
	}
	toLimit := func(page int, perPage int) int {
		return perPage
	}
	offset := toOffset(page, perPage)
	limit := toLimit(page, perPage)

	users := service.Registory.User().GetRange(limit, offset, onlyStarred)
	summaries := []*UserSummary{}
	for _, user := range users {
		links := service.Registory.UserSocialLink().GetAllByUser(user)
		stat, _ := service.Registory.UserStat().GetByUser(user)
		langStats := service.Registory.UserLanguageStat().GetAllByUser(user)
		scout, _ := service.Registory.UserScout().GetByUser(user)
		summaries = append(summaries, &UserSummary{
			User:      user,
			Links:     links,
			Stat:      stat,
			LangStats: langStats,
			Scout:     scout,
		})
	}

	prevPage := page - 1
	nextPage := page + 1
	if len(users) != limit {
		nextPage = 0
	}
	return summaries, prevPage, nextPage
}

func (service *Service) GetUserSummary(uid uint) (*UserSummary, bool) {
	user := &User{ID: uid}
	user, exist := service.Registory.User().Get(user)

	links := service.Registory.UserSocialLink().GetAllByUser(user)
	stat, _ := service.Registory.UserStat().GetByUser(user)
	langStats := service.Registory.UserLanguageStat().GetAllByUser(user)
	scout, _ := service.Registory.UserScout().GetByUser(user)

	return &UserSummary{
		User:      user,
		Links:     links,
		Stat:      stat,
		LangStats: langStats,
		Scout:     scout,
	}, exist
}

func (service *Service) GetUserItemSummary(uid uint) (*UserItemSummary, bool) {
	user := &User{ID: uid}
	user, exist := service.Registory.User().Get(user)

	items := service.Registory.UserItem().GetAllByUser(user)
	units := []*UserItemWithTags{}
	for _, item := range items {
		tags := service.Registory.UserItemTag().GetAllByItem(item)
		units = append(units, &UserItemWithTags{
			Body: item,
			Tags: tags,
		})
	}

	populars := service.Registory.UserPopularItem().GetAllByUser(user)
	recents := service.Registory.UserRecentItem().GetAllByUser(user)

	return &UserItemSummary{
		Items:    units,
		Populars: populars,
		Recents:  recents,
	}, exist
}

func (service *Service) UpdateUserStar(uid uint, star bool) {
	user := &User{ID: uid}
	next := &UserScout{
		UserID:  user.ID,
		Starred: &[]bool{star}[0],
	}

	scout, exist := service.Registory.UserScout().GetByUser(user)
	if exist {
		next.ID = scout.ID
		service.Registory.UserScout().Update(scout, next)
	} else {
		service.Registory.UserScout().Create(next)
	}
}
