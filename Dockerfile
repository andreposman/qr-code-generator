FROM golang:1.16 as builder
WORKDIR $GOPATH/src/go.k6.io/k6
ADD . .
RUN apk --no-cache add git
RUN CGO_ENABLED=0 go install -a -trimpath -ldflags \
"-s -w -X go.k6.io/k6/lib/consts.VersionDetails=$(date -u +\
"%FT%T%z")/$(git describe --always --long --dirty)"

COPY --from=builder /go/bin/k6 /usr/bin/k6

WORKDIR /home/k6
ENTRYPOINT ["k6"]