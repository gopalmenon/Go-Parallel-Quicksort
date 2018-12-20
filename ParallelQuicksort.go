package sorting

func Quicksort(input []int, startIndex, endIndex int) {

	if startIndex >= endIndex {
		return
	}

	pivotIndex := partition(input, startIndex, endIndex)
	Quicksort(input, startIndex, pivotIndex - 1)
	Quicksort(input, pivotIndex + 1, endIndex)

}

func ParallelQuicksort(input []int, startIndex, endIndex int, quit chan<-bool) {

	if startIndex >= endIndex {
		if quit != nil {
			quit <- true
		}
		return
	}

	quitChild  := make(chan bool)
	pivotIndex := partition(input, startIndex, endIndex)
	go ParallelQuicksort(input, startIndex, pivotIndex - 1, quitChild)
	ParallelQuicksort(input, pivotIndex + 1, endIndex, nil)

	<- quitChild

	if quit != nil {
		quit <- true
	}

}
func partition(input []int, startIndex, endIndex int) int {

	pivot := input[endIndex]
	slidingWindowLeft := startIndex - 1
	for slidingWindowRight := startIndex; slidingWindowRight < endIndex; slidingWindowRight++ {
		if input[slidingWindowRight] <= pivot {
			slidingWindowLeft += 1
			input[slidingWindowLeft], input[slidingWindowRight] = input[slidingWindowRight], input[slidingWindowLeft]
		}
	}
	input[slidingWindowLeft + 1], input[endIndex] = input[endIndex], input[slidingWindowLeft + 1]
	return slidingWindowLeft + 1

}