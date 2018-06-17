package domain

type InUserChip struct {
	Char int
	Page int
}

type OutUserChip struct {
	Chips   []*UserChip
	HasNext bool
	Error   error
}

type InUserPage struct {
	QiitaID string
}

type OutUserPage struct {
	User         *User
	SocialLinks  []*UserSocialLink
	FollowTags   []*UserFollowTag
	Followees    []*UserFollowee
	Stat         *UserStat
	LangStats    []*UserLanguageStat
	PopularItems []OutUserPageItem
	RecentItems  []OutUserPageItem
	Error        error
}

type OutUserPageItem struct {
	Item     *UserItem
	ItemTags []*UserItemTag
}

type Scraper interface {
	UserChip(InUserChip) OutUserChip
	UserPage(InUserPage) OutUserPage
}
