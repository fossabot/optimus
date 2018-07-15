FROM alpine:latest

RUN mkdir -p /pipelines

WORKDIR /pipelines

ADD _output/bin/pipelines /pipelines

RUN chown -R 1000 /pipelines

USER 1000

CMD ["pipelines"]
