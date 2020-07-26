package cloudflare

import "time"

type UserProfile struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result struct {
		ID                             string    `json:"id"`
		Email                          string    `json:"email"`
		FirstName                      string    `json:"first_name"`
		LastName                       string    `json:"last_name"`
		Username                       string    `json:"username"`
		Telephone                      string    `json:"telephone"`
		Country                        string    `json:"country"`
		Zipcode                        string    `json:"zipcode"`
		CreatedOn                      time.Time `json:"created_on"`
		ModifiedOn                     time.Time `json:"modified_on"`
		TwoFactorAuthenticationEnabled bool      `json:"two_factor_authentication_enabled"`
		Suspended                      bool      `json:"suspended"`
	} `json:"result"`
}
