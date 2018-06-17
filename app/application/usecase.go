package application

import (
	"n0couty/app/domain"
	"strconv"
	"time"
)

type App struct {
	Service *domain.Service

	State   *State
	Channel *Channel
}

type Channel struct {
	Crawl *CrawlChannel
}

func NewChannel() *Channel {
	return &Channel{
		Crawl: &CrawlChannel{
			Lock:    make(chan bool, 1),
			Stop:    make(chan bool, 1),
			Forward: nil,
		},
	}
}

type CrawlChannel struct {
	Lock    chan bool
	Stop    chan bool
	Forward chan *CrawlState
}

type State struct {
	Crawl *CrawlState
}

func NewState() *State {
	return &State{
		Crawl: &CrawlState{Stop: true},
	}
}

type CrawlState struct {
	Run         bool   `json:"run"`
	Stop        bool   `json:"stop"`
	Max         int    `json:"max"`
	Now         int    `json:"now"`
	QiitaID     string `json:"qiitaId"`
	Description string `json:"description"`
	Score       int    `json:"score"`
	Interval    int    `json:"interval"`
	Message     string `json:"message"`
}

type InStartCrawl struct {
	Start bool
}

type OutStartCrawl struct {
}

func (app *App) UcStartCrawl(in InStartCrawl) OutStartCrawl {
	app.Channel.Crawl.Lock <- true

	if in.Start {
		if app.State.Crawl.Run {
		} else {
			app.State.Crawl.Run = true

			go app.ProcCrawl()
			<-app.Channel.Crawl.Stop // wait for start
		}
	} else {
		app.State.Crawl.Run = false
		<-app.Channel.Crawl.Stop // wait for stop
	}

	<-app.Channel.Crawl.Lock
	return OutStartCrawl{}
}

func (app *App) ProcCrawl() {
	canContinue := func() bool {
		return app.State.Crawl.Run
	}

	onCount := func(cur int, size int) {
		app.State.Crawl.Max = size
		app.State.Crawl.Now = cur
		app.Channel.Crawl.Forward <- app.State.Crawl
	}
	onLoad := func(qiitaID string, description string, score int) {
		app.State.Crawl.QiitaID = qiitaID
		app.State.Crawl.Description = description
		app.State.Crawl.Score = score
	}
	onScrape := func() {
		interval := 2 * time.Second
		time.Sleep(interval)

		app.State.Crawl.Interval = int(interval)
	}
	onError := func(err error) {
		app.State.Crawl.Message = err.Error()
	}
	onComplete := func() {
		app.State.Crawl.Stop = true
		app.Channel.Crawl.Stop <- true
		app.Channel.Crawl.Forward <- app.State.Crawl
	}

	app.State.Crawl.Stop = false
	app.Channel.Crawl.Stop <- false
	app.Service.ScrapeUserPageAll(canContinue, onCount, onLoad, onScrape, onError, onComplete)
}

type InGetUsers struct {
	Page        int
	PerPage     int
	OnlyStarred bool
}

type OutGetUsers struct {
	Users []*domain.UserSummary `json:"users"`
	Page  string                `json:"page"`
	Prev  string                `json:"prev"`
	Next  string                `json:"next"`
}

func (app *App) UcGetUsers(in InGetUsers) OutGetUsers {
	summaries, prev, next := app.Service.GetUserSummaries(in.Page, in.PerPage, in.OnlyStarred)

	return OutGetUsers{
		Users: summaries,
		Page:  strconv.Itoa(in.Page),
		Prev:  strconv.Itoa(prev),
		Next:  strconv.Itoa(next),
	}
}

type InGetDetail struct {
	Id uint
}

type OutGetDetail struct {
	Summary *domain.UserSummary `json:"summary"`
	Exist   bool                `json:"exist"`
}

func (app *App) UcGetDetail(in InGetDetail) OutGetDetail {
	summary, exist := app.Service.GetUserSummary(in.Id)

	return OutGetDetail{
		Summary: summary,
		Exist:   exist,
	}
}

type InGetUserItems struct {
	Id uint
}

type OutGetUserItems struct {
	Summary *domain.UserItemSummary `json:"summary"`
	Exist   bool                    `json:"exist"`
}

func (app *App) UcGetUserItems(in InGetUserItems) OutGetUserItems {
	summary, exist := app.Service.GetUserItemSummary(in.Id)

	return OutGetUserItems{
		Summary: summary,
		Exist:   exist,
	}
}

type InUpdateUserScout struct {
	Id   uint
	Star bool
}

type OutpdateUserScout struct {
	Summary *domain.UserSummary `json:"summary"`
}

func (app *App) UcUpdateUserScout(in InUpdateUserScout) OutpdateUserScout {
	app.Service.UpdateUserStar(in.Id, in.Star)
	summary, _ := app.Service.GetUserSummary(in.Id)

	return OutpdateUserScout{
		Summary: summary,
	}
}
