package main

import (
	"go/go-xml/entity"
	xml "go/go-xml/xml_ops"
)


func main() {

	acts := []string{"lift","drop"}
	arm := &entity.ArmPayload{
		Device_id: "3428edafd97r",
		Payload_type: "arm",
		Actions: acts,
		Config: entity.Arm_config{Capacity: "light"},
	}
	a := new(xml.ArmXmlOps)
	//a.AddElementsToXML(arm)
	a.ReadFromXML(arm)
}
