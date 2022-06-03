package main

import (
	"fmt"
	"github.com/Darklabel91/Crawler_TRT04/CSV"
	"github.com/Darklabel91/Crawler_TRT04/Crawler"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strconv"
)

const (
	xpathDecisions = "//*[@id=\"id4\"]/div/div[2]/ul/li"
)

var initDate string
var endDate string

func main() {
	//Input the required dates
	fmt.Print("-> Data Inicial: ")
	_, err := fmt.Scanln(&initDate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("-> Data Final: ")
	_, err = fmt.Scanln(&endDate)
	if err != nil {
		fmt.Println(err)
	}

	//Create Driver
	driver, err := Crawler.SeleniumWebDriver()
	if err != nil {
		fmt.Println(err)
	}
	defer driver.Close()

	//Get the source code of the page
	htmlPgSrc, err := Crawler.Search(driver, initDate, endDate, xpathDecisions)
	if err != nil {
		fmt.Println(err)
	}

	//Calculates how many pages there are
	totalLoops, err := Crawler.GetTotalLoop(driver)
	if err != nil {
		fmt.Println(err)
	}

	//Crawler
	var decPages []Crawler.Decision
	for i := 0; i < totalLoops; i++ {

		//For every page != 0 we need to get the page source again
		if i != 0 {
			htmlPgSrc, err = Crawler.GetPageSource(driver)
			if err != nil {
				fmt.Println(err)
			}
		}

		//Count how many decisions are in the page
		decisionsNodes := htmlquery.Find(htmlPgSrc, xpathDecisions)

		//Get every decision on the page
		for _, decision := range decisionsNodes {
			dc, err := Crawler.GetDecision(decision)
			if err != nil {
				decPages = append(decPages, Crawler.Decision{})
			}
			decPages = append(decPages, dc)
		}

		err = clickNext(driver, htmlPgSrc)
		if err != nil {
			fmt.Println(err)
		}

	}

	//Create CSV file with array of results
	err = CSV.ExportCsv(decPages)
	if err != nil {
		fmt.Println(err)
	}
}

func clickNext(driver selenium.WebDriver, htmlPgSrc *html.Node) error {
	totalLI := htmlquery.Find(htmlPgSrc, "//*[@id=\"id4\"]/div/div[2]/div/span/nav/ul/li")

	textTotal := strconv.Itoa(len(totalLI) - 1)

	xpathNext := "//*[@id=\"id4\"]/div/div[2]/div/span/nav/ul/li[" + textTotal + "]/a"

	btt, err := driver.FindElement(selenium.ByXPATH, xpathNext)
	if err != nil {
		return err
	}

	err = btt.Click()
	if err != nil {
		return err
	}

	return nil
}
