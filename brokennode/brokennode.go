package brokennode

func FindBrokenNodes(brokenNodes int, reports []bool) string {
	var response string
	for index := range reports {
		if reports[(index+1)%len(reports)-1] {
			response += "W"
		} else {
			response += "B"
		}
	}
	return response
}
