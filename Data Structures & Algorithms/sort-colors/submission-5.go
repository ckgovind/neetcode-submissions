
type TrackerValues struct {
	inpArray  []int
	indxGreen int
	indxBlue  int
	indxRed   int
}

func (t *TrackerValues) seeRed(currIndx int) {
	t.indxRed++
}

func (t *TrackerValues) seeBlue(currIndx int) {
	fmt.Printf("Swapping indexes %v and %v\n", t.indxBlue+1, currIndx)
	t.inpArray[t.indxBlue+1], t.inpArray[currIndx] = t.inpArray[currIndx], t.inpArray[t.indxBlue+1]
	t.indxBlue++
	t.seeRed(currIndx)
}

func (t *TrackerValues) seeGreen(currIndx int) {
	fmt.Printf("Swapping indexes %v and %v\n", t.indxGreen+1, currIndx)
	t.inpArray[t.indxGreen+1], t.inpArray[currIndx] = t.inpArray[currIndx], t.inpArray[t.indxGreen+1]
	t.indxGreen++

	if t.indxGreen <= t.indxBlue {
		t.seeBlue(currIndx)
	} else {
		t.indxBlue++
		t.seeRed(currIndx)
	}
}

func sortColors(colors []int) []int {
	t := TrackerValues{inpArray: colors, indxGreen: -1, indxBlue: -1, indxRed: -1}
	for indx, color := range colors {
		fmt.Println(t.inpArray)
		fmt.Printf("Current value at indx %v is %v\n", indx, color)

		switch color {
		case 0:
			t.seeGreen(indx)
		case 1:
			t.seeBlue(indx)
		case 2:
			t.seeRed(indx)
		}

		fmt.Printf("Values are indx0 : %v indx1: %v indx2: %v\n", t.indxGreen, t.indxBlue, t.indxRed)

	}
	fmt.Println(t.inpArray)

	return t.inpArray
}
