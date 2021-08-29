package ttscore

/*
#cgo LDFLAGS: -lpython3.6m
#cgo LDFLAGS: -L${SRCDIR}/../../third/TTS-Core/build -lttscore_api
#cgo CFLAGS: -I${SRCDIR}/../../third/TTS-Core/include

#include "ttscore_api.h"

*/
import "C"

func TTSCoreInitModel(config_file string, model_file string, use_gpu int) *_Ctype_struct__object {
	var c_config_file = C.CString(config_file)
	var c_model_file = C.CString(model_file)
	var c_use_gpu = C.int(use_gpu)
	var model = C.getInstanceText2Speech(c_config_file, c_model_file, c_use_gpu)

	return model
}

func TTSCoreInference(model *_Ctype_struct__object, text string, path string, sample_rate int) float64 {
	var c_text = C.CString(text)
	var c_path = C.CString(path)
	var c_sample_rate = C.int(sample_rate)
	var length = C.inference(model, c_text, c_path, c_sample_rate)

	return float64(length)
}
