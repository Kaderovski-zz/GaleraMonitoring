package main

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/F00b4rch/SandBox/golang/GaleraMonitoring/galera"
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
	for srvName, db := range dbList {
		version, err := galera.GetVersion(db)
		if err != nil {
			log.Fatalf("Impossible to get version %v", err)
		}
		log.Printf("Serveur %s - version %s", srvName, version)
	}

	// Get Cluster State UUID
	muid := map[string]string{}

	for srvName, db := range dbList {
		_, uid, err := galera.GetClusterStateUUID(db)
		if err != nil {
			log.Fatalf("Impossible to get uid %v", err)
		}
		muid[srvName] = uid
		log.Printf("%s %s", srvName, uid)
	}

	// Check UUID
	err := checkUID(muid)
	if err != nil {
		log.Fatalf("%s : %v", err, muid)
	}

	// Get Total Nodes in map cnx
	nbSrv, err := numberNodes(cnx)
	if err != nil {
		log.Fatalf("Impossible to count total nodes %s", err)
	}
	log.Printf("Total Nodes : %v", nbSrv)

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
	err = checkNodesCount(mTotalNodes, nbSrv)
	if err != nil {
		fmt.Printf("Nodes count mismatched %s", err)
	}

	mStatusNodes := map[string]string{}
	// Get Cluster Status
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
	err = checkClusterStatus(mStatusNodes)
	if err != nil {
		fmt.Printf("Nodes are not Primary %v", err)
	}
}

func checkUID(uids map[string]string) error {

	lastUID := ""

	for srv, uid := range uids {
		if lastUID == "" {
			lastUID = uid
			continue
		}
		if lastUID == uid {
			continue
		}
		return fmt.Errorf("uid : %s of %s does not match", uid, srv)
	}

	return nil

}

func numberNodes(nodes map[string]string) (totalsrv int, err error) {

	totalsrv = len(nodes)
	return
}

func checkNodesCount(mapNodes map[string]int, totalNodes int) error {

	for _, numb := range mapNodes {
		if numb == totalNodes {
			continue
		}
		return fmt.Errorf("Number of connected Nodes is not the same, total = %v found = %v", totalNodes, numb)
	}

	return nil

}

func checkClusterStatus(mapStatus map[string]string) error {

	normalStatus := "Primary"

	for serverName, status := range mapStatus {
		if status == normalStatus {
			continue
		}

		return fmt.Errorf("Nodes status not primary on %s", serverName)

	}
	return nil
}
