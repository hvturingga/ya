package conf

type Provider struct {
	Repo    string `json:"repo"`
	Owner   string `json:"owner"`
	Version string `json:"version"`
}
