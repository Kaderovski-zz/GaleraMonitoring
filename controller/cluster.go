package controller

import "fmt"

// CheckClusterStatus verify if Nodes are Primary
func CheckClusterStatus(mapStatus map[string]string) error {

	normalStatus := "Primary"

	for serverName, status := range mapStatus {
		if status == normalStatus {
			continue
		}

		return fmt.Errorf("Nodes status not primary on %s", serverName)

	}
	return nil
}

// CheckON verify if all nodes answer OK to wsrep_ready
func CheckON(mapNodes map[string]string) (err error) {

	const ready = "ON"

	for srvName, values := range mapNodes {
		if values == ready {
			continue
		}
		err = fmt.Errorf("%v is not ready %v", srvName, err)
	}
	return nil
}

// CheckUID verify if all Nodes uuid are the same
func CheckUID(uids map[string]string) error {

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