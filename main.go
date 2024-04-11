package main

import (
	"aztecoByCash/iban"
	"aztecoByCash/scrap"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	const targetUrl = "https://azte.co/buy/azteco?country=%s&_rsc=aaejq"

	scrap.Wg.Add(len(iban.Countries))

	c := make(chan scrap.CountryCashAcceptance)
	for key, value := range iban.Countries {
		var targetUrlFmt = fmt.Sprintf(targetUrl, value)
		go scrap.CheckCashAcceptance(targetUrlFmt, key, c)
	}
	go func() {
		scrap.Wg.Wait()
		close(c)
	}()

	var tableBody string
	for msg := range c {
		tableBody += fmt.Sprintf("| _%s_ | %s |\n", msg.Country, msg.Url)
	}

	buildReadme(tableBody)
}

func buildReadme(content string) {
	dateScraped := time.Now().Format("2006-01-02") // Replace with dynamic date

	readmeTemplate := fmt.Sprintf(`# Azteco Cash Acceptance by Country

This script scrapes Azteco's website (https://azte.co/) to determine which countries allow cash 
payments for their services. The data is based on the information available on the website as of **_%s_**.

| **Country** | **URL** |
|---|---|
%s
`, dateScraped, content)

	err := ioutil.WriteFile("README.md", []byte(readmeTemplate), 0644)
	if err != nil {
		fmt.Println("Error writing to readme.md:", err)
	}
}
