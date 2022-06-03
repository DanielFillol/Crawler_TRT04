package Crawler

import (
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
)

const (
	webSite    = "https://pesquisatextual.trt4.jus.br/pesquisas/acordaos?0"
	xpathDate1 = "//*[@id=\"data_inicial\"]"
	xpathDate2 = "//*[@id=\"data_final\"]"
	xpathBtt   = "//*[@id=\"botaoSubmit\"]"
)

func Search(driver selenium.WebDriver, dateInit string, dateEnd string, xpathWait string) (*html.Node, error) {
	err := driver.Get(webSite)
	if err != nil {
		return nil, err
	}

	dt1, err := driver.FindElement(selenium.ByXPATH, xpathDate1)
	if err != nil {
		return nil, err
	}

	err = dt1.Clear()
	if err != nil {
		return nil, err
	}

	err = dt1.SendKeys(dateInit)
	if err != nil {
		return nil, err
	}

	dt2, err := driver.FindElement(selenium.ByXPATH, xpathDate2)
	if err != nil {
		return nil, err
	}

	err = dt2.Clear()
	if err != nil {
		return nil, err
	}

	err = dt2.SendKeys(dateEnd)
	if err != nil {
		return nil, err
	}

	btt, err := driver.FindElement(selenium.ByXPATH, xpathBtt)
	if err != nil {
		return nil, err
	}

	err = btt.Click()
	if err != nil {
		return nil, err
	}

	//infinity loop
	waitXpath(driver, xpathWait)

	htmlPgSrc, err := GetPageSource(driver)
	if err != nil {
		return nil, err
	}

	return htmlPgSrc, nil
}
