package aqi_test

import (
	"github.com/danesparza/aqicalc/pkg/aqi"
	"testing"
)

func Test_convertService_ConvertPPMFromUgm3(t *testing.T) {
	type args struct {
		pollutionName string
		concentration float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "Pollution not found",
			args: args{
				pollutionName: "so40",
				concentration: .24,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Pollution concentration less than 0",
			args: args{
				pollutionName: "o3",
				concentration: -.24,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Pollution doesn't have molecular weight",
			args: args{
				pollutionName: "pm2_5",
				concentration: 35.9,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "o3: Good air quality",
			args: args{
				pollutionName: "o3",
				concentration: 59.37,
			},
			want:    0.03,
			wantErr: false,
		},
		{
			name: "co: Good air quality",
			args: args{
				pollutionName: "co",
				concentration: 260.35,
			},
			want:    0.227,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := aqi.NewConvertServiceWithDefaults()
			got, err := c.ConvertPPMFromUgm3(tt.args.pollutionName, tt.args.concentration)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertPPMFromUgm3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertPPMFromUgm3() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertService_ConvertPPBFromUgm3(t *testing.T) {
	type args struct {
		pollutionName string
		concentration float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "so2: Good air quality",
			args: args{
				pollutionName: "so2",
				concentration: 2.77,
			},
			want:    1.057,
			wantErr: false,
		},
		{
			name: "no2: Good air quality",
			args: args{
				pollutionName: "no2",
				concentration: 3.17,
			},
			want:    1.685,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := aqi.NewConvertServiceWithDefaults()
			got, err := c.ConvertPPBFromUgm3(tt.args.pollutionName, tt.args.concentration)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertPPBFromUgm3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertPPBFromUgm3() got = %v, want %v", got, tt.want)
			}
		})
	}
}
