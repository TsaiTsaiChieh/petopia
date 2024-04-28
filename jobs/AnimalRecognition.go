package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseURL string

func DailyTask() {
	baseURL = os.Getenv("baseURL")
	if baseURL == "" {
		fmt.Println("Error: openDataUrl environment variable is not set.")
		return
	}
	pageCount := getTotalPage()
	fmt.Println(pageCount)
}

func getTotalPage() int {
	pageCount := 0
	for {
		url := fmt.Sprintf("%s&$top=1000&$skip=%d", baseURL, 1000*pageCount)
		fmt.Println(url)
		// send the request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data from API:", err)
			return pageCount
		}
		defer resp.Body.Close()

		// read the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return pageCount
		}

		// decode
		var result []interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return pageCount
		}

		if len(result) == 0 {
			break
		}
		pageCount++
	}
	return pageCount
}
