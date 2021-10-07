package genshin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	client       http.Client
	uid          string
	token        string
	cookieToken  string
	actionTicket string
}

func LoginBySToken(uid, sToken string) Client {
	return Client{
		client: http.Client{},
		uid:    uid,
		token:  sToken,
	}
}

func LoginByCookie(uid, cookieToken string) Client {
	return Client{
		client:      http.Client{},
		uid:         uid,
		cookieToken: cookieToken,
	}
}

func (c *Client) GetActionTicketBySToken() *ActionTicket {
	parse, _ := url.Parse(getActionTicket)
	parse.RawQuery = url.Values{
		"action_type": {"game_role"},
		"stoken":      {c.token},
		"uid":         {c.uid},
	}.Encode()
	content := c.appRequest(http.MethodGet, parse.String(), nil)
	res := &ActionTicket{}
	err := json.Unmarshal(content, res)
	if err != nil || res.Code != 0 {
		panic("unexpect response on getting action ticket. " + Byte2Str(content))
	}

	c.actionTicket = res.Data.Ticket

	return res
}

func (c *Client) GetUserGameRoles() *UserGameRoles {
	if c.actionTicket == "" {
		c.GetActionTicketBySToken()
	}

	parse, _ := url.Parse(getUserGameRoles)
	parse.RawQuery = url.Values{
		"action_ticket": {c.actionTicket},
	}.Encode()
	content := c.appRequest(http.MethodGet, parse.String(), nil)
	res := &UserGameRoles{}
	err := json.Unmarshal(content, res)
	if err != nil || res.Code != 0 {
		panic("unexpect response on getting user game roles. " + Byte2Str(content))
	}

	return res
}

func (c *Client) GetUserGameRolesByCookie() *UserGameRoles {
	if c.cookieToken == "" {
		c.GetCookieTokenBySToken()
	}

	parse, _ := url.Parse(getUserGameRolesByCookie)
	content := c.webRequest(http.MethodGet, parse.String(), nil)
	res := &UserGameRoles{}
	err := json.Unmarshal(content, res)
	if err != nil || res.Code != 0 {
		panic("unexpect response on getting user game roles. " + Byte2Str(content))
	}

	return res
}

func (c *Client) GetCookieTokenBySToken() *CookieAccount {
	parse, _ := url.Parse(getCookieAccount)
	parse.RawQuery = url.Values{
		"stoken": {c.token},
		"uid":    {c.uid},
	}.Encode()
	content := c.appRequest(http.MethodGet, parse.String(), nil)

	res := &CookieAccount{}
	err := json.Unmarshal(content, res)
	if err != nil || res.Code != 0 {
		panic("unexpect response on getting cookie account. " + Byte2Str(content))
	}

	c.cookieToken = res.Data.CookieToken

	return res
}

func (c *Client) GetRewardInfo(role *UserGameRole) *RewardInfo {
	parse, _ := url.Parse(signRewardInfo)
	parse.RawQuery = url.Values{
		"act_id": {activityId},
		"region": {role.Region},
		"uid":    {role.GameUid},
	}.Encode()
	content := c.webRequest(http.MethodGet, parse.String(), nil)

	res := &RewardInfo{}
	err := json.Unmarshal(content, res)
	if err != nil || res.Code != 0 {
		panic("unexpect response on getting reward info. " + Byte2Str(content))
	}

	return res
}

func (c *Client) Sign(role *UserGameRole) *Sign {
	body := Str2Byte(fmt.Sprintf(`{"act_id":"%s","region":"%s","uid":"%s"}`,
		activityId, role.Region, role.GameUid))

	content := c.webRequest(http.MethodPost, signRewardSign, bytes.NewBuffer(body))

	res := &Sign{}
	err := json.Unmarshal(content, res)
	if err != nil || res.Code != 0 {
		panic("unexpect response on signing. " + Byte2Str(content))
	}

	return res
}

/* internal */

func (c *Client) appRequest(method, url string, body io.Reader) []byte {
	request, _ := http.NewRequest(method, url, body)
	header := request.Header

	header.Add("DS", ds())

	header.Add("x-rpc-client_type", appClientType)
	header.Add("x-rpc-app_version", appVersion)

	header.Add("Referer", appReferer)
	header.Add("Accept-Encoding", acceptEncoding)
	header.Add("User-Agent", appUA)

	request.AddCookie(&http.Cookie{Name: "stuid", Value: c.uid})
	request.AddCookie(&http.Cookie{Name: "stoken", Value: c.token})

	resp, err := c.client.Do(request)
	if err != nil {
		panic("network error. " + err.Error())
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	return content
}

func (c *Client) webRequest(method, url string, body io.Reader) []byte {
	if c.cookieToken == "" {
		c.GetCookieTokenBySToken()
	}

	request, _ := http.NewRequest(method, url, body)
	header := request.Header

	header.Add("DS", ds())

	header.Add("x-rpc-client_type", webClientType)
	header.Add("x-rpc-app_version", appVersion)
	header.Add("x-rpc-device_id", uuid.New().String())

	header.Add("Referer", webReferer)
	header.Add("Accept-Encoding", acceptEncoding)
	header.Add("User-Agent", webUA)

	if method == http.MethodPost {
		header.Add("Content-Type", "application/json;charset=UTF-8")
	}

	request.AddCookie(&http.Cookie{Name: "account_id", Value: c.uid})
	request.AddCookie(&http.Cookie{Name: "cookie_token", Value: c.cookieToken})

	resp, err := c.client.Do(request)
	if err != nil {
		panic("network error. " + err.Error())
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	return content
}
