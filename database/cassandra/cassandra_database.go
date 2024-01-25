package cassandra

import (
	"fmt"
	"os"

	"github.com/gocql/gocql"
)

func Init() *gocql.Session {

	cluster := gocql.NewCluster(os.Getenv("DB_CASSANDRA_CLUSTER") + ":" + os.Getenv("DB_CASSANDRA_PORT")) // replace with your Cassandra host IP
	cluster.Keyspace = os.Getenv("DB_CASSANDRA_KEYSPACE")                                                 // replace with your keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Error while connecting to Cassandra:", err)
	}

	return session

}
