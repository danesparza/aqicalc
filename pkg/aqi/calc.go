package aqi

import (
	"fmt"
	"math"
)

type PollutantInfo struct {
	Name        string
	Breakpoints []BreakPoint
}

type BreakPoint struct {
	CategoryName      string
	AQIHigh           int
	AQILow            int
	ConcentrationHigh float32
	ConcentrationLow  float32
}

type CalcService interface {
	Calculate(pollutionType string, concentration float32) (int, error)
}

type calcService struct {
	Pollutants map[string]PollutantInfo
}

// NewCalcService returns a new CalcService from the passed map of
// pollutant information.
// The map key should be a pollutant name
// The map value should be a PollutantInfo struct
func NewCalcService(pollutants map[string]PollutantInfo) CalcService {
	return calcService{pollutants}
}

// Calculate returns the AQI for the passed pollutant name, given the passed pollution
// concentration.  If there is any error (like an unknown pollutant) an error is returned
func (c calcService) Calculate(pollutionName string, concentration float32) (int, error) {

	retval := 0

	// First lookup the passed pollutionName in the map of
	// PollutantInfo we have
	if _, exists := c.Pollutants[pollutionName]; !exists {
		return retval, fmt.Errorf("Pollution information can't be found for %s", pollutionName)
	}

	//	Next, validate that the concentration is not a negative number:
	if concentration < 0 {
		return retval, fmt.Errorf("Pollution concentration must be 0 or greater")
	}

	//	Next, let's find the breakpoint information for this pollutant
	for _, breakpoint := range c.Pollutants[pollutionName].Breakpoints {
		if breakpoint.ConcentrationLow <= concentration && concentration <= breakpoint.ConcentrationHigh {

			//	We found our breakpoint information
			//	Calculate our AQI using this breakpoint
			//	using equation 1 listed in the technical assistance document here:
			//	https://www.airnow.gov/sites/default/files/2020-05/aqi-technical-assistance-document-sept2018.pdf
			CPminusBPlo := concentration - breakpoint.ConcentrationLow
			IHIminusIlo := breakpoint.AQIHigh - breakpoint.AQILow
			BPHIminusBPlo := breakpoint.ConcentrationHigh - breakpoint.ConcentrationLow

			temp1 := float32(IHIminusIlo) / BPHIminusBPlo
			temp2 := temp1 * CPminusBPlo
			temp3 := temp2 + float32(breakpoint.AQILow)
			retval = int(math.Round(float64(temp3)))

			break // Get out
		}
	}

	//	If retval is still 0, then we need to use the highest AQI for this pollutant
	//	and return that.
	if retval == 0 {
		retval = c.Pollutants[pollutionName].Breakpoints[len(c.Pollutants[pollutionName].Breakpoints)-1].AQIHigh
	}

	return retval, nil
}
