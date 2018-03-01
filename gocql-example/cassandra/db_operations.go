package cassandra

import (
	"go/gocql-example/entity"
	"log"
	"github.com/gocql/gocql"
)

var (
	dbManager                 = GetInstance()
	query       *gocql.Query
	robotId     string
	payloadId   string
	name        string
	rtype       string
	host        string
	model       string
	vendor      string
	ptype       string
	confName    string
	confValue   string
	payloadList []entity.Payload
	actions     []string
)

const (
	//Query statements for robot
	RobotInsertQuery          = "INSERT INTO robot_basicInfo (robot_id, name, type, host_ip, vendor, model) VALUES (?,?,?,?,?,?)"
	RobotPayloadInsertQuery   = "INSERT INTO robot_payloadInfo (robot_id, payload_id) VALUES (?,?)"
	PayloadInsertQuery        = "INSERT INTO payload_info (type, properties, payload_id) VALUES (?,(?,?),?)"
	PayloadByPropQuery        = "SELECT payload_id from payload_info WHERE type = ? AND properties = (?,?)"
	RobotSearchByPayloadQuery = "SELECT robot_id from robot_payloadInfo WHERE payload_id = ?"
	FindRobotById 			  = "SELECT * from robot_basicInfo WHERE robot_id = ? ALLOW FILTERING"
	AllPayloadsQuery 		  = "SELECT type, properties from payload_info"
	RobotActionsInsertQuery   = "INSERT INTO robot_actionInfo (robot_id, actions) VALUES (?,?)"
	GetActionsByRobotQuery    = "SELECT actions FROM robot_actionInfo WHERE robot_id = ?"
)

func InsertRobotInfo(robot entity.RobotInfo) error {
	query = dbManager.Session.Query(RobotInsertQuery, robot.Robot_id, robot.Name, robot.Type, robot.Host_ip, robot.Vendor, robot.Model)
	if err := query.Exec(); err != nil {
		log.Println("Failed to insert robot data to database", err)
		return err
	}
	 if err := insertPayloadInfo(robot); err != nil {
	 	return err
	 }

	 if err := insertRobotPayloadInfo(robot); err != nil {
	 	return err
	 }
	 if err := insertRobotActionInfo(robot); err != nil {
		 return err
	 }
	log.Println("Robot insertion successful")
	return nil
}

func insertRobotActionInfo(robot entity.RobotInfo) error {
	query = dbManager.Session.Query(RobotActionsInsertQuery, robot.Robot_id, robot.Actions)
		if err := query.Exec(); err != nil {
			log.Println("Failed to insert robot_action data to database", err)
			return err
		}
	log.Println("robot_action insertion successful")
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
		query = dbManager.Session.Query(PayloadInsertQuery, ptype, props.ConfigName, props.ConfigValue, pid)
		if err := query.Exec(); err != nil {
			log.Println("Failed to insert payload data to database", err)
			return err
		}
		continue
	}
	log.Println("Payload insertion successful")
	return nil
}

func GetAllPayloads() []entity.Payload {
	query = dbManager.Session.Query(AllPayloadsQuery)
	iterator := query.Iter()
	for iterator.Scan(&ptype,&confName,&confValue) {
		pload := entity.Payload{
			PayloadType:ptype,
			Config: struct {
				ConfigName  string
				ConfigValue string
			}{ConfigName: confName, ConfigValue: confValue},
		}
		payloadList = append(payloadList,pload)
	}
	return payloadList
}

func GetActionsByRobot(robot entity.BasicRobotInfo) ([]string,error) {
	query = dbManager.Session.Query(GetActionsByRobotQuery, robot.Robot_id)
	if err := query.Scan(&actions); err != nil {
		log.Println("Failed to fetch actions for the robot")
		return nil,err
	}
	return actions,nil
}

func GetAllActions() []string {
	return nil
}

func GetRobots(payloads []entity.Payload) []string {
	var robots, tempRobots []string
	if len(payloads) > 1 {
		for _, payload := range payloads {
			temp := findAllRobotsWithRequestedPayloads(payload)
			tempRobots = append(tempRobots,temp...)
		}
		robots = getDuplicates(tempRobots)
	} else {
		robots = findAllRobotsWithRequestedPayloads(payloads[0])
	}
	return robots
}

func findAllRobotsWithRequestedPayloads(payload entity.Payload) []string {
	var payloadIdList []string
	payloadIds := searchPayloadWithConfig(payload)
	payloadIdList = append(payloadIdList, payloadIds...)
	robots := findAllRobotsWithPayloadIds(payloadIdList)
	return robots
}

func getDuplicates(ids []string) []string {
	result := []string{}
	for i := 0; i < len(ids); i++ {
		// Scan slice for a previous element of the same value.
		exists := false
		for v := 0; v < i; v++ {
			if ids[v] == ids[i] {
				exists = true
				break
			}
		}
		// If previous element exists, append.
		if exists {
			result = append(result, ids[i])
		}
	}
	return result
}

func searchPayloadWithConfig(payload entity.Payload) []string {
	var payloadList []string
	query = dbManager.Session.Query(PayloadByPropQuery, payload.PayloadType,payload.Config.ConfigName, payload.Config.ConfigValue)
	iterator := query.Iter()
	for iterator.Scan(&payloadId) {
		payloadList = append(payloadList, payloadId)
	}
	return payloadList
}

func findAllRobotsWithPayloadIds(payloadIds []string) []string {
	var robots []string
	for _, payloadId := range payloadIds {
		query = dbManager.Session.Query(RobotSearchByPayloadQuery, payloadId)
		if err := query.Scan(&robotId); err != nil {
			log.Println("Failed to find robot by payload data in database", err)
		}
		robots = append(robots, robotId)
	}
	return robots
}

func FindRobotByRobotId(robotId string) entity.BasicRobotInfo {
	query = dbManager.Session.Query(FindRobotById, robotId)
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



