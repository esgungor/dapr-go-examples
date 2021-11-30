FROM golang:1.17-buster 
WORKDIR /dir
COPY . .
RUN go get -d -v
CMD [ "./bin/air" ]

