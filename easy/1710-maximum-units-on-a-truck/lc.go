package maximumunitsonatruck

import (
	"sort"
)

func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	maxUnits := 0

	for _, boxType := range boxTypes {
		if truckSize == 0 {
			break
		}

		if boxType[0] <= truckSize {
			maxUnits += boxType[0] * boxType[1]
			truckSize -= boxType[0]
		} else {
			maxUnits += truckSize * boxType[1]
			truckSize = 0
		}
	}

	return maxUnits
}
