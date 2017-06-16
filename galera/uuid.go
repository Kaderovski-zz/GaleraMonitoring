package galera

import "database/sql"

// GetClusterStateUUID is makirg a query to get uuid on each nodes
func GetClusterStateUUID(db *sql.DB) (srv, uid string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_state_uuid'"
	err = db.QueryRow(q).Scan(&srv, &uid)

	return

}
