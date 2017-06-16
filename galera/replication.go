package galera

import (
	"database/sql"
)

// GetQueueAvg shows the average size of the local received queue since the last status query
func GetQueueAvg(db *sql.DB) (varName string, value []uint8, err error) {

	q := "SHOW STATUS LIKE 'wsrep_local_recv_queue_avg'"
	err = db.QueryRow(q).Scan(&varName, &value)

	return

}
