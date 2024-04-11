package scrap

import (
	"fmt"
	"net/http"
	"sync"
)

var Wg sync.WaitGroup

type CountryCashAcceptance struct {
	Country string
	Url     string
}

func CheckCashAcceptance(url, country string, c chan CountryCashAcceptance) {
	defer Wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s | %s", country, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		c <- CountryCashAcceptance{
			Country: country,
			Url:     url,
		}
	}
}
