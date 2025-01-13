package time

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "5h20m",
			args:    args{"5h20m"},
			want:    time.Hour*5 + 20*time.Minute,
			wantErr: false,
		},
		{
			name:    "1d5h20m",
			args:    args{"1d5h20m"},
			want:    24*time.Hour + time.Hour*5 + 20*time.Minute,
			wantErr: false,
		},
		{
			name:    "1d",
			args:    args{"1d"},
			want:    24 * time.Hour,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHumanDurationMillis(t *testing.T) {
	fmt.Println(ParseHumanDurationMillis(4886))               // 输出: "4 sec 886 ms"
	fmt.Println(ParseHumanDurationMillis(61020))              // 输出: "1 min 1 sec 20 msec"
	fmt.Println(ParseHumanDurationMillis(61000))              // 输出: "1 min 1 sec"
	fmt.Println(ParseHumanDurationMillis(3.6e+6))             // 输出: "1 hour"
	fmt.Println(ParseHumanDurationMillis(86400000 + 3600000)) // 输出: "1 day 1 hour"
	fmt.Println(ParseHumanDurationMillis(123456789))          // 输出: "1 day 10 hour 17 min 36 sec 789 msec"
}

func TestParseHumanTimeCost(t *testing.T) {
	start, _ := time.Parse("2006-01-02 15:04:05", "2024-11-07 12:00:00.008")
	end, _ := time.Parse("2006-01-02 15:04:05", "2024-11-07 13:12:09")
	fmt.Println(ParseHumanTimeCost(start, end))
}
