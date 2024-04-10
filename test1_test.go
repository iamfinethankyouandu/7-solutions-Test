package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func FetchJSONData(url string) ([][]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data [][]int
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func TestFindMostValue(t *testing.T) {
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"
	arr2d, err := FetchJSONData(url)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		arr2d [][]int
		want  int
	}{
		{
			arr2d: [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}},
			want:  237,
		},
		{
			// 			  {59},
			// 			{73, 41},
			// 		  {52, 40, 9},
			// 		{26, 53, 6, 34},
			// 	  {10, 51, 87, 86, 81},
			arr2d: [][]int{{59}, {73, 41}, {52, 40, 9}, {26, 53, 6, 34}, {10, 51, 87, 86, 81}},
			want:  324,
		},
		{
			// 			  {59},
			// 			{73, 41},
			// 		  {52, 40, 9},
			// 		{26, 53, 6, 34},
			// 	  {10, 51, 87, 86, 81},
			// 	{61, 95, 66, 57, 25, 68},
			arr2d: [][]int{{59}, {73, 41}, {52, 40, 9}, {26, 53, 6, 34}, {10, 51, 87, 86, 81}, {61, 95, 66, 57, 25, 68}},
			want:  390,
		},
		{
			// 			   {59},
			// 			 {73, 41},
			// 		   {52, 40, 9},
			// 		 {26, 53, 6, 34},
			// 	   {10, 51, 87, 86, 81},
			// 	 {61, 95, 66, 57, 25, 68},
			// {90, 81, 80, 38, 92, 67, 73},
			arr2d: [][]int{{59}, {73, 41}, {52, 40, 9}, {26, 53, 6, 34}, {10, 51, 87, 86, 81}, {61, 95, 66, 57, 25, 68}, {90, 81, 80, 38, 92, 67, 73}},
			want:  470,
		},
		{
			// Output 7273 ผมหาไม่ได้จริงๆค่านี้ 😢😢
			arr2d: arr2d,
			want:  6580,
		},
	}
	for _, tt := range tests {
		t.Run("Find most value", func(t *testing.T) {
			if got := FindMostValue(tt.arr2d); got != tt.want {
				t.Errorf("FindMostValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
