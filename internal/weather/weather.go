package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telegram_bot/internal/models"
	"time"
)

type GeneralResponse struct {
	Main     Main   `json:"main"`
	Wind     Wind   `json:"wind"`
	Clouds   Clouds `json:"clouds"`
	Sys      Sys    `json:"sys"`
	Timezone int    `json:"timezone"`
}

func (gr *GeneralResponse) String() string {
	sunrise := getTime(time.Unix(gr.Sys.Sunrise, 0).UTC(), gr.Timezone)
	sunset := getTime(time.Unix(gr.Sys.Sunset, 0).UTC(), gr.Timezone)

	responseString := fmt.Sprintf("Температура: %v\nКак ощущается: %v\nСкорость ветра: %v м/с\nОблачность: %v%%\nВлажность: %v\nРассвет: %v\nЗакат: %v\n",
		getTemp(gr.Main.Temp-273.15), getTemp(gr.Main.FeelsLike-273.15), gr.Wind.Speed, gr.Clouds.All, gr.Main.Humidity, sunrise, sunset)
	return responseString
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
}
type Wind struct {
	Speed float64 `json:"speed"`
}

type Clouds struct {
	All int `json:"all"`
}
type Sys struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}

func GetWeather(city string, openWeatherToken string) (*GeneralResponse, error) {
	res := &GeneralResponse{}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, openWeatherToken)
	resp, err := http.Get(url)
	if err != nil {
		return nil, models.ErrBadToken
	}

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, res)

	defer resp.Body.Close()
	return res, nil
}

func getTime(time time.Time, timezone int) string {
	return fmt.Sprintf(
		"%v:%v:%v",
		time.Hour()+(timezone/3600),
		time.Minute(),
		time.Second(),
	)
}

func getTemp(temp float64) string {
	return fmt.Sprintf("%.1f", temp)
}
