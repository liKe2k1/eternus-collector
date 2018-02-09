FROM alpine:latest

COPY ./workdir/eternus-collector /usr/bin/eternus-collector
COPY ./eternus-collector.conf.sample /etc/eternus-collector/eternus-collector.conf

CMD ["/usr/bin/eternus-collector"]
