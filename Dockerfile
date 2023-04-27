FROM ubuntu:latest

USER root
ENV DEBIAN_FRONTEND=noninteractive
WORKDIR /app/
COPY . .

RUN apt update && apt install -y golang
RUN go build -o app .

CMD ./app
