package challenges

import (
	"log/slog"
	"strings"

	"aoc/utilities"
)

// type ServerRack struct {
// 	serialNumber   string
// 	connectedRacks []string
// }

type ServerRacks map[string][]string

func Day11() {
	slog.Info("day11")

	serverRacks := processDay11File("day11.part1")
	numberOfConnectionPaths := getNumberOfConnectionPaths(serverRacks)

	slog.Info("Day 11", "Part 1", numberOfConnectionPaths)
}

func getNumberOfConnectionPaths(serverRacks ServerRacks) int {
	slog.Info("getNumberOfConnectionPaths", "serverRacks", serverRacks)

	connectedRacks := 0
	startingConnection := serverRacks["you"]
	for _, connection := range startingConnection {
		value := checkIfConnected(connection, serverRacks, 0)
		connectedRacks += value
	}
	return connectedRacks
}

func checkIfConnected(serialNumber string, serverRacks ServerRacks, paths int) int {
	serverRack := serverRacks[serialNumber]

	for _, connection := range serverRack {
		if connection == "out" {
			paths++
		} else {
			paths += checkIfConnected(connection, serverRacks, 0)
		}
	}

	slog.Info("no connection", "serialNumber", serialNumber, "connection", serverRack)
	return paths
}

func processDay11File(fileName string) ServerRacks {
	file := utilities.ReadFile(fileName)

	serverRacks := make(map[string][]string)
	for _, row := range file {
		rowParts := strings.Split(row, " ")

		var connections []string
		for _, part := range rowParts[1:] {
			if part != "" {
				connections = append(connections, part)
			}
		}

		serialNumber := rowParts[0]
		serverRacks[serialNumber[:len(serialNumber)-1]] = connections
	}

	return serverRacks
}
