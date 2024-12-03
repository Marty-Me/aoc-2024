package daytwo

import (
	_ "embed"
	"testing"
)

//go:embed test/day_two_example.txt
var part_1_input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part_1_input)

	if result != "2" {
		t.Errorf("day two Part 1 should solve with '2' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part_1_input)

	if result != "4" {
		t.Errorf("day two Part  should solve with '4' given input, not '%s'", result)
	}
}

func TestExample1(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "7", "6", "4", "2", "1")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestExample2(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "2", "7", "8", "9")

	if hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestExample3(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "9", "7", "6", "2", "1")

	if hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestExample4(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "3", "2", "4", "5")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestExample5(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "8", "6", "4", "4", "1")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestExample6(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "3", "6", "7", "9")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom1Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "9", "2", "3", "4", "5")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom2Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "9", "3", "4", "5")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom3Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "2", "9", "4", "5")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom4Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "2", "3", "9", "5")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom5Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "1", "2", "3", "4", "9")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom6Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "9", "4", "3", "2", "1")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom7Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "5", "9", "3", "2", "1")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom8Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "5", "4", "9", "2", "1")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom9Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "5", "4", "3", "9", "1")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom10Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "5", "4", "3", "2", "9")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom11InValid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "5", "5", "5", "5", "5")

	if hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom12Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "91", "92", "95", "93", "94")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}

func TestCustom13Valid(t *testing.T) {
	levels := make([]string, 0)
	levels = append(levels, "91", "92", "93", "94", "99")

	if !hasSafeLevelsWithDampener(&levels) {
		t.Errorf("expected to be valid")
	}
}
