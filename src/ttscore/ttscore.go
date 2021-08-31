package ttscore

/*
#cgo LDFLAGS: -lpython3.6m
#cgo LDFLAGS: -L${SRCDIR}/../../submodules/TTS-Core/build -lttscore_api
#cgo CFLAGS: -I${SRCDIR}/../../submodules/TTS-Core/include

#include "ttscore_api.h"

*/
import "C"
import (
	"unsafe"
)

func TTSCoreInitModel(model_conf string, model_ckpt string, vocoder_conf string, vocoder_ckpt string, use_gpu int) *C.PyObject {
	var c_model_conf = C.CString(model_conf)
	var c_model_ckpt = C.CString(model_ckpt)
	var c_vocoder_conf = C.CString(vocoder_conf)
	var c_vocoder_ckpt = C.CString(vocoder_ckpt)
	var c_use_gpu = C.int(use_gpu)
	var model = C.getInstanceText2Speech(c_model_conf, c_model_ckpt, c_vocoder_conf, c_vocoder_ckpt, c_use_gpu)

	C.free(unsafe.Pointer(c_model_conf))
	C.free(unsafe.Pointer(c_model_ckpt))
	C.free(unsafe.Pointer(c_vocoder_conf))
	C.free(unsafe.Pointer(c_vocoder_ckpt))
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
