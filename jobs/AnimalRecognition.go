package jobs

import (
	"fmt"
	"io"
	"net/http"
)

func DailyTask() {
	resp, err := http.Get("https://data.moa.gov.tw/api/v1/AnimalRecognition/?$top=1000&Page=1")
	if err != nil {
		fmt.Println("Error fetching data from API:", err)
		return
	}
	defer resp.Body.Close()

	// read the api data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("api", string(body))
}
