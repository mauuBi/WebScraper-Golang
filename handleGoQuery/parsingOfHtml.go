package handlegoquery

import (
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	errorhandling "github.com/mauuBi/NewScraperGo/errorHandling"
	handledatabase "github.com/mauuBi/NewScraperGo/handleDataBase"
)


func GoQueryParsing(resp *http.Response){
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorhandling.CheckError(err)

	doc.Find("ul.wp-block-post-template").Find("div.wp-block-techcrunch-card").Each(func(index int, item *goquery.Selection){
		h3 := item.Find("h3")
		title := strings.TrimSpace(h3.Text())
		url, _ := item.Find("a").Attr("href")
		author := item.Find("a.loop-card__author").Text()
		topics := item.Find("a.loop-card__cat").Text()
		linkImage, _ := item.Find("img").Attr("src")
		newArticle := handledatabase.Articles{
			Title: title,
			Url: url,
			Author: author,
			Topics: topics,
			LinkOfImage: linkImage,
		}
		handledatabase.CreatePosts(&newArticle)
	})
}

func CreateFile(data, filename string){
	newFile, err := os.Create(filename)
		errorhandling.CheckError(err)
	defer newFile.Close()
	newFile.WriteString(data)
}

func FormateDate(oldDate string) []string{
	newDate := []string{}
	for i, r:= range oldDate{
		if i == 10{
			break
		}
		newDate = append(newDate, string(r))
	}
	return newDate
}