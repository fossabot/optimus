FROM alpine:latest

RUN mkdir /ci-pipelines

WORKDIR /ci-pipelines

ADD _output/linux/amd64/cip /ci-pipelines

CMD ["cip"]
