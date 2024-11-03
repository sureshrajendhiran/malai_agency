package datetime

import (
	"log"
	"time"

	"github.com/nleeper/goment"
)

func CurrentDateTime() string {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	g, err := goment.New(now)
	if err != nil {
		log.Println("(CurrentDateTime) error", err)
		return ""
	}
	return g.Format("YYYY-MM-DD HH:mm:ss")
}

func CurrentDate() string {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	g, err := goment.New(now)
	if err != nil {
		log.Println("(CurrentDateTime) error", err)
		return ""
	}
	return g.Format("YYYY-MM-DD 00:00:00")
}

func DateFormat(date string, format string) string {
	dateObj, _ := goment.New(date)
	return dateObj.Format(format)
}

func CurrentDateWithFormat(format string) string {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	g, err := goment.New(now)
	if err != nil {
		log.Println("(CurrentDateWithFormat) error", err)
		return ""
	}
	return g.Format(format)
}
