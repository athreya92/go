package main

import (
	"github.com/gocql/gocql"
	"log"
	"gocql-example/entity"
	"fmt"
	"gocql-example/cassandra"
	//uuid "github.com/google/uuid"
)

var (
	dbManager = cassandra.GetInstance()
	name string
	rtype string
	serialNo string
	ipAddress string
	robotList []entity.Robot
	robot entity.Robot
	robotName string
)


func PerformOperations() {
	// Provide the cassandra cluster instance here.
	cluster := gocql.NewCluster("localhost")


	// gocql requires the keyspace to be provided before the session is created.
	// In future there might be provisions to do this later.
	cluster.Keyspace = "urs"

	cluster.ProtoVersion = 4
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("Could not connect to cassandra cluster: %v", err)
	}

	// Check if the table already exists. Create if table does not exist
	keySpaceMeta, _ := session.KeyspaceMetadata("urs")

	if _, exists := keySpaceMeta.Tables["person"]; exists != true {
		// Create a table
		session.Query("CREATE TABLE person (" +
			"id text, name text, phone text, " +
			"PRIMARY KEY (id))").Exec()
	}

	// DIY: Update table with something if it already exist.

	// Insert record into table using prepared statements
	session.Query("INSERT INTO person (id, name, phone) VALUES (?, ?, ?)",
		"shalabh", "Shalabh Aggarwal", "1234567890").Exec()

	// DIY: Update existing record

	// Select record and run some process on data fetched
	var name string
	var phone string
	if err := session.Query(
		"SELECT name, phone FROM person WHERE id='shalabh'").Scan(
		&name, &phone); err != nil {
		if err != gocql.ErrNotFound {
			log.Fatalf("Query failed: %v", err)
		}
	}
	log.Printf("Name: %v", name)
	log.Printf("Phone: %v", phone)

	// Fetch multiple rows and run process over them
	iter := session.Query("SELECT name, phone FROM person").Iter()
	for iter.Scan(&name, &phone) {
		log.Printf("Iter Name: %v", name)
		log.Printf("Iter Phone: %v", phone)
	}

	// DIY: Delete record
}

func main() {
	defer dbManager.Close()
/*	InsertRobot(dbManager, entity.Robot{uuid.New().String(),"xcd","xyz","1111"})
	InsertRobot(dbManager, entity.Robot{uuid.New().String(),"xxd","xyz","1111"})*/
	robotList = GetAllRobots(dbManager)
	FindRobot(dbManager)

	//PerformOperations()
}

func InsertRobot(dbm *cassandra.DbManager, robot entity.Robot) {
	query := dbm.Session.Query("INSERT INTO robots (serialNo, robotName, robotType, ipAddress) VALUES (?,?,?,?)", robot.SerialNo(), robot.Name(), robot.Type(), robot.IpAddress())
	if err := query.Exec(); err!=nil {
		log.Println("Failed to insert robot data to database",err)
	}else{
		log.Println("Robot insertion successful")
	}
}

func GetAllRobots(dbm *cassandra.DbManager) []entity.Robot {
	query := dbm.Session.Query("SELECT * FROM robots")
	iter := query.Iter()
	for iter.Scan(&serialNo, &name, &rtype, &ipAddress) {
		robot = entity.Robot{}
		robot.SetSerialNo(serialNo)
		robot.SetName(name)
		robot.SetType(rtype)
		robot.SetIpAddress(ipAddress)
		robotList = append(robotList, robot)
	}
	fmt.Println(robotList)
	return robotList
}

func FindRobot(dbm *cassandra.DbManager) {
	query := dbm.Session.Query("SELECT serialNo, robotName FROM robots WHERE serialNo = '1233' AND robotName ='abcd' ")
	if err := query.Scan(&serialNo); err!=nil {
		fmt.Println("Hi", err)
	}else{
		if err := query.Scan(&robotName); err!=nil {
			fmt.Println("Hello", err)
		}else {
			fmt.Println("get lost")
		}
		log.Println("success")
	}
}