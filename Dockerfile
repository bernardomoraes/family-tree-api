FROM golang:latest

WORKDIR /app

ENV PATH="/go/bin:${PATH}"

COPY . .

RUN go mod tidy && go install github.com/cosmtrek/air@latest

CMD ["air"]
