// Main generates index.html and does posts.
package main

import (
	"html/template"
	"os"

)

var tmpl = template.Must(template.New("index").Parse(`<!doctype html>
<meta charset="utf-8" />
<title>Weather Control</title>
<body>
  <h1>Weather Control App</h1>
  <hr />
  <p>WeatherType: <span>proto.weathertype goes here</span></p>
  <p>Intensity: <span>proto.intensity goes here</span></p>
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
</body>`)

func main() {
	
	tmpl.Execute(os.Stdout, map[any]string{})
}
