package responses

import "go.mongodb.org/mongo-driver/bson/primitive"

type InnerWeatherObject struct {
	Id                 primitive.ObjectID `bson:"_id"`
	WeatherRadar       int32              `json:"WeatherRadar"`
	Satellite          int32              `json:"Satellite"`
	Wind               []int32            `json:"Wind"`
	WindGusts          int32              `json:"WindGusts"`
	WindAccumulation   int32              `json:"WindAccumulation"`
	RainThunder        float64            `json:"Rain/Thunder"`
	RainAccumulation3D float64            `json:"RainAccumulation3D"`
	NewSnow            float64            `json:"NewSnow"`
	SnowDepth          []float64          `json:"SnowDepth"`
	Thunderstorms      int32              `json:"Thunderstorms"`
	Temperature        int32              `json:"Temperature"`
	DewPoint           int32              `json:"DewPoint"`
	Humidity           int32              `json:"Humidity"`
	FreezingAltitude   int32              `json:"FreezingAltitude"`
	Clouds             int32              `json:"Clouds"`
	HighClouds         int32              `json:"HighClouds"`
	MediumClouds       int32              `json:"MediumClouds"`
	LowClouds          int32              `json:"LowClouds"`
	CloudTops          int32              `json:"CloudTops"`
	CloudBase          int32              `json:"CloudBase"`
	Visibility         float64            `json:"Visibility"`
	CAPEIndex          int32              `json:"CAPEIndex"`
	Thermals           int32              `json:"Thermals"`
	SeaTemp            int32              `json:"SeaTemp"`
	NO2                float64            `json:"NO2"`
	PM25               int32              `json:"PM25"`
	Aerosol            float64            `json:"Aerosol"`
	OzoneLayer         int32              `json:"OzoneLayer"`
	SO2                float64            `json:"SO2"`
	SurfaceOzone       float64            `json:"SurfaceOzone"`
	COConcentration    int32              `json:"COConcentration"`
	DustMass           float64            `json:"DustMass"`
	FireIntensity      int32              `json:"FireIntensity"`
	Pressure           float64            `json:"Pressure"`
	ExtremeWind        int32              `json:"ExtremeWind"`
	ExtremeTemperature int32              `json:"ExtremeTemperature"`
	ExtremeRain        int32              `json:"ExtremeRain"`
	OutdoorMap         int32              `json:"OutdoorMap"`
}
