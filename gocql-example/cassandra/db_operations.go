package cassandra

import (
	"go/gocql-example/entity"
	"log"
	"github.com/gocql/gocql"
	ops "github.com/adam-hanna/arrayOperations"
)

var (
	dbManager          = GetInstance()
	query       *gocql.Query
	robotId     string
	payloadId   string
	name        string
	rtype       string
	host        string
	model string
	vendor string
)

const (
	//Query statements for robot
	RobotInsertQuery          = "INSERT INTO robot_basicInfo (robot_id, name, type, host_ip, vendor, model) VALUES (?,?,?,?,?,?)"
	RobotPayloadInsertQuery   = "INSERT INTO robot_payloadInfo (robot_id, payload_id) VALUES (?,?)"
	PayloadInsertQuery        = "INSERT INTO payload_info (type, properties, payload_id) VALUES (?,(?,?),?)"
	PayloadByPropQuery        = "SELECT payload_id from payload_info WHERE type = ? AND properties = (?,?)"
	RobotSearchByPayloadQuery = "SELECT robot_id from robot_payloadInfo WHERE payload_id = ?"
	FindRobotById 			  = "SELECT * from robot_basicInfo WHERE robot_id = ? ALLOW FILTERING"
)

func InsertRobotInfo(robot entity.RobotInfo) error {
	query = dbManager.Session.Query(RobotInsertQuery, robot.Robot_id, robot.Name, robot.Type, robot.Host_ip, robot.Vendor, robot.Model)
	if err := query.Exec(); err != nil {
		log.Println("Failed to insert robot data to database", err)
		return err
	}
	 if errr := insertPayloadInfo(robot); errr != nil {
	 	return errr
	 }

	 insertRobotPayloadInfo(robot)
	log.Println("Robot insertion successful")
	return nil

}

func insertRobotPayloadInfo(robot entity.RobotInfo) error {
	payloads := robot.Payloads
	for _, payload := range payloads {
		pid := payload.Payload_id
		query = dbManager.Session.Query(RobotPayloadInsertQuery, robot.Robot_id, pid)
		if err := query.Exec(); err != nil {
			log.Println("Failed to insert robot_payload data to database", err)
			return err
		}
		continue
	}
	log.Println("robot_Payload insertion successful")
	return nil

}

func insertPayloadInfo(robot entity.RobotInfo) error {
	payloads := robot.Payloads
	for _, payload := range payloads {
		ptype := payload.PayloadType
		props := payload.Config
		pid := payload.Payload_id
		if props.Capacity != "" {
			query = dbManager.Session.Query(PayloadInsertQuery, ptype, "capacity", props.Capacity, pid)
		} else {
			query = dbManager.Session.Query(PayloadInsertQuery, ptype, "resolution", props.Resolution, pid)
		}
		if err := query.Exec(); err != nil {
			log.Println("Failed to insert payload data to database", err)
			return err
		}
		continue
	}
	log.Println("Payload insertion successful")
	return nil
}

func searchPayloadByProp(payload entity.Payload) []string {
	var payloadList []string
	ptype := payload.PayloadType
	if ptype == "Arm" {
		query = dbManager.Session.Query(PayloadByPropQuery, payload.PayloadType,"capacity", payload.Config.Capacity)
	} else {
		query = dbManager.Session.Query(PayloadByPropQuery, payload.PayloadType,"resolution", payload.Config.Resolution)
	}
	iterator := query.Iter()
	for iterator.Scan(&payloadId) {
		payloadList = append(payloadList, payloadId)
	}
	return payloadList
}

func GetRobots(payloads []entity.Payload) []entity.BasicRobotInfo {
	var robotsWithArm []string
	var robotsWithCamera []string
	var robots []entity.BasicRobotInfo
	for _,payload := range payloads {
		if payload.PayloadType == "Arm" {
			armIds := searchPayloadByProp(payload)
			robotsWithArm = findRobotsByPayload(armIds)
		} else {
			cameraIds := searchPayloadByProp(payload)
			robotsWithCamera = findRobotsByPayload(cameraIds)
		}
	}
	if robotsWithArm == nil || robotsWithCamera == nil || (robotsWithArm == nil && robotsWithCamera == nil) {
		log.Println("No robot with these payload properties")
		return nil
	}
	temp, ok := ops.Intersect(robotsWithArm, robotsWithCamera)
	if !ok {
		log.Println("No robot with these payload properties")
		return nil
	}

	robotIds := temp.Interface().([]string)
	for _,robotId := range robotIds {
		robot := findRobotById(robotId)
		robots = append(robots,robot)
	}
	return robots
}

func findRobotById(id string) entity.BasicRobotInfo {
	query = dbManager.Session.Query(FindRobotById, id)
	if err := query.Scan(&rtype, &robotId, &host, &model, &name, &vendor); err != nil {
		log.Println("Failed to find robot", err)
	}
	robot := entity.BasicRobotInfo{
		Robot_id: robotId,
		Name:     name,
		Type:     rtype,
		Host_ip:  host,
	}
	return robot
}

func findRobotsByPayload(armIds []string) []string {
	var robots []string
	for _, id := range armIds {
		query = dbManager.Session.Query(RobotSearchByPayloadQuery, id)
		if err := query.Scan(&robotId); err != nil {
			log.Println("Failed to find robot by payload data to database", err)
		}
		robots = append(robots, robotId)
	}
	return robots
}


