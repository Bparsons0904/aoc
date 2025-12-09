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

type JunctionBoxGroup struct {
	connections []JunctionBox
}

func Day8() {
	log := logger.New("Day8")
	junctionBoxes, junctionBoxConnectionIndex := buildJunctionBoxes("day8.part1")
	connections := getSortedConnections(junctionBoxes)

	timer := log.Timer("Part 1 Timer")
	part1Results := buildJunctionBoxConnectionsPart1(
		connections,
		junctionBoxConnectionIndex,
		junctionBoxes,
		1000,
	)
	timer()

	junctionBoxes, junctionBoxConnectionIndex = buildJunctionBoxes("day8.part1")
	timer = log.Timer("Part 2 Timer")
	part2Results := buildJunctionBoxConnectionsPart2(
		connections,
		junctionBoxConnectionIndex,
		junctionBoxes,
	)
	timer()

	log.Info("part1", "Part 1", part1Results, "Part 2", part2Results)
}

type shortestConnection struct {
	distance int
	from     JunctionBox
	to       JunctionBox
}

func buildJunctionBoxConnectionsPart2(
	shortestConnections []shortestConnection,
	junctionBoxConnectionIndex map[JunctionBox]int,
	junctionBoxes []JunctionBox,
) int {
	connections := make([]JunctionBoxGroup, 0)

	connectionsMade := 0
	for i, connection := range shortestConnections {
		if len(connections) == 1 && len(connections[0].connections) == len(junctionBoxes) {
			return shortestConnections[i-1].from.X * shortestConnections[i-1].to.X
		}

		fromIndex := junctionBoxConnectionIndex[connection.from]
		toIndex := junctionBoxConnectionIndex[connection.to]

		connectionsMade++
		if fromIndex != -1 && fromIndex == toIndex {
			continue
		}

		if fromIndex == -1 && toIndex == -1 {
			connections = append(
				connections,
				JunctionBoxGroup{connections: []JunctionBox{connection.from, connection.to}},
			)
			junctionBoxConnectionIndex[connection.from] = len(connections) - 1
			junctionBoxConnectionIndex[connection.to] = len(connections) - 1
			continue
		}

		if fromIndex == -1 {
			connections[toIndex].connections = append(
				connections[toIndex].connections,
				connection.from,
			)
			junctionBoxConnectionIndex[connection.from] = toIndex
			continue
		}
		if toIndex == -1 {
			connections[fromIndex].connections = append(
				connections[fromIndex].connections,
				connection.to,
			)
			junctionBoxConnectionIndex[connection.to] = fromIndex
			continue
		}

		smallerIdx, largerIdx := fromIndex, toIndex
		if len(connections[fromIndex].connections) > len(connections[toIndex].connections) {
			smallerIdx, largerIdx = toIndex, fromIndex
		}

		for _, box := range connections[smallerIdx].connections {
			junctionBoxConnectionIndex[box] = largerIdx
		}

		connections[largerIdx].connections = append(
			connections[largerIdx].connections,
			connections[smallerIdx].connections...,
		)

		connections = append(connections[:smallerIdx], connections[smallerIdx+1:]...)

		for box, idx := range junctionBoxConnectionIndex {
			if idx > smallerIdx {
				junctionBoxConnectionIndex[box] = idx - 1
			}
		}

	}

	return 0
}

func getSortedConnections(junctionBoxes []JunctionBox) []shortestConnection {
	shortestConnections := make([]shortestConnection, 0)

	for i := range junctionBoxes {
		for k := i + 1; k < len(junctionBoxes); k++ {
			junctionBox := junctionBoxes[i]
			otherJunctionBox := junctionBoxes[k]

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

	return shortestConnections
}

func buildJunctionBoxConnectionsPart1(
	shortestConnections []shortestConnection,
	junctionBoxConnectionIndex map[JunctionBox]int,
	junctionBoxes []JunctionBox,
	limit int,
) int {
	connections := make([]JunctionBoxGroup, 0)

	connectionsMade := 0
	for _, connection := range shortestConnections {
		if connectionsMade >= limit {
			break
		}

		fromIndex := junctionBoxConnectionIndex[connection.from]
		toIndex := junctionBoxConnectionIndex[connection.to]

		connectionsMade++
		if fromIndex != -1 && fromIndex == toIndex {
			continue
		}

		if fromIndex == -1 && toIndex == -1 {
			connections = append(
				connections,
				JunctionBoxGroup{connections: []JunctionBox{connection.from, connection.to}},
			)
			junctionBoxConnectionIndex[connection.from] = len(connections) - 1
			junctionBoxConnectionIndex[connection.to] = len(connections) - 1
			continue
		}

		if fromIndex == -1 {
			connections[toIndex].connections = append(
				connections[toIndex].connections,
				connection.from,
			)
			junctionBoxConnectionIndex[connection.from] = toIndex
			continue
		}
		if toIndex == -1 {
			connections[fromIndex].connections = append(
				connections[fromIndex].connections,
				connection.to,
			)
			junctionBoxConnectionIndex[connection.to] = fromIndex
			continue
		}

		smallerIdx, largerIdx := fromIndex, toIndex
		if len(connections[fromIndex].connections) > len(connections[toIndex].connections) {
			smallerIdx, largerIdx = toIndex, fromIndex
		}

		for _, box := range connections[smallerIdx].connections {
			junctionBoxConnectionIndex[box] = largerIdx
		}

		connections[largerIdx].connections = append(
			connections[largerIdx].connections,
			connections[smallerIdx].connections...,
		)
	}

	for _, junctionBox := range junctionBoxes {
		if junctionBoxConnectionIndex[junctionBox] == -1 {
			connections = append(
				connections,
				JunctionBoxGroup{connections: []JunctionBox{junctionBox}},
			)
			junctionBoxConnectionIndex[junctionBox] = len(connections) - 1
		}
	}

	slices.SortFunc(connections, func(a, b JunctionBoxGroup) int {
		return len(b.connections) - len(a.connections)
	})

	result := 0
	for i, connection := range connections {
		if i >= 3 {
			break
		}
		if result == 0 {
			result = len(connection.connections)
		} else {
			result *= len(connection.connections)
		}
	}

	return result
}

func buildJunctionBoxes(fileName string) ([]JunctionBox, map[JunctionBox]int) {
	file := utilities.ReadFile(fileName)

	junctionBoxes := make([]JunctionBox, 0)
	junctionBoxConnectionIndex := make(map[JunctionBox]int)
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

		junctionBoxes = append(junctionBoxes, newJunctionBox)
		junctionBoxConnectionIndex[newJunctionBox] = -1
	}

	return junctionBoxes, junctionBoxConnectionIndex
}
