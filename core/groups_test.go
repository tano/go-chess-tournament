package core

import (
	"reflect"
	"sort"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		participants []string
	}
	tests := []struct {
		name               string
		args               args
		wantAmountOfGroups int
	}{
		{"6 participants", args{participants: []string{"Alex", "Igor", "Olga", "Max", "Vladimir", "Vadim"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Split(tt.args.participants); !reflect.DeepEqual(len(got), tt.wantAmountOfGroups) {
				t.Errorf("len of Split() = %v, wantAmountOfGroups %v", got, tt.wantAmountOfGroups)
			}
		})
	}
}

func TestSplitOddAmountOfParticipants(t *testing.T) {
	// given
	participants := []string{"Alex", "Igor", "Olga", "Max", "Vladimir"}

	// when
	_, error := Split(participants)

	// then
	if error == nil {
		t.Errorf("Wanted error but got nil for 5 participants %v", participants)
	}
}

func TestRandomOrder(t *testing.T) {
	// given
	participants := []string{"Alex", "Igor", "Olga", "Max", "Vladimir", "Vadim"}

	// when
	firstSplit, _ := Split(participants)
	secondSplit, _ := Split(participants)
	thirdSplit, _ := Split(participants)

	// then
	if len(firstSplit) != len(secondSplit) && len(firstSplit) != len(thirdSplit) {
		t.Errorf("Size of split for first, second and third split should be equal")
	}

	if !reflect.DeepEqual(sortAndMerge(firstSplit), sortAndMerge(secondSplit)) || !reflect.DeepEqual(sortAndMerge(secondSplit), sortAndMerge(thirdSplit)) {
		t.Errorf("Consequence runs produce different results (size)")
	}

	sameOrder := true

out:
	for i, firstSplitPart := range firstSplit {
		for j, elementInFirstSplitPart := range firstSplitPart {
			if elementInFirstSplitPart != secondSplit[i][j] ||
				elementInFirstSplitPart != thirdSplit[i][j] ||
				secondSplit[i][j] != thirdSplit[i][j] {
				sameOrder = false
				break out
			}
		}
	}

	if sameOrder == true {
		t.Errorf("Consequence runs produce same order results")
	}

}

func sortAndMerge(splitToSort [][]string) []string {
	result := make([]string, 0)
	totalSize := 0
	for i := range splitToSort {
		totalSize += len(splitToSort[i])
	}
	for _, element := range splitToSort {
		tmp := make([]string, len(element))
		copy(tmp, element)
		sort.Strings(tmp)
		for _, value := range tmp {
			result = append(result, value)
		}
	}
	return result
}
