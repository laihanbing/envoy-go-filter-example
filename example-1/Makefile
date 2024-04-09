.PHONY: build run
build:
        docker run --rm -v `pwd`:/go/src/go-filter -w /go/src/go-filter \
                -e GOPROXY=https://goproxy.cn \
                reg.docker.alibaba-inc.com/mosn/golang:1.18-mosn-thread \
                go build -v -o libgolang.so -buildmode=c-shared .

run:
        docker run --rm -v `pwd`/envoy.yaml:/etc/envoy/envoy.yaml \
                -v `pwd`/libgolang.so:/etc/envoy/libgolang.so \
                -p 10000:10000 \
                envoyproxy/envoy:contrib-dev-364d7838de614fcbe596877ec0fcb4675e038b4b \
                envoy -c /etc/envoy/envoy.yaml
