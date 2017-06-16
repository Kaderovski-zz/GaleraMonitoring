package controller

import "fmt"

// CheckON verify if all nodes answer ON to wsrep_ready
func CheckON(mapNodes map[string]string) (err error) {

	const ready = "ON"

	for srvName, values := range mapNodes {
		if values == ready {
			continue
		}
		err = fmt.Errorf("%v is not ready %v", srvName, values)
	}
	return err
}

// CheckConnected verify if all nodes answer ON to wsrep_connected
func CheckConnected(mapNodes map[string]string) (err error) {

	const ready = "ON"

	for srvName, values := range mapNodes {
		if values == ready {
			continue
		}
		err = fmt.Errorf("%v is not connected %v", srvName, values)
	}
	return err
}

// CheckNodesCount verify if total Nodes in program are equal to number Nodes inside each Nodes
func CheckNodesCount(mapNodes map[string]int, totalNodes int) error {

	for _, numb := range mapNodes {
		if numb == totalNodes {
			continue
		}
		return fmt.Errorf("Number of connected Nodes is not the same, total = %v found = %v", totalNodes, numb)
	}

	return nil

}
