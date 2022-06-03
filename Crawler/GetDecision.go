package Crawler

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Decision struct {
	Title        string
	Date         string
	Court        string
	Judge        string
	DecisionText string
	LinkDecision string
}

const (
	linkPart1         = "https://pesquisatextual.trt4.jus.br/pesquisas/"
	XpathTitle        = "/div[1]/p/a/span[2]"
	XpathDate         = "/div[2]/span/span"
	XpathCourt        = "/div[3]/span/span"
	XpathJudge        = "/div[4]/span/span"
	XpathDecisionText = "/div[5]/p/span"
	XpathLinkDecision = "/div[6]/a"
)

func GetDecision(pgSrc *html.Node) (Decision, error) {
	var title string
	tt := htmlquery.Find(pgSrc, XpathTitle)

	if len(tt) > 0 {
		title = htmlquery.InnerText(htmlquery.FindOne(pgSrc, XpathTitle))
	} else {
		return Decision{}, errors.New("could not find title")
	}

	var date string
	dt := htmlquery.Find(pgSrc, XpathDate)
	if len(dt) > 0 {
		date = htmlquery.InnerText(htmlquery.FindOne(pgSrc, XpathDate))
	} else {
		return Decision{}, errors.New("could not find date")
	}

	var court string
	ct := htmlquery.Find(pgSrc, XpathCourt)
	if len(ct) > 0 {
		court = htmlquery.InnerText(htmlquery.FindOne(pgSrc, XpathCourt))
	} else {
		return Decision{}, errors.New("could not find court")
	}

	var judge string
	jd := htmlquery.Find(pgSrc, XpathJudge)
	if len(jd) > 0 {
		judge = htmlquery.InnerText(htmlquery.FindOne(pgSrc, XpathJudge))
	} else {
		return Decision{}, errors.New("could not find judge")
	}

	var text string
	txt := htmlquery.Find(pgSrc, XpathDecisionText)
	if len(txt) > 0 {
		text = htmlquery.InnerText(htmlquery.FindOne(pgSrc, XpathDecisionText))
	} else {
		return Decision{}, errors.New("could not find text")
	}

	var link string
	lk := htmlquery.Find(pgSrc, XpathLinkDecision)
	if len(lk) > 0 {
		link = htmlquery.SelectAttr(htmlquery.FindOne(pgSrc, XpathLinkDecision), "href")
	} else {
		return Decision{}, errors.New("could not find link")
	}

	return Decision{
		Title:        title,
		Date:         date,
		Court:        court,
		Judge:        judge,
		DecisionText: text,
		LinkDecision: linkPart1 + link,
	}, nil
}
