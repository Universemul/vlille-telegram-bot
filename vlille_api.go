package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var DISTANCE_MAX = 1000
var MAX_ROWS = 5

func get_closest_bike(url string, lat float32, lng float32) ApiResult {
	url = fmt.Sprintf("%s&geofilter.distance=%f,%f,%d&rows=%d", url, lat, lng, DISTANCE_MAX, MAX_ROWS)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return ApiResult{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result ApiResult
	json.Unmarshal([]byte(body), &result)
	return result
}
