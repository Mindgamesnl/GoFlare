package cloudflare

type ZoneDeleteResponse struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result struct {
		ID string `json:"id"`
	} `json:"result"`
}
