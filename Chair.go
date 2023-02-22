package main

type Chair struct {
	isFree bool
}

func (c *Chair) SetAval(aval bool) {
	c.isFree = aval
}

func (c *Chair) GetAval() bool {
	return c.isFree
}
