package entities

import (
	"fmt"
	"strconv"
)

type Form struct {
	Search   string `json:"search"`
	Email    string `json:"email"`
	Radius   int    `json:"radius"`
	MinPrice int    `json:"min_price"`
	MaxPrice int    `json:"max_price"`
}

func (f *Form) Validate() error {
	if f.Search == "" {
		return fmt.Errorf("Missing Search")
	}
	if f.Email == "" {
		return fmt.Errorf("Missing Email")
	}
	return nil
}

func (f *Form) ToList() []string {
	return []string{
		f.Email,
		f.Search,
		fmt.Sprintf("%v", f.Radius),
		fmt.Sprintf("%v", f.MinPrice),
		fmt.Sprintf("%v", f.MaxPrice),
	}
}
func (f *Form) ToForm(s []string) {
	f.Email = s[0]
	f.Search = s[1]
	f.Radius, _ = strconv.Atoi(s[2])
	f.MinPrice, _ = strconv.Atoi(s[3])
	f.MaxPrice, _ = strconv.Atoi(s[4])
}
