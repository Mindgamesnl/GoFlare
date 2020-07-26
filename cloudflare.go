package cloudflare

import (
	"encoding/json"
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
	cf.doRequest("GET", "https://api.cloudflare.com/client/v4/zones?name=" + domain, response, emptyMap)
	return response
}

func (cf CloudFlare) RegisterZone(domain string) ZoneRegisterResponse {
	response := ZoneRegisterResponse{}
	cf.doRequest("POST", "https://api.cloudflare.com/client/v4/zones", response, map[string]string{
		"name": domain,
	})
	return response
}

func (cf CloudFlare) DeleteZone(domain string) ZoneDeleteResponse {
	response := ZoneDeleteResponse{}
	cf.doRequest("DELETE", "https://api.cloudflare.com/client/v4/zones/" + domain, response, emptyMap)
	return response
}

func (cf CloudFlare) SetZoneDevelopmentMode(domain string, inDevelopment bool) ZoneDevModeResponse {
	response := ZoneDevModeResponse{}
	cf.doRequest("PATH", "https://api.cloudflare.com/client/v4/zones/" + domain + "/settings/development_mode", response, map[string]string{
		"value": strconv.FormatBool(inDevelopment),
	})
	return response
}

func (cf CloudFlare) GetZoneDevelopmentMode(domain string) ZoneDevModeResponse {
	response := ZoneDevModeResponse{}
	cf.doRequest("GET", "https://api.cloudflare.com/client/v4/zones/" + domain + "/settings/development_mode", response, emptyMap)
	return response
}

func (cf CloudFlare) ZoneDnsList(domain string) ZoneDnsList {
	response := ZoneDnsList{}
	cf.doRequest("GET", "https://api.cloudflare.com/client/v4/zones/" + domain + "/dns_records", response, emptyMap)
	return response
}

func (cf CloudFlare) AddDns(domain string, recordType string, content string, name string, proxied bool) ZoneDnsAddedResponse {
	response := ZoneDnsAddedResponse{}
	cf.doRequest("POST", "https://api.cloudflare.com/client/v4/zones/" + domain + "/dns_records", response, map[string]string{
		"type": recordType,
		"name": name,
		"content": content,
		"proxied": strconv.FormatBool(proxied),
	})

	return response
}

func (cf CloudFlare) ListWorkers(profile UserProfile) WorkerListResponse {
	response := WorkerListResponse{}
	cf.doRequest("GET", "https://api.cloudflare.com/client/v4/accounts/" + profile.Result.ID + "/workers/scripts", response,emptyMap)
	return response
}

func (cf CloudFlare) doRequest(how string, endpoint string, what interface{}, values map[string]string) {
	client := http.Client{}
	request, err := http.NewRequest(how, endpoint, nil)
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
