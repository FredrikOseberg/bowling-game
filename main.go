package main

import "fmt"

// Rules of the game:
// 1. Each game of bowling consists of 10 rounds
// 2. Each round consists of two throws.
// 3. A strike or a spare adds bonus
// 4. A spare occurs when a total of 10 pins is knocked down in two throws.
// 5. Scoring for a spare is 10, plus the number of pins knocked down in the nest throw.
// 6. A strike occurs when a total of 10 pins is knocked down in one throw.
// 7. Scoring for a strike is 10, plus the number of pins knocked down in the next two throws.

// Inputs:
// a list denoting the amount of pins downed per throw.
// Example: 2, 3, 5, 4, 9, 1, 2, 5, 3, 2, 4, 2, 3, 3, 4, 6, 10, 3, 2

// Outputs:
// The final calculated score of the game.
// Example:
// |f1  |f2  |f3  |f4  |f5  |f6  |f7  |f8  |f9|   f10 |
// |-, 3|5, -|9, /|2, 5|3, 2|4, 2|3, 3|4, /|X |X, 2, 5|
// score: 103

type Frame struct {
	RollIndex       int
	Total           int
	FrameType       string
	CalculatedTotal *int
}

func main() {
	//input := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 9, 1, 4}
	//input := []int{2, 3, 5, 4, 9, 1, 2, 5, 3, 2, 4, 2, 3, 3, 4, 6, 10, 3, 2}
	// input := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	input := []int{1, 5, 8, 2, 2, 4, 10, 10, 5, 3, 9, 1, 10, 10, 10, 4, 3}
	//input := []int{1, 5, 8, 2, 2, 4, 10, 10, 5, 3, 9, 1, 10, 10, 9, 1, 4, 3}
	bowlingGame(input)
}

func bowlingGame(rolls []int) {
	frames := getFrames(rolls)

	lastFrame := frames[len(frames)-1]
	if lastFrame.FrameType == "strike" || lastFrame.FrameType == "spare" {
		lastItems := rolls[lastFrame.RollIndex+1:]

		if len(lastItems) != 2 && lastFrame.FrameType == "strike" || len(lastItems) != 1 && lastFrame.FrameType == "spare" {
			fmt.Print("Invalid input")
			return
		}

	}

	calculatedFrames, totalScore := calculateScore(frames, rolls)

	renderResult(calculatedFrames, totalScore, rolls)
}

func renderResult(frames []Frame, score int, rolls []int) {
	headers := ""
	results := ""

	for index, val := range frames {
		fmt.Printf("%v, calculatedScore: %v\n", val, *val.CalculatedTotal)
		if index == 0 {
			headers += fmt.Sprintf("|\tf%v\t|", (index + 1))
		} else {
			headers += fmt.Sprintf("\tf%v\t|", (index + 1))
		}

		rollString := formatRoll(val, rolls)
		if index == 0 {
			results += fmt.Sprintf("|\t%s\t|", rollString)
		} else {
			results += fmt.Sprintf("\t%s\t|", rollString)
		}
	}

	headers += "\n"
	results += "\n"

	fmt.Print(headers)
	fmt.Print(results)
	fmt.Printf("Score: %d\n", score)
}

func formatRoll(frame Frame, rolls []int) string {
	if frame.FrameType == "strike" {
		return "X"
	}

	if frame.FrameType == "spare" {
		return fmt.Sprintf("%d, %s", rolls[frame.RollIndex-1], "/")
	}

	return fmt.Sprintf("%d, %d", rolls[frame.RollIndex-1], rolls[frame.RollIndex])

}

func calculateScore(frames []Frame, rolls []int) ([]Frame, int) {
	totalScore := 0
	calculatedFrames := []Frame{}

	for _, val := range frames {

		if val.FrameType == "normal" {
			total := val.Total
			val.CalculatedTotal = &total

			totalScore = totalScore + *val.CalculatedTotal
		}

		if val.FrameType == "spare" {
			total := val.Total + rolls[val.RollIndex+1]
			val.CalculatedTotal = &total

			totalScore = totalScore + *val.CalculatedTotal
		}

		if val.FrameType == "strike" {
			total := val.Total + rolls[val.RollIndex+1] + rolls[val.RollIndex+2]
			val.CalculatedTotal = &total

			totalScore = totalScore + *val.CalculatedTotal
		}

		calculatedFrames = append(calculatedFrames, val)
	}

	return calculatedFrames, totalScore
}

func createFrame(rollIndex, total int, rollType string) Frame {
	return Frame{
		RollIndex:       rollIndex,
		Total:           total,
		FrameType:       rollType,
		CalculatedTotal: nil,
	}
}

func getFrames(rolls []int) []Frame {
	var frames []Frame

	if len(rolls) < 10 || len(rolls) > 23 {
		return frames
	}

	temp := []int{}
	for index, val := range rolls {
		if len(frames) == 10 {
			break
		}

		normalRoll := val < 10 && len(temp) < 2
		if normalRoll {
			temp = append(temp, val)
		}

		strike := val == 10 && len(temp) == 0

		if strike {
			frame := createFrame(index, val, "strike")
			frames = append(frames, frame)
			temp = []int{}
		}

		if len(temp) == 2 {
			total := temp[0] + temp[1]
			var frame Frame
			spare := total == 10
			if spare {
				frame = createFrame(index, total, "spare")
			} else {
				frame = createFrame(index, total, "normal")
			}

			frames = append(frames, frame)
			temp = []int{}
		}
	}

	return frames
}
