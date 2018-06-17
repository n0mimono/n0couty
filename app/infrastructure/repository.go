package infrastructure

import (
	"n0couty/app/domain"
)

func (repo *UserRepository) getByQiitaID(qiitaID string) (User, bool) {
	user := User{}
	notFound := repo.db.Where("qiita_id=?", qiitaID).First(&user).RecordNotFound()

	return user, !notFound
}

func (repo *UserRepository) GetByQiitaID(qiitaID string) (*domain.User, bool) {
	user, exist := repo.getByQiitaID(qiitaID)

	return user.out(), exist
}

func (repo *UserRepository) getRange(limit int, offset int, onlyStarred bool) []User {
	users := []User{}
	if onlyStarred {
		sub := repo.db.Table("user_scouts").Where("starred=?", onlyStarred).Select("user_id").QueryExpr()
		repo.db.Where("id in (?)", sub).Limit(limit).Offset(offset).Find(&users)
	} else {
		repo.db.Where("id between ? and ?", offset+1, offset+limit).Find(&users)
	}
	return users
}

func (repo *UserRepository) GetRange(limit int, offset int, onlyStarred bool) []*domain.User {
	us := []*domain.User{}
	users := repo.getRange(limit, offset, onlyStarred)
	for _, user := range users {
		us = append(us, user.out())
	}
	return us
}

func (repo *UserCrawlScoreRepository) getCandidate() (UserCrawlScore, bool) {
	tx := repo.db.Begin()

	var max int
	tx.Raw("select max(score) from user_crawl_scores where checked=false").Row().Scan(&max)

	score := UserCrawlScore{}
	notFound := tx.Where("score=? and checked=?", max, false).First(&score).RecordNotFound()

	if notFound {
		return score, false
	}

	next := score
	next.Checked = true
	tx.Model(&score).Update(&next)

	tx.Commit()
	return next, true
}

func (repo *UserCrawlScoreRepository) GetCandidate() (*domain.UserCrawlScore, bool) {
	score, exist := repo.getCandidate()
	return score.out(), exist
}

func (repo *UserCrawlScoreRepository) getCounts() (int, int) {
	var size int
	repo.db.Raw("select count(id) from user_crawl_scores").Row().Scan(&size)

	var cur int
	repo.db.Raw("select count(id) from user_crawl_scores where checked=true").Row().Scan(&cur)

	return cur, size
}

func (repo *UserCrawlScoreRepository) GetCounts() (int, int) {
	return repo.getCounts()
}

func (repo *UserCrawlScoreRepository) complete(score UserCrawlScore) {
	next := score
	next.Completed = true
	repo.db.Model(&score).Update(&next)
}

func (repo *UserCrawlScoreRepository) Complete(s *domain.UserCrawlScore) {
	score := UserCrawlScore{}.in(s)
	repo.complete(score)
}

func (repo *UserCrawlScoreRepository) DropAndCreate() {
	dropAndCreateUserCrawlScore(repo.db)
}

func (repo *UserSocialLinkRepository) getAllByUser(uid uint) []UserSocialLink {
	links := []UserSocialLink{}
	repo.db.Find(&links, "user_id=?", uid)

	return links
}

func (repo *UserSocialLinkRepository) GetAllByUser(u *domain.User) []*domain.UserSocialLink {
	ls := []*domain.UserSocialLink{}
	links := repo.getAllByUser(u.ID)
	for _, link := range links {
		ls = append(ls, link.out())
	}
	return ls
}

func (repo *UserSocialLinkRepository) updateAllByUser(uid uint, nexts []UserSocialLink) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserSocialLink{}
	tx.Find(&prevs, "user_id=?", uid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		next.UserID = uid
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserSocialLinkRepository) UpdateAllByUser(user *domain.User, ls []*domain.UserSocialLink) error {
	nexts := []UserSocialLink{}
	for _, l := range ls {
		nexts = append(nexts, UserSocialLink{}.in(l))
	}

	return repo.updateAllByUser(user.ID, nexts)
}

func (repo *UserStatRepository) getByUser(uid uint) (UserStat, bool) {
	stat := UserStat{}
	notFound := repo.db.Where("user_id=?", uid).First(&stat).RecordNotFound()
	return stat, !notFound
}

func (repo *UserStatRepository) GetByUser(u *domain.User) (*domain.UserStat, bool) {
	stat, exist := repo.getByUser(u.ID)

	return stat.out(), exist
}

func (repo *UserLanguageStatRepository) updateAllByUser(uid uint, nexts []UserLanguageStat) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserLanguageStat{}
	tx.Find(&prevs, "user_id=?", uid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		next.UserID = uid
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserLanguageStatRepository) getAllByUser(uid uint) []UserLanguageStat {
	stats := []UserLanguageStat{}
	repo.db.Find(&stats, "user_id=?", uid)

	return stats
}

func (repo *UserLanguageStatRepository) GetAllByUser(u *domain.User) []*domain.UserLanguageStat {
	ss := []*domain.UserLanguageStat{}
	stats := repo.getAllByUser(u.ID)
	for _, stat := range stats {
		ss = append(ss, stat.out())
	}
	return ss
}

func (repo *UserLanguageStatRepository) UpdateAllByUser(user *domain.User, ss []*domain.UserLanguageStat) error {
	nexts := []UserLanguageStat{}
	for _, s := range ss {
		nexts = append(nexts, UserLanguageStat{}.in(s))
	}

	return repo.updateAllByUser(user.ID, nexts)
}

func (repo *UserFollowTagRepository) updateAllByUser(uid uint, nexts []UserFollowTag) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserFollowTag{}
	tx.Find(&prevs, "user_id=?", uid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		next.UserID = uid
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserFollowTagRepository) UpdateAllByUser(user *domain.User, ts []*domain.UserFollowTag) error {
	nexts := []UserFollowTag{}
	for _, t := range ts {
		nexts = append(nexts, UserFollowTag{}.in(t))
	}

	return repo.updateAllByUser(user.ID, nexts)
}

func (repo *UserFolloweeRepository) updateAllByUser(uid uint, nexts []UserFollowee) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserFollowee{}
	tx.Find(&prevs, "user_id=?", uid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		next.UserID = uid
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserFolloweeRepository) UpdateAllByUser(user *domain.User, fs []*domain.UserFollowee) error {
	nexts := []UserFollowee{}
	for _, f := range fs {
		nexts = append(nexts, UserFollowee{}.in(f))
	}

	return repo.updateAllByUser(user.ID, nexts)
}

func (repo *UserItemRepository) getByArticleID(id string) (UserItem, bool) {
	item := UserItem{}
	notFound := repo.db.Where("article_id=?", id).First(&item).RecordNotFound()
	return item, !notFound
}

func (repo *UserItemRepository) GetByArticleID(id string) (*domain.UserItem, bool) {
	item, exist := repo.getByArticleID(id)

	return item.out(), exist
}

func (repo *UserItemRepository) getAllByUser(uid uint) []UserItem {
	items := []UserItem{}
	repo.db.Find(&items, "user_id=?", uid)

	return items
}

func (repo *UserItemRepository) GetAllByUser(u *domain.User) []*domain.UserItem {
	is := []*domain.UserItem{}
	items := repo.getAllByUser(u.ID)
	for _, item := range items {
		is = append(is, item.out())
	}
	return is
}

func (repo *UserItemTagRepository) getAllByItem(iid uint) []UserItemTag {
	tags := []UserItemTag{}
	repo.db.Find(&tags, "item_id=?", iid)

	return tags
}

func (repo *UserItemTagRepository) GetAllByItem(i *domain.UserItem) []*domain.UserItemTag {
	ts := []*domain.UserItemTag{}
	tags := repo.getAllByItem(i.ID)
	for _, tag := range tags {
		ts = append(ts, tag.out())
	}
	return ts
}

func (repo *UserItemTagRepository) updateAllByItem(iid uint, nexts []UserItemTag) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserItemTag{}
	tx.Find(&prevs, "item_id=?", iid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		next.ItemID = iid
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserItemTagRepository) UpdateAllByItem(item *domain.UserItem, ts []*domain.UserItemTag) error {
	nexts := []UserItemTag{}
	for _, t := range ts {
		nexts = append(nexts, UserItemTag{}.in(t))
	}

	return repo.updateAllByItem(item.ID, nexts)
}

func (repo *UserPopularItemRepository) getAllByUser(uid uint) []UserPopularItem {
	items := []UserPopularItem{}
	repo.db.Find(&items, "user_id=?", uid)

	return items
}

func (repo *UserPopularItemRepository) GetAllByUser(u *domain.User) []*domain.UserPopularItem {
	is := []*domain.UserPopularItem{}
	items := repo.getAllByUser(u.ID)
	for _, item := range items {
		is = append(is, item.out())
	}
	return is
}

func (repo *UserPopularItemRepository) updateAllByUser(uid uint, nexts []UserPopularItem) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserPopularItem{}
	tx.Find(&prevs, "user_id=?", uid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserPopularItemRepository) UpdateAllByUser(user *domain.User, is []*domain.UserItem) error {
	nexts := []UserPopularItem{}
	for _, i := range is {
		nexts = append(nexts, UserPopularItem{
			UserID: user.ID,
			ItemID: i.ID,
		})
	}

	return repo.updateAllByUser(user.ID, nexts)
}

func (repo *UserRecentItemRepository) getAllByUser(uid uint) []UserRecentItem {
	items := []UserRecentItem{}
	repo.db.Find(&items, "user_id=?", uid)

	return items
}

func (repo *UserRecentItemRepository) GetAllByUser(u *domain.User) []*domain.UserRecentItem {
	is := []*domain.UserRecentItem{}
	items := repo.getAllByUser(u.ID)
	for _, item := range items {
		is = append(is, item.out())
	}
	return is
}

func (repo *UserRecentItemRepository) updateAllByUser(uid uint, nexts []UserRecentItem) error {
	tx := repo.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	prevs := []UserRecentItem{}
	tx.Find(&prevs, "user_id=?", uid)
	for _, prev := range prevs {
		tx.Delete(&prev)
	}

	for _, next := range nexts {
		if err := tx.Create(&next).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *UserRecentItemRepository) UpdateAllByUser(user *domain.User, is []*domain.UserItem) error {
	nexts := []UserRecentItem{}
	for _, i := range is {
		nexts = append(nexts, UserRecentItem{
			UserID: user.ID,
			ItemID: i.ID,
		})
	}

	return repo.updateAllByUser(user.ID, nexts)
}

func (repo *UserScoutRepository) getByUser(uid uint) (UserScout, bool) {
	scout := UserScout{}
	notFound := repo.db.Where("user_id=?", uid).First(&scout).RecordNotFound()
	return scout, !notFound
}

func (repo *UserScoutRepository) GetByUser(u *domain.User) (*domain.UserScout, bool) {
	scout, exist := repo.getByUser(u.ID)

	return scout.out(), exist
}
