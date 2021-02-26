# Using a typical enterprise OS
FROM centos:7

RUN yum -y update \
        && yum clean all

# Using conservative official repo version
# For the latest stable uncomment the following two lines and rebuild
#RUN cd /etc/yum.repos.d \
#        && curl -LOJ http://rpms.adiscon.com/v8-stable/rsyslog.repo

RUN yum -y install \
            rsyslog \
        && yum clean all

EXPOSE 514/udp

CMD ["/usr/sbin/rsyslogd", "-n"]
