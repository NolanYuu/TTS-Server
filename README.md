# TTS-Server
Websocket based on go1.17

## Build
```bash
cd docker
docker build . -t pytorch_golang:v1.0
docker run -itd -p 3000:3000 --gpus all -v {CODE_DIE}:{WORD_DIR} --name pytorch-golang pytorch_golang:v1.0
docker exec -it pytorch-golang /bin/bash
```

## Setup
### TTS-Core Setup
Build TTS-Core first in ./submodules/TTS-Core, follow steps in ./submodules/TTS-Core/README.md
```bash
cd TTS-Server
git submodule init
git submodule update
```

```bash
cd submodules/TTS-Core
mkdir build
cd build
cmake ..
make
```

### Go Server Setup
```bash
cd src
export LD_LIBRARY_PATH=../submodules/TTS-Core/build/
go build
```

## Config
Modify `config/config.yaml`

## Run
```bash
./src
```