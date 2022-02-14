package field

import "gb_golang/internal/tools/stringWorker"

type Field struct {
	desk [3][3]string
}

func New() Field {
	var f Field
	return f
}

func (f *Field) checkField(x int, y int) bool {
	if f.desk[x][y] != "" {
		return false
	}
	return true
}

func (f *Field) SetCoordinate(goSide string, coordinate string) bool {
	var err error
	x_coordinate, y_coordinate, err := stringWorker.GetCoordinate(coordinate)
	if err != nil {
		panic(err)
	}
	if !f.checkField(x_coordinate, y_coordinate) {
		return false
	}
	f.desk[x_coordinate][y_coordinate] = goSide
	return true
}
