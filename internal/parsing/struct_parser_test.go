package parsing

import (
	"reflect"
	"testing"
)

const src = `
type T struct {
	// 第一个参数
	F1 bool
	F2 int
	F3 int8
	F4 int16
	F5 int32
	F6 int64
	F7 uint
	F8 uint8
	F9 uint16
	F10 uint32
	F11 uint64
	F12 uintptr
	F13 float32
	F14 float64
	F15 complex64
	F16 complex128
	F17 interface{}
	F18 map[string]string
	F19 string
	F20 []string
	F21 struct{}
	F22 name
}

type name struct {
}
`

func TestStructParser(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name           string
		args           args
		wantStructList []*StructFlat
		wantErr        bool
	}{
		{
			name: "1",
			args: args{
				src: src,
			},
			wantStructList: nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStructList, err := StructParser(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStructList, tt.wantStructList) {
				t.Errorf("StructParser() gotStructList = %v, want %v", gotStructList, tt.wantStructList)
			}
		})
	}
}
