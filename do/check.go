package do

import (
	"fmt"
	"unicode"
)

type Check struct {
	Input     string
	Condition bool
	Error     error
}

func NewCheck(input string) *Check {
	return &Check{Input: input, Condition: true}
}
func (c *Check) Min(want int) *Check {
	if c.Error != nil {
		return c
	}
	if len(c.Input) < want {
		c.Condition = false
		c.Error = fmt.Errorf("min wrong%v", len(c.Input))
		return c
	}
	return c
}
func (c *Check) Max(max int) *Check {
	if c.Error != nil {
		return c
	}
	if len(c.Input) > max {
		c.Condition = false
		c.Error = fmt.Errorf("max have bug%v", len(c.Input))
	}
	return c
}

type Rule func(*Check)

func Number(number int) Rule {
	 n:=0
	return func(c *Check) {
		for _, v := range c.Input {
			deal := unicode.IsDigit(v)
			if deal {
				n++
			}
		}
		 if n<number{
		c.Condition = false
		c.Error = fmt.Errorf("must have number")
		 }
	}
}
func Character(lows int, uppers int) Rule {
	return func(c *Check) {
		low := 0
		upper := 0
		for _, v := range c.Input {
			if unicode.IsLower(v) {
				low++
			} else if unicode.IsUpper(v) {
				upper++
			}

		}
		if low < lows || upper < uppers {
			c.Condition = false
			c.Error = fmt.Errorf("don't have enough want")
		}
	}
}
func (c *Check) Symbol() {

}
func (c *Check) CheckUser(min, max int) (bool, error) {
	c.Min(min).Max(max)
	if c.Error != nil || c.Condition != true {
		return false, c.Error
	}
	return true, nil
}
func (c *Check) CheckPassword(min int, max int, rules ...Rule) (bool, error) {
	c.Min(min).Max(max)
	if c.Error != nil || c.Condition != true {
		return false, c.Error
	}
	for _, rule := range rules {
		rule(c)
	}
	if c.Condition != true {
		return false, c.Error
	}
	return true, nil
}

func (c *Check) Hash(password string) string{
     return  ""
}
