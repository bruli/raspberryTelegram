FROM ubuntu:20.04
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN mkdir /app
WORKDIR /app

RUN apt update && apt upgrade -y
RUN apt install -y git make gcc gcc-arm-linux-gnueabi
COPY --from=golang:1.19-bullseye /usr/local/go/ /usr/local/go/
RUN echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc