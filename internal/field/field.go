package field

type Field struct {
	cross, zero [3][3]int
}

func New() Field {
	var f Field
	return f
}

func Set(f Field, isX bool, x_coordinate int, y_coordinate int) bool {
	checkField(f, x_coordinate, y_coordinate)
	return false
}

func checkField(f Field, x int, y int) bool {
	return false
}
