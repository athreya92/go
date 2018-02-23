package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/cassandra"
	_ "github.com/mattes/migrate/source/file"
	"log"
	"strconv"
	"sync"
)

type DbManager struct {
	Session *gocql.Session
}


const (
	createKeyspace = "CREATE KEYSPACE IF NOT EXISTS sample_urs WITH replication = {'class':'SimpleStrategy', 'replication_factor': 1}"
	keyspaceName   = "sample_urs"
	port           = 9042
	migrationFiles = "file://C:/Users/athreya/go/src/go/gocql-example/migration/cql"
	cassandra_host = "CASSANDRA_HOST"
)

var (
	instance *DbManager
	once sync.Once
)

func createInstance() *DbManager {
	/*	Instance = new(DbManager)
		osHostName, _ := os.Hostname()
		host := utils.GetEnv(cassandra_host, osHostName)*/

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
	}
	instance.Session = session
	log.Println("Cassandra session Initialized")

	return instance
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
	once.Do(func() {
		instance = createInstance()
	})
	return instance

}

func createCassandraKeyspace(host string) {
	//Create cassandra session
	cluster := gocql.NewCluster(host)
	cluster.Port = port
	dbSession, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Not able to connect to cassandra %v", err)
	}
	//Create keyspace if not exists
	err = dbSession.Query(createKeyspace).Exec()
	defer dbSession.Close()
	if err != nil {
		//Close the cassandra session
		dbSession.Close()
		log.Fatalf("Failed to create keyspace %v", err)
	}
}

func (db *DbManager) Close() {
	db.Session.Close()
	log.Println("Closing the cassandra session")
}
