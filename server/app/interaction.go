package app

type InteractionRequest struct {
	User        User     `json:"user"`
	ResponseURL string   `json:"response_url"`
	Actions     []Action `json:"actions"`
}

type Action struct {
	BlockID        string         `json:"block_id"`
	SelectedOption SelectedOption `json:"selected_option"`
}

type SelectedOption struct {
	Value string `json:"value"`
}

type User struct {
	Username string `json:"username"`
}
