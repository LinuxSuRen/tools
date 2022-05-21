package pkg

import "testing"

func TestJsonFormat(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "normal",
		args: args{
			txt: `{"a":"b"}`,
		},
		want: `{
  "a": "b"
}`,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := JsonFormat(tt.args.txt); got != tt.want {
				t.Errorf("JsonFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
