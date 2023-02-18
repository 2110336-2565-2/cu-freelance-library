package gosdk

import "github.com/pemistahl/lingua-go"

type UtilLang interface {
	CheckLang(string) lingua.Language
}

type langUtil struct {
	detector lingua.LanguageDetector
}

func NewUtilLang(languages []lingua.Language) UtilLang {
	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	return &langUtil{
		detector: detector,
	}
}

func (u *langUtil) CheckLang(text string) lingua.Language {
	if language, exists := u.detector.DetectLanguageOf(text); exists {
		return language
	}

	return lingua.Unknown
}
