package course

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func New(Name string, Price float64, Isfree bool) *course {
	if Price == 0 {
		Price = 30
	}

	return &course{
		Name:   Name,
		Price:  Price,
		IsFree: Isfree,
	}
}

func (c *course) ChangePrice(price float64) {
	c.Price = price
}

func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
