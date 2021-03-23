package meeting_rooms

import "sort"

func MinMeetingRooms(intervals [][]int) int {
	meetingRooms := map[int][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	roomsNumber := 1
	for _, meeting := range intervals {
		if len(meetingRooms) == 0 {
			meetingRooms[roomsNumber] = meeting
			continue
		}

		var meetingScheduled bool
		for room, meetingSchedule := range meetingRooms {
			if meeting[0] >= meetingSchedule[1] {
				meetingRooms[room] = meeting
				meetingScheduled = true
				break
			}
		}
		if meetingScheduled {
			continue
		}

		roomsNumber++
		meetingRooms[roomsNumber] = meeting
	}

	return roomsNumber
}
