package helpers

func DetectGeneDifferences(oldGene string, newGene string) int {
	var differences int
	if oldGene == "0" {
		return 0
	}
	// bigInt.String() removes leading zeroes so we have to recover them
	if len(oldGene) != len(newGene) {
		if len(oldGene) < len(newGene) {
			lenDiff := len(newGene) - len(oldGene)
			for i := 0; i < lenDiff; i++ {
				oldGene = "0" + oldGene
			}
		} else {
			lenDiff := len(oldGene) - len(newGene)
			for i := 0; i < lenDiff; i++ {
				newGene = "0" + newGene
			}

		}
	}

	for i := range oldGene {
		if oldGene[i] != newGene[i] {
			differences++
		}
	}

	return differences
}
