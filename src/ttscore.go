package main

/*
#cgo LDFLAGS: -lpython3.6m
#cgo LDFLAGS: -L${SRCDIR}/../third/TTS-Core/build -lttscore
#cgo CFLAGS: -I${SRCDIR}/../third/TTS-Core/include

#include "ttscore.h"

*/
import "C"
import "fmt"

func getInstanceText2Speech(config_file string, model_file string, use_gpu int) *_Ctype_struct__object {
	var c_config_file = C.CString(config_file)
	var c_model_file = C.CString(model_file)
	var c_use_gpu = C.int(use_gpu)
	var text2speech = C.getInstanceText2Speech(c_config_file, c_model_file, c_use_gpu)

	return text2speech
}

func inference(text2speech *_Ctype_struct__object, text string, path string, sample_rate int) float64 {
	var c_text = C.CString(text)
	var c_path = C.CString(path)
	var c_sample_rate = C.int(sample_rate)
	var length = C.inference(text2speech, c_text, c_path, c_sample_rate)

	return float64(length)
}

func main() {
	var text2speech = getInstanceText2Speech("/nolan/demo/TTS-Core/examples/ljspeech/fastspeech2.yaml", "/nolan/demo/TTS-Core/examples/ljspeech/fastspeech2.pth", 1)
	var length = inference(text2speech, "Just for test", "/nolan/inference/test.wav", 22050)
	fmt.Println(length)
}
