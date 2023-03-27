package fetcher

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type Fetcher struct {
	response response
}

func (r *Fetcher) Fetch() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.reddit.com/r/golang.json", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &r.response); err != nil {
		return err
	}
	return nil
}

func (r *Fetcher) Save(w io.Writer) error {
	for _, child := range r.response.Data.Children {
		if err := json.NewEncoder(w).Encode(child.Data.Title); err != nil {
			return err
		}
		if err := json.NewEncoder(w).Encode(child.Data.URL); err != nil {
			return err
		}
	}
	return nil
}
