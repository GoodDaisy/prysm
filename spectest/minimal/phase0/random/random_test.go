package random

import (
	"testing"

	"github.com/prysmaticlabs/prysm/spectest/shared/phase0/sanity"
)

func TestMinimal_Phase0_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "minimal", "random/random/pyspec_tests")
}