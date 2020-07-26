package cloudflare

import "time"

type WorkerUploadResponse struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result struct {
		Script     string    `json:"script"`
		Etag       string    `json:"etag"`
		Size       int       `json:"size"`
		ModifiedOn time.Time `json:"modified_on"`
	} `json:"result"`
}

