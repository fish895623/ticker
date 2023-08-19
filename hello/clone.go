package hello

import (
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func Hello() {
	service, err := selenium.NewChromeDriverService("./chromedriver.exe", 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"--disable-gpu",
		"--headless",
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	driver.Get("https://www.google.com/finance/quote/XRP-USD")
	ele, _ := driver.FindElement(
		selenium.ByXPATH,
		"/html/body/c-wiz[2]/div/div[4]/div/main/div[2]/div[1]/c-wiz/div/div[1]/div/div[1]/div/div[1]/div/span/div/div",
	)

	price, err := ele.Text()

	if err != nil {
		panic(err)
	}
	log.Println(price)

	driver.Close()
}
