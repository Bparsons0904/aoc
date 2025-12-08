package challenges

import (
	"fmt"
	"slices"
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
	junctionBoxes              []JunctionBox
	junctionBoxConnectionIndex map[JunctionBox]int
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
		if len(connection.connections) == 0 {
			continue
		}
		log.Info(
			"Part 1",
			"length",
			len(connection.connections),
			"connections",
			connection.connections,
		)
	}
}

func (j *JunctionBoxConnections) buildJunctionBoxConnections() {
	type shortestConnection struct {
		distance int
		from     JunctionBox
		to       JunctionBox
	}

	shortestConnections := make([]shortestConnection, 0)

	// Generate all unique pairs (i < j to avoid duplicates)
	for i := 0; i < len(j.junctionBoxes); i++ {
		for k := i + 1; k < len(j.junctionBoxes); k++ {
			junctionBox := j.junctionBoxes[i]
			otherJunctionBox := j.junctionBoxes[k]

			dx := junctionBox.X - otherJunctionBox.X
			dy := junctionBox.Y - otherJunctionBox.Y
			dz := junctionBox.Z - otherJunctionBox.Z

			// Use squared distance to avoid floating point truncation issues
			distance := dx*dx + dy*dy + dz*dz

			shortestConnections = append(shortestConnections, shortestConnection{
				distance: distance,
				from:     junctionBox,
				to:       otherJunctionBox,
			})
		}
	}

	slices.SortFunc(shortestConnections, func(a, b shortestConnection) int {
		return a.distance - b.distance
	})

	// Debug: print first 15 connections
	for i := 0; i < 15 && i < len(shortestConnections); i++ {
		c := shortestConnections[i]
		fmt.Printf("Pair %d: dist=%d from=%v to=%v\n", i, c.distance, c.from, c.to)
	}

	connectionsMade := 0
	for _, connection := range shortestConnections {
		if connectionsMade >= 10 {
			break
		}

		fromIndex := j.junctionBoxConnectionIndex[connection.from]
		toIndex := j.junctionBoxConnectionIndex[connection.to]

		connectionsMade++

		// Already in same circuit - still counts as an attempt
		if fromIndex != -1 && fromIndex == toIndex {
			continue
		}

		// Both unassigned - create new circuit
		if fromIndex == -1 && toIndex == -1 {
			j.connections = append(
				j.connections,
				JunctionBoxGroup{connections: []JunctionBox{connection.from, connection.to}},
			)
			j.junctionBoxConnectionIndex[connection.from] = len(j.connections) - 1
			j.junctionBoxConnectionIndex[connection.to] = len(j.connections) - 1
			continue
		}

		// One is unassigned - add to existing circuit
		if fromIndex == -1 {
			j.connections[toIndex].connections = append(
				j.connections[toIndex].connections,
				connection.from,
			)
			j.junctionBoxConnectionIndex[connection.from] = toIndex
			continue
		}
		if toIndex == -1 {
			j.connections[fromIndex].connections = append(
				j.connections[fromIndex].connections,
				connection.to,
			)
			j.junctionBoxConnectionIndex[connection.to] = fromIndex
			continue
		}

		// Both in different circuits - merge them
		smallerIdx, largerIdx := fromIndex, toIndex
		if len(j.connections[fromIndex].connections) > len(j.connections[toIndex].connections) {
			smallerIdx, largerIdx = toIndex, fromIndex
		}

		// Update index for every box in the smaller circuit
		for _, box := range j.connections[smallerIdx].connections {
			j.junctionBoxConnectionIndex[box] = largerIdx
		}

		// Move all boxes to the larger circuit
		j.connections[largerIdx].connections = append(
			j.connections[largerIdx].connections,
			j.connections[smallerIdx].connections...,
		)

		// Clear the smaller circuit
		j.connections[smallerIdx].connections = nil
	}

	// loop over all the junction boxes and see which still have an index of -1 and add them to the connections
	for _, junctionBox := range j.junctionBoxes {
		if j.junctionBoxConnectionIndex[junctionBox] == -1 {
			j.connections = append(
				j.connections,
				JunctionBoxGroup{connections: []JunctionBox{junctionBox}},
			)
			j.junctionBoxConnectionIndex[junctionBox] = len(j.connections) - 1
		}
	}
}

func (j *JunctionBoxConnections) buildJunctionBoxDistanceMatrix() {
	// j.junctionBoxDistanceMatrix = make(map[JunctionBox]map[JunctionBox]int)
	// for _, junctionBox := range j.junctionBoxes {
	// 	for _, otherJunctionBox := range j.junctionBoxes {
	// 		if junctionBox.X == otherJunctionBox.X && junctionBox.Y == otherJunctionBox.Y &&
	// 			junctionBox.Z == otherJunctionBox.Z {
	// 			continue
	// 		}
	//
	// 		dx := float64(junctionBox.X - otherJunctionBox.X)
	// 		dy := float64(junctionBox.Y - otherJunctionBox.Y)
	// 		dz := float64(junctionBox.Z - otherJunctionBox.Z)
	//
	// 		distance := int(math.Sqrt(dx*dx + dy*dy + dz*dz))
	//
	// 		if _, ok := j.junctionBoxDistanceMatrix[junctionBox]; !ok {
	// 			j.junctionBoxDistanceMatrix[junctionBox] = make(map[JunctionBox]int)
	// 		}
	//
	// 		j.junctionBoxDistanceMatrix[junctionBox][otherJunctionBox] = distance
	// 	}
	// }
	//
	// for key, junctionBox := range j.junctionBoxDistanceMatrix {
	// }
}

func (j *JunctionBoxConnections) buildJunctionBoxes(fileName string) JunctionBoxConnections {
	file := utilities.ReadFile(fileName)

	j.junctionBoxConnectionIndex = make(map[JunctionBox]int)
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

		j.junctionBoxes = append(j.junctionBoxes, newJunctionBox)
		j.junctionBoxConnectionIndex[newJunctionBox] = -1
	}

	return JunctionBoxConnections{}
}
