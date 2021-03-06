FROM 2fkk/upx AS build-upx

FROM golang:1.14-alpine AS build-env

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk --no-cache add build-base git

# build
ARG BIN_NAME=bot
WORKDIR /${BIN_NAME}
ADD go.mod .
ADD go.sum .
RUN export GOPROXY=https://goproxy.cn go mod download
ADD . .
RUN make build

# upx
WORKDIR /data
COPY --from=build-upx /bin/upx /bin/upx
RUN cp /${BIN_NAME}/bin/${BIN_NAME} /data/main
RUN upx -k --best --ultra-brute /data/main

FROM alpine:3.11

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

COPY --from=build-env /data/main /home/main

CMD ["/home/main"]