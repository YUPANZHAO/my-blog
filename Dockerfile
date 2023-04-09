FROM node:18

RUN sed -i s/archive.ubuntu.com/mirrors.aliyun.com/g /etc/apt/sources.list \
    && apt update -y \
    && apt install -y git

RUN cd ~ \ 
    && wget https://golang.google.cn/dl/go1.19.8.linux-amd64.tar.gz \
    && tar xfz go1.19.8.linux-amd64.tar.gz -C /usr/local \
    && echo 'export GOROOT=/usr/local/go' >> /etc/profile \
    && echo 'export GOPATH=$HOME/gowork' >> /etc/profile \
    && echo 'export GOBIN=$GOPATH/bin' >> /etc/profile \
    && echo 'export PATH=$GOPATH:$GOBIN:$GOROOT/bin:$PATH' >> /etc/profile \
    && echo 'source /etc/profile' >> ~/.bashrc \
    && /bin/bash -c "source /etc/profile" \
    && /usr/local/go/bin/go env -w GOPROXY="https://goproxy.cn" \
    && /usr/local/go/bin/go env -w GO111MODULE=on

RUN cd ~ \
    && git clone https://github.com/YUPANZHAO/my-blog.git \
    && cd my-blog \
    && npm install \
    && cd server