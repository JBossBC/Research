package main

func main() {

}
func bestHand(ranks []int, suits []byte) string {
	result := map[int]string{0: "Flush", 1: "Three of a Kind", 2: "Pair", 3: "High Card"}
	var suit = suits[0]
	var equal bool = true
	for i := 1; i < len(suits); i++ {
		if suit != suits[i] {
			equal = false
			break
		}
	}
	if equal {
		return result[0]
	}
	var maxRepeat int
	var times = make(map[int]int)
	for i := 0; i < len(ranks); i++ {
		times[ranks[i]]++
		if times[ranks[i]] > maxRepeat {
			maxRepeat = times[ranks[i]]
		}
	}
	if maxRepeat >= 3 {
		return result[1]
	} else if maxRepeat == 2 {
		return result[2]
	}
	return result[3]
}
