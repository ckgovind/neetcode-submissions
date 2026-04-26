type TrackerValues struct {
	inpArray  []int
	indxGreen int
	indxBlue  int
	indxRed   int
}

func (t *TrackerValues) sawRed(currIndx int) {
	t.indxRed++
}

func (t *TrackerValues) sawBlue(currIndx int) {
	t.inpArray[t.indxBlue+1], t.inpArray[currIndx] = t.inpArray[currIndx], t.inpArray[t.indxBlue+1]
	t.indxBlue++
	t.sawRed(currIndx)
}
func (t *TrackerValues) sawGreen(currIndx int) {
    t.inpArray[t.indxGreen+1], t.inpArray[currIndx] = t.inpArray[currIndx], t.inpArray[t.indxGreen+1]
    t.indxGreen++
    if t.indxGreen <= t.indxBlue {
        // A 1 was displaced to currIndx — move it into the 1s region
        t.sawBlue(currIndx)
    } else {
        // No 1s between green and blue; just advance blue to stay in sync
        t.indxBlue++
        t.sawRed(currIndx)
    }
}
func sortColors(colors []int) []int {
	t := TrackerValues{inpArray: colors, indxGreen: -1, indxBlue: -1, indxRed: -1}
	for indx, color := range colors {

		switch color {
		case 0:
			t.sawGreen(indx)
		case 1:
			t.sawBlue(indx)
		case 2:
			t.sawRed(indx)
		}

	}

	return t.inpArray
}