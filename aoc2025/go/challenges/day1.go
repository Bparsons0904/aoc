package challenges

import (
	"log"
	"log/slog"
	"strconv"

	"aoc/utilities"
)

type DialInstruction struct {
	Direction string
	Step      int
}

type DialInstructions []DialInstruction

type Dial struct {
	Value int
	Next  *Dial
	Prev  *Dial
}

type DialList struct {
	Head *Dial
	Tail *Dial
	Size int
}

func Day1() {
	files := utilities.ReadFile("day1.part1")

	var dialInstructions DialInstructions
	dialInstructions.parseInstructions(files)

	var dialList DialList
	dialList.buildLinkedList()
	for {
		if dialList.Head.Value == 50 {
			break
		}
		dialList.next()
	}

	day1part1Count := 0
	day1part2Count := 0
	for _, dialInstruction := range dialInstructions {
		switch dialInstruction.Direction {
		case "R":
			for j := 0; j < dialInstruction.Step; j++ {
				dialList.next()
				if dialList.Head.Value == 0 {
					day1part2Count++
				}
			}

		case "L":
			for j := 0; j < dialInstruction.Step; j++ {
				dialList.prev()
				if dialList.Head.Value == 0 {
					day1part2Count++
				}
			}
		}
		if dialList.Head.Value == 0 {
			day1part1Count++
		}
	}

	slog.Info("day1", "part1", day1part1Count, "part2", day1part2Count)
}

func (d *DialList) next() {
	d.Head = d.Head.Next
}

func (d *DialList) prev() {
	d.Head = d.Head.Prev
}

func (d *DialList) buildLinkedList() {
	var prevNode *Dial
	for i := range 100 {
		node := &Dial{
			Value: i,
			Next:  nil,
			Prev:  prevNode,
		}

		if d.Head == nil {
			d.Head = node
		}

		if prevNode != nil {
			prevNode.Next = node
		}

		d.Tail = node
		prevNode = node
		d.Size++
	}

	d.Tail.Next = d.Head
	d.Head.Prev = d.Tail
}

func (di *DialInstructions) parseInstructions(file []string) {
	for i := range file {
		var instruction DialInstruction
		instruction.Direction = string(file[i][0])
		step, err := strconv.Atoi(file[i][1:])
		if err != nil {
			log.Fatal(err)
		}
		instruction.Step = step

		*di = append(*di, instruction)
	}
}
