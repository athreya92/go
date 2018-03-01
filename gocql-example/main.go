package main


import (
	"go/gocql-example/entity"
	"go/gocql-example/cassandra"
	"fmt"
)

var payloadsC []entity.Payload
var actionList []string
var err error

func main() {
	armConfigC := entity.PayloadConfig{
		ConfigName:"capacity",
		ConfigValue:"heavy",
	}
	armPayloadC := entity.Payload{
		PayloadType: "Arm",
		Config:      armConfigC,
	}
	cameraConfigC := entity.PayloadConfig{
		ConfigName:"resolution",
		ConfigValue:"720",
	}
	cameraPayloadC := entity.Payload{
		PayloadType: "Camera",
		Config:      cameraConfigC,
	}

	payloadsC = append(payloadsC, cameraPayloadC, armPayloadC)
/*	robot := entity.BasicRobotInfo{
		Robot_id:"1",
	}

	if actionList,err = cassandra.GetActionsByRobot(robot); err != nil {
		log.Println(err)
	}
	fmt.Println(actionList)*/

	robots := cassandra.GetRobots(payloadsC)
	fmt.Println(robots)
}


