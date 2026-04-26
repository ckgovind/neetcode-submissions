

type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	var answer string
	for _, str := range strs {
		val := strconv.Itoa(len(str))
		answer += "#" + val + "#"
		answer += str
	}
	return answer
}

func readNumberAt(rencoded string, indx int) (int, int) {
	for i := indx + 1; ; i++ {
		if rencoded[i] == '#' {
			val, _ := strconv.Atoi(rencoded[indx+1 : i])
			return val, i + 1
		}
	}
}

func (s *Solution) Decode(encoded string) []string {
	var answer []string

	indx := 0
	for indx < len(encoded) {
		number, startIndx := readNumberAt(encoded, indx)
		currAnwser := encoded[startIndx : startIndx+number]
		answer = append(answer, string(currAnwser))
		indx = startIndx + number
	}
	return answer
}
