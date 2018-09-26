FROM golang:1.11.0-alpine3.8

COPY .gitconfig /root

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk upgrade && \
  apk add mdocml-apropos coreutils gcc g++ libc-dev ca-certificates curl wget autoconf automake file libtool pkgconfig \
  bc tree vim git fish dialog less dos2unix diffutils unzip make cmake tzdata linux-headers

RUN sed -i "s/bin\/ash/usr\/bin\/fish/" /etc/passwd

ENV SHELL /usr/bin/fish

RUN echo "set mouse-=a" >> ~/.vimrc

RUN mkdir -p /root/.config/fish && \
  echo "set -gx PATH {\$PATH} /usr/local/go/bin" >> /root/.config/fish/config.fish && \
  echo "set -gx GOPATH /root/gocode" >> /root/.config/fish/config.fish && \
  echo "set -gx GO111MODULE off" >> /root/.config/fish/config.fish

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /root

VOLUME /root/data

CMD ["/usr/bin/fish"]
