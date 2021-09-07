package main

/*
#include <python3.6m/Python.h>
*/
import "C"
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"unsafe"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v2"

	ttscore "TTS-Server/src/ttscore"
)

type Config struct {
	Port      string `yaml:"port"`
	Data_path string `yaml:"data_path"`
	Html      string `yaml:"html"`
	Model     struct {
		Model_conf string `yaml:"model_conf"`
		Model_ckpt string `yaml:"model_ckpt"`
	}
	Vocoder struct {
		Vocoder_conf string `yaml:"vocoder_conf"`
		Vocoder_ckpt string `yaml:"vocoder_ckpt"`
	}
}

type TTSRequest struct {
	Text        string  `json:"text"`
	Language    string  `json:"language"`
	Speaker     string  `json:"speaker"`
	Sample_rate int     `json:"sample_rate"`
	Format      string  `json:"format"`
	Volume      float32 `json:"volume"`
	Speed       float32 `json:"speed"`
}

func (c *Config) getConf(config_path string) {
	yamlFile, err := ioutil.ReadFile(config_path)
	if err != nil {
		log.Fatal("yamlFile.Get: ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal("Unmarshal: ", err)
	}
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	mutex = sync.Mutex{}
)

func getTTSRequest(data interface{}) TTSRequest {
	var tts_request TTSRequest
	databyte, err := json.Marshal(data)
	if err != nil {
		log.Fatal("json marshal: ", err)
	}
	err = json.Unmarshal(databyte, &tts_request)
	if err != nil {
		log.Fatal("json unmarshal: ", err)
	}

	return tts_request
}

func serveHome(config_ptr *Config, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, (*config_ptr).Html)
}

func serveWs(config_ptr *Config, model_map_ptr *map[string](unsafe.Pointer), w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrader.Upgrade: ", err)
		return
	}
	defer ws.Close()

	for {
		var data interface{}
		err := ws.ReadJSON(&data)
		if err != nil {
			log.Fatal("ws.ReadMessage: ", err)
		}

		tts_request := getTTSRequest(data)

		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), tts_request.Text)
		wav_uuid := uuid.NewString()
		path := (*config_ptr).Data_path + wav_uuid + ".wav"
		fmt.Println(wav_uuid)

		mutex.Lock()
		ttscore.TTSCoreInference((*model_map_ptr)["tts_en_lj_0"], tts_request.Text, path, 22050)
		mutex.Unlock()

		fmt.Println("finish")

		err = ws.WriteMessage(websocket.TextMessage, []byte(tts_request.Text))

		if err != nil {
			log.Fatal("ws.WriteMessage: ", err)
		}
	}
}

func main() {
	var config Config
	config.getConf("../config/config.yaml")
	os.Mkdir(config.Data_path, 0777)
	addr := flag.String("addr", ":"+config.Port, "http service address")

	model_map := make(map[string](unsafe.Pointer))

	ttscore.TTSCoreInit()
	tts_en_lj_0 := ttscore.TTSCoreGetHandle(
		config.Model.Model_conf,
		config.Model.Model_ckpt,
		config.Vocoder.Vocoder_conf,
		config.Vocoder.Vocoder_ckpt,
		1,
	)

	// for further features, model name: tts_{language}_{speaker}
	model_map["tts_en_lj_0"] = tts_en_lj_0
	fmt.Println("model loaded")
	// FIXME: delete it, how to clear VRAM?
	ttscore.TTSCoreInference(model_map["tts_en_lj_0"], "test", "../data/test.wav", 22050)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveHome(&config, w, r)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(&config, &model_map, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	ttscore.TTSCoreFinalize()
}
