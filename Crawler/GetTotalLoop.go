package Crawler

import (
	"errors"
	"github.com/tebeka/selenium"
	"strconv"
	"strings"
)

const (
	xpathTotalReturn   = "//*[@id=\"id4\"]/div/div[1]/div/div[1]/h2/small/span"
	expectedAmountPage = 10
)

func GetTotalLoop(driver selenium.WebDriver) (int, error) {
	elemText, err := driver.FindElement(selenium.ByXPATH, xpathTotalReturn)
	if err != nil {
		return 0, err
	}

	textOriginal, err := elemText.Text()
	if err != nil {
		return 0, err
	}

	textSplit := strings.Split(textOriginal, "Exibindo ")

	if len(textSplit) > 0 {
		totalResultsText := strings.Split(textSplit[1], " resultados")

		totalResults, err := strconv.Atoi(totalResultsText[0])
		if err != nil {
			return 0, errors.New("could not convert string to int")
		}

		totalLoop := totalResults / expectedAmountPage

		if totalResults%expectedAmountPage != 0 {
			totalLoop = totalLoop + 1
		}

		return totalLoop, nil

	} else {
		return 0, errors.New("could not find the total results text")
	}
}
