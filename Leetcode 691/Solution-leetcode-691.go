package main

import "math"

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func minStickers(stickers []string, target string) int {
	stickCount := map[string]map[rune]int{}

	for _, st := range stickers {
		stickCount[st] = map[rune]int{}
		for _, c := range st {
			stickCount[st][c]++
		}
	}

	charInString := func(s string, r rune) bool {
		for _, c := range s {
			if c == r {
				return true
			}
		}
		return false
	}

	dp := map[string]int{}
	var dfs func(t string, stick map[rune]int) int
	dfs = func(t string, stick map[rune]int) int {
		if v, f := dp[t]; f {
			return v
		}

		res := 0
		if stick != nil {
			res = 1
		}
		remainT := []rune{}
		for _, c := range t {
			if cnt, f := stick[c]; f && cnt > 0 {
				stick[c]--
			} else {
				remainT = append(remainT, c)
			}
		}

		if len(remainT) > 0 {
			usedSticks := math.MaxInt32
			for st, stCount := range stickCount {
				if !charInString(st, remainT[0]) {
					continue
				}
				newStick := map[rune]int{}
				for key, value := range stCount {
					newStick[key] = value
				}

				usedSticks = min(usedSticks, dfs(string(remainT), newStick))
			}
			dp[string(remainT)] = usedSticks
			res += usedSticks
		}

		return res
	}

	if res := dfs(target, nil); res != math.MaxInt32 {
		return res
	}
	return -1
}