package galera

import "database/sql"

// GetClusterStatus is getting cluster status (def : PRIMARY)
func GetClusterStatus(db *sql.DB) (varName, value string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_status'"
	err = db.QueryRow(q).Scan(&varName, &value)

	return

}
