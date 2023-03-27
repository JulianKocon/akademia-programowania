package main

import (
	"io"
	"log"
	"os"
	"reddit/fetcher"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	f = &fetcher.Fetcher{}
	if err := f.Fetch(); err != nil {
		log.Fatal("Could not download data: ", err)
	}

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Error while creating file", err)
	}
	defer file.Close()

	w = file
	if err := f.Save(w); err != nil {
		log.Fatal("Error while saving: ", err)
	}
}
