package field

import (
	"errors"
)

var ValidateCoordinate error = errors.New("bad coordinates")
var FieldIsNotEmpty error = errors.New("field is not empty")

type Field struct {
	moves uint8
	desk  [3][3]string
}

func New() Field {
	var f Field
	return f
}

func (f *Field) CheckValidateCoordinate(x_coordinate uint8, y_coordinate uint8) error {
	if x_coordinate <= 2 && y_coordinate <= 2 {
		return nil
	}
	return ValidateCoordinate
}

func (f *Field) isFieldEmpty(x_coordinate uint8, y_coordinate uint8) error {
	if f.desk[x_coordinate][y_coordinate] == "" {
		return nil
	}
	return FieldIsNotEmpty
}

func (f *Field) IsDeskFull() bool {
	if f.moves >= 9 {
		return true
	}
	return false
}

func (f *Field) SetCoordinate(goSide string, x_coordinate uint8, y_coordinate uint8) error {
	err := f.CheckValidateCoordinate(x_coordinate, y_coordinate)
	if err != nil {
		return err
	}
	err = f.isFieldEmpty(uint8(x_coordinate), uint8(y_coordinate))
	if err != nil {
		return err
	}

	f.desk[x_coordinate][y_coordinate] = goSide
	f.moves++
	return nil
}

func (f *Field) GetField() [3][3]string {
	return f.desk
}

// func (f *Field) ResetField() { //Правильно ли? newField же останеться свободной для сборщика мусора?
// 	var newField [3][3]string
// 	f.desk = newField
// }
