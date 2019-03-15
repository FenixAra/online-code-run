FROM ubuntu:18.04

RUN apt-get update -y && apt-get install -y golang-go && apt install -y build-essential && apt-get install -y openjdk-8-jdk && apt-get install -y python && apt-get install -y ruby-full && apt-get install -y php && apt-get install -y python3