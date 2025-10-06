package services

var offsets = map[string]int16{
	"Ireland": 0, "Portugal": 0, "United Kingdom": 0,
	"Albania": 60, "Austria": 60, "Belgium": 60, "Bosnia": 60, "France": 60, "Germany": 60, "Italy": 60, "Netherlands": 60, "Norway": 60, "Poland": 60, "Serbia": 60, "Slovakia": 60, "Slovenia": 60, "Spain": 60, "Sweden": 60, "Switzerland": 60,
	"Croatia": 120, "Czechia": 120, "Denmark": 120, "Egypt": 120, "Hungary": 120, "South Africa": 120,
	"Belarus": 180, "Bulgaria": 180, "Estonia": 180, "Finland": 180, "Greece": 180, "Latvia": 180, "Lithuania": 180, "Romania": 180, "Russia": 180, "Turkey": 180, "Ukraine": 180,
	"Chile": -240, "Dominican Republic": -240,
	"Colombia": -300, "Peru": -300,
	"Mexico":   -360,
	"Cambodia": 420, "Indonesia": 420, "Thailand": 420, "Vietnam": 420,
	"China": 480, "Hong Kong": 480, "Malaysia": 480, "Singapore": 480, "Taiwan": 480, "Philippines": 480,
	"Japan": 540, "Korea": 540,
	"Australia":   600,
	"New Zealand": 720,
	"USA":         -480,
	"Argentina":   -180, "Brazil": -180,
	"India": 330, "Canada": -210,
}

func TimeConverter(hour, minute int, country string) (int, int, bool) {
	offset, exists := offsets[country]
	if !exists {
		return 0, 0, false
	}
	total := int16(hour*60+minute) + offset
	total = (total%1440 + 1440) % 1440
	return int(total / 60), int(total % 60), true
}
