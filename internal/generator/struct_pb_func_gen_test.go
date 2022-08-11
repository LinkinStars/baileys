package generator

import (
	"testing"

	"github.com/LinkinStars/baileys/internal/parsing"
)

func TestGenerateStruct2PBFunc(t *testing.T) {
	type args struct {
		structList []*parsing.StructFlat
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				structList: []*parsing.StructFlat{
					{
						Name:    "User",
						Comment: "用户",
						Fields: []*parsing.StructField{
							{
								Name:    "ID",
								Type:    "int64",
								Comment: "唯一标识",
							},
							{
								Name:    "Age",
								Type:    "int",
								Comment: "年龄",
							},
							{
								Name:    "Name",
								Type:    "string",
								Comment: "用户名",
							},
							{
								Name:    "CreatedAt",
								Type:    "time.Time",
								Comment: "创建时间",
							},
						},
					},
				},
			},
			wantRes: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GenerateStruct2PBFunc(tt.args.structList)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateStruct2PBFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("GenerateStruct2PBFunc() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
