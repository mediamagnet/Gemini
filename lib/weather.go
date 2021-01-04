package lib

import (
	"time"
)

type Weather struct {
	CurrentCondition []struct {
		FeelsLikeC       string `json:"FeelsLikeC"`
		FeelsLikeF       string `json:"FeelsLikeF"`
		Cloudcover       string `json:"cloudcover"`
		Humidity         string `json:"humidity"`
		LocalObsDateTime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		PrecipMM         string `json:"precipMM"`
		Pressure         string `json:"pressure"`
		TempC            string `json:"temp_C"`
		TempF            string `json:"temp_F"`
		UvIndex          string `json:"uvIndex"`
		Visibility       string `json:"visibility"`
		WeatherCode      string `json:"weatherCode"`
		WeatherDesc      []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WeatherIconURL []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		Winddir16Point string `json:"winddir16Point"`
		WinddirDegree  string `json:"winddirDegree"`
		WindspeedKmph  string `json:"windspeedKmph"`
		WindspeedMiles string `json:"windspeedMiles"`
	} `json:"current_condition"`
	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Population string `json:"population"`
		Region     []struct {
			Value string `json:"value"`
		} `json:"region"`
		WeatherURL []struct {
			Value string `json:"value"`
		} `json:"weatherUrl"`
	} `json:"nearest_area"`
	Request []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
		} `json:"astronomy"`
		Date   string `json:"date"`
		Hourly []struct {
			DewPointC        string `json:"DewPointC"`
			DewPointF        string `json:"DewPointF"`
			FeelsLikeC       string `json:"FeelsLikeC"`
			FeelsLikeF       string `json:"FeelsLikeF"`
			HeatIndexC       string `json:"HeatIndexC"`
			HeatIndexF       string `json:"HeatIndexF"`
			WindChillC       string `json:"WindChillC"`
			WindChillF       string `json:"WindChillF"`
			WindGustKmph     string `json:"WindGustKmph"`
			WindGustMiles    string `json:"WindGustMiles"`
			Chanceoffog      string `json:"chanceoffog"`
			Chanceoffrost    string `json:"chanceoffrost"`
			Chanceofhightemp string `json:"chanceofhightemp"`
			Chanceofovercast string `json:"chanceofovercast"`
			Chanceofrain     string `json:"chanceofrain"`
			Chanceofremdry   string `json:"chanceofremdry"`
			Chanceofsnow     string `json:"chanceofsnow"`
			Chanceofsunshine string `json:"chanceofsunshine"`
			Chanceofthunder  string `json:"chanceofthunder"`
			Chanceofwindy    string `json:"chanceofwindy"`
			Cloudcover       string `json:"cloudcover"`
			Humidity         string `json:"humidity"`
			PrecipMM         string `json:"precipMM"`
			Pressure         string `json:"pressure"`
			TempC            string `json:"tempC"`
			TempF            string `json:"tempF"`
			Time             string `json:"time"`
			UvIndex          string `json:"uvIndex"`
			Visibility       string `json:"visibility"`
			WeatherCode      string `json:"weatherCode"`
			WeatherDesc      []struct {
				Value string `json:"value"`
			} `json:"weatherDesc"`
			WeatherIconURL []struct {
				Value string `json:"value"`
			} `json:"weatherIconUrl"`
			Winddir16Point string `json:"winddir16Point"`
			WinddirDegree  string `json:"winddirDegree"`
			WindspeedKmph  string `json:"windspeedKmph"`
			WindspeedMiles string `json:"windspeedMiles"`
		} `json:"hourly"`
		MaxtempC    string `json:"maxtempC"`
		MaxtempF    string `json:"maxtempF"`
		MintempC    string `json:"mintempC"`
		MintempF    string `json:"mintempF"`
		SunHour     string `json:"sunHour"`
		TotalSnowCm string `json:"totalSnow_cm"`
		UvIndex     string `json:"uvIndex"`
	} `json:"weather"`
}

type Metar struct {
	Meta                  Meta          `json:"meta"`
	Altimeter             Altimeter     `json:"altimeter"`
	Clouds                []interface{} `json:"clouds"`
	FlightRules           string        `json:"flight_rules"`
	Other                 []interface{} `json:"other"`
	Sanitized             string        `json:"sanitized"`
	Visibility            Visibility    `json:"visibility"`
	WindDirection         WindDirection `json:"wind_direction"`
	WindGust              WindGust      `json:"wind_gust"`
	WindSpeed             WindSpeed     `json:"wind_speed"`
	WxCodes               []interface{} `json:"wx_codes"`
	Raw                   string        `json:"raw"`
	Station               string        `json:"station"`
	Time                  Time          `json:"time"`
	Remarks               string        `json:"remarks"`
	Dewpoint              Dewpoint      `json:"dewpoint"`
	RemarksInfo           RemarksInfo   `json:"remarks_info"`
	RunwayVisibility      []interface{} `json:"runway_visibility"`
	Temperature           Temperature   `json:"temperature"`
	WindVariableDirection []interface{} `json:"wind_variable_direction"`
	Units                 Units         `json:"units"`
}
type Meta struct {
	Timestamp       time.Time `json:"timestamp"`
	StationsUpdated string    `json:"stations_updated"`
}
type Altimeter struct {
	Repr   string  `json:"repr"`
	Value  float64 `json:"value"`
	Spoken string  `json:"spoken"`
}
type Visibility struct {
	Repr   string `json:"repr"`
	Value  int    `json:"value"`
	Spoken string `json:"spoken"`
}
type WindDirection struct {
	Repr   string `json:"repr"`
	Value  int    `json:"value"`
	Spoken string `json:"spoken"`
}
type WindGust struct {
	Repr   string `json:"repr"`
	Value  int    `json:"value"`
	Spoken string `json:"spoken"`
}
type WindSpeed struct {
	Repr   string `json:"repr"`
	Value  int    `json:"value"`
	Spoken string `json:"spoken"`
}
type Time struct {
	Repr string    `json:"repr"`
	Dt   time.Time `json:"dt"`
}
type Dewpoint struct {
	Repr   string `json:"repr"`
	Value  int    `json:"value"`
	Spoken string `json:"spoken"`
}
type DewpointDecimal struct {
	Repr   string  `json:"repr"`
	Value  float64 `json:"value"`
	Spoken string  `json:"spoken"`
}
type TemperatureDecimal struct {
	Repr   string  `json:"repr"`
	Value  float64 `json:"value"`
	Spoken string  `json:"spoken"`
}
type RemarksInfo struct {
	DewpointDecimal    DewpointDecimal    `json:"dewpoint_decimal"`
	TemperatureDecimal TemperatureDecimal `json:"temperature_decimal"`
}
type Temperature struct {
	Repr   string `json:"repr"`
	Value  int    `json:"value"`
	Spoken string `json:"spoken"`
}
type Units struct {
	Altimeter   string `json:"altimeter"`
	Altitude    string `json:"altitude"`
	Temperature string `json:"temperature"`
	Visibility  string `json:"visibility"`
	WindSpeed   string `json:"wind_speed"`
}