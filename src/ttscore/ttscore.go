package ttscore

/*
#cgo LDFLAGS: -lpython3.8
#cgo LDFLAGS: -L${SRCDIR}/../../submodules/TTS-Core/build -lttscore_api
#cgo CFLAGS: -I${SRCDIR}/../../submodules/TTS-Core/include
#cgo CFLAGS: -I/usr/include/python3.8

#include "ttscore_api.h"

*/
import "C"
import (
	"unsafe"
)

func TTSCoreInitialize() {
	C.initialize()
	return
}

func TTSCoreGetHandle(model_conf string, model_ckpt string, vocoder_conf string, vocoder_ckpt string, use_gpu int) unsafe.Pointer {
	var handle unsafe.Pointer

	var c_model_conf = C.CString(model_conf)
	var c_model_ckpt = C.CString(model_ckpt)
	var c_vocoder_conf = C.CString(vocoder_conf)
	var c_vocoder_ckpt = C.CString(vocoder_ckpt)
	var c_use_gpu = C.int(use_gpu)

	C.getInstanceHandle(&handle, c_model_conf, c_model_ckpt, c_vocoder_conf, c_vocoder_ckpt, c_use_gpu)

	C.free(unsafe.Pointer(c_model_conf))
	C.free(unsafe.Pointer(c_model_ckpt))
	C.free(unsafe.Pointer(c_vocoder_conf))
	C.free(unsafe.Pointer(c_vocoder_ckpt))
	return handle
}

func TTSCoreInference(handle unsafe.Pointer, text string, path string, sample_rate int) {
	var c_text = C.CString(text)
	var c_path = C.CString(path)
	var c_sample_rate = C.int(sample_rate)
	C.inference(handle, c_text, c_path, c_sample_rate)

	C.free(unsafe.Pointer(c_text))
	C.free(unsafe.Pointer(c_path))
	return
}

func TTSCoreFinalize() {
	C.finalize()
	return
}
