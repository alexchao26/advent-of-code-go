package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`

func Test_twoFA(t *testing.T) {
	type args struct {
		instructions  string
		height, width int
	}
	tests := []struct {
		name             string
		args             args
		wantCount        int
		wantOutputString string
	}{
		{"example", args{example, 3, 7}, 6, " #  # #\n# #    \n #     \n"},
		{"actual", args{util.ReadFile("input.txt"), 6, 50}, 115, `#### #### #### #   ##  # #### ###  ####  ###   ## 
#    #    #    #   ## #  #    #  # #      #     # 
###  ###  ###   # # ##   ###  #  # ###    #     # 
#    #    #      #  # #  #    ###  #      #     # 
#    #    #      #  # #  #    # #  #      #  #  # 
#### #    ####   #  #  # #    #  # #     ###  ##  
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotString := twoFA(tt.args.instructions, tt.args.height, tt.args.width)
			if gotCount != tt.wantCount {
				t.Errorf("twoFA().count = %v, want %v", gotCount, tt.wantCount)
			}
			if gotString != tt.wantOutputString {
				t.Errorf("twoFA().outputString = \n%q, want \n%q", gotString, tt.wantOutputString)
			}
		})
	}
}
