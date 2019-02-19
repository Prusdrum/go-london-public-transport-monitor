package main

type DisruptionLine struct {
	name string
	id   string
}

type getTubesType func() []TubeStatusResponse

func GetDisruptions(getTubes getTubesType) []DisruptionLine {
	disruptions := make([]DisruptionLine, 0)
	tubeStatuses := getTubes()

	for _, tubeStatus := range tubeStatuses {
		if len(tubeStatus.Disruptions) > 0 {
			disruption := DisruptionLine{
				id:   tubeStatus.Id,
				name: tubeStatus.Name}
			disruptions = append(disruptions, disruption)
		}
	}

	return disruptions
}
