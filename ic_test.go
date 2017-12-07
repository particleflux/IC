package IC

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"regexp"
)

func TestOneShot(t *testing.T) {
	cases := []struct {
		in   string
		want float32
	}{
		{"Hello, world", 0.08888888888888889},
		{"Defend the east wall of the castle", 0.082010582010582},
		{"RIP VAN WINKLE", 0.030303030303030304},
		{"wmzfxtdhzfngfwxwnwxjevxdmzoxfkvxdmzowmkwmkfgzzexenfzpjotkebmneloz" +
			"lfjpbzkofxwvjefxfwfjpfngfwxwnwxeszyzobdhkxewzawvmkokvwzopjoklxppz" +
			"ozewvxdmzowzawvmkokvwzoxwlxppzofpojtvkzfkovxdmzoxewmkwwmzvxdmzokh" +
			"dmkgzwxfejwfxtdhbwmzkhdmkgzwfmxpwzlxwxfvjtdhzwzhbrntghzl",
			0.06667729},
		{"vptzmdrttzysubxaykkwcjmgjmgpwreqeoiivppalrujtlrzpchljftupucywvsyi" +
			"uuwufirtaxagfpaxzxjqnhbfjvqibxzpotciiaxahmevmmagyczpjxvtndyeuknul" +
			"vvpbrptygzilbkeppyetvmgpxuknulvjhzdtgrgapygzrptymevppaxygkxwlvtia" +
			"wlrdmipweqbhpqgngioirnxwhfvvawpjkglxamjewbwpvvmafnlojalh",
			0.042167332},
		{"", 0.0},
		{"ţņŗ\n", 0.0},
	}
	for _, c := range cases {
		got := OneShot(c.in)
		if got != c.want {
			t.Errorf("OneShot(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBegin(t *testing.T) {
	got := Begin()

	assert.Zero(t, got.totalLength)
	assert.IsType(t, regexp.Regexp{}, *got.cleanerRegex)
}

func TestCompute(t *testing.T) {
	ic := Begin()
	calculated := Compute(&ic)
	assert.Equal(t, float32(0.0), calculated, "Empty input IC should yield an IC of 0.0")

	Update(&ic, "Hello, world")
	calculated = Compute(&ic)
	assert.Equal(t, float32(0.08888888888888889), calculated)
}

func TestUpdate(t *testing.T) {
	ic := Begin()
	Update(&ic, "a")
	assert.Equal(t, 1, ic.characterCounts[0])
	assert.Equal(t, 0, ic.characterCounts[1])

	Update(&ic, "ab")
	assert.Equal(t, 2, ic.characterCounts[0])
	assert.Equal(t, 1, ic.characterCounts[1])

	Update(&ic, "zzz")
	assert.Equal(t, 2, ic.characterCounts[0])
	assert.Equal(t, 1, ic.characterCounts[1])
	assert.Equal(t, 3, ic.characterCounts[25])

	Update(&ic, "")
	assert.Equal(t, 2, ic.characterCounts[0])
	assert.Equal(t, 1, ic.characterCounts[1])
	assert.Equal(t, 3, ic.characterCounts[25])
}
