package client

import (
	"fmt"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
	"github.com/dannyhankk/cognitive/util"
	"os"
	"time"

	pb "github.com/dannyhankk/cognitive/proto"
	"strings"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
)

const (
	ssmlHeaderStart = "<speak xmlns=\"http://www.w3.org/2001/10/synthesis\" " +
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
		for index, _ := range partSecond {
			if partSecond[index] == "" {
				continue
			}
			if index == 0 {
				ssml += ssmlVoiceHeaderStart
				ssml += getVoiceName(lang, true)
				ssml += ssmlVoiceHeaderEnd
				ssml += partSecond[0]
				ssml += ssmlVoiceEnd
				continue
			}
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

func CreateVoice(text string, lang pb.LangType, file string) {
	var err error
	err = nil
	defer func(err error) {
		if err != nil {
			util.Logger.Errorf("voice gen error, remove file")
			os.Remove(file)
		}
	}(err)
	audioConfig, err := audio.NewAudioConfigFromWavFileOutput(file)
	if err != nil {
		util.Logger.Errorf("init audio config error, %s", err.Error())
		return
	}
	defer audioConfig.Close()
	speechConfig, err := speech.NewSpeechConfigFromSubscription(
		util.RootConfig.SpeechKey, util.RootConfig.SpeechRegion)
	if err != nil {
		util.Logger.Errorf("init speech config error, %s", err.Error())
		return
	}
	defer speechConfig.Close()
	speechSynthesizer, err := speech.NewSpeechSynthesizerFromConfig(speechConfig, audioConfig)
	if err != nil {
		util.Logger.Errorf("init speech synthesizer config error, %s", err.Error())
		return
	}
	defer speechSynthesizer.Close()
	ssmlString, err := transText2SSML(text, lang)
	if err != nil {
		util.Logger.Errorf("format ssml failed, %s", err.Error())
		return
	}

	task := speechSynthesizer.SpeakSsmlAsync(ssmlString)
	var outcome speech.SpeechSynthesisOutcome
	select {
	case outcome = <-task:
	case <-time.After(60 * time.Second):
		err = fmt.Errorf("time out")
		util.Logger.Errorf("generate time out")
		return
	}
	defer outcome.Close()
	if outcome.Error != nil {
		err = outcome.Error
		util.Logger.Errorf("generate error, %s", outcome.Error.Error())
		return
	}

	if outcome.Result.Reason == common.SynthesizingAudioCompleted {
		util.Logger.Infof("Speech synthesized to speaker for text\n")
	} else {
		cancellation, _ := speech.NewCancellationDetailsFromSpeechSynthesisResult(outcome.Result)
		util.Logger.Infof("CANCELED: Reason=%d.\n", cancellation.Reason)

		if cancellation.Reason == common.Error {
			util.Logger.Infof("CANCELED: ErrorCode=%d\nCANCELED: ErrorDetails=[%s]\nCANCELED: Did you set the speech resource key and region values?\n",
				cancellation.ErrorCode,
				cancellation.ErrorDetails)
		}
	}
}
