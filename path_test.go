package grandet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatAndCheck(t *testing.T) {

	test_cases := []string{
		"/github.com/setekhid/grandet/test.txt",
		"github.com/setekhid/grandet/test.txt",
		"github.com/./setekhid/grandet/test.txt",
		"github.com/setekhid/blabla/../grandet/test.txt",
		"/github.com/setekhid/../setekhid/grandet/test.txt",
		"github.com/setekhid/grandet/test.txt/",
		"github.com/setekhid/grandet/test.txt/.",
		"github.com/setekhid/grandet/test.txt/blabla/..",
	}

	for _, test_case := range test_cases {
		assert.EqualValues(t,
			"/github.com/setekhid/grandet/test.txt",
			pathFormatAndCheck(test_case))
	}
}
