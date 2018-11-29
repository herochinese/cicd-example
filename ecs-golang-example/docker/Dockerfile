FROM golang:alpine as builder
# Install git
RUN apk update && apk add git

COPY ./json.go $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/


#get dependancies
#you can also use dep
RUN go get -d -v

#build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/buzz

# STEP 2 build a small image
# start from scratch
FROM scratch

# Copy our static executable
COPY --from=builder /go/bin/buzz /go/bin/buzz
EXPOSE 8080
CMD ["/go/bin/buzz"]
