FROM alpine:edge

ADD .gitconfig /root

WORKDIR /root

# RUN sed -i 's/http:\/\/dl-cdn.alpinelinux.org/https:\/\/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk update && apk upgrade && \
  apk add mdocml-apropos coreutils gcc g++ libc-dev ca-certificates curl wget autoconf automake file libtool pkgconfig \
  bc tree vim git fish dialog less dos2unix diffutils unzip make cmake tzdata linux-headers go

RUN sed -i "s/bin\/ash/usr\/bin\/fish/" /etc/passwd

RUN echo "set mouse-=a" >> ~/.vimrc

RUN mkdir -p /root/.config/fish && \
  echo "set -gx GOPATH /data/gocode" >> /root/.config/fish/config.fish && \
  echo "set -gx GO111MODULE off" >> /root/.config/fish/config.fish

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

VOLUME /root/data

ENV SHELL /usr/bin/fish

CMD ["/usr/bin/fish"]
