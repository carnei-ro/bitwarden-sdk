package main

import (
	"fmt"
	"os"

	sdk "github.com/bitwarden/sdk/languages/go"
	"github.com/gofrs/uuid"
)

func main() {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "https://api.bitwarden.com"
	}
	identityURL := os.Getenv("IDENTITY_URL")
	if identityURL == "" {
		identityURL = "https://identity.bitwarden.com"
	}

	bitwardenClient, _ := sdk.NewBitwardenClient(&apiURL, &identityURL)

	accessToken := os.Getenv("ACCESS_TOKEN")
	organizationIDStr := os.Getenv("ORGANIZATION_ID")
	projectName := os.Getenv("PROJECT_NAME")
	secretKey := os.Getenv("SECRET_KEY")

	if projectName == "" {
		projectName = "NewTestProject" // default value
	}

	err := bitwardenClient.AccessTokenLogin(accessToken)
	if err != nil {
		panic(err)
	}

	organizationID, err := uuid.FromString(organizationIDStr)
	if err != nil {
		panic(err)
	}

	var projectID string

	projectList, err := bitwardenClient.Projects.List(organizationID.String())
	if err != nil {
		panic(err)
	}
	for _, project := range projectList.Data {
		if project.Name == projectName {
			projectID = project.ID
		}
	}

	secretsList, err := bitwardenClient.Secrets.List(organizationID.String())
	if err != nil {
		panic(err)
	}
	var secretValue string
	for _, secret := range secretsList.Data {
		if secret.Key == secretKey {
			s, err := bitwardenClient.Secrets.Get(secret.ID)
			if err != nil {
				panic(err)
			}
			if *s.ProjectID == projectID {
				secretValue = s.Value
			}
		}
	}

	fmt.Println(secretValue)

	defer bitwardenClient.Close()
}
