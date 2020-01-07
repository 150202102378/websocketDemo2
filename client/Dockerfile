FROM ubuntu:18.04
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/home/demo
RUN apt-get update
RUN apt-get install --yes vim
RUN apt-get install --yes wget
RUN apt-get install --yes htop
RUN apt-get install --yes tmux
RUN wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz

ADD ./conf /home/demo/src/conf/
ADD ./client/client /home/demo/src/
WORKDIR /home/demo/src
CMD ["/bin/bash"]