package cloudflare

import "time"

type WorkerListResponse struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result []struct {
		ID         string    `json:"id"`
		Etag       string    `json:"etag"`
		CreatedOn  time.Time `json:"created_on"`
		ModifiedOn time.Time `json:"modified_on"`
	} `json:"result"`
}
