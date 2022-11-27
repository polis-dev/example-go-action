ARG USER=runner
ARG UID=1000

FROM golang:1 AS builder
RUN echo "${USER}:x:${UID}:${UID}:${USER}:/:" > /etc_passwd
RUN apt-get -qq update && apt-get -yqq install upx
WORKDIR /src
COPY . .
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build -ldflags "-s -w -extldflags '-static'" -o /bin/action ./cmd && strip /bin/action && upx -q -9 /bin/action

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder --chown=${UID}:0 /bin/action /action
ENTRYPOINT ["/action"]
USER ${USER}
