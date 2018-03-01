package entity

type BasicRobotInfo struct {
	Robot_id string
	Name     string
	Type     string
	Host_ip  string
}

type RobotInfo struct {
	Robot_id string
	Name     string
	Type     string
	Host_ip  string
	Vendor   string
	Model    string
	Payloads []Payload
	Actions []string
}

type PayloadConfig struct {
	ConfigName string
	ConfigValue string
}

type Payload struct {
	Payload_id  string
	PayloadType string
	Config      PayloadConfig
}