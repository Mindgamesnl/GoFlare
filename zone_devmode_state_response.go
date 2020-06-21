package cloudflare

import "time"

type ZoneDevModeResponse struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result struct {
		ID            string    `json:"id"`
		Value         string    `json:"value"`
		Editable      bool      `json:"editable"`
		ModifiedOn    time.Time `json:"modified_on"`
		TimeRemaining int       `json:"time_remaining"`
	} `json:"result"`
}
