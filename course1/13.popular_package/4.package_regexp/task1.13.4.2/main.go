package main

import (
	"fmt"
	"regexp"
)

type Ad struct {
	Title       string
	Description string
}

func main() {
	ads := []Ad{
		{
			Title:       "Куплю велосипед MeRiDa",
			Description: "Куплю велосипед meriDA в хорошем состоянии.",
		},
		{
			Title:       "Продам ВаЗ 2101",
			Description: "Продам ваз 2101 в хорошем состоянии.",
		},
		{
			Title:       "Продам БМВ",
			Description: "Продам бМв в хорошем состоянии.",
		},
		{
			Title:       "Продам macBook pro",
			Description: "Продам macBook PRO в хорошем состоянии.",
		},
	}

	ads = censorAds(ads, map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	})

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Description)
		fmt.Println()
	}
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	for keyWord, word := range censor {
		pattern := fmt.Sprintf(`(?i)((\B|\b)%s(\B|\b))`, regexp.QuoteMeta(keyWord))
		re := regexp.MustCompile(pattern)
		for adID, ad := range ads {
			ads[adID].Title = re.ReplaceAllString(ad.Title, word)
			ads[adID].Description = re.ReplaceAllString(ad.Description, word)
		}
	}
	return ads
}
