FROM alpine:latest

RUN mkdir -p /cloudflavor

WORKDIR /cloudflavor

ADD _output/bin/pipelines /cloudflavor/

RUN chown -R 1000:1000 /cloudflavor

RUN chmod +x pipelines

USER 1000

CMD ["/cloudflavor/pipelines"]
