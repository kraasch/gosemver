
package semv

import (

  // this is a test.
  "testing"

  // printing and formatting.
  "fmt"

  // other imports.
  "github.com/kraasch/godiff/godiff"
)

var (
  NL = fmt.Sprintln()
)

type TestList struct {
  testName          string
  isMulti           bool
  inputArr          []string
  expectedValue     string
}

type TestSuite struct {
  testingFunction   func(in TestList) string
  tests             []TestList
}

var suites = []TestSuite{
  /*
  * Test for the function SemverToUrl().
  */
  {
    testingFunction:
    func(in TestList) (out string) {
      inputValue := in.inputArr[0]
      out = SemverToUrl(inputValue)
      return
    },
    tests:
    []TestList{
      {
        testName:      "semver_to-url_00",
        isMulti:       false,
        inputArr:      []string{"0.0"},
        expectedValue: "https://go.dev/doc/devel/release#pre.go1",
      },
      {
        testName:      "semver_to-url_01",
        isMulti:       false,
        inputArr:      []string{"1.0"},
        expectedValue: "https://go.dev/doc/devel/release#go1",
      },
      {
        testName:      "semver_to-url_02",
        isMulti:       false,
        inputArr:      []string{"1.19"},
        expectedValue: "https://go.dev/doc/devel/release#go1.19",
      },
      {
        testName:      "semver_to-url_03",
        isMulti:       false,
        inputArr:      []string{"1.20"},
        expectedValue: "https://go.dev/doc/devel/release#go1.20",
      },
      {
        testName:      "semver_to-url_04",
        isMulti:       false,
        inputArr:      []string{"1.21"},
        expectedValue: "https://go.dev/doc/devel/release#go1.21.0",
      },
      {
        testName:      "semver_to-url_05",
        isMulti:       false,
        inputArr:      []string{"1.22"},
        expectedValue: "https://go.dev/doc/devel/release#go1.22.0",
      },
      {
        testName:      "semver_to-url_06",
        isMulti:       false,
        inputArr:      []string{"1.23"},
        expectedValue: "https://go.dev/doc/devel/release#go1.23.0",
      },
      {
        testName:      "semver_to-url_07",
        isMulti:       false,
        inputArr:      []string{"1.24"},
        expectedValue: "https://go.dev/doc/devel/release",
      },
      {
        testName:      "semver_to-url_08",
        isMulti:       false,
        inputArr:      []string{"1.99"},
        expectedValue: "https://go.dev/doc/devel/release",
      },
    },
  },
  /*
  * Test for the function DateToSemver().
  */
  {
    testingFunction:
    func(in TestList) (out string) {
      inputValue := in.inputArr[0]
      out = DateToSemver(inputValue)
      return
    },
    tests:
    []TestList{
      {
        testName:      "date_to-url_direct-hit_00",
        isMulti:       false,
        inputArr:      []string{"2024-08-13"},
        expectedValue: "1.23",
      },
      {
        testName:      "date_to-url_direct-hit_01",
        isMulti:       false,
        inputArr:      []string{"2024-02-06"},
        expectedValue: "1.22",
      },
      {
        testName:      "date_to-url_direct-hit_02",
        isMulti:       false,
        inputArr:      []string{"2023-08-08"},
        expectedValue: "1.21",
      },
      {
        testName:      "date_to-url_direct-hit_03",
        isMulti:       false,
        inputArr:      []string{"2023-02-01"},
        expectedValue: "1.20",
      },
      {
        testName:      "date_to-url_direct-hit_04",
        isMulti:       false,
        inputArr:      []string{"2022-08-02"},
        expectedValue: "1.19",
      },
      {
        testName:      "date_to-url_direct-hit_05",
        isMulti:       false,
        inputArr:      []string{"2013-05-13"},
        expectedValue: "1.1",
      },
      {
        testName:      "date_to-url_direct-hit_06",
        isMulti:       false,
        inputArr:      []string{"2012-03-28"},
        expectedValue: "1.0",
      },
      {
        testName:      "date_to-url_before-hit_00",
        isMulti:       false,
        inputArr:      []string{"2024-08-12"},
        expectedValue: "1.22",
      },
      {
        testName:      "date_to-url_before-hit_01",
        isMulti:       false,
        inputArr:      []string{"2024-02-05"},
        expectedValue: "1.21",
      },
      {
        testName:      "date_to-url_before-hit_02",
        isMulti:       false,
        inputArr:      []string{"2023-08-07"},
        expectedValue: "1.20",
      },
      {
        testName:      "date_to-url_before-hit_03",
        isMulti:       false,
        inputArr:      []string{"2023-02-00"},
        expectedValue: "1.19",
      },
      {
        testName:      "date_to-url_before-hit_04",
        isMulti:       false,
        inputArr:      []string{"2022-08-01"},
        expectedValue: "1.18",
      },
      {
        testName:      "date_to-url_before-hit_05",
        isMulti:       false,
        inputArr:      []string{"2013-05-12"},
        expectedValue: "1.0",
      },
      {
        testName:      "date_to-url_before-hit_06",
        isMulti:       false,
        inputArr:      []string{"2012-03-27"},
        expectedValue: "0.0",
      },
      {
        testName:      "date_to-url_after-hit_00",
        isMulti:       false,
        inputArr:      []string{"2024-08-14"},
        expectedValue: "1.23",
      },
      {
        testName:      "date_to-url_after-hit_01",
        isMulti:       false,
        inputArr:      []string{"2024-02-07"},
        expectedValue: "1.22",
      },
      {
        testName:      "date_to-url_after-hit_02",
        isMulti:       false,
        inputArr:      []string{"2023-08-09"},
        expectedValue: "1.21",
      },
      {
        testName:      "date_to-url_after-hit_03",
        isMulti:       false,
        inputArr:      []string{"2023-02-02"},
        expectedValue: "1.20",
      },
      {
        testName:      "date_to-url_after-hit_04",
        isMulti:       false,
        inputArr:      []string{"2022-08-03"},
        expectedValue: "1.19",
      },
      {
        testName:      "date_to-url_after-hit_05",
        isMulti:       false,
        inputArr:      []string{"2013-05-14"},
        expectedValue: "1.1",
      },
      {
        testName:      "date_to-url_after-hit_06",
        isMulti:       false,
        inputArr:      []string{"2012-03-29"},
        expectedValue: "1.0",
      },
    },
  },
}

func TestAll(t *testing.T) {
  for _, suite := range suites {
    for _, test := range suite.tests {
      name := test.testName
      t.Run(name, func(t *testing.T) {
        exp := test.expectedValue
        got := suite.testingFunction(test)
        if exp != got {
          if test.isMulti {
            t.Errorf("In '%s':\n", name)
            diff := godiff.CDiff(exp, got)
            t.Errorf("\nExp: '%#v'\nGot: '%#v'\n", exp, got)
            t.Errorf("exp/got:\n%s\n", diff)
          } else {
            t.Errorf("In '%s':\n  Exp: '%#v'\n  Got: '%#v'\n", name, exp, got)
          }
        }
      })
    }
  }
}

