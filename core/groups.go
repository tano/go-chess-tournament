package core

import (
	"errors"
	"math/rand"
	"time"
)

func Split(participants []string) ([][]string, error) {
	middle := len(participants) / 2
	firstHalf := participants[:middle]
	secondHalf := participants[middle:]
	shuffle(&firstHalf)
	shuffle(&secondHalf)
	if participantsAreOK(participants) == true {
		firstCopy := make([]string, len(firstHalf))
		secondCopy := make([]string, len(secondHalf))
		copy(firstCopy, firstHalf)
		copy(secondCopy, secondHalf)
		return [][]string{firstCopy, secondCopy}, nil
	} else {
		return nil, errors.New("participants count should be even")
	}

}

func participantsAreOK(participants []string) bool {
	if len(participants)%2 == 0 {
		return true
	}
	return false
}

func shuffle(participants *[]string) {
	rand.Seed(time.Now().UnixNano())
	temp := *participants
	rand.Shuffle(len(temp), func(i, j int) {
		(*participants)[i], (*participants)[j] = (*participants)[j], (*participants)[i]
	})
}
