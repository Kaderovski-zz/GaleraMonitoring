package controller

import "fmt"

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
