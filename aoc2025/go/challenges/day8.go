package challenges

import (
	"math"
	"strconv"
	"strings"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

type JunctionBox struct {
	X int
	Y int
	Z int
}

type JunctionBoxConnections struct {
	connections                []JunctionBoxGroup
	allJunctionBoxes           []JunctionBox
	junctionBoxConnectionIndex map[JunctionBox]int
	junctionBoxDistanceMatrix  map[JunctionBox]map[JunctionBox]int
	junctionBoxConnections     []JunctionBoxConnection
}

type JunctionBoxConnection struct {
	Distance int
	From     JunctionBox
	To       JunctionBox
}

type JunctionBoxGroup struct {
	connections []JunctionBox
}

func Day8() {
	log := logger.New("Day8")
	var junctionBoxConnections JunctionBoxConnections
	junctionBoxConnections.buildJunctionBoxes("day8.part0")
	junctionBoxConnections.buildJunctionBoxDistanceMatrix()
	junctionBoxConnections.buildJunctionBoxConnections()

	// timer := log.Timer("Build Junction Box Distance Matrix Timer")
	// timer()
	for _, connection := range junctionBoxConnections.connections {
		log.Info("Part 1", "length", len(connection.connections))
	}
}

func (j *JunctionBoxConnections) buildJunctionBoxConnections() {
	// Loop through each of the junction box matrices and find the shortest path for each junction box
	// for _, junctionBox := range j.allJunctionBoxes {
	// 	if _, ok := j.junctionBoxConnectionIndex[junctionBox]; ok {
	// 		slog.Info("junctionBox already has a connectionIndex", "junctionBox", junctionBox)
	// 		continue
	// 	}
	// 	shortestPath := 0
	// 	shortestPathJunctionBox := JunctionBox{}
	// 	for _, otherJunctionBox := range j.allJunctionBoxes {
	// 		if junctionBox.X == otherJunctionBox.X && junctionBox.Y == otherJunctionBox.Y &&
	// 			junctionBox.Z == otherJunctionBox.Z {
	// 			continue
	// 		}
	//
	// 		distance := j.junctionBoxDistanceMatrix[junctionBox][otherJunctionBox]
	// 		if shortestPath == 0 || distance < shortestPath {
	// 			shortestPath = distance
	// 			shortestPathJunctionBox = otherJunctionBox
	// 		}
	// 	}
	// 	// Check if otherJunctionBox  already has a connectionIndex
	// 	// If it does, we add junctionBox to that slice and update the index
	// 	// Else, we create a new slice and add junctionBox and otherJunctionBox and update both indexes
	// 	if _, ok := j.junctionBoxConnectionIndex[shortestPathJunctionBox]; ok {
	// 		j.connections[j.junctionBoxConnectionIndex[shortestPathJunctionBox]].connections = append(
	// 			j.connections[j.junctionBoxConnectionIndex[shortestPathJunctionBox]].connections,
	// 			junctionBox,
	// 		)
	// 		j.junctionBoxConnectionIndex[junctionBox] = j.junctionBoxConnectionIndex[shortestPathJunctionBox]
	// 	} else {
	// 		j.connections = append(j.connections, JunctionBoxGroup{connections: []JunctionBox{junctionBox, shortestPathJunctionBox}})
	// 		j.junctionBoxConnectionIndex[junctionBox] = len(j.connections) - 1
	// 		j.junctionBoxConnectionIndex[shortestPathJunctionBox] = len(j.connections) - 1
	// 	}
	// }
}

func (j *JunctionBoxConnections) buildJunctionBoxDistanceMatrix() {
	j.junctionBoxDistanceMatrix = make(map[JunctionBox]map[JunctionBox]int)
	for _, junctionBox := range j.allJunctionBoxes {
		for _, otherJunctionBox := range j.allJunctionBoxes {
			if junctionBox.X == otherJunctionBox.X && junctionBox.Y == otherJunctionBox.Y &&
				junctionBox.Z == otherJunctionBox.Z {
				continue
			}

			dx := float64(junctionBox.X - otherJunctionBox.X)
			dy := float64(junctionBox.Y - otherJunctionBox.Y)
			dz := float64(junctionBox.Z - otherJunctionBox.Z)

			distance := int(math.Sqrt(dx*dx + dy*dy + dz*dz))

			if _, ok := j.junctionBoxDistanceMatrix[junctionBox]; !ok {
				j.junctionBoxDistanceMatrix[junctionBox] = make(map[JunctionBox]int)
			}

			j.junctionBoxDistanceMatrix[junctionBox][otherJunctionBox] = distance
		}
	}

	for key, junctionBox := range j.junctionBoxDistanceMatrix {
	}
}

func (j *JunctionBoxConnections) buildJunctionBoxes(fileName string) JunctionBoxConnections {
	file := utilities.ReadFile(fileName)

	for _, row := range file {
		coordinates := strings.Split(row, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])

		newJunctionBox := JunctionBox{
			X: x,
			Y: y,
			Z: z,
		}

		j.allJunctionBoxes = append(j.allJunctionBoxes, newJunctionBox)
	}

	j.junctionBoxConnectionIndex = make(map[JunctionBox]int)
	return JunctionBoxConnections{}
}
