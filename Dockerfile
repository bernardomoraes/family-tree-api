FROM golang:latest

WORKDIR /app

ENV PATH="/go/bin:${PATH}"

RUN go mod tidy && go install github.com/cosmtrek/air@latest

CMD [ "tail", "-f", "/dev/null" ]
