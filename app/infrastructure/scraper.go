package infrastructure

import (
	"n0couty/app/domain"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const QiitaURL = "https://qiita.com/"

type Scraper struct {
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func idxToChar(i int) string {
	if i <= 25 {
		return string(i + 65)
	} else if i <= 35 {
		return strconv.Itoa(i - 26)
	} else if i <= 36 {
		return "_"
	}
	return ""
}

func (s *Scraper) UserChip(in domain.InUserChip) domain.OutUserChip {
	// fetch
	values := url.Values{}
	values.Set("char", idxToChar(in.Char))
	values.Set("page", strconv.Itoa(in.Page))
	doc, err := goquery.NewDocument(QiitaURL + "users?" + values.Encode())
	if err != nil {
		return domain.OutUserChip{Error: err}
	}

	// scrape: users
	chips := []*domain.UserChip{}
	doc.Find(".js-hovercard").Each(func(i int, s *goquery.Selection) {
		qiitaID, _ := s.Attr("data-hovercard-target-name")
		imgURL, _ := s.Find("img").Attr("src")
		description := s.Find("p").Next().Text()

		chip := &domain.UserChip{
			QiitaID:     qiitaID,
			ImageURL:    imgURL,
			Description: description,
			Char:        in.Char,
			Page:        in.Page,
		}
		chips = append(chips, chip)
	})

	// scrape: next
	hasNext := doc.Find("a").FilterFunction(func(i int, s *goquery.Selection) bool {
		next, b := s.Attr("rel")
		return b && next == "next"
	}).Length() == 1

	// result
	return domain.OutUserChip{
		Chips:   chips,
		HasNext: hasNext,
		Error:   err,
	}
}

func (s *Scraper) UserPage(in domain.InUserPage) domain.OutUserPage {
	// fetch
	doc, err := goquery.NewDocument(QiitaURL + in.QiitaID)
	if err != nil {
		return domain.OutUserPage{Error: err}
	}

	// scrape: check ban
	if doc.Find(".er-SuspendedUser").Length() > 0 {
		return domain.OutUserPage{
			User: &domain.User{
				QiitaID: in.QiitaID,
				Ban:     true,
			},
		}
	}

	// scrape: user
	user := &domain.User{
		QiitaID:           in.QiitaID,
		Name:              doc.Find(".newUserPageProfile_fullName").Text(),
		ImageURL:          doc.Find(".newUserPageProfile_avatar").Find("img").AttrOr("src", ""),
		Description:       doc.Find(".newUserPageProfile_description").Text(),
		Mail:              unwrapUserMail(doc.Find(".fa-envelope").Parent().Find("a").Attr("href")),
		Link:              doc.Find(".fa-link").Parent().Find("a").AttrOr("href", ""),
		Organization:      doc.Find(".fa-building-o").Parent().Text(),
		Place:             doc.Find(".fa-map-marker").Parent().Text(),
		QiitaOrganization: doc.Find(".newUserPageProfile_organizations").Find("meta").AttrOr("content", ""),
		Ban:               false,
	}

	// scrape; social link
	socialLinks := []*domain.UserSocialLink{}
	keys := []string{"github", "twitter", "facebook", "linkedin", "googlePlus"}
	for i, k := range keys {
		if v, ok := doc.Find(".newUserPageProfile_socialLink-" + k).Find("a").Attr("href"); ok {
			socialLinks = append(socialLinks, &domain.UserSocialLink{
				UserID:    0,
				ServiceID: i + 1,
				URL:       v,
			})
		}
	}

	// scrape: follow tag
	followTags := []*domain.UserFollowTag{}
	doc.Find(".fa-tags").Parent().Parent().Find(".TagList__label").Each(func(i int, s *goquery.Selection) {
		followTags = append(followTags, &domain.UserFollowTag{
			UserID: 0,
			Name:   s.Text(),
		})
	})

	// scrape: follow user
	followees := []*domain.UserFollowee{}
	doc.Find(".newUserPageProfile_followees").Find("meta").Each(func(i int, s *goquery.Selection) {
		followees = append(followees, &domain.UserFollowee{
			UserID:     0,
			FolloweeID: s.AttrOr("content", ""),
		})
	})

	// scrape: stat
	statCounts := []int{}
	doc.Find(".userActivityChart_statCount").Each(func(i int, s *goquery.Selection) {
		statCounts = append(statCounts, toint(s.Text()))
	})
	if len(statCounts) < 2 {
		return domain.OutUserPage{
			User: &domain.User{
				QiitaID: in.QiitaID,
				Ban:     true,
			},
		}
	}
	stat := &domain.UserStat{
		UserID:        0,
		Items:         statCounts[0],
		Contributions: statCounts[1],
		Followers:     statCounts[2],
		Followees:     len(followees),
	}

	// scrape: laungauge stat
	langStats := []*domain.UserLanguageStat{}
	for k, v := range decodeLangDataProps(doc.Find(".js-userActivityChart").Attr("data-props")) {
		langStats = append(langStats, &domain.UserLanguageStat{
			UserID:   0,
			Name:     k,
			Quantity: v,
		})
	}

	// scrape: popular items
	populars := []domain.OutUserPageItem{}
	doc.Find(".userPopularItems_item").Each(func(i int, s *goquery.Selection) {
		item := &domain.UserItem{
			UserID:        0,
			ArticleID:     unwrapUserItemArticleID(s.Find(".userPopularItems_title").Attr("href")),
			Contributions: toint(s.Find(".userPopularItems_likeCount").Text()),
			Comments:      toint(s.Find(".fa-comment-o").Parent().Find("li").Text()),
			Title:         s.Find(".userPopularItems_title").Text(),
			Date:          totimeOnPopular(s.Find(".userPopularItems_notes").Find("li").Text()),
		}

		tags := []*domain.UserItemTag{}
		s.Find(".tagList_item").Each(func(ii int, ss *goquery.Selection) {
			tag := &domain.UserItemTag{
				Name: ss.Find("a").Text(),
			}
			tags = append(tags, tag)
		})

		populars = append(populars, domain.OutUserPageItem{
			Item:     item,
			ItemTags: tags,
		})
	})

	// scrape: recent items
	recents := []domain.OutUserPageItem{}
	doc.Find(".tableList").Find("article").Each(func(i int, s *goquery.Selection) {
		item := &domain.UserItem{
			UserID:        0,
			ArticleID:     unwrapUserItemArticleID(s.Find(".ItemLink__title").Find("a").Attr("href")),
			Contributions: toint(s.Find(".fa-like").Parent().Text()),
			Comments:      toint(s.Find(".fa-comment-o").Parent().Find("a").Text()),
			Title:         s.Find(".ItemLink__title").Find("a").Text(),
			Date:          totimeOnRecent(s.Find(".ItemLink__info").Text()),
		}

		tags := []*domain.UserItemTag{}
		s.Find(".TagList__item").Each(func(ii int, ss *goquery.Selection) {
			tag := &domain.UserItemTag{
				Name: ss.Find("a").Text(),
			}
			tags = append(tags, tag)
		})

		recents = append(recents, domain.OutUserPageItem{
			Item:     item,
			ItemTags: tags,
		})
	})

	return domain.OutUserPage{
		User:         user,
		SocialLinks:  socialLinks,
		FollowTags:   followTags,
		Followees:    followees,
		Stat:         stat,
		LangStats:    langStats,
		PopularItems: populars,
		RecentItems:  recents,
		Error:        nil,
	}
}
