package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"n0couty/app/domain"
)

type UserChipRepository struct {
	db *gorm.DB
}

func (n UserChip) in(i *domain.UserChip) UserChip {
	r := UserChip{}
	CopyFrom(&r, i)
	return r
}

func (r UserChip) out() *domain.UserChip {
	i := &domain.UserChip{}
	CopyTo(&r, i)
	return i
}

func (repo *UserChipRepository) create(chip UserChip) (UserChip, error) {
	err := repo.db.Create(&chip).Error
	return chip, err
}

func (repo *UserChipRepository) Create(u *domain.UserChip) (*domain.UserChip, error) {
	chip := UserChip{}.in(u)
	chip, err := repo.create(chip)
	return chip.out(), err
}

func (repo *UserChipRepository) update(chipOld UserChip, chipNext UserChip) (UserChip, error) {
	err := repo.db.Model(&chipOld).Update(&chipNext).Error
	return chipNext, err
}

func (repo *UserChipRepository) Update(uOld *domain.UserChip, uNext *domain.UserChip) (*domain.UserChip, error) {
	chipOld := UserChip{}.in(uOld)
	chipNext := UserChip{}.in(uNext)

	chip, err := repo.update(chipOld, chipNext)
	return chip.out(), err
}

func (repo *UserChipRepository) get(chip UserChip) (UserChip, bool) {
	notFound := repo.db.First(&chip).RecordNotFound()
	return chip, !notFound
}

func (repo *UserChipRepository) Get(u *domain.UserChip) (*domain.UserChip, bool) {
	chip := UserChip{}.in(u)
	chip, exist := repo.get(chip)
	return chip.out(), exist
}

func (repo *UserChipRepository) delete(chip UserChip) {
	repo.db.Delete(&chip)
}

func (repo *UserChipRepository) Delete(u *domain.UserChip) {
	chip := UserChip{}.in(u)
	repo.delete(chip)
}

type UserCrawlScoreRepository struct {
	db *gorm.DB
}

func (n UserCrawlScore) in(i *domain.UserCrawlScore) UserCrawlScore {
	r := UserCrawlScore{}
	CopyFrom(&r, i)
	return r
}

func (r UserCrawlScore) out() *domain.UserCrawlScore {
	i := &domain.UserCrawlScore{}
	CopyTo(&r, i)
	return i
}

func (repo *UserCrawlScoreRepository) create(score UserCrawlScore) (UserCrawlScore, error) {
	err := repo.db.Create(&score).Error
	return score, err
}

func (repo *UserCrawlScoreRepository) Create(u *domain.UserCrawlScore) (*domain.UserCrawlScore, error) {
	score := UserCrawlScore{}.in(u)
	score, err := repo.create(score)
	return score.out(), err
}

func (repo *UserCrawlScoreRepository) update(scoreOld UserCrawlScore, scoreNext UserCrawlScore) (UserCrawlScore, error) {
	err := repo.db.Model(&scoreOld).Update(&scoreNext).Error
	return scoreNext, err
}

func (repo *UserCrawlScoreRepository) Update(uOld *domain.UserCrawlScore, uNext *domain.UserCrawlScore) (*domain.UserCrawlScore, error) {
	scoreOld := UserCrawlScore{}.in(uOld)
	scoreNext := UserCrawlScore{}.in(uNext)

	score, err := repo.update(scoreOld, scoreNext)
	return score.out(), err
}

func (repo *UserCrawlScoreRepository) get(score UserCrawlScore) (UserCrawlScore, bool) {
	notFound := repo.db.First(&score).RecordNotFound()
	return score, !notFound
}

func (repo *UserCrawlScoreRepository) Get(u *domain.UserCrawlScore) (*domain.UserCrawlScore, bool) {
	score := UserCrawlScore{}.in(u)
	score, exist := repo.get(score)
	return score.out(), exist
}

func (repo *UserCrawlScoreRepository) delete(score UserCrawlScore) {
	repo.db.Delete(&score)
}

func (repo *UserCrawlScoreRepository) Delete(u *domain.UserCrawlScore) {
	score := UserCrawlScore{}.in(u)
	repo.delete(score)
}

type UserRepository struct {
	db *gorm.DB
}

func (n User) in(i *domain.User) User {
	r := User{}
	CopyFrom(&r, i)
	return r
}

func (r User) out() *domain.User {
	i := &domain.User{}
	CopyTo(&r, i)
	return i
}

func (repo *UserRepository) create(user User) (User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

func (repo *UserRepository) Create(u *domain.User) (*domain.User, error) {
	user := User{}.in(u)
	user, err := repo.create(user)
	return user.out(), err
}

func (repo *UserRepository) update(userOld User, userNext User) (User, error) {
	err := repo.db.Model(&userOld).Update(&userNext).Error
	return userNext, err
}

func (repo *UserRepository) Update(uOld *domain.User, uNext *domain.User) (*domain.User, error) {
	userOld := User{}.in(uOld)
	userNext := User{}.in(uNext)

	user, err := repo.update(userOld, userNext)
	return user.out(), err
}

func (repo *UserRepository) get(user User) (User, bool) {
	notFound := repo.db.First(&user).RecordNotFound()
	return user, !notFound
}

func (repo *UserRepository) Get(u *domain.User) (*domain.User, bool) {
	user := User{}.in(u)
	user, exist := repo.get(user)
	return user.out(), exist
}

func (repo *UserRepository) delete(user User) {
	repo.db.Delete(&user)
}

func (repo *UserRepository) Delete(u *domain.User) {
	user := User{}.in(u)
	repo.delete(user)
}

type UserSocialLinkRepository struct {
	db *gorm.DB
}

func (n UserSocialLink) in(i *domain.UserSocialLink) UserSocialLink {
	r := UserSocialLink{}
	CopyFrom(&r, i)
	return r
}

func (r UserSocialLink) out() *domain.UserSocialLink {
	i := &domain.UserSocialLink{}
	CopyTo(&r, i)
	return i
}

func (repo *UserSocialLinkRepository) create(link UserSocialLink) (UserSocialLink, error) {
	err := repo.db.Create(&link).Error
	return link, err
}

func (repo *UserSocialLinkRepository) Create(u *domain.UserSocialLink) (*domain.UserSocialLink, error) {
	link := UserSocialLink{}.in(u)
	link, err := repo.create(link)
	return link.out(), err
}

func (repo *UserSocialLinkRepository) update(linkOld UserSocialLink, linkNext UserSocialLink) (UserSocialLink, error) {
	err := repo.db.Model(&linkOld).Update(&linkNext).Error
	return linkNext, err
}

func (repo *UserSocialLinkRepository) Update(uOld *domain.UserSocialLink, uNext *domain.UserSocialLink) (*domain.UserSocialLink, error) {
	linkOld := UserSocialLink{}.in(uOld)
	linkNext := UserSocialLink{}.in(uNext)

	link, err := repo.update(linkOld, linkNext)
	return link.out(), err
}

func (repo *UserSocialLinkRepository) get(link UserSocialLink) (UserSocialLink, bool) {
	notFound := repo.db.First(&link).RecordNotFound()
	return link, !notFound
}

func (repo *UserSocialLinkRepository) Get(u *domain.UserSocialLink) (*domain.UserSocialLink, bool) {
	link := UserSocialLink{}.in(u)
	link, exist := repo.get(link)
	return link.out(), exist
}

func (repo *UserSocialLinkRepository) delete(link UserSocialLink) {
	repo.db.Delete(&link)
}

func (repo *UserSocialLinkRepository) Delete(u *domain.UserSocialLink) {
	link := UserSocialLink{}.in(u)
	repo.delete(link)
}

type UserStatRepository struct {
	db *gorm.DB
}

func (n UserStat) in(i *domain.UserStat) UserStat {
	r := UserStat{}
	CopyFrom(&r, i)
	return r
}

func (r UserStat) out() *domain.UserStat {
	i := &domain.UserStat{}
	CopyTo(&r, i)
	return i
}

func (repo *UserStatRepository) create(stat UserStat) (UserStat, error) {
	err := repo.db.Create(&stat).Error
	return stat, err
}

func (repo *UserStatRepository) Create(u *domain.UserStat) (*domain.UserStat, error) {
	stat := UserStat{}.in(u)
	stat, err := repo.create(stat)
	return stat.out(), err
}

func (repo *UserStatRepository) update(statOld UserStat, statNext UserStat) (UserStat, error) {
	err := repo.db.Model(&statOld).Update(&statNext).Error
	return statNext, err
}

func (repo *UserStatRepository) Update(uOld *domain.UserStat, uNext *domain.UserStat) (*domain.UserStat, error) {
	statOld := UserStat{}.in(uOld)
	statNext := UserStat{}.in(uNext)

	stat, err := repo.update(statOld, statNext)
	return stat.out(), err
}

func (repo *UserStatRepository) get(stat UserStat) (UserStat, bool) {
	notFound := repo.db.First(&stat).RecordNotFound()
	return stat, !notFound
}

func (repo *UserStatRepository) Get(u *domain.UserStat) (*domain.UserStat, bool) {
	stat := UserStat{}.in(u)
	stat, exist := repo.get(stat)
	return stat.out(), exist
}

func (repo *UserStatRepository) delete(stat UserStat) {
	repo.db.Delete(&stat)
}

func (repo *UserStatRepository) Delete(u *domain.UserStat) {
	stat := UserStat{}.in(u)
	repo.delete(stat)
}

type UserLanguageStatRepository struct {
	db *gorm.DB
}

func (n UserLanguageStat) in(i *domain.UserLanguageStat) UserLanguageStat {
	r := UserLanguageStat{}
	CopyFrom(&r, i)
	return r
}

func (r UserLanguageStat) out() *domain.UserLanguageStat {
	i := &domain.UserLanguageStat{}
	CopyTo(&r, i)
	return i
}

func (repo *UserLanguageStatRepository) create(stat UserLanguageStat) (UserLanguageStat, error) {
	err := repo.db.Create(&stat).Error
	return stat, err
}

func (repo *UserLanguageStatRepository) Create(u *domain.UserLanguageStat) (*domain.UserLanguageStat, error) {
	stat := UserLanguageStat{}.in(u)
	stat, err := repo.create(stat)
	return stat.out(), err
}

func (repo *UserLanguageStatRepository) update(statOld UserLanguageStat, statNext UserLanguageStat) (UserLanguageStat, error) {
	err := repo.db.Model(&statOld).Update(&statNext).Error
	return statNext, err
}

func (repo *UserLanguageStatRepository) Update(uOld *domain.UserLanguageStat, uNext *domain.UserLanguageStat) (*domain.UserLanguageStat, error) {
	statOld := UserLanguageStat{}.in(uOld)
	statNext := UserLanguageStat{}.in(uNext)

	stat, err := repo.update(statOld, statNext)
	return stat.out(), err
}

func (repo *UserLanguageStatRepository) get(stat UserLanguageStat) (UserLanguageStat, bool) {
	notFound := repo.db.First(&stat).RecordNotFound()
	return stat, !notFound
}

func (repo *UserLanguageStatRepository) Get(u *domain.UserLanguageStat) (*domain.UserLanguageStat, bool) {
	stat := UserLanguageStat{}.in(u)
	stat, exist := repo.get(stat)
	return stat.out(), exist
}

func (repo *UserLanguageStatRepository) delete(stat UserLanguageStat) {
	repo.db.Delete(&stat)
}

func (repo *UserLanguageStatRepository) Delete(u *domain.UserLanguageStat) {
	stat := UserLanguageStat{}.in(u)
	repo.delete(stat)
}

type UserFollowTagRepository struct {
	db *gorm.DB
}

func (n UserFollowTag) in(i *domain.UserFollowTag) UserFollowTag {
	r := UserFollowTag{}
	CopyFrom(&r, i)
	return r
}

func (r UserFollowTag) out() *domain.UserFollowTag {
	i := &domain.UserFollowTag{}
	CopyTo(&r, i)
	return i
}

func (repo *UserFollowTagRepository) create(tag UserFollowTag) (UserFollowTag, error) {
	err := repo.db.Create(&tag).Error
	return tag, err
}

func (repo *UserFollowTagRepository) Create(u *domain.UserFollowTag) (*domain.UserFollowTag, error) {
	tag := UserFollowTag{}.in(u)
	tag, err := repo.create(tag)
	return tag.out(), err
}

func (repo *UserFollowTagRepository) update(tagOld UserFollowTag, tagNext UserFollowTag) (UserFollowTag, error) {
	err := repo.db.Model(&tagOld).Update(&tagNext).Error
	return tagNext, err
}

func (repo *UserFollowTagRepository) Update(uOld *domain.UserFollowTag, uNext *domain.UserFollowTag) (*domain.UserFollowTag, error) {
	tagOld := UserFollowTag{}.in(uOld)
	tagNext := UserFollowTag{}.in(uNext)

	tag, err := repo.update(tagOld, tagNext)
	return tag.out(), err
}

func (repo *UserFollowTagRepository) get(tag UserFollowTag) (UserFollowTag, bool) {
	notFound := repo.db.First(&tag).RecordNotFound()
	return tag, !notFound
}

func (repo *UserFollowTagRepository) Get(u *domain.UserFollowTag) (*domain.UserFollowTag, bool) {
	tag := UserFollowTag{}.in(u)
	tag, exist := repo.get(tag)
	return tag.out(), exist
}

func (repo *UserFollowTagRepository) delete(tag UserFollowTag) {
	repo.db.Delete(&tag)
}

func (repo *UserFollowTagRepository) Delete(u *domain.UserFollowTag) {
	tag := UserFollowTag{}.in(u)
	repo.delete(tag)
}

type UserFolloweeRepository struct {
	db *gorm.DB
}

func (n UserFollowee) in(i *domain.UserFollowee) UserFollowee {
	r := UserFollowee{}
	CopyFrom(&r, i)
	return r
}

func (r UserFollowee) out() *domain.UserFollowee {
	i := &domain.UserFollowee{}
	CopyTo(&r, i)
	return i
}

func (repo *UserFolloweeRepository) create(followee UserFollowee) (UserFollowee, error) {
	err := repo.db.Create(&followee).Error
	return followee, err
}

func (repo *UserFolloweeRepository) Create(u *domain.UserFollowee) (*domain.UserFollowee, error) {
	followee := UserFollowee{}.in(u)
	followee, err := repo.create(followee)
	return followee.out(), err
}

func (repo *UserFolloweeRepository) update(followeeOld UserFollowee, followeeNext UserFollowee) (UserFollowee, error) {
	err := repo.db.Model(&followeeOld).Update(&followeeNext).Error
	return followeeNext, err
}

func (repo *UserFolloweeRepository) Update(uOld *domain.UserFollowee, uNext *domain.UserFollowee) (*domain.UserFollowee, error) {
	followeeOld := UserFollowee{}.in(uOld)
	followeeNext := UserFollowee{}.in(uNext)

	followee, err := repo.update(followeeOld, followeeNext)
	return followee.out(), err
}

func (repo *UserFolloweeRepository) get(followee UserFollowee) (UserFollowee, bool) {
	notFound := repo.db.First(&followee).RecordNotFound()
	return followee, !notFound
}

func (repo *UserFolloweeRepository) Get(u *domain.UserFollowee) (*domain.UserFollowee, bool) {
	followee := UserFollowee{}.in(u)
	followee, exist := repo.get(followee)
	return followee.out(), exist
}

func (repo *UserFolloweeRepository) delete(followee UserFollowee) {
	repo.db.Delete(&followee)
}

func (repo *UserFolloweeRepository) Delete(u *domain.UserFollowee) {
	followee := UserFollowee{}.in(u)
	repo.delete(followee)
}

type UserItemRepository struct {
	db *gorm.DB
}

func (n UserItem) in(i *domain.UserItem) UserItem {
	r := UserItem{}
	CopyFrom(&r, i)
	return r
}

func (r UserItem) out() *domain.UserItem {
	i := &domain.UserItem{}
	CopyTo(&r, i)
	return i
}

func (repo *UserItemRepository) create(item UserItem) (UserItem, error) {
	err := repo.db.Create(&item).Error
	return item, err
}

func (repo *UserItemRepository) Create(u *domain.UserItem) (*domain.UserItem, error) {
	item := UserItem{}.in(u)
	item, err := repo.create(item)
	return item.out(), err
}

func (repo *UserItemRepository) update(itemOld UserItem, itemNext UserItem) (UserItem, error) {
	err := repo.db.Model(&itemOld).Update(&itemNext).Error
	return itemNext, err
}

func (repo *UserItemRepository) Update(uOld *domain.UserItem, uNext *domain.UserItem) (*domain.UserItem, error) {
	itemOld := UserItem{}.in(uOld)
	itemNext := UserItem{}.in(uNext)

	item, err := repo.update(itemOld, itemNext)
	return item.out(), err
}

func (repo *UserItemRepository) get(item UserItem) (UserItem, bool) {
	notFound := repo.db.First(&item).RecordNotFound()
	return item, !notFound
}

func (repo *UserItemRepository) Get(u *domain.UserItem) (*domain.UserItem, bool) {
	item := UserItem{}.in(u)
	item, exist := repo.get(item)
	return item.out(), exist
}

func (repo *UserItemRepository) delete(item UserItem) {
	repo.db.Delete(&item)
}

func (repo *UserItemRepository) Delete(u *domain.UserItem) {
	item := UserItem{}.in(u)
	repo.delete(item)
}

type UserItemTagRepository struct {
	db *gorm.DB
}

func (n UserItemTag) in(i *domain.UserItemTag) UserItemTag {
	r := UserItemTag{}
	CopyFrom(&r, i)
	return r
}

func (r UserItemTag) out() *domain.UserItemTag {
	i := &domain.UserItemTag{}
	CopyTo(&r, i)
	return i
}

func (repo *UserItemTagRepository) create(tag UserItemTag) (UserItemTag, error) {
	err := repo.db.Create(&tag).Error
	return tag, err
}

func (repo *UserItemTagRepository) Create(u *domain.UserItemTag) (*domain.UserItemTag, error) {
	tag := UserItemTag{}.in(u)
	tag, err := repo.create(tag)
	return tag.out(), err
}

func (repo *UserItemTagRepository) update(tagOld UserItemTag, tagNext UserItemTag) (UserItemTag, error) {
	err := repo.db.Model(&tagOld).Update(&tagNext).Error
	return tagNext, err
}

func (repo *UserItemTagRepository) Update(uOld *domain.UserItemTag, uNext *domain.UserItemTag) (*domain.UserItemTag, error) {
	tagOld := UserItemTag{}.in(uOld)
	tagNext := UserItemTag{}.in(uNext)

	tag, err := repo.update(tagOld, tagNext)
	return tag.out(), err
}

func (repo *UserItemTagRepository) get(tag UserItemTag) (UserItemTag, bool) {
	notFound := repo.db.First(&tag).RecordNotFound()
	return tag, !notFound
}

func (repo *UserItemTagRepository) Get(u *domain.UserItemTag) (*domain.UserItemTag, bool) {
	tag := UserItemTag{}.in(u)
	tag, exist := repo.get(tag)
	return tag.out(), exist
}

func (repo *UserItemTagRepository) delete(tag UserItemTag) {
	repo.db.Delete(&tag)
}

func (repo *UserItemTagRepository) Delete(u *domain.UserItemTag) {
	tag := UserItemTag{}.in(u)
	repo.delete(tag)
}

type UserPopularItemRepository struct {
	db *gorm.DB
}

func (n UserPopularItem) in(i *domain.UserPopularItem) UserPopularItem {
	r := UserPopularItem{}
	CopyFrom(&r, i)
	return r
}

func (r UserPopularItem) out() *domain.UserPopularItem {
	i := &domain.UserPopularItem{}
	CopyTo(&r, i)
	return i
}

func (repo *UserPopularItemRepository) create(item UserPopularItem) (UserPopularItem, error) {
	err := repo.db.Create(&item).Error
	return item, err
}

func (repo *UserPopularItemRepository) Create(u *domain.UserPopularItem) (*domain.UserPopularItem, error) {
	item := UserPopularItem{}.in(u)
	item, err := repo.create(item)
	return item.out(), err
}

func (repo *UserPopularItemRepository) update(itemOld UserPopularItem, itemNext UserPopularItem) (UserPopularItem, error) {
	err := repo.db.Model(&itemOld).Update(&itemNext).Error
	return itemNext, err
}

func (repo *UserPopularItemRepository) Update(uOld *domain.UserPopularItem, uNext *domain.UserPopularItem) (*domain.UserPopularItem, error) {
	itemOld := UserPopularItem{}.in(uOld)
	itemNext := UserPopularItem{}.in(uNext)

	item, err := repo.update(itemOld, itemNext)
	return item.out(), err
}

func (repo *UserPopularItemRepository) get(item UserPopularItem) (UserPopularItem, bool) {
	notFound := repo.db.First(&item).RecordNotFound()
	return item, !notFound
}

func (repo *UserPopularItemRepository) Get(u *domain.UserPopularItem) (*domain.UserPopularItem, bool) {
	item := UserPopularItem{}.in(u)
	item, exist := repo.get(item)
	return item.out(), exist
}

func (repo *UserPopularItemRepository) delete(item UserPopularItem) {
	repo.db.Delete(&item)
}

func (repo *UserPopularItemRepository) Delete(u *domain.UserPopularItem) {
	item := UserPopularItem{}.in(u)
	repo.delete(item)
}

type UserRecentItemRepository struct {
	db *gorm.DB
}

func (n UserRecentItem) in(i *domain.UserRecentItem) UserRecentItem {
	r := UserRecentItem{}
	CopyFrom(&r, i)
	return r
}

func (r UserRecentItem) out() *domain.UserRecentItem {
	i := &domain.UserRecentItem{}
	CopyTo(&r, i)
	return i
}

func (repo *UserRecentItemRepository) create(item UserRecentItem) (UserRecentItem, error) {
	err := repo.db.Create(&item).Error
	return item, err
}

func (repo *UserRecentItemRepository) Create(u *domain.UserRecentItem) (*domain.UserRecentItem, error) {
	item := UserRecentItem{}.in(u)
	item, err := repo.create(item)
	return item.out(), err
}

func (repo *UserRecentItemRepository) update(itemOld UserRecentItem, itemNext UserRecentItem) (UserRecentItem, error) {
	err := repo.db.Model(&itemOld).Update(&itemNext).Error
	return itemNext, err
}

func (repo *UserRecentItemRepository) Update(uOld *domain.UserRecentItem, uNext *domain.UserRecentItem) (*domain.UserRecentItem, error) {
	itemOld := UserRecentItem{}.in(uOld)
	itemNext := UserRecentItem{}.in(uNext)

	item, err := repo.update(itemOld, itemNext)
	return item.out(), err
}

func (repo *UserRecentItemRepository) get(item UserRecentItem) (UserRecentItem, bool) {
	notFound := repo.db.First(&item).RecordNotFound()
	return item, !notFound
}

func (repo *UserRecentItemRepository) Get(u *domain.UserRecentItem) (*domain.UserRecentItem, bool) {
	item := UserRecentItem{}.in(u)
	item, exist := repo.get(item)
	return item.out(), exist
}

func (repo *UserRecentItemRepository) delete(item UserRecentItem) {
	repo.db.Delete(&item)
}

func (repo *UserRecentItemRepository) Delete(u *domain.UserRecentItem) {
	item := UserRecentItem{}.in(u)
	repo.delete(item)
}

type UserScoutRepository struct {
	db *gorm.DB
}

func (n UserScout) in(i *domain.UserScout) UserScout {
	r := UserScout{}
	CopyFrom(&r, i)
	return r
}

func (r UserScout) out() *domain.UserScout {
	i := &domain.UserScout{}
	CopyTo(&r, i)
	return i
}

func (repo *UserScoutRepository) create(scout UserScout) (UserScout, error) {
	err := repo.db.Create(&scout).Error
	return scout, err
}

func (repo *UserScoutRepository) Create(u *domain.UserScout) (*domain.UserScout, error) {
	scout := UserScout{}.in(u)
	scout, err := repo.create(scout)
	return scout.out(), err
}

func (repo *UserScoutRepository) update(scoutOld UserScout, scoutNext UserScout) (UserScout, error) {
	err := repo.db.Model(&scoutOld).Update(&scoutNext).Error
	return scoutNext, err
}

func (repo *UserScoutRepository) Update(uOld *domain.UserScout, uNext *domain.UserScout) (*domain.UserScout, error) {
	scoutOld := UserScout{}.in(uOld)
	scoutNext := UserScout{}.in(uNext)

	scout, err := repo.update(scoutOld, scoutNext)
	return scout.out(), err
}

func (repo *UserScoutRepository) get(scout UserScout) (UserScout, bool) {
	notFound := repo.db.First(&scout).RecordNotFound()
	return scout, !notFound
}

func (repo *UserScoutRepository) Get(u *domain.UserScout) (*domain.UserScout, bool) {
	scout := UserScout{}.in(u)
	scout, exist := repo.get(scout)
	return scout.out(), exist
}

func (repo *UserScoutRepository) delete(scout UserScout) {
	repo.db.Delete(&scout)
}

func (repo *UserScoutRepository) Delete(u *domain.UserScout) {
	scout := UserScout{}.in(u)
	repo.delete(scout)
}
