package field

import (
	"errors"
)

type Field struct {
	moves uint8
	desk  [3][3]string
}

func New() Field {
	var f Field
	return f
}

func (f *Field) ValidateCoordinate(x_coordinate uint8, y_coordinate uint8) error {
	if x_coordinate <= 2 && y_coordinate <= 2 {
		return nil
	}
	return errors.New("Некорректные координаты! Попробуйте другие координаты")
}

func (f *Field) ifFieldEmpty(x_coordinate uint8, y_coordinate uint8) error {
	if f.desk[x_coordinate][y_coordinate] == "" {
		return nil
	}
	return errors.New("Это поле занято, попробуйте другие координаты.")
}

func (f *Field) CheckFullDesk() bool {
	if f.moves >= 9 {
		return true
	}
	return false
}

func (f *Field) SetCoordinate(goSide string, x_coordinate uint8, y_coordinate uint8) error {
	err := f.ValidateCoordinate(x_coordinate, y_coordinate)
	if err != nil {
		return err
	}
	err = f.ifFieldEmpty(uint8(x_coordinate), uint8(y_coordinate))
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
