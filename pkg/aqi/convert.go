package aqi

import (
	"fmt"
	"math"
)

const (
	// StandardMoleVolume is the volume (litres) of a mole (gram molecular weight)
	// of a gas when the temperature is at 25°C and the pressure is at 1 atmosphere
	// (1 atm = 1.01325 bar).
	StandardMoleVolume float32 = 24.45
)

type ConvertService interface {
	ConvertPPMFromUgm3(pollutionName string, concentration float32) (float32, error)
	ConvertPPBFromUgm3(pollutionName string, concentration float32) (float32, error)
}

type convertService struct {
	Pollutants map[string]PollutantInfo
}

// NewConvertService returns a new ConvertService from the passed map of
// pollutant information.
// The map key should be a pollutant name
// The map value should be a PollutantInfo struct
func NewConvertService(pollutants map[string]PollutantInfo) ConvertService {
	return convertService{pollutants}
}

// NewConvertServiceWithDefaults returns a ConvertService that
// uses the pollutant molecular weight from the US EPA
func NewConvertServiceWithDefaults() ConvertService {
	return convertService{
		Pollutants: GetDefaultPollutantInfo(),
	}
}

// ConvertPPBFromUgm3 returns ppb concentration from ugm3 concentration information
func (c convertService) ConvertPPBFromUgm3(pollutionName string, concentration float32) (float32, error) {
	retval := float32(0)

	// First lookup the passed pollutionName in the map of
	// PollutantInfo we have
	if _, exists := c.Pollutants[pollutionName]; !exists {
		return retval, fmt.Errorf("Pollution information can't be found for %s", pollutionName)
	}

	// Next, validate that the concentration is not a negative number:
	if concentration < 0 {
		return retval, fmt.Errorf("Pollution concentration must be 0 or greater")
	}

	// Next, validate that we have a molecular weight for the requested pollutant
	// (some pollutants -- like particulate matter (example: pm10) don't have
	// a fixed molecular weight.
	// So this conversion would be impossible.
	if c.Pollutants[pollutionName].MolecularWeight == 0 {
		return retval, fmt.Errorf("Molecular weight information can't be found for %s", pollutionName)
	}

	//	Concentration (ppb) = 24.45 x concentration (µg/m3) ÷ molecular weight
	ppb := (StandardMoleVolume * concentration) / c.Pollutants[pollutionName].MolecularWeight

	//	Round to 3 decimal places
	retval = float32(roundFloat(float64(ppb), 3))

	return retval, nil
}

// ConvertPPMFromUgm3 returns ppm concentration from ugm3 concentration information
func (c convertService) ConvertPPMFromUgm3(pollutionName string, concentration float32) (float32, error) {
	retval := float32(0)

	ppb, err := c.ConvertPPBFromUgm3(pollutionName, concentration)
	if err != nil {
		return retval, err
	}

	//  xppm = xppb / 1000
	ppm := ppb / 1000

	//	Round to 3 decimal places
	retval = float32(roundFloat(float64(ppm), 3))

	return retval, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
