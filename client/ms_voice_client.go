package client

import (
	"fmt"
	pb "github.com/dannyhankk/cognitive/proto"
	"strings"
)

const (
	ssmlHeaderStart = "speak xmlns=\"http://www.w3.org/2001/10/synthesis\" " +
		"xmlns:mstts=\"http://www.w3.org/2001/mstts\" xmlns:emo=\"http://www.w3.org/2009/10/emotionml\" " +
		"version=\"1.0\" xml:lang=\""
	ssmlHeaderEnd = "\">"
	ssmlEnd       = "</speak>"

	ssmlVoiceHeaderStart = "<voice name=\""
	ssmlVoiceHeaderEnd   = "\">"
	ssmlVoiceEnd         = "</voice>"
)

type MSText2SpeechOpt struct {
	output string
}

func transText2SSML(text string, lang pb.LangType) (string, error) {
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
func getVoiceName(lang pb.LangType, isMale bool) string {
	switch lang {
	case pb.LangType_ENGLISH:
		if isMale {
			return "en-US-JasonNeural"
		}
		return "en-US-JennyNeural"
	case pb.LangType_RUSSIA:
		if isMale {
			return "ru-RU-DmitryNeural"
		}
		return "ru-RU-DariyaNeural"
	case pb.LangType_FRENCH:
		if isMale {
			return "fr-FR-ClaudeNeural"
		}
		return "fr-FR-DeniseNeural"
	case pb.LangType_SPANISH:
		if isMale {
			return "es-ES-DarioNeural"
		}
		return "es-ES-ElviraNeural"
	case pb.LangType_PORTUGAL:
		if isMale {
			return "pt-PT-DuarteNeural"
		}
		return "pt-PT-FernandaNeural"
	case pb.LangType_JAPANESE:
		if isMale {
			return "ja-JP-KeitaNeural"
		}
		return "ja-JP-NanamiNeural"
	case pb.LangType_KOREAN:
		if isMale {
			return "ko-KR-InJoonNeural"
		}
		return "ko-KR-SoonBokNeural"
	case pb.LangType_ARABIC:
		if isMale {
			return "ar-SA-HamedNeural"
		}
		return "ar-SA-ZariyahNeural"
	}
	return "en-US"
}

func getLang(lang pb.LangType) string {
	switch lang {
	case pb.LangType_ENGLISH:
		return "en-US"
	case pb.LangType_RUSSIA:
		return "ru-RU"
	case pb.LangType_FRENCH:
		return "fr-FR"
	case pb.LangType_SPANISH:
		return "es-ES"
	case pb.LangType_PORTUGAL:
		return "pt-PT"
	case pb.LangType_JAPANESE:
		return "ja-JP"
	case pb.LangType_KOREAN:
		return "ko-KR"
	case pb.LangType_ARABIC:
		return "ar-SA"
	}
	return "en-US"
}
