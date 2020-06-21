package cloudflare

import "time"

type ZoneRegisterResponse struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result struct {
		ID                  string    `json:"id"`
		Name                string    `json:"name"`
		DevelopmentMode     int       `json:"development_mode"`
		OriginalNameServers []string  `json:"original_name_servers"`
		OriginalRegistrar   string    `json:"original_registrar"`
		OriginalDnshost     string    `json:"original_dnshost"`
		CreatedOn           time.Time `json:"created_on"`
		ModifiedOn          time.Time `json:"modified_on"`
		ActivatedOn         time.Time `json:"activated_on"`
		Owner               struct {
			ID struct {
			} `json:"id"`
			Email struct {
			} `json:"email"`
			Type string `json:"type"`
		} `json:"owner"`
		Account struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"account"`
		Permissions []string `json:"permissions"`
		Plan        struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			Price        int    `json:"price"`
			Currency     string `json:"currency"`
			Frequency    string `json:"frequency"`
			LegacyID     string `json:"legacy_id"`
			IsSubscribed bool   `json:"is_subscribed"`
			CanSubscribe bool   `json:"can_subscribe"`
		} `json:"plan"`
		PlanPending struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			Price        int    `json:"price"`
			Currency     string `json:"currency"`
			Frequency    string `json:"frequency"`
			LegacyID     string `json:"legacy_id"`
			IsSubscribed bool   `json:"is_subscribed"`
			CanSubscribe bool   `json:"can_subscribe"`
		} `json:"plan_pending"`
		Status      string   `json:"status"`
		Paused      bool     `json:"paused"`
		Type        string   `json:"type"`
		NameServers []string `json:"name_servers"`
	} `json:"result"`
}
