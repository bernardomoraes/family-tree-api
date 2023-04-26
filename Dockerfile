FROM golang:1.19-alpine

WORKDIR /app

ENV PATH="/go/bin:${PATH}"

COPY . .

RUN go mod tidy && go install github.com/cosmtrek/air@latest

CMD ["air"]
