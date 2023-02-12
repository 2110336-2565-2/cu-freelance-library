package gosdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
	"time"
)

type UtilTest struct {
	suite.Suite
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilTest))
}

func (t *UtilTest) TestGetStudyYearSuccess() {
	currentYear := time.Now().Year()

	testGetStudyYearSuccess(t.T(), "62xxxxxxxx", strconv.Itoa(currentYear-2019))
	testGetStudyYearSuccess(t.T(), "63xxxxxxxx", strconv.Itoa(currentYear-2020))
	testGetStudyYearSuccess(t.T(), "64xxxxxxxx", strconv.Itoa(currentYear-2021))
	testGetStudyYearSuccess(t.T(), "65xxxxxxxx", strconv.Itoa(currentYear-2022))
}

func testGetStudyYearSuccess(t *testing.T, sid string, expect string) {
	want := expect

	actual, err := CalYearFromID(sid)

	assert.Nil(t, err)
	assert.Equal(t, want, actual)
}

func (t *UtilTest) TestCalStudyYearInvalidFormat() {
	testCalStudyYearInvalidInput(t.T(), "")
	testCalStudyYearInvalidInput(t.T(), "65xxx")
	testCalStudyYearInvalidInput(t.T(), "xx24xxxxxx")
	testCalStudyYearInvalidInput(t.T(), "65xxxxxxxxxxx")
}

func (t *UtilTest) TestCalStudyYearInvalidYear() {
	testCalStudyYearInvalidInput(t.T(), "66xxxxxxxxxxx")
	testCalStudyYearInvalidInput(t.T(), "68xxxxxxxxxxx")
	testCalStudyYearInvalidInput(t.T(), "99xxxxxxxxxxx")
}

func testCalStudyYearInvalidInput(t *testing.T, sid string) {
	want := "Invalid student id"

	actual, err := CalYearFromID(sid)

	assert.Equal(t, actual, "")
	assert.Equal(t, want, err.Error())
}

func (t *UtilTest) TestIsExistedTrue() {
	e := map[string]struct{}{
		"A": {},
		"B": {},
	}

	ok := IsExisted(e, "A")

	assert.True(t.T(), ok)
}

func (t *UtilTest) TestIsExistedFalse() {
	e := map[string]struct{}{
		"A": {},
		"B": {},
	}

	ok := IsExisted(e, "C")

	assert.False(t.T(), ok)
}

func (t *UtilTest) TestMergeStringSlice() {
	var nilSlice []string

	testMergeStringSlice(t.T(), []string{"1"}, []string{"2"}, []string{"1", "2"})
	testMergeStringSlice(t.T(), []string{"1"}, []string{"2", "3"}, []string{"1", "2", "3"})
	testMergeStringSlice(t.T(), []string{}, []string{"2"}, []string{"2"})
	testMergeStringSlice(t.T(), []string{"1"}, []string{}, []string{"1"})
	testMergeStringSlice(t.T(), nilSlice, nilSlice, nilSlice)
}

func testMergeStringSlice(t *testing.T, s1 []string, s2 []string, want []string) {
	actual := MergeStringSlice(s1, s2)

	assert.Equal(t, want, actual)
}

func (t *UtilTest) TestTrimInList() {
	trimList := map[string]struct{}{
		"v1": {},
		"v2": {},
	}

	testTrimInList(t.T(), "/path/1", "/v1/path/1", "/", trimList)
	testTrimInList(t.T(), "/path/something/1", "/v2/path/something/1", "/", trimList)
	testTrimInList(t.T(), "/exclude/1", "/v1/exclude/1", "/", trimList)
	testTrimInList(t.T(), "/path/1", "/path/1", "/", trimList)
}

func testTrimInList(t *testing.T, want string, word string, sep string, trimList map[string]struct{}) {
	actual := TrimInList(word, sep, trimList)

	assert.Equal(t, want, actual)
}
