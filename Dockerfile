FROM alpine:latest

RUN mkdir /ci-pipelines

WORKDIR /ci-pipelines

ADD _output/linux/amd64/cipip /ci-pipelines

CMD ["cipip"]
