FROM pytorch/pytorch:1.7.0-cuda11.0-cudnn8-runtime
MAINTAINER nolan.yu@zoom.us
WORKDIR /tmp
COPY ./requirements.txt requirements.txt
RUN apt-get update && \
    apt-get install -y gcc && \
    apt-get install -y g++ && \
    apt-get install -y wget && \
    apt-get install -y curl && \
    apt-get install -y tar && \
    apt-get install -y git && \
    apt-get install -y libsndfile1 && \
    apt-get install -y ffmpeg && \
    apt-get install -y psmisc && \
    apt-get install -y python3.8-dev && \
    apt-get install -y cmake && \
    pip install -r requirements.txt && \
    ln -s /opt/conda/bin/python3.8 /usr/bin/python && \
    ln -s /opt/conda/bin/pip /usr/bin/pip && \
    wget https://golang.org/dl/go1.17.linux-amd64.tar.gz && \
    tar xvf ./go1.17.linux-amd64.tar.gz && \
    mv ./go /usr/lib/ && \
    rm -rf ./go1.17.linux-amd64.tar.gz && \
    ln -s /usr/lib/go/bin/go /usr/bin/go && \
    ln -s /usr/lib/go/bin/gofmt /usr/bin/gofmt