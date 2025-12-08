package challenges

import (
	"math"
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
	junctionBoxConnections.buildJunctionBoxes("day8.part1")
	junctionBoxConnections.buildJunctionBoxDistanceMatrix()
	junctionBoxConnections.buildJunctionBoxConnections()

	// timer := log.Timer("Build Junction Box Distance Matrix Timer")
	// timer()

	result := 0
	for i, connection := range junctionBoxConnections.connections {
		if i >= 3 {
			break
		}
		if result == 0 {
			result = len(connection.connections)
		} else {
			result *= len(connection.connections)
		}

	}

	log.Info("part1", "Part 1", result)
}

func (j *JunctionBoxConnections) buildJunctionBoxConnections() {
	type shortestConnection struct {
		distance int
		from     JunctionBox
		to       JunctionBox
	}

	shortestConnections := make([]shortestConnection, 0)

	for i := 0; i < len(j.junctionBoxes); i++ {
		for k := i + 1; k < len(j.junctionBoxes); k++ {
			junctionBox := j.junctionBoxes[i]
			otherJunctionBox := j.junctionBoxes[k]

			dx := float64(junctionBox.X - otherJunctionBox.X)
			dy := float64(junctionBox.Y - otherJunctionBox.Y)
			dz := float64(junctionBox.Z - otherJunctionBox.Z)

			distance := int(math.Sqrt(dx*dx + dy*dy + dz*dz))

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

	connectionsMade := 0
	for _, connection := range shortestConnections {
		if connectionsMade >= 1000 {
			break
		}

		fromIndex := j.junctionBoxConnectionIndex[connection.from]
		toIndex := j.junctionBoxConnectionIndex[connection.to]

		connectionsMade++
		if fromIndex != -1 && fromIndex == toIndex {
			continue
		}

		if fromIndex == -1 && toIndex == -1 {
			j.connections = append(
				j.connections,
				JunctionBoxGroup{connections: []JunctionBox{connection.from, connection.to}},
			)
			j.junctionBoxConnectionIndex[connection.from] = len(j.connections) - 1
			j.junctionBoxConnectionIndex[connection.to] = len(j.connections) - 1
			continue
		}

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

		smallerIdx, largerIdx := fromIndex, toIndex
		if len(j.connections[fromIndex].connections) > len(j.connections[toIndex].connections) {
			smallerIdx, largerIdx = toIndex, fromIndex
		}

		for _, box := range j.connections[smallerIdx].connections {
			j.junctionBoxConnectionIndex[box] = largerIdx
		}

		j.connections[largerIdx].connections = append(
			j.connections[largerIdx].connections,
			j.connections[smallerIdx].connections...,
		)

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

	slices.SortFunc(j.connections, func(a, b JunctionBoxGroup) int {
		return len(b.connections) - len(a.connections)
	})
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
