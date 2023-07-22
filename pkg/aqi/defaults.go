package aqi

// NewCalcServiceWithDefaults returns a CalcService that
// uses the pollutant breakpoints from the US EPA
// Defaults taken from this technical document:
// https://www.airnow.gov/sites/default/files/2020-05/aqi-technical-assistance-document-sept2018.pdf
func NewCalcServiceWithDefaults() CalcService {
	return calcService{
		Pollutants: map[string]PollutantInfo{
			"o3": {
				Name: "o3",
				Breakpoints: []BreakPoint{
					{
						CategoryName:      "Good",
						AQIHigh:           50,
						AQILow:            0,
						ConcentrationHigh: .054,
						ConcentrationLow:  0,
					},
					{
						CategoryName:      "Moderate",
						AQIHigh:           100,
						AQILow:            51,
						ConcentrationHigh: .070,
						ConcentrationLow:  .055,
					},
					{
						CategoryName:      "Unhealthy for Sensitive Groups",
						AQIHigh:           150,
						AQILow:            101,
						ConcentrationHigh: .085,
						ConcentrationLow:  .071,
					},
					{
						CategoryName:      "Unhealthy",
						AQIHigh:           200,
						AQILow:            151,
						ConcentrationHigh: .105,
						ConcentrationLow:  .086,
					},
					{
						CategoryName:      "Very unhealthy",
						AQIHigh:           300,
						AQILow:            201,
						ConcentrationHigh: .200,
						ConcentrationLow:  .106,
					},
					{
						CategoryName:      "Hazardous",
						AQIHigh:           400,
						AQILow:            301,
						ConcentrationHigh: .300,
						ConcentrationLow:  .200,
					},
				},
			},
			"pm2_5": {
				Name: "pm2_5",
				Breakpoints: []BreakPoint{
					{
						CategoryName:      "Good",
						AQIHigh:           50,
						AQILow:            0,
						ConcentrationHigh: 12.0,
						ConcentrationLow:  0,
					},
					{
						CategoryName:      "Moderate",
						AQIHigh:           100,
						AQILow:            51,
						ConcentrationHigh: 35.4,
						ConcentrationLow:  12.1,
					},
					{
						CategoryName:      "Unhealthy for Sensitive Groups",
						AQIHigh:           150,
						AQILow:            101,
						ConcentrationHigh: 55.4,
						ConcentrationLow:  35.5,
					},
					{
						CategoryName:      "Unhealthy",
						AQIHigh:           200,
						AQILow:            151,
						ConcentrationHigh: 150.4,
						ConcentrationLow:  55.5,
					},
					{
						CategoryName:      "Very unhealthy",
						AQIHigh:           300,
						AQILow:            201,
						ConcentrationHigh: 250.4,
						ConcentrationLow:  150.5,
					},
					{
						CategoryName:      "Hazardous",
						AQIHigh:           400,
						AQILow:            301,
						ConcentrationHigh: 250.5,
						ConcentrationLow:  350.4,
					},
				},
			},
			"pm10": {
				Name: "pm10",
				Breakpoints: []BreakPoint{
					{
						CategoryName:      "Good",
						AQIHigh:           50,
						AQILow:            0,
						ConcentrationHigh: 54,
						ConcentrationLow:  0,
					},
					{
						CategoryName:      "Moderate",
						AQIHigh:           100,
						AQILow:            51,
						ConcentrationHigh: 154,
						ConcentrationLow:  55,
					},
					{
						CategoryName:      "Unhealthy for Sensitive Groups",
						AQIHigh:           150,
						AQILow:            101,
						ConcentrationHigh: 254,
						ConcentrationLow:  155,
					},
					{
						CategoryName:      "Unhealthy",
						AQIHigh:           200,
						AQILow:            151,
						ConcentrationHigh: 354,
						ConcentrationLow:  255,
					},
					{
						CategoryName:      "Very unhealthy",
						AQIHigh:           300,
						AQILow:            201,
						ConcentrationHigh: 424,
						ConcentrationLow:  355,
					},
					{
						CategoryName:      "Hazardous",
						AQIHigh:           400,
						AQILow:            301,
						ConcentrationHigh: 504,
						ConcentrationLow:  425,
					},
				},
			},
			"co": {
				Name: "co",
				Breakpoints: []BreakPoint{
					{
						CategoryName:      "Good",
						AQIHigh:           50,
						AQILow:            0,
						ConcentrationHigh: 4.4,
						ConcentrationLow:  0,
					},
					{
						CategoryName:      "Moderate",
						AQIHigh:           100,
						AQILow:            51,
						ConcentrationHigh: 9.4,
						ConcentrationLow:  4.5,
					},
					{
						CategoryName:      "Unhealthy for Sensitive Groups",
						AQIHigh:           150,
						AQILow:            101,
						ConcentrationHigh: 12.4,
						ConcentrationLow:  9.5,
					},
					{
						CategoryName:      "Unhealthy",
						AQIHigh:           200,
						AQILow:            151,
						ConcentrationHigh: 15.4,
						ConcentrationLow:  12.5,
					},
					{
						CategoryName:      "Very unhealthy",
						AQIHigh:           300,
						AQILow:            201,
						ConcentrationHigh: 30.4,
						ConcentrationLow:  15.4,
					},
					{
						CategoryName:      "Hazardous",
						AQIHigh:           400,
						AQILow:            301,
						ConcentrationHigh: 40.4,
						ConcentrationLow:  30.5,
					},
				},
			},
			"so2": {
				Name: "so2",
				Breakpoints: []BreakPoint{
					{
						CategoryName:      "Good",
						AQIHigh:           50,
						AQILow:            0,
						ConcentrationHigh: 35,
						ConcentrationLow:  0,
					},
					{
						CategoryName:      "Moderate",
						AQIHigh:           100,
						AQILow:            51,
						ConcentrationHigh: 75,
						ConcentrationLow:  36,
					},
					{
						CategoryName:      "Unhealthy for Sensitive Groups",
						AQIHigh:           150,
						AQILow:            101,
						ConcentrationHigh: 185,
						ConcentrationLow:  76,
					},
					{
						CategoryName:      "Unhealthy",
						AQIHigh:           200,
						AQILow:            151,
						ConcentrationHigh: 304,
						ConcentrationLow:  186,
					},
					{
						CategoryName:      "Very unhealthy",
						AQIHigh:           300,
						AQILow:            201,
						ConcentrationHigh: 604,
						ConcentrationLow:  305,
					},
					{
						CategoryName:      "Hazardous",
						AQIHigh:           400,
						AQILow:            301,
						ConcentrationHigh: 804,
						ConcentrationLow:  605,
					},
				},
			},
			"no2": {
				Name: "no2",
				Breakpoints: []BreakPoint{
					{
						CategoryName:      "Good",
						AQIHigh:           50,
						AQILow:            0,
						ConcentrationHigh: 53,
						ConcentrationLow:  0,
					},
					{
						CategoryName:      "Moderate",
						AQIHigh:           100,
						AQILow:            51,
						ConcentrationHigh: 100,
						ConcentrationLow:  54,
					},
					{
						CategoryName:      "Unhealthy for Sensitive Groups",
						AQIHigh:           150,
						AQILow:            101,
						ConcentrationHigh: 360,
						ConcentrationLow:  101,
					},
					{
						CategoryName:      "Unhealthy",
						AQIHigh:           200,
						AQILow:            151,
						ConcentrationHigh: 649,
						ConcentrationLow:  361,
					},
					{
						CategoryName:      "Very unhealthy",
						AQIHigh:           300,
						AQILow:            201,
						ConcentrationHigh: 1249,
						ConcentrationLow:  650,
					},
					{
						CategoryName:      "Hazardous",
						AQIHigh:           400,
						AQILow:            301,
						ConcentrationHigh: 1649,
						ConcentrationLow:  1250,
					},
				},
			},
		},
	}
}
