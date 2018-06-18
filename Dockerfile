FROM alpine:latest

RUN mkdir /lifecycle-hooks

WORKDIR /lifecycle-hooks

ADD _output/linux/amd64/lfhooks /lifecycle-hooks

CMD ["lfhooks"]
