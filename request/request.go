package request

import (
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

var client = &http.Client{
	Timeout: time.Minute,
}

func GetVideoUrls(url string, token string) ([]string, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Cookie", token)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	utf8Reader, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(utf8Reader)
	if err != nil {
		return nil, err
	}

	var videos []string
	document.Find("#lessons-list link[itemprop='contentUrl']").Each(func(index int, element *goquery.Selection) {
		attr, exists := element.Attr("href")
		if exists {
			videos = append(videos, attr)
		}
	})

	return videos, nil
}

func DownloadVideo(url string) (io.ReadCloser, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
