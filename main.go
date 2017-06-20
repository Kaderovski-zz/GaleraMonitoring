package main

import (
	"database/sql"
	"log"
	"encoding/binary"
	//"fmt"
	"github.com/F00b4rch/GaleraMonitoring/controller"
	"github.com/F00b4rch/GaleraMonitoring/galera"
	"github.com/F00b4rch/GaleraMonitoring/slackApp"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"time"
)

func main() {

	// Starting loop
	for {

		// Define here your nodes connexions settings
		cnx := map[string]string{
			"n1": "root:@(172.17.0.2:3307)/",
			"n2": "root:@(172.17.0.3:3306)/",
			"n3": "root:@(172.17.0.4:3306)/",
		}

		dbList := map[string]*sql.DB{}

		// Initialize mysql connexions
		for key, con := range cnx {
			db, err := sql.Open("mysql", con)
			if err != nil {
				slackApp.PayloadSlack("\n### [ERROR] ###\n Can't connect to galera Nodes" + err.Error())
				log.Printf("Can't connect to galera Nodes %v", err)
			}
			defer db.Close()
			dbList[key] = db
		}

		// Get MariaDB version
		//fmt.Print("### Version ####\n")
		for _, db := range dbList {
			_, err := galera.GetVersion(db)
			if err != nil {
				slackApp.PayloadSlack("\n### [WARNING] ###\nImpossible to get version" + err.Error())
				log.Printf("Impossible to get version", err)
			}
			//log.Printf("Serveur %s - version %s", srvName, version)
		}

		// Get Cluster State UUID
		//fmt.Print("### UUID ####\n")
		muid := map[string]string{}
		for srvName, db := range dbList {
			_, uid, err := galera.GetClusterStateUUID(db)
			if err != nil {
				slackApp.PayloadSlack("\n### [FATAL] UUID ###\nImpossible to get uid " + err.Error())
				log.Printf("Impossible to get uuid %v", err)
			}
			muid[srvName] = uid
			//log.Printf("%s %s", srvName, uid)
		}

		// Check UUID
		err := controller.CheckUID(muid)
		if err != nil {
			log.Printf("%s : %v", err, muid)
		}

		// Get Total Nodes in map cnx
		//fmt.Print("### NODES ####\n")
		nbSrv, err := numberNodes(cnx)
		if err != nil {
			slackApp.PayloadSlack("\n### [FATAL] Nodes ###\nImpossible to count total nodes " + err.Error())
			log.Printf("Impossible to count total nodes %v", err)
		}
		//log.Printf("Total Nodes : %v", nbSrv)

		mTotalNodes := map[string]int{}

		// If total Nodes is not equal nbSrv
		for srvName, db := range dbList {
			_, numb, err := galera.GetNumbNodes(db)
			if err != nil {
				slackApp.PayloadSlack("\n### [FATAL] Nodes ###\nImpossible to get total nodes" + err.Error() + " total Nodes " + string(nbSrv) + " Nodes get " + string(numb))
				log.Printf("Impossible to get total nodes %v", err)
			} /*else {
			log.Printf("Number of Nodes counts : %v", numb)
		}*/
			mTotalNodes[srvName] = numb
		}

		// Diff between count nodes connexion and get nodes SQL
		err = controller.CheckNodesCount(mTotalNodes, nbSrv)
		if err != nil {
			slackApp.PayloadSlack("\n### [FATAL] Nodes ###\nNodes count mismatched" + err.Error())
			log.Printf("Nodes count mismatched ! %v", err)
		}

		// Get Cluster Status
		//fmt.Print("### STATUS ####\n")
		mStatusNodes := map[string]string{}
		for srvName, db := range dbList {
			_, status, err := galera.GetClusterStatus(db)
			if err != nil {
				slackApp.PayloadSlack("\n### [FATAL] STATUS ###\nImpossible to get cluster status " + err.Error())
				log.Printf("Impossible to get cluster status %v", err)
			} /*else {
			log.Printf("%v status : %v", srvName, status)
		}*/
			mStatusNodes[srvName] = status
		}

		// Check if status is != Primary
		err = controller.CheckClusterStatus(mStatusNodes)
		if err != nil {
			slackApp.PayloadSlack("\n" + err.Error())
			log.Printf("Node not primary %v", err)
		}

		mNodesReady := map[string]string{}
		// Get Nodes wsrep_ready
		for srvName, db := range dbList {
			_, values, err := galera.GetReady(db)
			if err != nil {
				slackApp.PayloadSlack("\n### [FATAL] Nodes ###\nImpossible to get Nodes wsrep_ready " + err.Error())
				log.Printf("Impossible to get wsrep_ready %v", err)
			} /*else {
			log.Printf("%v is ready : [%v]", srvName, values)
		}*/
			mNodesReady[srvName] = values
		}

		// Check if wsrep_ready is ON
		err = controller.CheckON(mNodesReady)
		if err != nil {
			slackApp.PayloadSlack("\n" + err.Error())
			log.Printf("Nodes not ready %v", err)
		}

		// Get Nodes wsrep_connected
		mNodesCon := map[string]string{}
		for srvName, db := range dbList {
			_, values, err := galera.GetConnected(db)
			if err != nil {
				slackApp.PayloadSlack("\n### [FATAL] Nodes ###\nImpossible to get Nodes wsrep_connected " + err.Error())
				log.Printf("Impossible to get wsrep_connected %v", err)
			} /*else {
			log.Printf("%v is connected : [%v]", srvName, values)
		}*/
			mNodesCon[srvName] = values
		}

		// Check if wsrep_connected is ON
		err = controller.CheckConnected(mNodesCon)
		if err != nil {
			slackApp.PayloadSlack("\n" + err.Error())
		}

		// Get wsrep_local_recv_queue_avg
		//fmt.Print("### AVERAGE REPLICATION ####\n")
		for srvName, db := range dbList {
			_, values, err := galera.GetQueueAvg(db)
			if err != nil {
				slackApp.PayloadSlack("\nImpossible to get average replication " + err.Error())
				log.Printf("Impossible to get average rep %v", err)
				continue
			}
			if Float64frombytes(values) >= 0.100000 {
				slackApp.PayloadSlack("\n### [WARNING] Average Replication ###\nAverage on " + srvName + ":" + string(values))
				//log.Printf("Replication warning on %s %v", srvName, values)
			}
		}

		// Looping every minutes
		time.Sleep(60 * time.Second)
	}
}
func numberNodes(nodes map[string]string) (totalsrv int, err error) {

	totalsrv = len(nodes)
	return
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
