package domain

import (
	"time"
)

const (
	SocialServiceQiita      = 0
	SocialServiceGitHub     = 1
	SocialServiceTwitter    = 2
	SocialServiceFaceBook   = 3
	SocialServiceLinkedIn   = 4
	SocialServiceGooglePlus = 5
)

type UserChip struct {
	ID          uint
	QiitaID     string
	ImageURL    string
	Description string
	Char        int
	Page        int
}

type UserCrawlScore struct {
	ID        uint
	ChipID    uint
	Score     int
	Checked   bool
	Completed bool
}

type UserSummary struct {
	User      *User               `json:"user"`
	Links     []*UserSocialLink   `json:"links"`
	Stat      *UserStat           `json:"stat"`
	LangStats []*UserLanguageStat `json:"langStats"`
	Scout     *UserScout          `json:"scout"`
}

type User struct {
	ID                uint   `json:"id"`
	QiitaID           string `json:"qiitaId"`
	Name              string `json:"name"`
	ImageURL          string `json:"imageUrl"`
	Description       string `json:"description"`
	Mail              string `json:"mail"`
	Link              string `json:"link"`
	Organization      string `json:"organization"`
	Place             string `json:"place"`
	QiitaOrganization string `json:"qiitaOrganization"`
	Ban               bool   `json:"ban"`
}

type UserSocialLink struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userId"`
	ServiceID int    `json:"serviceId"`
	URL       string `json:"url"`
}

type UserStat struct {
	ID            uint `json:"id"`
	UserID        uint `json:"userId"`
	Items         int  `json:"items"`
	Contributions int  `json:"contributions"`
	Followers     int  `json:"followers"`
	Followees     int  `json:"followees"`
}

type UserLanguageStat struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"userId"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type UserFollowTag struct {
	ID     uint
	UserID uint
	Name   string
}

type UserFollowee struct {
	ID         uint
	UserID     uint
	FolloweeID string
}

type UserItemSummary struct {
	Items    []*UserItemWithTags `json:"items"`
	Populars []*UserPopularItem  `json:"populars"`
	Recents  []*UserRecentItem   `json:"recents"`
}

type UserItemWithTags struct {
	Body *UserItem      `json:"body"`
	Tags []*UserItemTag `json:"tags"`
}

type UserItem struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"userId"`
	ArticleID     string    `json:"articleId"`
	Contributions int       `json:"contributions"`
	Comments      int       `json:"comments"`
	Title         string    `json:"title"`
	Date          time.Time `json:"date"`
}

type UserItemTag struct {
	ID     uint   `json:"id"`
	ItemID uint   `json:"itemId"`
	Name   string `json:"name"`
}

type UserPopularItem struct {
	ID     uint `json:"id"`
	UserID uint `json:"userId"`
	ItemID uint `json:"itemId"`
}

type UserRecentItem struct {
	ID     uint `json:"id"`
	UserID uint `json:"userId"`
	ItemID uint `json:"itemId"`
}

type UserScout struct {
	ID      uint  `json:"id"`
	UserID  uint  `json:"userId"`
	Starred *bool `json:"starred"`
}

type Registory interface {
	UserChip() UserChipRepository
	UserCrawlScore() UserCrawlScoreRepository
	User() UserRepository
	UserSocialLink() UserSocialLinkRepository
	UserStat() UserStatRepository
	UserLanguageStat() UserLanguageStatRepository
	UserFollowTag() UserFollowTagRepository
	UserFollowee() UserFolloweeRepository
	UserItem() UserItemRepository
	UserItemTag() UserItemTagRepository
	UserPopularItem() UserPopularItemRepository
	UserRecentItem() UserRecentItemRepository
	UserScout() UserScoutRepository
}

type UserChipRepository interface {
	Create(*UserChip) (*UserChip, error)
	Update(*UserChip, *UserChip) (*UserChip, error)
	Get(*UserChip) (*UserChip, bool)
}

type UserCrawlScoreRepository interface {
	Create(*UserCrawlScore) (*UserCrawlScore, error)
	Update(*UserCrawlScore, *UserCrawlScore) (*UserCrawlScore, error)
	Get(*UserCrawlScore) (*UserCrawlScore, bool)

	GetCandidate() (*UserCrawlScore, bool)
	GetCounts() (int, int)
	Complete(*UserCrawlScore)
	DropAndCreate()
}

type UserRepository interface {
	Create(*User) (*User, error)
	Update(*User, *User) (*User, error)
	Get(*User) (*User, bool)

	GetByQiitaID(string) (*User, bool)
	GetRange(int, int, bool) []*User
}

type UserSocialLinkRepository interface {
	Create(*UserSocialLink) (*UserSocialLink, error)
	Update(*UserSocialLink, *UserSocialLink) (*UserSocialLink, error)
	Get(*UserSocialLink) (*UserSocialLink, bool)

	GetAllByUser(*User) []*UserSocialLink
	UpdateAllByUser(*User, []*UserSocialLink) error
}

type UserStatRepository interface {
	Create(*UserStat) (*UserStat, error)
	Update(*UserStat, *UserStat) (*UserStat, error)
	Get(*UserStat) (*UserStat, bool)

	GetByUser(*User) (*UserStat, bool)
}

type UserLanguageStatRepository interface {
	Create(*UserLanguageStat) (*UserLanguageStat, error)
	Update(*UserLanguageStat, *UserLanguageStat) (*UserLanguageStat, error)
	Get(*UserLanguageStat) (*UserLanguageStat, bool)

	GetAllByUser(*User) []*UserLanguageStat
	UpdateAllByUser(*User, []*UserLanguageStat) error
}

type UserFollowTagRepository interface {
	Create(*UserFollowTag) (*UserFollowTag, error)
	Update(*UserFollowTag, *UserFollowTag) (*UserFollowTag, error)
	Get(*UserFollowTag) (*UserFollowTag, bool)

	UpdateAllByUser(*User, []*UserFollowTag) error
}

type UserFolloweeRepository interface {
	Create(*UserFollowee) (*UserFollowee, error)
	Update(*UserFollowee, *UserFollowee) (*UserFollowee, error)
	Get(*UserFollowee) (*UserFollowee, bool)

	UpdateAllByUser(*User, []*UserFollowee) error
}

type UserItemRepository interface {
	Create(*UserItem) (*UserItem, error)
	Update(*UserItem, *UserItem) (*UserItem, error)
	Get(*UserItem) (*UserItem, bool)

	GetByArticleID(string) (*UserItem, bool)
	GetAllByUser(*User) []*UserItem
}

type UserItemTagRepository interface {
	Create(*UserItemTag) (*UserItemTag, error)
	Update(*UserItemTag, *UserItemTag) (*UserItemTag, error)
	Get(*UserItemTag) (*UserItemTag, bool)

	GetAllByItem(*UserItem) []*UserItemTag
	UpdateAllByItem(*UserItem, []*UserItemTag) error
}

type UserPopularItemRepository interface {
	Create(*UserPopularItem) (*UserPopularItem, error)
	Update(*UserPopularItem, *UserPopularItem) (*UserPopularItem, error)
	Get(*UserPopularItem) (*UserPopularItem, bool)

	GetAllByUser(*User) []*UserPopularItem
	UpdateAllByUser(*User, []*UserItem) error
}

type UserRecentItemRepository interface {
	Create(*UserRecentItem) (*UserRecentItem, error)
	Update(*UserRecentItem, *UserRecentItem) (*UserRecentItem, error)
	Get(*UserRecentItem) (*UserRecentItem, bool)

	GetAllByUser(*User) []*UserRecentItem
	UpdateAllByUser(*User, []*UserItem) error
}

type UserScoutRepository interface {
	Create(*UserScout) (*UserScout, error)
	Update(*UserScout, *UserScout) (*UserScout, error)
	Get(*UserScout) (*UserScout, bool)

	GetByUser(*User) (*UserScout, bool)
}
