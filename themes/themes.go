package themes

import (
	"encoding/json"
	"log"
)

func SetupThemes(defaultThemes string) (ts Themes) {
	if err := json.Unmarshal([]byte(defaultThemes), &ts); err != nil {
		log.Println("Error:", err)
		return
	}

	return
}
