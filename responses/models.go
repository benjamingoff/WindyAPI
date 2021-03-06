package responses

// InnerWeatherObject Struct for the weather object that is returned from the database
// This isn't technically being used right now, TODO: Go back in and type this stronger.
type InnerWeatherObject struct {
	Time               float64   `json:"time"`
	WeatherRadar       float64   `json:"WeatherRadar"`
	Satellite          float64   `json:"Satellite"`
	Wind               []float64 `json:"Wind"`
	WindGusts          float64   `json:"WindGusts"`
	WindAccumulation   float64   `json:"WindAccumulation"`
	RainThunder        float64   `json:"Rain/Thunder"`
	RainAccumulation3D float64   `json:"RainAccumulation3D"`
	NewSnow            float64   `json:"NewSnow"`
	SnowDepth          []float64 `json:"SnowDepth"`
	Thunderstorms      float64   `json:"Thunderstorms"`
	Temperature        float64   `json:"Temperature"`
	DewPoint           float64   `json:"DewPoint"`
	Humidity           float64   `json:"Humidity"`
	FreezingAltitude   float64   `json:"FreezingAltitude"`
	Clouds             float64   `json:"Clouds"`
	HighClouds         float64   `json:"HighClouds"`
	MediumClouds       float64   `json:"MediumClouds"`
	LowClouds          float64   `json:"LowClouds"`
	CloudTops          float64   `json:"CloudTops"`
	CloudBase          float64   `json:"CloudBase"`
	Visibility         float64   `json:"Visibility"`
	CAPEIndex          float64   `json:"CAPEIndex"`
	Thermals           float64   `json:"Thermals"`
	SeaTemp            float64   `json:"SeaTemp"`
	NO2                float64   `json:"NO2"`
	PM25               float64   `json:"PM25"`
	Aerosol            float64   `json:"Aerosol"`
	OzoneLayer         float64   `json:"OzoneLayer"`
	SO2                float64   `json:"SO2"`
	SurfaceOzone       float64   `json:"SurfaceOzone"`
	COConcentration    float64   `json:"COConcentration"`
	DustMass           float64   `json:"DustMass"`
	FireIntensity      float64   `json:"FireIntensity"`
	Pressure           float64   `json:"Pressure"`
	ExtremeWind        float64   `json:"ExtremeWind"`
	ExtremeTemperature float64   `json:"ExtremeTemperature"`
	ExtremeRain        float64   `json:"ExtremeRain"`
	OutdoorMap         float64   `json:"OutdoorMap"`
}
