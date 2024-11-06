// Weather shows and set weather.
package weather

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

const backend = "https://player1-backend-438714028168.us-central1.run.app/"

var tmpl = template.Must(template.New("index").Parse(`<!doctype html>
<meta charset="utf-8" />
<title>Weather Control</title>
<body>
  <h1>Weather Control App</h1>
  <hr />
  <p>WeatherType: <span>{{ .Type}}</span></p>
  <p>Intensity: <span>{{ .Intensity}}</span></p>
  <button>Click to update with the current weather</button>
  <hr />
  <form action="/" method="post">
  <p>
    Weather Type:
    <input type="text" name="type" />
  </p>
  <p>
    Weather Intensity:
    <input type="text" name="intensity" />
  </p>
  <p>
    <button>Set the current weather</button>
  </p>
  </form>
</body>`))

func handleWeather(w http.ResponseWriter, r *http.Request) {
	cur, err := getWeather()
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get weather: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, cur)
}

type weather struct {
	Type      string
	Intensity int
}

func getWeather() (weather, error) {
	resp, err := http.Get(backend)
	if err != nil {
		return weather{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather{}, err
	}
	var w weather
	if err := json.Unmarshal(b, &w); err != nil {
		return weather{}, err
	}

	return w, nil
}

func setWeather(w weather) error {
	return nil
}

func init() {
	functions.HTTP("FRONTEND", handleWeather)
}
