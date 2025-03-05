package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TVShow struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image Image  `json:"image"`
}

type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

func main() {
	// Used for requesting each page with a list of tvshows when sending the request to the TVMaze api
	page := 0

	// Create a map to store TVShows
	tvShowMap := make(map[int]TVShow)

	// Start timer for how long it takes to get the shows into the map from when we send the GET request to having the map populated.
	downloadStart := time.Now()

	//"while"
	for {
		page++
		// NOTE TO SELF: The first item has id 250. I spend way too long not understanding why items below id 250 were empty 😆.
		// I learned it was because of the above *Facepalm*
		url := fmt.Sprintf("https://api.tvmaze.com/shows?page=%d", page)

		resp, err := http.Get(url)
		if resp.StatusCode != 200 {
			break
		}

		if err != nil {
			// apparently if something is nil in Go, it means it succeeded. Gotta get used to that ^^
			fmt.Println("Error", err)
			return
		}

		defer resp.Body.Close() // Make sure we close the body to prevent resource leakage

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading the response", err)
			return
		}

		var allShows []TVShow
		err = json.Unmarshal(body, &allShows)
		if err != nil {
			fmt.Println("Error parsing the JSON", err)
			return
		}

		// Check that the array contains shows and perform logic
		if len(allShows) != 0 {
			// Populate the map. The "_," means we ignore the index. Alternatively we could do "for i, show := range tvShows"
			for _, show := range allShows {
				tvShowMap[show.ID] = show
			}
		}
	}

	downloadElapse := time.Since(downloadStart).Milliseconds()
	fmt.Println("Number of TVShows in map: ", len(tvShowMap))
	fmt.Printf("Total time to download in milliseconds: %d ms\n", downloadElapse)

	fmt.Println()
	uploadStart := time.Now()
	go func() {
		for _, show := range tvShowMap {
			UploadShowToMyHeartcoreProject(show)
		}
	}()
	uploadElapse := time.Since(uploadStart).Milliseconds()
	fmt.Printf("Total time to upload in milliseconds: %d ms\n", uploadElapse)

}

func UploadShowToMyHeartcoreProject(show TVShow) {
	return
}
