func getHint(secret string, guess string) string {
    bulls := 0
	cows := 0

	secretCount := make([]int, 10)
	guessCount:= make([]int, 10) 

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretCount[secret[i]-'0']++
			guessCount[guess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		cows += min(secretCount[i], guessCount[i])
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}