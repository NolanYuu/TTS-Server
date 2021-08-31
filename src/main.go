package main

import (
	"fmt"
	"time"

	ttscore "TTS-Server/src/ttscore"
)

func main() {
	var model = ttscore.TTSCoreInitModel(
		"/nolan/demo/TTS-Server/submodules/TTS-Core/examples/ljspeech/fastspeech2.yaml",
		"/nolan/demo/TTS-Server/submodules/TTS-Core/examples/ljspeech/fastspeech2.pth",
		"/nolan/demo/TTS-Server/submodules/TTS-Core/examples/vocoder/melgan.yaml",
		"/nolan/demo/TTS-Server/submodules/TTS-Core/examples/vocoder/melgan.pth",
		1,
	)
	var start = time.Now().UnixNano() / 1e6

	var length = ttscore.TTSCoreInference(model, "Just for test", "/nolan/inference/test.wav", 22050)
	var end = time.Now().UnixNano() / 1e6
	fmt.Println(end - start)
	fmt.Println(length)
}

// package main

// import (
// 	"os"
// 	"flag"
// 	"net/http"
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"gopkg.in/yaml.v2"
// 	"github.com/gorilla/websocket"
// )

// type Config struct {
// 	Port      string    `yaml:"port"`
// 	Data_path string `yaml:"data_path"`
// }

// func (c *Config) getConf(config_path string) {
// 	yamlFile, err := ioutil.ReadFile(config_path)
// 	if err != nil {
// 		log.Printf("yamlFile.Get err   #%v ", err)
// 	}

// 	err = yaml.Unmarshal(yamlFile, c)
// 	if err != nil {
// 		log.Fatalf("Unmarshal: %v", err)
// 	}
// }

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
// }

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "../submodules/TTS-Web/home.html")
// }

// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	conn, _ := upgrader.Upgrade(w, r, nil)

// 	for {
// 		msgType, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			return
// 		}

// 		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

// 		if err = conn.WriteMessage(msgType, msg); err != nil {
// 			return
// 		}
// 	}
// }

// func main() {
// 	var config Config
// 	config.getConf("../config/config.yaml")
// 	os.Mkdir(config.Data_path, 0777)
// 	var addr = flag.String("addr", ":" + config.Port, "http service address")

// 	http.HandleFunc("/", serveHome)
// 	http.HandleFunc("/ws", serveWs)
// 	err := http.ListenAndServe(*addr, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
