package main

import (
	"backend/model"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

var (
	//go:embed fixtures/create_member.json
	memberJson string
	//go:embed fixtures/member.json
	member string
)

func (t *IntegrationTest) Test_Walkthrough() {
	t.T().Skip() // Add a room first to make this test work
	token := adminToken()
	t.Run("add member", func() {
		resp, err := t.client.R().
			SetBody(memberJson).
			SetAuthToken(token).
			Post("/api/members")
		t.Require().NoError(err)
		t.Require().Equal(http.StatusOK, resp.StatusCode())

		t.JSONEq(member, RemoveTimestamps(t.T(), resp.Body()))
	})
}

type IntegrationTest struct {
	suite.Suite

	config model.Configuration
	server *httptest.Server
	client *resty.Client
}

func Test_Integration(t *testing.T) {
	suite.Run(t, &IntegrationTest{})
}

func (t *IntegrationTest) SetupSuite() {
	config, err := loadConfig()
	t.Require().NoError(err)
	app, err := newApplication(config)
	t.Require().NoError(err)

	t.config = config
	t.server = httptest.NewServer(app.router)
	t.client = resty.New()
	t.client.BaseURL = t.server.URL
}

func (t *IntegrationTest) SetupTest() {
	err := cleanTables(t.config)
	t.Require().NoError(err)
}
