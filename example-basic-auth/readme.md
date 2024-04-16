
## 参考文档 https://uncledou.site/2023/moe-extend-envoy-using-golang-3/

## 基本操作
### 安装k8s
kind create cluster
### Istio
#### 下载最新版的 istioctl
$ export ISTIO_VERSION=1.18.0-alpha.0
$ curl -L https://istio.io/downloadIstio | sh -

#### 将 istioctl 加入 PATH
$ cd istio-$ISTIO_VERSION/
$ export PATH=$PATH:$(pwd)/bin

#### 安装，包括 istiod 和 ingressgateway
$ istioctl install

