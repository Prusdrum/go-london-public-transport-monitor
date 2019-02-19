package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetDisruptions(t *testing.T) {
	getTubes := func() []TubeStatusResponse {
		disruption := TubeStatusResponse{
			Id:       "waterloo",
			Name:     "Waterloo",
			ModeName: "tube",
			Type:     "",
			Disruptions: []TubeDisruptionResponse{
				TubeDisruptionResponse{
					Category: "disaster",
				},
			},
		}
		return []TubeStatusResponse{disruption}
	}

	disruptions := GetDisruptions(getTubes)

	assert.Equal(t, len(disruptions), 1)
	assert.Equal(t, disruptions[0].name, "Waterloo")
}
