package main

func main() {

}

type AuthenticationManager struct {
	defaultTimeToLive int
	tokenToTime       map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		defaultTimeToLive: timeToLive,
		tokenToTime:       make(map[string]int),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokenToTime[tokenId] = currentTime + this.defaultTimeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if value, ok := this.tokenToTime[tokenId]; ok {
		if value > currentTime {
			this.tokenToTime[tokenId] = currentTime + this.defaultTimeToLive
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var result int
	for _, value := range this.tokenToTime {
		if value > currentTime {
			result++
		}
	}
	return result
}
