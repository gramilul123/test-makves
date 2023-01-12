package makves_repo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gramilul123/test-makves/config"
	"github.com/gramilul123/test-makves/internal/models"
	"github.com/gramilul123/test-makves/pkg/logger"
	"gotest.tools/v3/assert"
)

func TestDoStuffWithTestServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`#,id,uid,domain,cn,department,title,who,logon_count,num_logons7,num_share7,num_file7,num_ad7,num_n7,num_logons14,num_share14,num_file14,num_ad14,num_n14,num_logons30,num_share30,num_file30,num_ad30,num_n30,num_logons150,num_share150,num_file150,num_ad150,num_n150,num_logons365,num_share365,num_file365,num_ad365,num_n365,has_user_principal_name,has_mail,has_phone,flag_disabled,flag_lockout,flag_password_not_required,flag_password_cant_change,flag_dont_expire_password,owned_files,num_mailboxes,num_member_of_groups,num_member_of_indirect_groups,member_of_indirect_groups_ids,member_of_groups_ids,is_admin,is_service
		1,872,S-1-5-21-3686381713-1037878038-1682765610-1877,dev.makves.ru,Цезарь Чикольба,Склад,Кладовщик/экспедитор,Ts.Chikolba,69,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,4,2,10;11,25;26;31;32,0,0
		2,1534,S-1-5-21-3686381713-1037878038-1682765610-2544,dev.makves.ru,Тарас Шуфрич,Отдел продаж,Менеджер по продажам,T.Shufrich,56,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,4,2,10;11,25;26;49;50,0,0`))
	}))
	defer server.Close()

	cfg := config.New()
	log, err := logger.NewZapLogger(cfg)
	assert.NilError(t, err)
	api := MakvesRepo{server.Client(), log}
	result, err := api.Download(context.Background(), server.URL)
	assert.NilError(t, err)

	expect := []*models.User{
		{
			Id:                        "872",
			Uid:                       "S-1-5-21-3686381713-1037878038-1682765610-1877",
			Domain:                    "dev.makves.ru",
			Cn:                        "Цезарь Чикольба",
			Department:                "Склад",
			Title:                     "Кладовщик/экспедитор",
			Who:                       "Ts.Chikolba",
			LogonCount:                "69",
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
			MemberOfGroupsIds:         "25;26;31;32",
			IsAdmin:                   "0",
			IsService:                 "0",
		},
		{
			Id:                        "1534",
			Uid:                       "S-1-5-21-3686381713-1037878038-1682765610-2544",
			Domain:                    "dev.makves.ru",
			Cn:                        "Тарас Шуфрич",
			Department:                "Отдел продаж",
			Title:                     "Менеджер по продажам",
			Who:                       "T.Shufrich",
			LogonCount:                "56",
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
		},
	}
	assert.DeepEqual(t, expect, result)
}
