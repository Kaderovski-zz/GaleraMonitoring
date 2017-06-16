package galera

import "database/sql"

// GetReady gets wsrep_ready on each Nodes
func GetReady(db *sql.DB) (varName, value string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_ready'"
	err = db.QueryRow(q).Scan(&varName, &value)

	return
}

// GetConnected gets wsrep_connected on each Nodes
func GetConnected(db *sql.DB) (varName, value string, err error) {

	q := "SHOW GLOBAL STATUS LIKE 'wsrep_connected'"
	err = db.QueryRow(q).Scan(&varName, &value)

	return
}

// GetVersion is getting version mariadb on each nodes
func GetVersion(db *sql.DB) (version string, err error) {

	q := "select version()"
	err = db.QueryRow(q).Scan(&version)

	return
}
