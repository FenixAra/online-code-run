FROM ubuntu:18.04

RUN apt-get update -y && apt-get install -y golang-go && apt install -y build-essential && apt-get install -y openjdk-8-jdk && apt-get install -y python && apt-get install -y ruby-full && apt-get install -y python3 && apt-get install -y nodejs npm

ADD . /home/src/github.com/FenixAra/online-code-run

ENV GOPATH /home

RUN go install github.com/FenixAra/online-code-run

ENTRYPOINT /go/bin/online-code-run

EXPOSE 3000