package dockerhub

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// OrganizationService Type Service
type OrganizationService service

// CreateOrganizationRequest struct
type CreateOrganizationRequest struct {
	Orgname string `json:"orgname"`
	Company string `json:"company"`
}

// Organization struct
type Organization struct {
	ID            string    `json:"id"`
	Orgname       string    `json:"orgname"`
	FullName      string    `json:"full_name"`
	Location      string    `json:"location"`
	Company       string    `json:"company"`
	ProfileURL    string    `json:"profile_url"`
	DateJoined    time.Time `json:"date_joined"`
	GravatarURL   string    `json:"gravatar_url"`
	GravatarEmail string    `json:"gravatar_email"`
	Type          string    `json:"type"`
}

// OrganizationList Struct
type OrganizationList struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID            string    `json:"id"`
		Orgname       string    `json:"orgname"`
		FullName      string    `json:"full_name"`
		Location      string    `json:"location"`
		Company       string    `json:"company"`
		ProfileURL    string    `json:"profile_url"`
		DateJoined    time.Time `json:"date_joined"`
		GravatarURL   string    `json:"gravatar_url"`
		GravatarEmail string    `json:"gravatar_email"`
		Type          string    `json:"type"`
	} `json:"results"`
}

// CreateOrganization Create new Organization
func (s *OrganizationService) CreateOrganization(ctx context.Context, organization, company string) (*Organization, error) {
	url := "/orgs/"
	org := CreateOrganizationRequest{
		Orgname: organization,
		Company: company,
	}

	req, err := s.client.NewRequest(http.MethodPost, url, org)

	if err != nil {
		return nil, err
	}

	res := &Organization{}

	if _, err := s.client.Do(ctx, req, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetOrganizations all organizations of user
func (s *OrganizationService) GetOrganizations(ctx context.Context, pageSize int) (*OrganizationList, error) {
	slug := fmt.Sprintf("/user/orgs/?page_size=%d", pageSize)
	req, err := s.client.NewRequest(http.MethodGet, slug, nil)
	if err != nil {
		return nil, err
	}

	res := &OrganizationList{}
	if _, err := s.client.Do(ctx, req, res); err != nil {
		return nil, err
	}
	return res, nil
}
