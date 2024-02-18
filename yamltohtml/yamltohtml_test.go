package yamltohtml_test

import (
	"goModDemo/yamltohtml"
	"testing"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestYamltohtml(t *testing.T) {

	testcases := []TestCase{
		{"case 1", "testdata/test_01.yml", "<html><head><title>My Awesome Page</title></head><body>This is my awesome content</body></html>"},
		{"case 2", "testdata/test_02.yml", "<html><head><title>My Second Page</title></head><body>This is my awesome content</body></html>"},
	}

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			res, err := yamltohtml.Yamltohtml(test.path)
			if err != nil {
				t.Fail()
			}
			if res != test.expected {
				t.Fail()
			}
		})
	}

}
