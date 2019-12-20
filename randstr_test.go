package randstr

import (
	"math/rand"
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
			res := New(100, WithChars(char))
			b.Logf("generated string: %s", res)
		})
	}
}

func TestNew(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	type args struct {
		l    int
		opts []func(*Config)
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithRand(r),
				},
			},
			want: `dtpsBCLKwhCGHLF9EoWyo1KFHeio1r`,
		},
		{
			name: "OnlyNumber",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithChars(`1234567890`),
					WithRand(r),
				},
			},
			want: `408027112802971279976969250732`,
		},
		{
			name: "WithSpecialCharacter",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithChars(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
					WithRand(r),
				},
			},
			want: `mN\FyCD7eM/r\!zm1d:KpGue1/wtLu`,
		},
		{
			name: "WithJapanese",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithChars(`あいうえおカキクケコ東京都大阪府ABCDEfghjk12345@%+\/`),
					WithRand(r),
				},
			},
			want: `コhECB2A4D%おBコhk\ケf+京うhキえククE東Eh`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.l, tt.args.opts...); got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWithConfig(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	type args struct {
		l    int
		conf *Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "DefaultCharacters",
			args: args{
				l: 30,
				conf: &Config{
					chars: defaultChars,
					rand:  r,
				},
			},
			want: `dtpsBCLKwhCGHLF9EoWyo1KFHeio1r`,
		},
		{
			name: "OnlyNumber",
			args: args{
				l: 30,
				conf: &Config{
					chars: []rune(`1234567890`),
					rand:  r,
				},
			},
			want: `408027112802971279976969250732`,
		},
		{
			name: "WithSpecialCharacter",
			args: args{
				l: 30,
				conf: &Config{
					chars: []rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
					rand:  r,
				},
			},
			want: `lM+ExBC6eL\q+'ylZd?JoFteZ\vsKt`,
		},

		{
			name: "WithJapanese",
			args: args{
				l: 30,
				conf: &Config{
					chars: []rune(`あいうえおカキクケコ東京都大阪府ABCDEfghjk12345@%+\/`),
					rand:  r,
				},
			},
			want: `コhECB2A4D%おBコhk\ケf+京うhキえククE東Eh`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWithConfig(tt.args.l, tt.args.conf); got != tt.want {
				t.Errorf("newWithConfig() = %v, want %v", got, tt.want)
			}
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
			want: 7,
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
			want: 127,
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
			want: 9,
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
