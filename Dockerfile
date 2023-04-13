FROM golang:latest

WORKDIR /app

ENV PATH="/go/bin:${PATH}"

CMD [ "tail", "-f", "/dev/null" ]
