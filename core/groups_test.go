package core

import (
	"reflect"
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
			if got := Split(tt.args.participants); !reflect.DeepEqual(got, tt.wantAmountOfGroups) {
				t.Errorf("Split() = %v, wantAmountOfGroups %v", got, tt.wantAmountOfGroups)
			}
		})
	}
}
