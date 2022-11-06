package style

import (
	"fmt"
	"strconv"
	"strings"
)

type Hex string

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (r RGB) toRGB() (RGB, error) {
	return r, nil
}

func (h Hex) toRGB() (RGB, error) {
	return Hex2RGB(h)
}

func Hex2RGB(hex Hex) (RGB, error) {
	var rgb RGB

	values, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		Red:   uint8(values >> 16),
		Green: uint8((values >> 8) & 0xFF),
		Blue:  uint8(values & 0xFF),
	}

	return rgb, nil
}

func (r RGB) String() string {
	var ss []string
	template := "38;2;%d;%d;%d"
	ss = append(ss, fmt.Sprintf(template, r.Red, r.Green, r.Blue))

	return strings.Join(ss, ";")
}
