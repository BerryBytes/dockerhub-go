package dockerhub

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestOrganizationService_CreateOrganization(t *testing.T) {
	client, mux, teardown := makeMockClient()
	defer teardown()
	organizationName := "hamroOrg"
	companyName := "hamroCompany"

	newOrg := &Organization{}

	uri := "/orgs/"

	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, r, http.MethodPost)
		assertNoHeader(t, r, "Authorization")
		assertBody(t, r, string(mustJSONMarshal(&CreateOrganizationRequest{
			Orgname: organizationName,
			Company: companyName,
		},
		)))

		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSONMarshal(&CreateOrganizationRequest{}))
	})

	res, err := client.Organization.CreateOrganization(context.Background(), organizationName, companyName)
	if err != nil {
		t.Errorf("Organization.CreateOrganization returned error: %v", err)
	}

	if !reflect.DeepEqual(res, newOrg) {
		t.Errorf("organization is %v; want %v", res, newOrg)
	}
}

func TestOrganizationService_GetOrganizations(t *testing.T) {
	client, mux, teardown := makeMockClient()
	defer teardown()

	orgs := &OrganizationList{}
	pageSize := 100

	uri := "/user/orgs/"
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSONMarshal(orgs))
	})

	res, err := client.Organization.GetOrganizations(context.Background(), pageSize)
	if err != nil {
		t.Errorf("Organization.GetOrganizations returned error: %v", err)
	}

	if !reflect.DeepEqual(res, orgs) {
		t.Errorf("organization is %v; want %v", res, orgs)
	}

}
