package gun

import "fmt"

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}

	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed.")
}
