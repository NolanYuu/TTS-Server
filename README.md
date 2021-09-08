# TTS-Server
Websocket based on go1.17

## Build


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