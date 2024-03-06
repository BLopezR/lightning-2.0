FROM golang:1.20 AS build

WORKDIR /usr/src/bin

COPY ./main.go ./go.mod ./

RUN go build -v -o /usr/local/bin/l2sm-br ./... 

FROM ubuntu:latest 
RUN  apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y nano sudo iproute2 iputils-ping tcpdump traceroute ssh # buildkit
RUN useradd -m -s /bin/bash student
RUN echo "student ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/student && chmod 0440 /etc/sudoers.d/student # buildkit
USER student
WORKDIR /home/student
RUN mkdir $HOME/shared-folder # buildkit

COPY --from=build /usr/local/bin/ .

COPY ./vswitch.ovsschema /tmp/

RUN sudo apt-get update && \
  sudo apt-get install -y net-tools iproute2 netcat-openbsd dnsutils curl iputils-ping iptables nmap tcpdump openvswitch-switch && \
  sudo mkdir /var/run/openvswitch && sudo mkdir -p /etc/openvswitch && sudo ovsdb-tool create /etc/openvswitch/conf.db /tmp/vswitch.ovsschema 

RUN sudo apt-get install -y python3 python3-pip && pip3 install --user "simulaqron==3.0.16" # buildkit

COPY ./setup_switch.sh .

RUN sudo chmod +x ./setup_switch.sh


CMD [ "./setup_switch.sh" ]