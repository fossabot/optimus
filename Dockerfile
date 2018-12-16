FROM debian:buster-slim

RUN mkdir -p /cloudflavor

WORKDIR /cloudflavor

ADD _output/bin/optimus /cloudflavor/

RUN chown -R 1000:1000 /cloudflavor

RUN chmod +x /cloudflavor/optimus

USER 1000

CMD ["/cloudflavor/optimus"]
