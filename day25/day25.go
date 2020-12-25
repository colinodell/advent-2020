package main

func transformSubjectNumber(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = (value * subjectNumber) % 20201227
	}

	return value
}

func findLoopSize(subjectNumber, publicKey int) int {
	i, value := 1, 1
	for {
		value = (value * subjectNumber) % 20201227

		if value == publicKey {
			return i
		}

		i++
	}
}

func Part1(cardPublicKey, doorPublicKey int) int {
	return transformSubjectNumber(doorPublicKey, findLoopSize(7, cardPublicKey))
}
