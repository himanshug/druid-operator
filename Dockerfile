# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /workspace

# Install kube-apiserver and etcd for envtest based tests
ARG K8S_VER=v1.18.14
RUN curl -L -O https://dl.k8s.io/${K8S_VER}/kubernetes-server-linux-arm64.tar.gz
RUN tar -xvzf kubernetes-server-linux-arm64.tar.gz

ARG ETCD_VER=v3.3.11
RUN curl -L -O https://github.com/etcd-io/etcd/releases/download/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz
RUN tar -xvzf etcd-${ETCD_VER}-linux-amd64.tar.gz

RUN mkdir k8sbin
RUN cp kubernetes/server/bin/kube-apiserver k8sbin/
RUN cp etcd-${ETCD_VER}-linux-amd64/etcd k8sbin/

COPY . .

ENV KUBEBUILDER_ASSETS=/workspace/k8sbin
ENV KUBEBUILDER_ATTACH_CONTROL_PLANE_OUTPUT=true

RUN ls -l ${KUBEBUILDER_ASSETS}

# Build
RUN go mod download
RUN go fmt ./...
RUN go vet ./...
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
