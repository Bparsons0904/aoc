package challenges

import (
	"log"
	"log/slog"
	"strconv"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
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
}

func Day1() {
	log := logger.New("Day1")
	done := log.Timer("Day1 Timer")
	defer done()

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

func Day1_1() {
	log := logger.New("Day1_1")
	done := log.Timer("Timer")
	defer done()

	files := utilities.ReadFile("day1.part1")

	var dialInstructions DialInstructions
	dialInstructions.parseInstructions(files)

	step1Count := 0
	step2Count := 0
	currentValue := 50

	for _, dialInstruction := range dialInstructions {
		switch dialInstruction.Direction {
		case "R":
			newValue := currentValue + dialInstruction.Step
			if newValue >= 100 {
				if currentValue == 0 {
					step2Count += dialInstruction.Step / 100
				} else {
					step2Count += (dialInstruction.Step + currentValue) / 100
				}
			}
			currentValue = newValue % 100
		case "L":
			newValue := currentValue - dialInstruction.Step
			if currentValue == 0 {
				step2Count += dialInstruction.Step / 100
			} else if dialInstruction.Step >= currentValue {
				step2Count += (dialInstruction.Step-currentValue)/100 + 1
			}
			currentValue = ((newValue % 100) + 100) % 100
		}
		if currentValue == 0 {
			step1Count++
		}
	}

	slog.Info("day1", "part1", step1Count, "part2", step2Count)
}
