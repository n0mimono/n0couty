package infrastructure

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"n0couty/app/config"
)

func Open() (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	return gorm.Open("mysql", args)
}

func Migrate(db *gorm.DB) {

	if !db.HasTable(&UserChip{}) {
		db.CreateTable(&UserChip{})
	}

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})

		db.CreateTable(&UserSocialLink{})
		db.CreateTable(&UserStat{})
		db.CreateTable(&UserLanguageStat{})
		db.CreateTable(&UserFollowTag{})
		db.CreateTable(&UserFollowee{})

		db.CreateTable(&UserItem{})
		db.CreateTable(&UserItemTag{})
		db.CreateTable(&UserPopularItem{})
		db.CreateTable(&UserRecentItem{})

		db.CreateTable(&UserScout{})

		db.Table("user_social_links").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Table("user_stats").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Table("user_language_stats").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Table("user_follow_tags").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Table("user_followees").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

		db.Table("user_items").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Table("user_item_tags").
			AddForeignKey("item_id", "user_items(id)", "RESTRICT", "RESTRICT")
		db.Table("user_popular_items").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").
			AddForeignKey("item_id", "user_items(id)", "RESTRICT", "RESTRICT")
		db.Table("user_recent_items").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").
			AddForeignKey("item_id", "user_items(id)", "RESTRICT", "RESTRICT")

		db.Table("user_scouts").
			AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	}
}

func dropAndCreateUserCrawlScore(db *gorm.DB) {
	if db.HasTable(&UserCrawlScore{}) {
		db.DropTable(&UserCrawlScore{})
	}

	db.CreateTable(&UserCrawlScore{})
	db.Table("user_crawl_scores").
		AddForeignKey("chip_id", "user_chips(id)", "RESTRICT", "RESTRICT")
}

func DropUserTables(db *gorm.DB) {
	if db.HasTable(&User{}) {
		db.DropTable(&UserSocialLink{})
		db.DropTable(&UserStat{})
		db.DropTable(&UserLanguageStat{})
		db.DropTable(&UserFollowTag{})
		db.DropTable(&UserFollowee{})

		db.DropTable(&UserItemTag{})
		db.DropTable(&UserPopularItem{})
		db.DropTable(&UserRecentItem{})

		db.DropTable(&UserScout{})

		db.DropTable(&UserItem{})
		db.DropTable(&User{})
	}
}

type UserChip struct {
	gorm.Model
	QiitaID     string `gorm:"unique"`
	ImageURL    string
	Description string
	Char        int
	Page        int
}

type UserCrawlScore struct {
	gorm.Model
	ChipID    uint
	Score     int
	Checked   bool
	Completed bool
}

type User struct {
	gorm.Model
	QiitaID           string `gorm:"unique"`
	Name              string
	ImageURL          string
	Description       string
	Mail              string
	Link              string
	Organization      string
	Place             string
	QiitaOrganization string
	Ban               bool
}

type UserSocialLink struct {
	gorm.Model
	UserID    uint
	ServiceID int
	URL       string
}

type UserStat struct {
	gorm.Model
	UserID        uint
	Items         int
	Contributions int
	Followers     int
	Followees     int
}

type UserLanguageStat struct {
	gorm.Model
	UserID   uint
	Name     string
	Quantity int
}

type UserFollowTag struct {
	gorm.Model
	UserID uint
	Name   string
}

type UserFollowee struct {
	gorm.Model
	UserID     uint
	FolloweeID string
}

type UserItem struct {
	gorm.Model
	UserID        uint
	ArticleID     string `gorm:"unique"`
	Contributions int
	Comments      int
	Title         string
	Date          time.Time
}

type UserItemTag struct {
	gorm.Model
	ItemID uint
	Name   string
}

type UserPopularItem struct {
	gorm.Model
	UserID uint
	ItemID uint
}

type UserRecentItem struct {
	gorm.Model
	UserID uint
	ItemID uint
}

type UserScout struct {
	gorm.Model
	UserID  uint
	Starred *bool
}
