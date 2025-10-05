package services

var timezones = map[string]int8{
	"Ireland": 0, "Portugal": 0, "United Kingdom": 0,
	"Albania": 1, "Austria": 1, "Belgium": 1, "Bosnia": 1, "France": 1, "Germany": 1, "Italy": 1, "Netherlands": 1, "Norway": 1, "Poland": 1, "Serbia": 1, "Slovakia": 1, "Slovenia": 1, "Spain": 1, "Sweden": 1, "Switzerland": 1,
	"Croatia": 2, "Czechia": 2, "Denmark": 2, "Egypt": 2, "Hungary": 2, "South Africa": 2,
	"Belarus": 3, "Bulgaria": 3, "Estonia": 3, "Finland": 3, "Greece": 3, "Latvia": 3, "Lithuania": 3, "Romania": 3, "Russia": 3, "Turkey": 3, "Ukraine": 3,
	"Chile": -4, "Dominican Republic": -4,
	"Colombia": -5, "Peru": -5,
	"Mexico":   -6,
	"Cambodia": 7, "Indonesia": 7, "Thailand": 7, "Vietnam": 7,
	"China": 8, "Hong Kong": 8, "Malaysia": 8, "Singapore": 8, "Taiwan": 8, "Philippines": 8,
	"Japan": 9, "Korea": 9,
	"Australia":   10,
	"New Zealand": 12,
	"USA":         -8,
	"Argentina":   -3, "Brazil": -3,
}

var specialOffsets = map[string]int16{"India": 330, "Canada": -210}

func TimeConverter(hour, minute int, country string) (int, int, bool) {
	var offsetMinutes int16
	if special, exists := specialOffsets[country]; exists {
		offsetMinutes = special
	} else if offset, exists := timezones[country]; exists {
		offsetMinutes = int16(offset) * 60
	} else {
		return 0, 0, false
	}
	totalMinutes := int16(hour*60+minute) + offsetMinutes
	if totalMinutes < 0 {
		totalMinutes += 1440
	} else if totalMinutes >= 1440 {
		totalMinutes -= 1440
	}
	return int(totalMinutes / 60), int(totalMinutes % 60), true
}
