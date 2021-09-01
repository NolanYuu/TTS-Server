package main

/*
#include <python3.6m/Python.h>
*/
import "C"
import (
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

func serveHome(config_ptr *Config, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, (*config_ptr).Html)
}

func serveWs(config_ptr *Config, model_map_ptr *map[string](unsafe.Pointer), w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrader.Upgrade: ", err)
		return
	}
	// defer ws.Close()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal("conn.ReadMessage: ", err)
		}

		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))
		wav_uuid := uuid.NewString()
		fmt.Printf(wav_uuid + "\n")

		mutex.Lock()
		ttscore.TTSCoreInference((*model_map_ptr)["tts_en_lj_0"], string(msg), (*config_ptr).Data_path+wav_uuid+".wav", 22050)
		mutex.Unlock()

		fmt.Println("finish")

		err = ws.WriteMessage(msgType, msg)

		if err != nil {
			log.Fatal("conn.WriteMessage: ", err)
		}
	}
}

func main() {
	var config Config
	config.getConf("../config/config.yaml")
	os.Mkdir(config.Data_path, 0777)
	addr := flag.String("addr", ":"+config.Port, "http service address")

	var model_map = make(map[string](unsafe.Pointer))

	var tts_en_lj_0 = ttscore.TTSCoreGetHandle(
		config.Model.Model_conf,
		config.Model.Model_ckpt,
		config.Vocoder.Vocoder_conf,
		config.Vocoder.Vocoder_ckpt,
		0,
	)

	// for further features
	model_map["tts_en_lj_0"] = tts_en_lj_0
	fmt.Println("loaded")
	ttscore.TTSCoreInference(model_map["tts_en_lj_0"], "test", config.Data_path+"test"+".wav", 22050)
	fmt.Println("finish")

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
}
