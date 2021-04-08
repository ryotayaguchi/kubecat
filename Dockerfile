FROM golang:latest

RUN useradd --create-home --uid 1000 catuser
WORKDIR /go/src
USER catuser

COPY src/kubecat/*.go ./
RUN go mod init github.com/ryotayaguchi/kubecat
RUN go build -o /go/bin
CMD ["kubecat"]
