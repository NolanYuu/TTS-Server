package ttscore

/*
#cgo LDFLAGS: -lpython3.6m
#cgo LDFLAGS: -L${SRCDIR}/../../submodules/TTS-Core/build -lttscore_api
#cgo CFLAGS: -I${SRCDIR}/../../submodules/TTS-Core/include

#include "ttscore_api.h"

*/
import "C"
import "unsafe"

func TTSCoreInitModel(config_file string, model_file string, use_gpu int) *C.PyObject {
	var c_config_file = C.CString(config_file)
	var c_model_file = C.CString(model_file)
	var c_use_gpu = C.int(use_gpu)
	var model = C.getInstanceText2Speech(c_config_file, c_model_file, c_use_gpu)

	C.free(unsafe.Pointer(c_config_file))
	C.free(unsafe.Pointer(c_model_file))
	return model
}

func TTSCoreInference(model *C.PyObject, text string, path string, sample_rate int) float64 {
	var c_text = C.CString(text)
	var c_path = C.CString(path)
	var c_sample_rate = C.int(sample_rate)
	var length = C.inference(model, c_text, c_path, c_sample_rate)

	C.free(unsafe.Pointer(c_text))
	C.free(unsafe.Pointer(c_path))
	return float64(length)
}
