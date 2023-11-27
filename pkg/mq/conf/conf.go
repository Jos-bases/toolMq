package conf

import "fmt"

type MqConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
}

func (t *MqConf) NewDsn() string {
	return fmt.Sprintf("amqp://%v:%v@%v:%v/", t.Username, t.Password, t.Ip, t.Port)
}
