package conf

type Subscribe struct {
	Name      string                 `json:"name"`
	AutoAck   bool                   `json:"auto_ack"`
	Exclusive bool                   `json:"exclusive"`
	NoLocal   bool                   `json:"noLocal"`
	NoWait    bool                   `json:"no_wait"`
	Args      map[string]interface{} `json:"args"`
}
