package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

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

func serveWs(config_ptr *Config, w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("conn.ReadMessage: ", err)
		}

		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		err = conn.WriteMessage(msgType, msg)

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

	var en_lj_0 = ttscore.TTSCoreInitModel(
		config.Model.Model_conf,
		config.Model.Model_ckpt,
		config.Vocoder.Vocoder_conf,
		config.Vocoder.Vocoder_ckpt,
		1,
	)

	// for further features
	ttscore.Model_map["en_lj_0"] = en_lj_0
	fmt.Println("ok")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveHome(&config, w, r)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(&config, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
