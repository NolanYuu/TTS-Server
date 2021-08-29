package main

import (
	"fmt"

	ttscore "./ttscore"
)

func main() {
	var model = ttscore.TTSCoreInitModel("/nolan/demo/TTS-Core/examples/ljspeech/fastspeech2.yaml", "/nolan/demo/TTS-Core/examples/ljspeech/fastspeech2.pth", 1)
	var length = ttscore.TTSCoreInference(model, "Just for test", "/nolan/inference/test.wav", 22050)
	fmt.Println(length)
}
