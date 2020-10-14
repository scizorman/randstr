package randstr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithCharacters(t *testing.T) {
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
				characters: []rune(`abcdefghijklmnopqrstuvwxyz`),
			},
		},
		{
			name: "Duplicates",
			args: args{
				str: `123456789012345`,
			},
			want: &Config{
				characters: []rune(`1234567890`),
			},
		},
	}
	opt := cmp.AllowUnexported(Config{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Config{}
			if WithCharacters(tt.args.str)(got); !cmp.Equal(tt.want, got, opt) {
				t.Errorf("WithChars() mismatch (-want +got):\n%s", cmp.Diff(tt.want, got, opt))
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
			if got := removeDuplicates(tt.args.chars); !cmp.Equal(tt.want, got) {
				t.Errorf("removeDuplicates() mismatch (-want +got):\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}
