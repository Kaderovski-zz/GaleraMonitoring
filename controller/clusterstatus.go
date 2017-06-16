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
