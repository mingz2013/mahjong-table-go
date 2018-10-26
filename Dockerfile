FROM centos
MAINTAINER mingz2013
COPY mahjong-table-go /sbin/mahjong-table-go
RUN chmod 755 /sbin/mahjong-table-go
ENTRYPOINT ["/sbin/mahjong-table-go"]
