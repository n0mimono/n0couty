package infrastructure

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func unwrapUserMail(v string, b bool) string {
	if b {
		return strings.Split(v, ":")[1]
	} else {
		return v
	}
}

func unwrapUserItemArticleID(v string, b bool) string {
	if b {
		return strings.Split(v, "/")[3]
	} else {
		return v
	}
}

func totimeOnPopular(v string) time.Time {
	s := regexp.MustCompile(`(.*) .`).FindStringSubmatch(v)

	if len(s) == 2 {
		t, _ := time.Parse("Jan 02, 2006", s[1])
		return t
	}
	return time.Time{}
}

func totimeOnRecent(v string) time.Time {
	s := regexp.MustCompile(`.* posted on (.*) `).FindStringSubmatch(v)

	if len(s) == 2 {
		t, _ := time.Parse("Jan 02, 2006", s[1])
		return t
	}
	return time.Time{}
}

func toint(v string) int {
	v = strings.Replace(v, " ", "", -1)
	s, _ := strconv.Atoi(v)
	return s
}

func decodeLangDataProps(code string, b bool) map[string]int {
	type decoder struct {
		Data struct {
			Columns [][]json.Number
		}
	}
	dat := decoder{}

	err := json.Unmarshal([]byte(code), &dat)
	if err != nil {
		log.Fatalln(err)
	}

	props := map[string]int{}
	for _, v := range dat.Data.Columns {
		name := v[0].String()
		quantity, _ := v[1].Int64()
		props[name] = int(quantity)
	}
	return props
}
