package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/cassandra"
	_ "github.com/mattes/migrate/source/file"
	"log"
	"strconv"
)

type DbManager struct {
	Session *gocql.Session
}


const (
	createKeyspace = "CREATE KEYSPACE IF NOT EXISTS urs WITH replication = {'class':'SimpleStrategy', 'replication_factor': 1}"
	keyspaceName   = "urs"
	port           = 9042
	migrationFiles = "file://C:/Users/athreya/go/src/gocql-example/migration/cql"
	cassandra_host = "CASSANDRA_HOST"
)

var (
	instance *DbManager
)

//This function is called automatically when ever the package is imported
func init() {
	instance = new(DbManager)
	host := "localhost"

	instance.handleMigrations(host)

	//create cassandra session
	cluster := gocql.NewCluster(host)
	cluster.Port = port
	cluster.Keyspace = keyspaceName

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Unable to connect to cassandra session: %v", err)
	} else {
		instance.Session = session
		log.Println("Cassandra session Initialized")
	}
}

func (db *DbManager) handleMigrations(host string) {

	//create cassandra keyspace
	createCassandraKeyspace(host)

	cassandra := cassandra.Cassandra{}

	url := "cassandra://" + host + ":" + strconv.Itoa(port) + "/" + keyspaceName
	driver, err := cassandra.Open(url)
	if err != nil {
		log.Fatalf("Unable to connect to cassandra:%v", err)
	}
	mig, err := migrate.NewWithDatabaseInstance(
		migrationFiles,
		"cassandra", driver)
	if err != nil {
		log.Fatalf("Failed to validate migration:%v", err)
	}
	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migration %v", err)
	}

	//Close the connection and handle error
	if srcErr, dbErr := mig.Close(); srcErr != nil || dbErr != nil {
		log.Fatalf("Error while executing the migrations %v %v", srcErr, dbErr)
	}
}

func GetInstance() *DbManager {
	return instance
}

func createCassandraKeyspace(host string) {
	//Create cassandra session
	cluster := gocql.NewCluster(host)
	cluster.Port = port
	if dbSession, err := cluster.CreateSession(); err != nil {
		log.Fatalf("Not able to connect to cassandra %v", err)
	} else {
		//Create keyspace if not exists
		if err := dbSession.Query(createKeyspace).Exec(); err != nil {
			log.Fatalf("Failed to create keyspace %v", err)
		}

		//Close the cassandra session
		dbSession.Close()
	}
}

func (db *DbManager) Close() {
	db.Session.Close()
	log.Println("Closing the cassandra session")
}
