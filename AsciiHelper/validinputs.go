package ascii

func AreStringValid(runes []rune) bool {
	for _, char := range runes {
		if (char < 32 || char > 126) && char != '\n'  {
			return false
		}
	}
	return true
}
