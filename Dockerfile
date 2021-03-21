# Using a typical enterprise OS
FROM centos:7

RUN yum -y install deltarpm \
        && yum -y update \
        && yum clean all

RUN cd /etc/yum.repos.d \
        && curl -LOJ http://rpms.adiscon.com/v8-stable/rsyslog.repo

RUN yum -y install \
            rsyslog \
            python3 \
        && yum clean all

EXPOSE 514/udp

CMD ["/usr/sbin/rsyslogd", "-n"]
