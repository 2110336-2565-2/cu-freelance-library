package gosdk

import (
	"github.com/pemistahl/lingua-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilLangTest struct {
	suite.Suite
}

func TestUtilLang(t *testing.T) {
	suite.Run(t, new(UtilTest))
}

func (t *UtilLangTest) TestCheckLangSuccess() {
	testCheckLangSuccess(t.T(), lingua.English, "hello world")
	testCheckLangSuccess(t.T(), lingua.English, "cu-freelance")
	testCheckLangSuccess(t.T(), lingua.English, "how old are you")
	testCheckLangSuccess(t.T(), lingua.Thai, "เจอกันเมื่อชาติต้องการ")
	testCheckLangSuccess(t.T(), lingua.Thai, "AI กับโลกยุคใหม่")
	testCheckLangSuccess(t.T(), lingua.Thai, "ชมเรม Thinc Web 3.0")
	testCheckLangSuccess(t.T(), lingua.Thai, "มาฝึก deploy ด้วย ArgoCD กันเถอะ")
}

func testCheckLangSuccess(t *testing.T, want lingua.Language, text string) {
	utils := NewUtilLang([]lingua.Language{lingua.English, lingua.Thai})

	actual := utils.CheckLang(text)

	assert.Equal(t, want, actual)
}

func (t *UtilLangTest) TestCheckLangUnknown() {
	testCheckLangUnknown(t.T(), "然而，第二天，在美国回复之前，赫鲁晓夫发出了另一封更加紧急的信件，加上了更多的条件")
	testCheckLangUnknown(t.T(), "尼克松说，将有专人翻译赫鲁晓夫先生的讲话内容，并且他希望他的演讲也可以同样被翻译成俄文，并且在苏联播出。")
	testCheckLangUnknown(t.T(), "赫鲁晓夫：…那么尼克松以前是做律师的？")
	testCheckLangUnknown(t.T(), "เจอกันเมื่อชาติต้องการ")
	testCheckLangUnknown(t.T(), "AI กับโลกยุคใหม่")
	testCheckLangUnknown(t.T(), "ชมเรม Thinc Web 3.0")
	testCheckLangUnknown(t.T(), "มาฝึก deploy ด้วย ArgoCD กันเถอะ")
}

func testCheckLangUnknown(t *testing.T, text string) {
	utils := NewUtilLang([]lingua.Language{lingua.English})

	actual := utils.CheckLang(text)

	assert.Equal(t, lingua.Unknown, actual)
}
