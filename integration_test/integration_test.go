package integration_test

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/gramilul123/test-makves/internal/models"
	"github.com/stretchr/testify/suite"
)

type e2eTestSuite struct {
	suite.Suite
	client     *http.Client
	serverHost string
}

func TestEndToEndTestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	s.serverHost = os.Getenv("SERVICE")

	s.client = &http.Client{
		//Timeout: 200 * time.Second,
	}

}

func (s e2eTestSuite) TearDownSuite() { // nolint:govet

}

func (s *e2eTestSuite) TestServicePattern() {

	// set
	req, err := http.NewRequest(http.MethodGet, "http://service:8080/set", nil)
	s.NoError(err)

	resp, err := s.client.Do(req)
	s.NoError(err)
	s.Equal(200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)

	var answer string
	err = json.Unmarshal(body, &answer)
	s.NoError(err)
	s.Equal("OK", answer)

	// get
	req, err = http.NewRequest(http.MethodGet, "http://service:8080/get-items/713,4972,ttt", nil)
	s.NoError(err)

	resp, err = s.client.Do(req)
	s.NoError(err)
	s.Equal(200, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	s.NoError(err)

	var items []models.User
	err = json.Unmarshal(body, &items)
	s.NoError(err)
	if s.Equal(2, len(items)) {
		for _, item := range items {
			if item.Id == "713" {
				s.Equal(models.User{
					Id:                        "713",
					Uid:                       "S-1-5-21-3686381713-1037878038-1682765610-1715",
					Domain:                    "dev.makves.ru",
					Cn:                        "Ефим Шамрыло",
					Department:                "Отдел продаж",
					Title:                     "Младший специалист",
					Who:                       "E.Shamrylo",
					LogonCount:                "58",
					NumLogons7:                "0",
					NumShare7:                 "0",
					NumFile7:                  "0",
					NumAd7:                    "0",
					NumN7:                     "0",
					NumLogons14:               "0",
					NumShare14:                "0",
					NumFile14:                 "0",
					NumAd14:                   "0",
					NumN14:                    "0",
					NumLogons30:               "0",
					NumShare30:                "0",
					NumFile30:                 "0",
					NumAd30:                   "0",
					NumN30:                    "0",
					NumLogons150:              "0",
					NumShare150:               "0",
					NumFile150:                "0",
					NumAd150:                  "0",
					NumN150:                   "0",
					NumLogons365:              "0",
					NumShare365:               "0",
					NumFile365:                "0",
					NumAd365:                  "0",
					NumN365:                   "0",
					HasUserPrincipalName:      "0",
					HasMail:                   "1",
					HasPhone:                  "0",
					FlagDisabled:              "0",
					FlagLockout:               "0",
					FlagPasswordNotRequired:   "0",
					FlagPasswordCantChange:    "0",
					FlagDontExpirePassword:    "0",
					OwnedFiles:                "0",
					NumMailboxes:              "0",
					NumMemberOfGroups:         "4",
					NumMemberOfIndirectGroups: "2",
					MemberOfIndirectGroupsIds: "10;11",
					MemberOfGroupsIds:         "25;26;49;50",
					IsAdmin:                   "0",
					IsService:                 "0",
				}, item)
			}
			if item.Id == "4972" {
				s.Equal(models.User{
					Id:                        "4972",
					Uid:                       "S-1-5-21-3686381713-1037878038-1682765610-5993",
					Domain:                    "dev.makves.ru",
					Cn:                        "Тимур Белоусов",
					Department:                "ОУП",
					Title:                     "Младший специалист",
					Who:                       "T.Belousov",
					LogonCount:                "57",
					NumLogons7:                "0",
					NumShare7:                 "0",
					NumFile7:                  "0",
					NumAd7:                    "0",
					NumN7:                     "0",
					NumLogons14:               "0",
					NumShare14:                "0",
					NumFile14:                 "0",
					NumAd14:                   "0",
					NumN14:                    "0",
					NumLogons30:               "0",
					NumShare30:                "0",
					NumFile30:                 "0",
					NumAd30:                   "0",
					NumN30:                    "0",
					NumLogons150:              "0",
					NumShare150:               "0",
					NumFile150:                "0",
					NumAd150:                  "0",
					NumN150:                   "0",
					NumLogons365:              "0",
					NumShare365:               "0",
					NumFile365:                "0",
					NumAd365:                  "0",
					NumN365:                   "0",
					HasUserPrincipalName:      "0",
					HasMail:                   "1",
					HasPhone:                  "0",
					FlagDisabled:              "0",
					FlagLockout:               "0",
					FlagPasswordNotRequired:   "0",
					FlagPasswordCantChange:    "0",
					FlagDontExpirePassword:    "0",
					OwnedFiles:                "2",
					NumMailboxes:              "0",
					NumMemberOfGroups:         "4",
					NumMemberOfIndirectGroups: "2",
					MemberOfIndirectGroupsIds: "10;11",
					MemberOfGroupsIds:         "25;26;39;40",
					IsAdmin:                   "0",
					IsService:                 "0",
				}, item)
			}
		}
	}
}
