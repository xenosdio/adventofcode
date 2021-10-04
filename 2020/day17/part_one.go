package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cube struct {
	IsActive       bool
	ActiveNeigbors int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pocketDimensions := make(map[string]Cube, 0)

	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()
		for x, state := range input {

			if string(state) == "#" {
				key := fmt.Sprintf("%d,%d,0", x, y)

				pocketDimensions[key] = Cube{
					IsActive:       true,
					ActiveNeigbors: 0,
				}
			}
		}
		y++
	}

	for cycle := 0; cycle < 6; cycle++ {
		countActiveNeigbors(pocketDimensions)
		changeState(pocketDimensions)
	}

	log.Println(len(pocketDimensions))
}

func countActiveNeigbors(pocketDimensions map[string]Cube) {

	for key, cube := range pocketDimensions {

		if !cube.IsActive {
			continue
		}

		var x, y, z int
		fmt.Sscanf(key, "%d,%d,%d", &x, &y, &z)

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {

					neighborKey := fmt.Sprintf("%d,%d,%d", x+i, y+j, z+k)

					if neighborKey != key {
						neighborCube := pocketDimensions[neighborKey]
						neighborCube.ActiveNeigbors++
						pocketDimensions[neighborKey] = neighborCube
					}
				}
			}
		}
	}
}

func changeState(pocketDimensions map[string]Cube) {

	for key, cube := range pocketDimensions {

		if (cube.IsActive && (cube.ActiveNeigbors == 2 || cube.ActiveNeigbors == 3)) || (!cube.IsActive && cube.ActiveNeigbors == 3) {
			pocketDimensions[key] = Cube{
				IsActive:       true,
				ActiveNeigbors: 0,
			}
		} else {
			delete(pocketDimensions, key)
		}
	}
}
