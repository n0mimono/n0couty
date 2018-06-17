package infrastructure

import (
	"n0couty/app/domain"

	"github.com/jinzhu/gorm"
)

type Registory struct {
	userChip         *UserChipRepository
	userCrawlScore   *UserCrawlScoreRepository
	user             *UserRepository
	userSocialLink   *UserSocialLinkRepository
	userStat         *UserStatRepository
	userLanguageStat *UserLanguageStatRepository
	userFollowTag    *UserFollowTagRepository
	userFollowee     *UserFolloweeRepository
	userItem         *UserItemRepository
	userItemTag      *UserItemTagRepository
	userPopularItem  *UserPopularItemRepository
	userRecentItem   *UserRecentItemRepository
	userScout        *UserScoutRepository
}

func NewRegistory(db *gorm.DB) domain.Registory {
	return &Registory{
		userChip:         &UserChipRepository{db: db},
		userCrawlScore:   &UserCrawlScoreRepository{db: db},
		user:             &UserRepository{db: db},
		userSocialLink:   &UserSocialLinkRepository{db: db},
		userStat:         &UserStatRepository{db: db},
		userLanguageStat: &UserLanguageStatRepository{db: db},
		userFollowTag:    &UserFollowTagRepository{db: db},
		userFollowee:     &UserFolloweeRepository{db: db},
		userItem:         &UserItemRepository{db: db},
		userItemTag:      &UserItemTagRepository{db: db},
		userPopularItem:  &UserPopularItemRepository{db: db},
		userRecentItem:   &UserRecentItemRepository{db: db},
		userScout:        &UserScoutRepository{db: db},
	}
}

func (registory *Registory) UserChip() domain.UserChipRepository {
	return registory.userChip
}

func (registory *Registory) UserCrawlScore() domain.UserCrawlScoreRepository {
	return registory.userCrawlScore
}

func (registory *Registory) User() domain.UserRepository {
	return registory.user
}

func (registory *Registory) UserSocialLink() domain.UserSocialLinkRepository {
	return registory.userSocialLink
}

func (registory *Registory) UserStat() domain.UserStatRepository {
	return registory.userStat
}

func (registory *Registory) UserLanguageStat() domain.UserLanguageStatRepository {
	return registory.userLanguageStat
}

func (registory *Registory) UserFollowTag() domain.UserFollowTagRepository {
	return registory.userFollowTag
}

func (registory *Registory) UserFollowee() domain.UserFolloweeRepository {
	return registory.userFollowee
}

func (registory *Registory) UserItem() domain.UserItemRepository {
	return registory.userItem
}

func (registory *Registory) UserItemTag() domain.UserItemTagRepository {
	return registory.userItemTag
}

func (registory *Registory) UserPopularItem() domain.UserPopularItemRepository {
	return registory.userPopularItem
}

func (registory *Registory) UserRecentItem() domain.UserRecentItemRepository {
	return registory.userRecentItem
}

func (registory *Registory) UserScout() domain.UserScoutRepository {
	return registory.userScout
}
