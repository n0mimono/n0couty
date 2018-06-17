package main

import (
	"fmt"
	"log"

	"n0couty/app/application"
	"n0couty/app/config"
	"n0couty/app/domain"
	"n0couty/app/infrastructure"
)

func main() {
	fmt.Println("n0couty: start")

	// config
	config.Init()

	// db: init
	db, err := infrastructure.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db: migrate
	//infrastructure.DropUserTables(db)
	infrastructure.Migrate(db)

	// di
	registory := infrastructure.NewRegistory(db)
	scraper := infrastructure.NewScraper()
	service := &domain.Service{
		Registory: registory,
		Scraper:   scraper,
	}
	app := &application.App{
		Service: service,
		State:   application.NewState(),
		Channel: application.NewChannel(),
	}

	// app
	app.Listen()

	//app.ScrapeUserChips()
	//app.ScrapeUserPageTest()
	//app.InitCrawlScores()
	//app.ScrapeUserPageAll()

	fmt.Println("n0couty: fin..")
}
