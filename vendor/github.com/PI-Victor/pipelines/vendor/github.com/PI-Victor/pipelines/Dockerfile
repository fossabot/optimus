FROM alpine:latest

RUN mkdir -p /ci-pipelines

WORKDIR /ci-pipelines

ADD _output/bin/ci-pipelines /ci-pipelines

RUN chown -R 1000 /ci-pipelines

USER 1000

CMD ["cip"]
