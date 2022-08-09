package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	geojson "github.com/paulmach/go.geojson"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	data := map[string]interface{}{
			"title": "Learning Golang Web Map",
			"name":  "Marker Coordinate",
			"modaltitle": "Marker Coordinate",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerApi(w http.ResponseWriter, r *http.Request) {
	fc := geojson.NewFeatureCollection()

	// waypoint
	fc.AddFeature(geojson.NewPointFeature([]float64{110.3648114, -7.8013631}))
	fc.AddFeature(geojson.NewPointFeature([]float64{110.6018168, -7.7049554}))
	fc.AddFeature(geojson.NewPointFeature([]float64{110.8288121, -7.5746805}))
	fc.AddFeature(geojson.NewPointFeature([]float64{110.6014252, -7.5323503}))
	fc.AddFeature(geojson.NewPointFeature([]float64{110.4997587, -7.3304062}))
	fc.AddFeature(geojson.NewPointFeature([]float64{110.2182126, -7.4772031}))
	fc.AddFeature(geojson.NewPointFeature([]float64{110.2834439, -7.5780838}))

	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(rawJSON))
}

func HandlerGallery(w http.ResponseWriter, r *http.Request) {
	message := "Gallery"
	w.Write([]byte(message))
}