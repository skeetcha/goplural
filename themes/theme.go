package themes

import (
	"image/color"
)

type ThemeType int

const (
	Type_Light = iota
	Type_Dark
)

type ThemeColors struct {
	Primary   color.Color
	Secondary color.Color
	Success   color.Color
	Info      color.Color
	Warning   color.Color
	Danger    color.Color
	Light     color.Color
	Dark      color.Color
	Bg        color.Color
	Fg        color.Color
	SelectBg  color.Color
	SelectFg  color.Color
	Border    color.Color
	InputFg   color.Color
	InputBg   color.Color
}

type Theme struct {
	ThemeType ThemeType
	Colors    ThemeColors
	Name      string
}
