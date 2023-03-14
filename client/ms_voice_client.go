package client

import (
	"fmt"
	"strings"
)

const (
	ssmlHeaderStart = "<speak version=\"1.0\" xmlns=\"http://www.w3.org/2001/10/synthesis\" xml:lang=\""
	ssmlHeaderEnd   = "\">"
	ssmlEnd         = "</speak>"

	ssmlVoiceHeaderStart = "<voice name=\""
	ssmlVoiceHeaderEnd   = "\">"
	ssmlVoiceEnd         = "</voice>"
)

type MSText2SpeechOpt struct {
	output string
}

func transText2SSML(text string, lang int32) (string, error) {
	ssml := ssmlHeaderStart + getLang(lang) + ssmlHeaderEnd
	first := strings.Split(text, "Male:")
	for _, firstItem := range first {
		partSecond := strings.Split(firstItem, "Female")
		if len(partSecond) != 2 {
			return "", fmt.Errorf("handle gen voice failed")
		}
		if partSecond[0] != "" {
			ssml += ssmlVoiceHeaderStart
			ssml += getVoiceName(lang, true)
			ssml += ssmlVoiceHeaderEnd
			ssml += partSecond[0]
			ssml += ssmlVoiceEnd
		}
		if partSecond[1] != "" {
			ssml += ssmlVoiceHeaderStart
			ssml += getVoiceName(lang, false)
			ssml += ssmlVoiceHeaderEnd
			ssml += partSecond[1]
			ssml += ssmlVoiceEnd
		}
	}
	ssml += ssmlEnd
	return ssml, nil
}

// return value: female male
func getVoiceName(lang int32, isMale bool) string {
	switch lang {
	case 1:
		return ""

	}

	return ""
}

func getLang(lang int32) string {
	return "en-us"
}
