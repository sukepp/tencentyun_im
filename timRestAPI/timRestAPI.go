package timRestAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	sdkappid   = 1400048262
	identifier = "admin"
	usersig    = "eJxlj11LwzAYhe-7K0JvK5qkHxRhF3PMVo0Xad2Y3pTYpO5F*7E0rivD-*6sAwOe2*c5HM7RQQi5Tyy-FGXZfjamMGOnXHSNXOxe-MGuA1kIU-ha-oPq0IFWhaiM0hMkYRhSjG0HpGoMVHA2hKyhsXAv34tp47cfnMpBTCNqK-A2wcclX9ylPFpsdHWVb-eBZFmWP6hArEjcxnPZsdeP9bBpDymvPMrmcLNNa6JGM6xXCS13Owb3L2TwuP-MhefdJrzny3BMcJ9pPptZkwZqdT5EMI5iEoUW3SvdQ9tMAsUnhfr4J67z5XwD2mxccg__"
)

type TimRestAPI struct {
	sdkappid   int
	identifier string
	usersig    string
}

func CreateRestAPI() *TimRestAPI {
	return &TimRestAPI{
		sdkappid,
		identifier,
		usersig,
	}
}

func (timRestAPI *TimRestAPI) api(service_name string, cmd_name string, req_data []byte) string {
	url_part := []string{"https://console.tim.qq.com/v4/", service_name, "/", cmd_name, "?usersig=",
		timRestAPI.usersig, "&identifier=", timRestAPI.identifier, "&sdkappid=", strconv.Itoa(timRestAPI.sdkappid),
		"&random=", strconv.Itoa(rand.Int()), "&contenttype=json"}
	url := strings.Join(url_part, "")

	body_type := "application/json;charset=utf-8"

	req := bytes.NewBuffer(req_data)
	resp, _ := http.Post(url, body_type, req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return string(body)
}

func (timRestAPI *TimRestAPI) Account_import(identifier, nick, face_url string) {
	msg := struct{ Identifier, Nick, FaceUrl string }{identifier, nick, face_url}
	if req_data, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("im_open_login_svc", "account_import", req_data)
	}
}

func (timRestAPI *TimRestAPI) Profile_portrait_set(account_id, new_name string) {
	profile_nick := struct{ Tag, Value string }{"Tag_Profile_IM_Nick", new_name}
	profile_list := []struct{ Tag, Value string }{profile_nick}
	msg := struct {
		From_Account string
		ProfileItem  []struct{ Tag, Value string }
	}{account_id, profile_list}

	if req_data, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("profile", "portrait_set", req_data)
	}
}

func (timRestAPI *TimRestAPI) Group_create_group(group_type, group_name, owner_id string) {
	msg := struct{ Type, Name, Owner_Account string }{group_type, group_name, owner_id}
	if req_data, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		timRestAPI.api("group_open_http_svc", "create_group", req_data)
	}
}