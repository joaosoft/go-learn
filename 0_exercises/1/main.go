// I did some updated in the code, based on the problem description

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Segment struct {
	Id   int
	Name string
}

type Campaign struct {
	Id       int
	Name     string
	Segments []Segment
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a list of random campaigns
	campaigns := make([]Campaign, 0)

	// updated from 1000 to 10,000 campaigns
	for i := 1; i <= 10000; i++ {
		segments := make([]Segment, 0)
		// updated to 100 segments
		for j := 1; j <= 100; j++ {
			segment := Segment{
				Id:   rand.Intn(10000),
				Name: fmt.Sprintf("Segment %d", j),
			}
			segments = append(segments, segment)
		}

		campaign := Campaign{
			Id:       i,
			Name:     fmt.Sprintf("Campaign %d", i),
			Segments: segments,
		}

		campaigns = append(campaigns, campaign)
	}

	// Segment list to compare ( Random list )
	// updated to 1000 segments
	segmentsList := make([]Segment, 0)
	for i := 1; i <= 1000; i++ {
		segmentId := rand.Intn(10000)

		segment := Segment{
			Id:   segmentId,
			Name: fmt.Sprintf("Segment %d", segmentId),
		}

		segmentsList = append(segmentsList, segment)
	}

	// Get matching segments for each campaign
	for _, campaign := range campaigns {
		match := getMatchingSegments(campaign.Segments, segmentsList)
		if len(match) > 0 {
			// print match segments
			fmt.Println("Campaign ", campaign.Id, " - ", len(match), "Segment(s) matching", match)
		}
	}
}

func getMatchingSegments(campaignSegments []Segment, requestSegments []Segment) []int {
	result := make([]int, 0)

	for _, campaignSegment := range campaignSegments {
		for _, requestSegment := range requestSegments {
			if requestSegment.Id == campaignSegment.Id {
				result = append(result, requestSegment.Id)
			}
		}
	}

	return result
}
