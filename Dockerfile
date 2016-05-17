FROM golang:alpine

RUN mkdir /planet/
ADD . /planet/
WORKDIR /planet/

RUN go build
CMD /planet/planet -p $PORT
