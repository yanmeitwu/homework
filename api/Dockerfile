From centos:centos7
MAINTAINER Yanmei Wu

# Add GO LANG
RUN yum -y install curl build-essential ca-certificates git mercurial bzr gcc gcc-c++ make
RUN mkdir /goroot && curl https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1
RUN mkdir -p /go
RUN mkdir -p /go/bin

ENV GOROOT /goroot
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN curl https://glide.sh/get | sh
ENV API_PATH /go/src/github.com/wyiemay/convoy-api/api

WORKDIR ${API_PATH}
# ADD . .
# RUN glide install && \
#    go build main.go

CMD cd ${API_PATH} && go run ./main.go

#ENTRYPOINT ./main
