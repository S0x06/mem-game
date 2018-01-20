FROM golang:latest

CMD mkdir -p /usr/app
CMD mkdir -p /usr/app/conf
WORKDIR /usr/app
RUN apt update -y && apt install ca-certificates
COPY mem-game-server mem-game-server
COPY conf/app.conf conf/app.conf

CMD "./mem-game-server"