# aqicalc
AQI calculator (using US EPA standards)

[Link: US EPA Air Quality Index calculations](https://www.airnow.gov/sites/default/files/2020-05/aqi-technical-assistance-document-sept2018.pdf)

### Example
```` go
import (
   "github.com/danesparza/aqicalc/pkg/aqi"
)


//	Create our calculation helper
c := aqi.NewCalcServiceWithDefaults()

// Call Calculate with our pollution name and the concentration of pollution
// For pm 2.5 particulate pollution this is in parts per million
result, _ := c.Calculate("pm2_5", 35.9)

//	EPA says that the AQI should be 102 for this pollutant in this case
if result != 102 {
    t.Errorf("Calculate() got = %v, want %v", result, 102)
}
````