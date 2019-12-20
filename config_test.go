package randstr

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestWithChars(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "NoDuplicates",
			args: args{
				str: `abcdefghijklmnopqrstuvwxyz`,
			},
			want: &Config{
				chars: []rune(`abcdefghijklmnopqrstuvwxyz`),
			},
		},
		{
			name: "Duplicates",
			args: args{
				str: `123456789012345`,
			},
			want: &Config{
				chars: []rune(`1234567890`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Config{}
			if WithChars(tt.args.str)(got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "NoDuplicates",
			args: args{
				chars: []rune(`abcdefghijklmnopqrstuvwxyz`),
			},
			want: []rune(`abcdefghijklmnopqrstuvwxyz`),
		},
		{
			name: "Duplicates",
			args: args{
				chars: []rune(`123456789012345`),
			},
			want: []rune(`1234567890`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.chars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
