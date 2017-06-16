package galera

import "database/sql"

// GetNumbNodes function is getting number of nodes
func GetNumbNodes(db *sql.DB) (varName string, number int, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_cluster_size'"
	err = db.QueryRow(q).Scan(&varName, &number)

	return
}
