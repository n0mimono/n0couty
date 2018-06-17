package application

import (
	"fmt"
	"log"
	"n0couty/app/domain"
	"time"
)

func (app *App) ScrapeUserChips() {
	for i := 0; i < 37; i++ {
		char := i

		canContinue := func(page int, hasNext bool) bool {
			return hasNext
		}
		onError := func(err error) {
			log.Println(err)
		}
		onNext := func(page int, chips []*domain.UserChip, hasNext bool) {
			log.Println(fmt.Sprintf("%d, %d: %d -> %v", char, page, len(chips), hasNext))

			time.Sleep(2 * time.Second)
		}

		app.Service.ScrapeUserChips(char, canContinue, onError, onNext)
	}
}

func (app *App) ScrapeUserPageTest() {
	qiitaID := "n0mimono"

	app.Service.ScrapeUserPage(qiitaID, func(err error) {
		log.Println(err)
	})
}

func (app *App) ScrapeUserPageAll() {
	canContinue := func() bool {
		return true
	}
	onCount := func(cur int, size int) {
		log.Println(fmt.Sprintf("%d / %d", cur, size))
	}
	onLoad := func(qiitaID string, description string, score int) {
		log.Println(fmt.Sprintf("%s -> %d: %s", qiitaID, score, description))
	}
	onScrape := func() {
		time.Sleep(2 * time.Second)
	}
	onError := func(err error) {
		log.Println(err)
	}
	onComplete := func() {
		log.Println("complete.")
	}

	app.Service.ScrapeUserPageAll(canContinue, onCount, onLoad, onScrape, onError, onComplete)
}

func (app *App) InitCrawlScores() {
	onError := func(err error) {
		log.Println(err)
	}

	app.Service.InitCrawlScores(onError)
}
