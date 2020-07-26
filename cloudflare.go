package cloudflare

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

var emptyMap = make(map[string]string)

type CloudFlare struct {
	Email  string
	ApiKey string
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCloudFlare(email string, key string) CloudFlare {
	return CloudFlare{
		Email:  email,
		ApiKey: key,
	}
}

func (cf CloudFlare) GetZoneInfo(domain string) ZoneInfoResponse {
	response := ZoneInfoResponse{}
	cf.doBasicRequest("GET", "https://api.cloudflare.com/client/v4/zones?name=" + domain, response)
	return response
}

func (cf CloudFlare) RegisterZone(domain string) ZoneRegisterResponse {
	response := ZoneRegisterResponse{}
	cf.doHeaderRequest("POST", "https://api.cloudflare.com/client/v4/zones", response, map[string]string{
		"name": domain,
	})
	return response
}

func (cf CloudFlare) DeleteZone(domain string) ZoneDeleteResponse {
	response := ZoneDeleteResponse{}
	cf.doBasicRequest("DELETE", "https://api.cloudflare.com/client/v4/zones/" + domain, response)
	return response
}

func (cf CloudFlare) SetZoneDevelopmentMode(domain string, inDevelopment bool) ZoneDevModeResponse {
	response := ZoneDevModeResponse{}
	cf.doHeaderRequest("PATH", "https://api.cloudflare.com/client/v4/zones/" + domain + "/settings/development_mode", response, map[string]string{
		"value": strconv.FormatBool(inDevelopment),
	})
	return response
}

func (cf CloudFlare) GetZoneDevelopmentMode(domain string) ZoneDevModeResponse {
	response := ZoneDevModeResponse{}
	cf.doBasicRequest("GET", "https://api.cloudflare.com/client/v4/zones/" + domain + "/settings/development_mode", response)
	return response
}

func (cf CloudFlare) ZoneDnsList(domain string) ZoneDnsList {
	response := ZoneDnsList{}
	cf.doBasicRequest("GET", "https://api.cloudflare.com/client/v4/zones/" + domain + "/dns_records", response)
	return response
}

func (cf CloudFlare) AddDns(domain string, recordType string, content string, name string, proxied bool) ZoneDnsAddedResponse {
	response := ZoneDnsAddedResponse{}
	cf.doHeaderRequest("POST", "https://api.cloudflare.com/client/v4/zones/" + domain + "/dns_records", response, map[string]string{
		"type": recordType,
		"name": name,
		"content": content,
		"proxied": strconv.FormatBool(proxied),
	})

	return response
}

func (cf CloudFlare) ListWorkers(profile UserProfile) WorkerListResponse {
	response := WorkerListResponse{}
	cf.doBasicRequest("GET", "https://api.cloudflare.com/client/v4/accounts/" + profile.Result.ID + "/workers/scripts", response)
	return response
}

func (cf CloudFlare) doBasicRequest(how string, endpoint string, what interface{}) {
	cf.doRequest(how, endpoint, what, emptyMap, nil)
}

func (cf CloudFlare) doHeaderRequest(how string, endpoint string, what interface{}, headers map[string]string) {
	cf.doRequest(how, endpoint, what, headers, nil)
}

func (cf CloudFlare) doUploadRequest(how string, endpoint string, what interface{}, headers map[string]string, body string) {
	cf.doRequest(how, endpoint, what, headers, bytes.NewBuffer([]byte(body)))
}

func (cf CloudFlare) doRequest(how string, endpoint string, what interface{}, values map[string]string, body io.Reader) {
	client := http.Client{}
	request, err := http.NewRequest(how, endpoint, body)
	request.Header.Add("X-Auth-Key", cf.ApiKey)
	request.Header.Add("X-Auth-Email", cf.Email)

	for s := range values {
		request.PostForm.Add(s, values[s])
	}

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(resp.Body).Decode(&what)
}
