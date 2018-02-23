package main

import (
	"go/gocql-example/entity"
	"go/gocql-example/cassandra"
	/*"fmt"*/
	"fmt"
)

var payloadsC []entity.Payload

func main() {
	armConfigC := entity.PayloadConfig{
		Capacity:"light",
	}
	cameraConfigB := entity.PayloadConfig{
		Resolution:"720",
	}
	armPayloadC := entity.Payload{
		Payload_id:  "6",
		PayloadType: "Arm",
		Actions:     []string{"lift","drop"},
		Config:      armConfigC,
	}
	cameraPayloadC := entity.Payload{
		Payload_id:  "9",
		PayloadType: "Camera",
		Actions:     []string{"start_video","stop_video","take_snapshot"},
		Config:      cameraConfigB,
	}

	payloadsC = append(payloadsC, armPayloadC, cameraPayloadC)

	/*robotC := entity.RobotInfo{
		Robot_id: "3",
		Name:     "Robot-C",
		Type:     "Innok",
		Host_ip:  "127.1.1.1",
		Vendor:   "Innok",
		Model:    "Heroes",
		Payloads: payloadsC,
	}*/

	//cassandra.InsertRobotInfo(robotC)
	robotList := cassandra.GetRobots(payloadsC)
	fmt.Println(robotList)

}
