package responses

type ResponseAddSessionKey struct {
	NodeType   float64 `json:"node_type"`
	UID        []byte  `json:"uid,omitempty"`
	PrivateKey string  `json:"private_key,omitempty"`
	Result     string  `json:"result"`
}
