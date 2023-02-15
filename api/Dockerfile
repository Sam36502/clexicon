FROM alpine:latest

RUN mkdir /data
COPY clex_api /data/

COPY install_ob.sh /data/
RUN /data/install_ob.sh

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /data
CMD ["./clex_api"]
