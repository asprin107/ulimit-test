FROM amazoncorretto:8u242
RUN yum install -y git tar
RUN curl -O https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.8.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
#RUN git clone https://github.com/asprin107/ulimit-test.git
#RUN /usr/local/go/bin/go build /ulimit-test/ulimit.go
#RUN ln -s /ulimit-test/ulimit /ulimit
ADD ./ulimit /ulimit
EXPOSE 8000
ENTRYPOINT /ulimit
