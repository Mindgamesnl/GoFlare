package cloudflare

import "time"

type Worker struct {
	Name         string    `json:"id"`
	Etag       string    `json:"etag"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}
