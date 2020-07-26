package cloudflare

type WorkerListResponse struct {
	Success  bool          `json:"success"`
	Errors   []Error       `json:"errors"`
	Result []Worker `json:"result"`
}
