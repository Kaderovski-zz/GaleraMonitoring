package galera

import "database/sql"

// GetVersion is getting version mariadb on each nodes
func GetVersion(db *sql.DB) (version string, err error) {

	q := "select version()"
	err = db.QueryRow(q).Scan(&version)

	return
}
