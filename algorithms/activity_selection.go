package algorithms

import (
	"fmt"
	"sort"
)

type Activity struct {
	start int
	end   int
}

func activitySelection(activities []Activity) []Activity {

	sort.Slice(activities, func(i, j int) bool {
		return activities[i].end < activities[j].end
	})

	selected := []Activity{}

	selected = append(selected, activities[0])
	lastEndTime := activities[0].end

	for i := 1; i < len(activities); i++ {

		if activities[i].start >= lastEndTime {
			selected = append(selected, activities[i])
			lastEndTime = activities[i].end
		}
	}

	return selected
}

func main7() {

	activities := []Activity{
		{start: 1, end: 3},
		{start: 2, end: 5},
		{start: 4, end: 7},
		{start: 1, end: 8},
		{start: 5, end: 9},
		{start: 8, end: 10},
		{start: 9, end: 11},
		{start: 11, end: 14},
	}

	selected := activitySelection(activities)

	fmt.Println("Selected activities:")
	for _, activity := range selected {
		fmt.Printf("Start: %d, End: %d\n", activity.start, activity.end)
	}
}
