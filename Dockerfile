FROM golang:latest
RUN mkdir -p /go/src/snippets.page-backend
WORKDIR  /go/src/snippets.page-backend
COPY .  /go/src/snippets.page-backend
RUN go get -u github.com/kardianos/govendor
RUN govendor sync
#RUN go get -u github.com/labstack/echo/...
#RUN go get -u github.com/globalsign/mgo
#RUN go get -u github.com/dgrijalva/jwt-go
#RUN go get -u github.com/globalsign/mgo/bson
#RUN go get -u github.com/go-playground/validator
#RUN go get -u github.com/labstack/echo/middleware
CMD ["go", "run", "main.go"]