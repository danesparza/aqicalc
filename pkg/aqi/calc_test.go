package aqi_test

import (
	"github.com/danesparza/aqicalc/pkg/aqi"
	"testing"
)

func Test_NewCalcService_Calculate(t *testing.T) {

	type fields struct {
		Pollutants map[string]aqi.PollutantInfo
	}
	type args struct {
		pollutionName string
		concentration float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Pollution not found",
			fields: fields{
				Pollutants: map[string]aqi.PollutantInfo{
					"o3": {
						Name: "o3",
						Breakpoints: []aqi.BreakPoint{
							{
								CategoryName:      "Good",
								AQIHigh:           50,
								AQILow:            0,
								ConcentrationHigh: .054,
								ConcentrationLow:  0,
							},
						},
					},
				},
			},
			args: args{
				pollutionName: "so2",
				concentration: .24,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Pollution concentration less than 0",
			fields: fields{
				Pollutants: map[string]aqi.PollutantInfo{
					"o3": {
						Name: "o3",
						Breakpoints: []aqi.BreakPoint{
							{
								CategoryName:      "Good",
								AQIHigh:           50,
								AQILow:            0,
								ConcentrationHigh: .054,
								ConcentrationLow:  0,
							},
						},
					},
				},
			},
			args: args{
				pollutionName: "so2",
				concentration: -.24,
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := aqi.NewCalcService(tt.fields.Pollutants)

			got, err := c.Calculate(tt.args.pollutionName, tt.args.concentration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewCalcServiceWithDefaults_Calculate(t *testing.T) {

	type args struct {
		pollutionName string
		concentration float32
	}
	tests := []struct {
		name    string
		args    args
		want    int
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
			name: "pm2_5: Unhealthy for sensitive groups",
			args: args{
				pollutionName: "pm2_5",
				concentration: 35.9,
			},
			want:    102,
			wantErr: false,
		},
		{
			name: "o3: Unhealthy for sensitive groups",
			args: args{
				pollutionName: "o3",
				concentration: .078,
			},
			want:    126,
			wantErr: false,
		},
		{
			name: "co: Moderate",
			args: args{
				pollutionName: "co",
				concentration: 8.4,
			},
			want:    90,
			wantErr: false,
		},
		{
			name: "so2: Beyond hazardous",
			args: args{
				pollutionName: "so2",
				concentration: 1100,
			},
			want:    400,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := aqi.NewCalcServiceWithDefaults()

			got, err := c.Calculate(tt.args.pollutionName, tt.args.concentration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
