package conf

type ExchangeT string

const (
	Direct  ExchangeT = "direct"
	Fanout  ExchangeT = "fanout"
	Headers ExchangeT = "headers"
	Topic   ExchangeT = "topic"
)

type Exchange struct {
	Name       string                 `json:"name"`
	Types      ExchangeT              `json:"types"`
	Durable    bool                   `json:"durable"`
	AutoDelete bool                   `json:"auto_delete"`
	Internal   bool                   `json:"internal"`
	NoWait     bool                   `json:"no_wait"`
	IfUnused   bool                   `json:"if_unused"`
	IfEmpty    bool                   `json:"if_empty"`
	Args       map[string]interface{} `json:"args"`
}
