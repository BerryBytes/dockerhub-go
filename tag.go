package dockerhub

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// TagService Type
type TagService service

// Tags response
type Tags struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Creator int         `json:"creator"`
		ID      int         `json:"id"`
		ImageID interface{} `json:"image_id"`
		Images  []struct {
			Architecture string      `json:"architecture"`
			Features     string      `json:"features"`
			Variant      interface{} `json:"variant"`
			Digest       string      `json:"digest"`
			Os           string      `json:"os"`
			OsFeatures   string      `json:"os_features"`
			OsVersion    interface{} `json:"os_version"`
			Size         int         `json:"size"`
			Status       string      `json:"status"`
			LastPulled   time.Time   `json:"last_pulled"`
			LastPushed   time.Time   `json:"last_pushed"`
		} `json:"images"`
		LastUpdated         time.Time `json:"last_updated"`
		LastUpdater         int       `json:"last_updater"`
		LastUpdaterUsername string    `json:"last_updater_username"`
		Name                string    `json:"name"`
		Repository          int       `json:"repository"`
		FullSize            int       `json:"full_size"`
		V2                  bool      `json:"v2"`
		TagStatus           string    `json:"tag_status"`
		TagLastPulled       time.Time `json:"tag_last_pulled"`
		TagLastPushed       time.Time `json:"tag_last_pushed"`
	} `json:"results"`
}

// GetTags of the repo
func (s *TagService) GetTags(ctx context.Context, namespace, repo string, page int) (*Tags, error) {
	slug := fmt.Sprintf("/repositories/%v/%v/tags/?page_size=%d&ordering=last_updated", namespace, repo, page)

	req, err := s.client.NewRequest(http.MethodGet, slug, nil)
	if err != nil {
		return nil, err
	}

	res := &Tags{}
	if _, err := s.client.Do(ctx, req, res); err != nil {
		return nil, err
	}
	return res, nil
}
