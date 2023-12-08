package functions

func GetEmoji(absenceType int) string {
	emojis := map[int]string{
		1:  "🏠",
		2:  "",
		3:  "✈",
		4:  "✈",
		5:  "🌡",
		6:  "🌡",
		7:  "",
		8:  "",
		9:  "🎓",
		10: "🏠",
		11: "☀",
		12: "☀",
		13: "☀",
	}
	return emojis[absenceType]
}
