package gogiven

import (
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/then"
	"testing"
	"github.com/corbym/gogiven/base"
	"github.com/corbym/gogiven/testdata"
)

type ActualExpected struct {
	actual   string
	expected int
}

func TestMyFirst(testing *testing.T) {
	var someRange = []ActualExpected{
		{actual: "", expected: 0},
		{actual: "a", expected: 1},
	}
	for _, test := range someRange {
		Given(testing, someDataSetup, func(givens *testdata.InterestingGivens) {
			givens.Givens["actual"] = test.actual
		}).
			When(someAction(test)).
			Then(func(t base.TestingT, actual *testdata.CapturedIO, givens *testdata.InterestingGivens) {
			//do assertions
			then.AssertThat(t, test.actual, has.Length(test.expected))
		})
	}
}
func someAction(data ActualExpected) base.CapturedIOGivenData {
	return func(capturedIO *testdata.CapturedIO, givens *testdata.InterestingGivens) {
		capturedIO.CapturedIO[data.actual] = data.expected
	}
}

func someDataSetup(givens *testdata.InterestingGivens) {
	givens.Givens["foofar"] = "faff"
}