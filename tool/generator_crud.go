package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Replacer struct {
	name      string
	shortName string
}

func GenerateCRUDAll() {
	prev := Replacer{name: "UserChip", shortName: "chip"}

	GenerateCRUD(prev, Replacer{name: "UserChip", shortName: "chip"})
	GenerateCRUD(prev, Replacer{name: "UserCrawlScore", shortName: "score"})

	GenerateCRUD(prev, Replacer{name: "User", shortName: "user"})
	GenerateCRUD(prev, Replacer{name: "UserSocialLink", shortName: "link"})
	GenerateCRUD(prev, Replacer{name: "UserStat", shortName: "stat"})
	GenerateCRUD(prev, Replacer{name: "UserLanguageStat", shortName: "stat"})
	GenerateCRUD(prev, Replacer{name: "UserFollowTag", shortName: "tag"})
	GenerateCRUD(prev, Replacer{name: "UserFollowee", shortName: "followee"})
	GenerateCRUD(prev, Replacer{name: "UserItem", shortName: "item"})
	GenerateCRUD(prev, Replacer{name: "UserItemTag", shortName: "tag"})
	GenerateCRUD(prev, Replacer{name: "UserPopularItem", shortName: "item"})
	GenerateCRUD(prev, Replacer{name: "UserRecentItem", shortName: "item"})

	GenerateCRUD(prev, Replacer{name: "UserScout", shortName: "scout"})
}

func GenerateCRUD(prev Replacer, next Replacer) {
	bytes, err := ioutil.ReadFile("./tool/template_crud")
	if err != nil {
		log.Fatalln(err)
		return
	}

	lines := string(bytes)
	lines = strings.Replace(lines, prev.name, next.name, -1)
	lines = strings.Replace(lines, prev.shortName, next.shortName, -1)

	fmt.Println(lines)
	fmt.Println()
}

func main() {
	GenerateCRUDAll()
}
