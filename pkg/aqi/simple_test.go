package aqi_test

import (
	"github.com/danesparza/aqicalc/pkg/aqi"
	"testing"
)

func Test_Calculate_ReturnsAQI(t *testing.T) {
	//	Arrange
	c := aqi.NewCalcServiceWithDefaults()
	pollutionName := "pm2_5"
	pollutionConcentration := float32(35.9) // for pm 2.5 this is in parts per million
	wantAQI := 102

	//	Act - Call Calculate with our pollution name and the concentration of pollution
	result, _ := c.Calculate(pollutionName, pollutionConcentration)

	//	Assert - EPA says that the AQI should be 102 for this pollutant in this case
	if result != wantAQI {
		t.Errorf("Calculate() got = %v, want %v", result, wantAQI)
	}
}
