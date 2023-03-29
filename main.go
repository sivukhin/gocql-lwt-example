package main

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "ks"
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy(), gocql.ShuffleReplicas())
	for {
		session, err := gocqlx.WrapSession(gocql.NewSession(*cluster))
		if err != nil {
			log.Printf("err: %v", err)
		} else {
			query := qb.Select("tbl").Where(qb.Eq("id")).Query(session).BindMap(map[string]interface{}{"id": "test"})
			log.Printf("err: %v", query.Exec()) // LWT flag will be set after prepare request - so we need to Exec query before it
			log.Printf("isLwt: %v", query.IsLWT())
			query.Release()
		}
		session.Close()
		time.Sleep(1 * time.Second)
	}
}
