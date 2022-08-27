package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	commands := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}
	payload := [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}
	for idx := 0; idx < b.N; idx++ {
		RunTest(commands, payload)
	}
}
func TestRunTest(t *testing.T) {
	type args struct {
		commands []string
		payload  [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{commands: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}, payload: [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}}},
			want: []string{"null", "null", "null", "null", "-3", "null", "0", "-2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunTest(tt.args.commands, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
