package conf

type QueueT string

const (
	Classic QueueT = "classic"
	Quorum  QueueT = "quorum"
)

type Queue struct {
	Name       string                 `json:"name"`
	Types      QueueT                 `json:"types"`
	Durable    bool                   `json:"durable"`
	AutoDelete bool                   `json:"auto_delete"`
	Internal   bool                   `json:"internal"`
	NoWait     bool                   `json:"no_wait"`
	IfUnused   bool                   `json:"if_unused"`
	IfEmpty    bool                   `json:"if_empty"`
	Args       map[string]interface{} `json:"args"`
}
