package randstr

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func BenchmarkNew(b *testing.B) {
	benches := []struct {
		name  string
		chars string
	}{
		{
			name:  "Default",
			chars: `abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`,
		},
		{
			name:  "OnlyNumber",
			chars: `1234567890`,
		},
		{
			name:  "WithSpecialCharacters",
			chars: `abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`,
		},
		{
			name:  "WithJapanese",
			chars: `あいうえおカキクケコ東京都大阪府ABCDEfghjk12345@%+\/`,
		},
	}
	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			b.ResetTimer()
			New(30, WithCharacters(bb.chars))
		})
	}
}

func TestNew(t *testing.T) {
	Seed(1)

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
			},
			want: `dtpsBCLKwhCGHLF9EoWyo1KFHeio1r`,
		},
		{
			name: "OnlyNumber",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithCharacters(`1234567890`),
				},
			},
			want: `408027112802971279976969250732`,
		},
		{
			name: "WithSpecialCharacters",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithCharacters(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
				},
			},
			want: `mN\FyCD7eM/r\!zm1d:KpGue1/wtLu`,
		},
		{
			name: "WithJapanese",
			args: args{
				l: 30,
				opts: []func(*Config){
					WithCharacters(`あいうえおカキクケコ東京都大阪府ABCDEfghjk12345@%+\/`),
				},
			},
			want: `コhECB2A4D%おBコhk\ケf+京うhキえククE東Eh`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.l, tt.args.opts...); got != tt.want {
				t.Errorf("New() mismatch (-want +got):\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func Test_newWithConfig(t *testing.T) {
	Seed(1)

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
					characters: defaultCharacters,
				},
			},
			want: `dtpsBCLKwhCGHLF9EoWyo1KFHeio1r`,
		},
		{
			name: "OnlyNumber",
			args: args{
				l: 30,
				conf: &Config{
					characters: []rune(`1234567890`),
				},
			},
			want: `408027112802971279976969250732`,
		},
		{
			name: "WithSpecialCharacters",
			args: args{
				l: 30,
				conf: &Config{
					characters: []rune(`abcdefgihijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@%+\/'!#$^?:(){}[]~-_`),
				},
			},
			want: `lM+ExBC6eL\q+'ylZd?JoFteZ\vsKt`,
		},

		{
			name: "WithJapanese",
			args: args{
				l: 30,
				conf: &Config{
					characters: []rune(`あいうえおカキクケコ東京都大阪府ABCDEfghjk12345@%+\/`),
				},
			},
			want: `コhECB2A4D%おBコhk\ケf+京うhキえククE東Eh`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWithConfig(tt.args.l, tt.args.conf); got != tt.want {
				t.Errorf("newWithConfig() mismatch (-want +got):\n%s", cmp.Diff(tt.want, got))
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
			name: "WithSpecialCharacters",
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
			name: "WithSpecialCharacters",
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
			name: "WithSpecialCharacters",
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
