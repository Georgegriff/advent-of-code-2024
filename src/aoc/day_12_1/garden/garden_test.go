package garden

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestLoadGarden(t *testing.T) {

	garden := LoadGarden("../simple.txt")
	fmt.Println(garden)
	fmt.Print("\n")
	fmt.Println(garden.regions)

	region := garden.regions[0]
	testutils.ExpectToMatchString(t, region.RegionType, "A")

	region = garden.regions[1]
	testutils.ExpectToMatchString(t, region.RegionType, "B")

	region = garden.regions[2]
	testutils.ExpectToMatchString(t, region.RegionType, "C")

	region = garden.regions[3]
	testutils.ExpectToMatchString(t, region.RegionType, "D")

	region = garden.regions[4]
	testutils.ExpectToMatchString(t, region.RegionType, "E")
}
