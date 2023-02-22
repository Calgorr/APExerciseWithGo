package main

import "errors"

type Row struct {
	leftChairs, rightChair [3]Chair
	middleChair            [4]Chair
}

func NewRow() *Row {
	return new(Row)
}
func (r *Row) GetChair(c int) (bool, error) {
	if c < 4 {
		return r.rightChair[c-1].isFree, nil
	} else if c > 3 && c < 8 {
		return r.middleChair[c-4].isFree, nil
	} else if c > 7 && c < 11 {
		return r.rightChair[c-8].isFree, nil
	}
	return false, errors.New("Wrong Input")
}

func (r *Row) SetChair(c int) error {
	if c < 4 {
		r.leftChairs[c-1].SetAval(!r.leftChairs[c-1].GetAval())
		return nil
	} else if c > 3 && c < 8 {
		r.leftChairs[c-4].SetAval(!r.leftChairs[c-4].GetAval())
		return nil
	} else if c > 7 && c < 11 {
		r.leftChairs[c-8].SetAval(!r.leftChairs[c-8].GetAval())
		return nil
	}
	return errors.New("Wrong Input")
}
