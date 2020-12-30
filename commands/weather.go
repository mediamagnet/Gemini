package commands

import (
	"Bolts/lib"
	"Bolts/tools"
	"context"
	"encoding/json"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/joho/godotenv"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Weather struct{ Command }

func InitWeather() Weather {
	return Weather{Init(&CommandItem{
		Name:        "weather",
		Description: "Returns the weather for a City/Postal Code or Coordinate Pair from wttr.in",
		Aliases: []string{"wttr", "w"},
		Usage:       "]weather <location>",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "Location you're wanting to view the weather for.",
				Required:    true,
			},
		},
	})}
}

func (c Weather) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {

		// Load .env files
		err := godotenv.Load()
		if err != nil {
			logrus.Fatalln("Error loading .env file")
		}

		avwxKey := os.Getenv("AVWX_KEY")

		location := strings.ReplaceAll(strings.TrimPrefix(ctx.Message.Content, "]weather ")," ", "_")
		url := fmt.Sprintf("https://wttr.in/%v?format=j1", location)
		fmt.Println(url)
		client := http.Client{
			Timeout: time.Second * 2,
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			logrus.Fatal(err)
		}

		res, getErr := client.Do(req)
		if getErr != nil {
			logrus.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			logrus.Fatal(readErr)
		}

		weather := lib.Weather{}
		jsonErr := json.Unmarshal(body, &weather)
		if jsonErr != nil {
			logrus.Fatal(jsonErr)
		}


		url2 := fmt.Sprintf("https://avwx.rest/api/metar/%v,%v", weather.NearestArea[0].Latitude, weather.NearestArea[0].Longitude)
		req2, err := http.NewRequest(http.MethodGet, url2, nil)
		if err != nil {
			logrus.Fatal(err)
		}
		req2.Header.Set("Authorization", avwxKey)
		res2, getErr := client.Do(req2)
		if getErr != nil {
			logrus.Fatal(getErr)
		}
		logrus.Infoln(res2.StatusCode)

		if res2.Body != nil {
			defer res2.Body.Close()
		}

		body2, readErr := ioutil.ReadAll(res2.Body)
		if readErr != nil {
			logrus.Fatal(readErr)
		}
		metar := lib.Metar{}
		jsonErr = json.Unmarshal(body2, &metar)
		if jsonErr != nil {
			logrus.Error(jsonErr)
		}

		//fmt.Println(weather.CurrentCondition[0].FeelsLikeC)
		fmt.Println(metar.Sanitized)

		err = atlas.Disgord.DeleteMessage(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, ctx.Message.ID)
		if err != nil {
			logrus.Warnln(err)
		}

			_, err = ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:     fmt.Sprintf("The Weather in %v", strings.ReplaceAll(location, "_", " ")),
				URL:       fmt.Sprintf("https://wttr.in/%v", location),
				Timestamp: disgord.Time{},
				Color:     0x007FFF,
				Thumbnail: &disgord.EmbedThumbnail{
					URL: tools.WeatherIcon(weather.CurrentCondition[0].WeatherCode),
				},
				Video:    nil,
				Provider: nil,
				Author: &disgord.EmbedAuthor{
					Name: "wttr.in",
				},
				Fields: []*disgord.EmbedField{
					{
						Name:   "Currently:",
						Value:  fmt.Sprintf("%v° C / %v° F", weather.CurrentCondition[0].TempC, weather.CurrentCondition[0].TempF),
						Inline: true,
					},
					{
						Name: "Feels Like:",
						Value: fmt.Sprintf("%v° C / %v° F", weather.CurrentCondition[0].FeelsLikeC, weather.CurrentCondition[0].FeelsLikeF),
						Inline: true,
					},
					{
						Name: "Humidity:",
						Value: weather.CurrentCondition[0].Humidity,
						Inline: false,
					},
					{
						Name: "Pressure:",
						Value: weather.CurrentCondition[0].Pressure,
						Inline: false,
					},
					{
						Name: "Wind Speed",
						Value: fmt.Sprintf("%v KMpH, %v MPH", weather.CurrentCondition[0].WindspeedKmph, weather.CurrentCondition[0].WindspeedMiles),
						Inline: true,
					},
					{
						Name: "Wind Direction",
						Value: fmt.Sprintf("%v° %v", weather.CurrentCondition[0].WinddirDegree, weather.CurrentCondition[0].Winddir16Point),
						Inline: true,
					},
					{
						Name: "Sunset/Sunrise",
						Value: fmt.Sprintf( "%v, %v", weather.Weather[0].Astronomy[0].Sunset, weather.Weather[0].Astronomy[0].Sunrise),
						Inline: false,
					},
				},
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("%v", metar.Sanitized),
				},
			},
		})
		if err != nil {
			logrus.Warnln(err)
		}
	}

	return c.CommandInterface
}
