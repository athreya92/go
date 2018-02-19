package entity


type ArmPayload struct {
	Device_id string
	Payload_type string
	Actions []string
	Config Arm_config
}

type Arm_config struct {
	Capacity string
}
