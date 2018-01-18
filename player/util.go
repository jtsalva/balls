package player

// Convert int array to float64 array
func PosToFloat64(pos [2]int) [2]float64 {
	return [2]float64{float64(pos[0]), float64(pos[1])}
}

// Convert float64 array to int array
func VectorToInt(pos [2]float64) [2]int {
	return [2]int{int(pos[0]), int(pos[1])}
}

// Add each value from one array to corresponding value in next array
func MovePosByVector(pos [2]int, vector [2]int) [2]int {
	return [2]int{pos[0] + vector[0], pos[1] + vector[1]}
}
