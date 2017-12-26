package entity

type Robot struct {
	Serialno  string
	RobotName string
	RobotType string
	Ipaddress string
}

func (r *Robot) SetSerialNo(serialNo string) {
	r.Serialno = serialNo
}

func (r Robot) SerialNo() string {
	return r.Serialno
}

func (r *Robot) SetName(robotName string) {
	r.RobotName = robotName
}

func (r Robot) Name() string {
	return r.RobotName
}

func (r *Robot) SetType(robotType string) {
	r.RobotType = robotType
}

func (r Robot) Type() string {
	return r.RobotType
}

func (r *Robot) SetIpAddress(ipAddress string) {
	r.Ipaddress = ipAddress
}

func (r Robot) IpAddress() string {
	return r.Ipaddress
}
