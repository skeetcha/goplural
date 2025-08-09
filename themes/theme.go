package themes

import (
	"encoding/json"
	"errors"
	"log"
)

type ThemeType int

const (
	Type_Light = iota
	Type_Dark
)

func (tt *ThemeType) MarshalJSON() ([]byte, error) {
	var out string

	switch *tt {
	case Type_Light:
		out = "light"
	case Type_Dark:
		out = "dark"
	default:
		out = ""
	}

	return json.Marshal(out)
}

func (tt *ThemeType) UnmarshalJSON(data []byte) error {
	var in string

	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}

	switch in {
	case "light":
		*tt = Type_Light
	case "dark":
		*tt = Type_Dark
	default:
		panic("This really shouldn't happen, so someone fucked up")
	}

	return nil
}

type ThemeColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var errInvalidFormat = errors.New("invalid format")

func parseHexColor(s string) (c ThemeColor, err error) {
	c.A = 0xff

	if s[0] != '#' {
		log.Println(s)
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}

		log.Println("hextobyte", b)
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 9:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		c.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		log.Println("switch", s)
		err = errInvalidFormat
	}
	return
}

func (tc *ThemeColor) UnmarshalJSON(data []byte) error {
	var in string

	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}

	color, err := parseHexColor(in)

	if err != nil {
		return err
	}

	*tc = color
	return nil
}

type ThemeColors struct {
	Primary   ThemeColor `json:"primary"`
	Secondary ThemeColor `json:"secondary"`
	Success   ThemeColor `json:"success"`
	Info      ThemeColor `json:"info"`
	Warning   ThemeColor `json:"warning"`
	Danger    ThemeColor `json:"danger"`
	Light     ThemeColor `json:"light"`
	Dark      ThemeColor `json:"dark"`
	Bg        ThemeColor `json:"bg"`
	Fg        ThemeColor `json:"fg"`
	SelectBg  ThemeColor `json:"selectbg"`
	SelectFg  ThemeColor `json:"selectfg"`
	Border    ThemeColor `json:"border"`
	InputFg   ThemeColor `json:"inputfg"`
	InputBg   ThemeColor `json:"inputbg"`
}

type Theme struct {
	ThemeType ThemeType   `json:"type"`
	Colors    ThemeColors `json:"colors"`
	Name      string      `json:"name"`
}

type Themes []Theme
