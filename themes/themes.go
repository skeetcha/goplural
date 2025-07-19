package themes

import "image/color"

var (
	ThemeCosmo     Theme
	ThemeFlatly    Theme
	ThemeLitera    Theme
	ThemeMinty     Theme
	ThemeLumen     Theme
	ThemeSandstone Theme
	ThemeYeti      Theme
	ThemePulse     Theme
	ThemeUnited    Theme
	ThemeMorph     Theme
	ThemeJournal   Theme
	ThemeDarkly    Theme
	ThemeSuperhero Theme
	ThemeSolar     Theme
	ThemeCyborg    Theme
	ThemeVapor     Theme
	ThemeSimplex   Theme
	ThemeCerculean Theme
)

func SetupThemes() {
	ThemeCosmo = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeFlatly = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeLitera = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeMinty = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeLumen = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeSandstone = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeYeti = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemePulse = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeUnited = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeMorph = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeJournal = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeDarkly = Theme{
		ThemeType: Type_Dark,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeSuperhero = Theme{
		ThemeType: Type_Dark,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeSolar = Theme{
		ThemeType: Type_Dark,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeCyborg = Theme{
		ThemeType: Type_Dark,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeVapor = Theme{
		ThemeType: Type_Dark,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeSimplex = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
	ThemeCerculean = Theme{
		ThemeType: Type_Light,
		Colors: ThemeColors{
			Primary:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Secondary: color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Success:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Info:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Warning:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Danger:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Light:     color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Dark:      color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Bg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Fg:        color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectBg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			SelectFg:  color.RGBA{R: 248, G: 249, B: 250, A: 255},
			Border:    color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputFg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
			InputBg:   color.RGBA{R: 248, G: 249, B: 250, A: 255},
		},
	}
}
