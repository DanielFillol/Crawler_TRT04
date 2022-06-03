package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strings"
)

func GetPageSource(driver selenium.WebDriver) (*html.Node, error) {

	pageSource, err := driver.PageSource()
	if err != nil {
		return nil, errors.New("could not get page source")
	}

	htmlPgSrc, err := htmlquery.Parse(strings.NewReader(pageSource))
	if err != nil {
		return nil, errors.New("could not convert string to node html")
	}

	return htmlPgSrc, nil
}
