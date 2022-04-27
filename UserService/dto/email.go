package dto

type Email struct {
	Subject string       `json:"subject"`
	From    string       `json:"from"`
	To      string       `json:"to"`
	Message EmailMessage `json:"message"`
}

type EmailMessage struct {
	Template string                 `json:"template"`
	Params   map[string]interface{} `json:"params"`
}
