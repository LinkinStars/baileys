package generator

import (
	"testing"

	"github.com/LinkinStars/baileys/internal/converter"
)

func TestGenPBMessage(t *testing.T) {
	type args struct {
		flatList []*converter.PBFlat
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				flatList: []*converter.PBFlat{
					{
						Name: "Test",
						PBFieldList: []*converter.PBField{
							{
								Type:  "string",
								Name:  "Name",
								Index: 1,
							}, {
								Type:  "int",
								Name:  "Age",
								Index: 2,
							},
						},
					},
					{
						Name: "Test",
						PBFieldList: []*converter.PBField{
							{
								Type:  "string",
								Name:  "Name",
								Index: 1,
							}, {
								Type:  "int",
								Name:  "Age",
								Index: 2,
							},
						},
					},
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GenPBMessage(tt.args.flatList); got != tt.want {
				t.Errorf("GenPBMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
