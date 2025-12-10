import time
import math
from utilities import read_file


class JunctionBox:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y and self.z == other.z

    def __hash__(self):
        return hash((self.x, self.y, self.z))

    def __repr__(self):
        return f"JunctionBox({self.x}, {self.y}, {self.z})"


class JunctionBoxGroup:
    def __init__(self, connections=None):
        self.connections = connections if connections else []


class ShortestConnection:
    def __init__(self, distance, from_box, to_box):
        self.distance = distance
        self.from_box = from_box
        self.to_box = to_box


def build_junction_boxes(filename):
    lines = read_file(filename)
    junction_boxes = []
    junction_box_connection_index = {}

    for row in lines:
        coordinates = row.split(",")
        x = int(coordinates[0])
        y = int(coordinates[1])
        z = int(coordinates[2])

        new_junction_box = JunctionBox(x, y, z)
        junction_boxes.append(new_junction_box)
        junction_box_connection_index[new_junction_box] = -1

    return junction_boxes, junction_box_connection_index


def get_sorted_connections(junction_boxes):
    shortest_connections = []

    for i in range(len(junction_boxes)):
        for k in range(i + 1, len(junction_boxes)):
            junction_box = junction_boxes[i]
            other_junction_box = junction_boxes[k]

            dx = junction_box.x - other_junction_box.x
            dy = junction_box.y - other_junction_box.y
            dz = junction_box.z - other_junction_box.z

            distance = int(math.sqrt(dx * dx + dy * dy + dz * dz))

            shortest_connections.append(
                ShortestConnection(distance, junction_box, other_junction_box)
            )

    shortest_connections.sort(key=lambda c: c.distance)
    return shortest_connections


def build_junction_box_connections_part1(
    shortest_connections, junction_box_connection_index, junction_boxes, limit
):
    connections = []
    connections_made = 0

    for connection in shortest_connections:
        if connections_made >= limit:
            break

        from_index = junction_box_connection_index[connection.from_box]
        to_index = junction_box_connection_index[connection.to_box]

        connections_made += 1
        if from_index != -1 and from_index == to_index:
            continue

        if from_index == -1 and to_index == -1:
            connections.append(
                JunctionBoxGroup([connection.from_box, connection.to_box])
            )
            junction_box_connection_index[connection.from_box] = len(connections) - 1
            junction_box_connection_index[connection.to_box] = len(connections) - 1
            continue

        if from_index == -1:
            connections[to_index].connections.append(connection.from_box)
            junction_box_connection_index[connection.from_box] = to_index
            continue

        if to_index == -1:
            connections[from_index].connections.append(connection.to_box)
            junction_box_connection_index[connection.to_box] = from_index
            continue

        smaller_idx = from_index
        larger_idx = to_index
        if len(connections[from_index].connections) > len(
            connections[to_index].connections
        ):
            smaller_idx = to_index
            larger_idx = from_index

        for box in connections[smaller_idx].connections:
            junction_box_connection_index[box] = larger_idx

        connections[larger_idx].connections.extend(
            connections[smaller_idx].connections
        )

    for junction_box in junction_boxes:
        if junction_box_connection_index[junction_box] == -1:
            connections.append(JunctionBoxGroup([junction_box]))
            junction_box_connection_index[junction_box] = len(connections) - 1

    connections.sort(key=lambda g: len(g.connections), reverse=True)

    result = 0
    for i, connection in enumerate(connections):
        if i >= 3:
            break
        if result == 0:
            result = len(connection.connections)
        else:
            result *= len(connection.connections)

    return result


def build_junction_box_connections_part2(
    shortest_connections, junction_box_connection_index, junction_boxes
):
    connections = []
    connections_made = 0

    for i, connection in enumerate(shortest_connections):
        if (
            len(connections) == 1
            and len(connections[0].connections) == len(junction_boxes)
        ):
            return (
                shortest_connections[i - 1].from_box.x
                * shortest_connections[i - 1].to_box.x
            )

        from_index = junction_box_connection_index[connection.from_box]
        to_index = junction_box_connection_index[connection.to_box]

        connections_made += 1
        if from_index != -1 and from_index == to_index:
            continue

        if from_index == -1 and to_index == -1:
            connections.append(
                JunctionBoxGroup([connection.from_box, connection.to_box])
            )
            junction_box_connection_index[connection.from_box] = len(connections) - 1
            junction_box_connection_index[connection.to_box] = len(connections) - 1
            continue

        if from_index == -1:
            connections[to_index].connections.append(connection.from_box)
            junction_box_connection_index[connection.from_box] = to_index
            continue

        if to_index == -1:
            connections[from_index].connections.append(connection.to_box)
            junction_box_connection_index[connection.to_box] = from_index
            continue

        smaller_idx = from_index
        larger_idx = to_index
        if len(connections[from_index].connections) > len(
            connections[to_index].connections
        ):
            smaller_idx = to_index
            larger_idx = from_index

        for box in connections[smaller_idx].connections:
            junction_box_connection_index[box] = larger_idx

        connections[larger_idx].connections.extend(
            connections[smaller_idx].connections
        )

        connections.pop(smaller_idx)

        for box in list(junction_box_connection_index.keys()):
            if junction_box_connection_index[box] > smaller_idx:
                junction_box_connection_index[box] -= 1

    return 0


def day8():
    junction_boxes, junction_box_connection_index = build_junction_boxes("day8.part1")
    connections = get_sorted_connections(junction_boxes)

    start = time.time()
    part1_results = build_junction_box_connections_part1(
        connections, junction_box_connection_index, junction_boxes, 1000
    )
    part1_time = time.time() - start

    junction_boxes, junction_box_connection_index = build_junction_boxes("day8.part1")
    start = time.time()
    part2_results = build_junction_box_connections_part2(
        connections, junction_box_connection_index, junction_boxes
    )
    part2_time = time.time() - start

    print(f"Day 8:")
    print(f"  Part 1: {part1_results} ({part1_time:.4f}s)")
    print(f"  Part 2: {part2_results} ({part2_time:.4f}s)")


if __name__ == "__main__":
    day8()
