go 1.17
```bash
cd TTS-Server
export LD_LIBRARY_PATH=./submodules/TTS-Core/build/
git submodule init
git submodule update
cd src
go build

```