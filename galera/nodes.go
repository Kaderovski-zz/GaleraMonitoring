package galera

import "database/sql"

// GetReady gets wsrep_ready on each Nodes
func GetReady(db *sql.DB) (varName, value string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_ready'"
	err = db.QueryRow(q).Scan(&varName, &value)

	return
}

// GetNumbNodes function is getting number of nodes
func GetNumbNodes(db *sql.DB) (varName string, number int, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_size'"
	err = db.QueryRow(q).Scan(&varName, &number)

	return
}

// GetClusterStatus is getting cluster status (def : PRIMARY)
func GetClusterStatus(db *sql.DB) (varName, value string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_status'"
	err = db.QueryRow(q).Scan(&varName, &value)

	return

}

// GetClusterStateUUID is makirg a query to get uuid on each nodes
func GetClusterStateUUID(db *sql.DB) (srv, uid string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_state_uuid'"
	err = db.QueryRow(q).Scan(&srv, &uid)

	return

}

// GetVersion is getting version mariadb on each nodes
func GetVersion(db *sql.DB) (version string, err error) {

	q := "select version()"
	err = db.QueryRow(q).Scan(&version)

	return
}
