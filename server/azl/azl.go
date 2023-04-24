package azl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Response struct {
	Artistname string `json:"artist"`
	Featartist string `json:"feat,omitempty"`
	Songname   string `json:"song"`
	Lyrics     string `json:"lyrics"`
}

func Format(text string) string {
	t := strings.ToLower(text)
	t = strings.ReplaceAll(t, " ", "")
	return t
}

func FetchLyrics(artistname, songname string) (response *Response) {

	resp := &Response{}
	c := colly.NewCollector()

	c.OnHTML(".col-xs-12.col-lg-8.text-center", func(e *colly.HTMLElement) {

		var depth int = 8

		artistname := e.DOM.Find(".lyricsh").Text()
		resp.Artistname = strings.ReplaceAll(strings.TrimSpace(artistname), "Lyrics", "")

		songname := strings.ReplaceAll(e.DOM.Find(".col-xs-12.col-lg-8.text-center > b").Text(), "\"", "")
		resp.Songname = strings.TrimSpace(songname)

		feat := e.DOM.Find(".feat")

		if feat.Length() > 0 {
			depth = 10
			resp.Featartist = feat.Text()
		}

		lyrics := e.DOM.Find(".col-xs-12.col-lg-8.text-center > div:nth-child(" + strconv.Itoa(depth) + ")").Text()
		resp.Lyrics = strings.TrimSpace(lyrics)

	})

	url := fmt.Sprintf("https://www.azlyrics.com/lyrics/%s/%s.html", artistname, songname)
	c.Visit(url)
	return resp
}
