FROM golang:1.9-stretch

WORKDIR /go/src/github.com/koryaz/eskit-todo/
COPY ./services/todos/config.yaml .
COPY ./bin/todos /bin/todos
COPY ./bin/todosgw /bin/todosgw
ENTRYPOINT ["/bin/todos"]
