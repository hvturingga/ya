package conf

type Conf struct {
	ClashAPI string `json:"clash_api"`
}

type Provider struct {
	Repo    string   `json:"repo"`
	Owner   string   `json:"owner"`
	Version []string `json:"version"`
}
