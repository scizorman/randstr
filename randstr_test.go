package randstr

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	bt := map[string]string{
		"Default":              `abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`,
		"OnlyNumber":           `1234567890`,
		"WithSpecialCharacter": `abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`,
		"WithJapanese":         `あいうえおカキクケコ東京都大阪府ABCDEfghjk12345@%+\/`,
	}

	for n, char := range bt {
		b.Run(n, func(b *testing.B) {
			b.ResetTimer()
			res := New(100, WithCharacters(char))
			b.Logf("generated string: %s", res)
		})
	}
}

func Test_bitsIndex(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Default",
			args: args{
				chars: []rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`),
			},
			want: 6,
		},
		{
			name: "OnlyNumber",
			args: args{
				chars: []rune(`1234567890`),
			},
			want: 4,
		},
		{
			name: "WithSpecialCharacter",
			args: args{
				[]rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitsIndex(tt.args.chars); got != tt.want {
				t.Errorf("bitsIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskIndex(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Default",
			args: args{
				chars: []rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`),
			},
			want: 63,
		},
		{
			name: "OnlyNumber",
			args: args{
				chars: []rune(`1234567890`),
			},
			want: 15,
		},
		{
			name: "WithSpecialCharacter",
			args: args{
				[]rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
			},
			want: 63,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskIndex(tt.args.chars); got != tt.want {
				t.Errorf("maskIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxIndex(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Default",
			args: args{
				chars: []rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`),
			},
			want: 10,
		},
		{
			name: "OnlyNumber",
			args: args{
				chars: []rune(`1234567890`),
			},
			want: 15,
		},
		{
			name: "WithSpecialCharacter",
			args: args{
				[]rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIndex(tt.args.chars); got != tt.want {
				t.Errorf("maxIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
