FROM ubuntu:latest

ADD clex_api /clex_api
ADD static /static
ADD objectboxlib/lib/libobjectbox.so /usr/local/lib/libobjectbox.so
RUN ldconfig

EXPOSE 34573
CMD ["/clex_api"]
