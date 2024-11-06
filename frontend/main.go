// Main generates index.html and does posts.
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const backend = "http://"

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
</body>`))

func main() {
	http.HandleFunc("/", handleIndex)
	port, _ := os.LookupEnv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
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
	resp, err := http.Get(backend + "/weather")
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

func setWeather() {

}
