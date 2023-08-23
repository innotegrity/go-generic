package generic_test

import (
	"testing"

	"go.innotegrity.dev/generic"
)

// TODO: implement testing and benchmarks

func TestSet1(t *testing.T) {
	requiredLangs := generic.NewSet("go", "javascript", "C#")
	t.Logf("required languages: %s", requiredLangs)
	knownLangs := generic.NewSet("java", "C++", "go")
	t.Logf("known languages: %s", knownLangs)
	knownLangs.Add("javascript", "python", "shell")
	t.Logf("known languages now: %s", knownLangs)

	t.Logf("matching languages: %s", requiredLangs.Intersection(knownLangs))
	t.Logf("is Python known: %t", knownLangs.Contains("python"))
}
