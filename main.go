package main

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/F00b4rch/Galera_Monitoring/controller"
	"github.com/F00b4rch/Galera_Monitoring/galera"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Define here your nodes connexions settings
	cnx := map[string]string{
		"n1": "root:@(172.17.0.2:3306)/",
		"n2": "root:@(172.17.0.3:3306)/",
		"n3": "root:@(172.17.0.4:3306)/",
	}

	dbList := map[string]*sql.DB{}

	// Initialize mysql connexions
	for key, con := range cnx {
		db, err := sql.Open("mysql", con)
		if err != nil {
			log.Fatalf("Error while connecting to server %s : %v", key, err)
		}
		defer db.Close()
		dbList[key] = db
	}

	// Get MariaDB version
	fmt.Print("### Version ####\n")
	for srvName, db := range dbList {
		version, err := galera.GetVersion(db)
		if err != nil {
			log.Fatalf("Impossible to get version %v", err)
		}
		log.Printf("Serveur %s - version %s", srvName, version)
	}

	// Get Cluster State UUID
	muid := map[string]string{}
	fmt.Print("### UUID ###\n")
	for srvName, db := range dbList {
		_, uid, err := galera.GetClusterStateUUID(db)
		if err != nil {
			log.Fatalf("Impossible to get uid %v", err)
		}
		muid[srvName] = uid
		log.Printf("%s %s", srvName, uid)
	}

	// Check UUID
	err := controller.CheckUID(muid)
	if err != nil {
		log.Fatalf("%s : %v", err, muid)
	}

	// Get Total Nodes in map cnx
	fmt.Print("### Nodes ###\n")
	nbSrv, err := numberNodes(cnx)
	if err != nil {
		log.Fatalf("Impossible to count total nodes %s", err)
	}
	log.Printf("Total Nodes : %v", nbSrv)

	mTotalNodes := map[string]int{}

	// If total Nodes is not equal nbSrv
	for srvName, db := range dbList {
		_, numb, err := galera.GetNumbNodes(db)
		if err != nil {
			log.Fatalf("Impossible to get total nodes %v total Nodes = %v Nodes get = %v", err, nbSrv, numb)
		} else {
			log.Printf("Number of Nodes counts : %v", numb)
		}
		mTotalNodes[srvName] = numb
	}

	// Diff between count nodes connexion and get nodes SQL
	err = controller.CheckNodesCount(mTotalNodes, nbSrv)
	if err != nil {
		fmt.Printf("Nodes count mismatched %s", err)
	}

	// Get Cluster Status
	fmt.Print("### STATUS ###\n")
	mStatusNodes := map[string]string{}
	for srvName, db := range dbList {
		_, status, err := galera.GetClusterStatus(db)
		if err != nil {
			log.Fatalf("Impossible to get cluster status %s", err)
		} else {
			log.Printf("%v status : %v", srvName, status)
		}
		mStatusNodes[srvName] = status
	}

	// Check if status is != Primary
	err = controller.CheckClusterStatus(mStatusNodes)
	if err != nil {
		fmt.Print(err)
	}

	mNodesReady := map[string]string{}
	// Get Nodes wsrep_ready
	for srvName, db := range dbList {
		_, values, err := galera.GetReady(db)
		if err != nil {
			log.Fatalf("Impossible to get Nodes wsrep_ready %s", err)
		} else {
			log.Printf("%v is ready : [%v]", srvName, values)
		}
		mNodesReady[srvName] = values
	}

	// Check if wsrep_ready is ON
	err = controller.CheckON(mNodesReady)
	if err != nil {
		log.Fatal(err)
	}

	// Get Nodes wsrep_connected
	mNodesCon := map[string]string{}
	for srvName, db := range dbList {
		_, values, err := galera.GetConnected(db)
		if err != nil {
			log.Fatalf("Impossible to get Nodes wsrep_connected %s", err)
		} else {
			log.Printf("%v is connected : [%v]", srvName, values)
		}
		mNodesCon[srvName] = values
	}

	// Check if wsrep_connected is ON
	err = controller.CheckConnected(mNodesCon)
	if err != nil {
		log.Fatal(err)
	}

	// Get wsrep_local_recv_queue_avg
	fmt.Print("### Average Replication ###\n")
	for srvName, db := range dbList {
		_, values, err := galera.GetQueueAvg(db)
		if err != nil {
			log.Fatalf("Impossible to get average replication %s", err)
		} else {
			log.Printf("Average on %v : %v", srvName, string(values))
		}
	}

}

func numberNodes(nodes map[string]string) (totalsrv int, err error) {

	totalsrv = len(nodes)
	return
}
