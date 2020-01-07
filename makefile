CONFIGPATH=./conf/config.toml
#LISTENPORT=`cat ${CONFIGPATH} | grep 'listenPort1' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
#REACTPORT= `cat ${CONFIGPATH} | grep 'reactPort2' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
LISTENPORT1=`cat ${CONFIGPATH} | grep 'listenPort1' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
LISTENPORT2=`cat ${CONFIGPATH} | grep 'listenPort2' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
REACTPORT1= `cat ${CONFIGPATH} | grep 'reactPort1' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
REACTPORT2= `cat ${CONFIGPATH} | grep 'reactPort2' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
POSTGRESPORT = `cat ${CONFIGPATH} | grep 'postgresPort' | sed -e 's/\(.*\)\"\(.*\)\"\(.*\)/\2/'`
#webSocketDemo2
build:
	docker build -t websocket:server -f ./server/Dockerfile .
	#docker pull postgres:11.3
	#docker build -t websocket:jmeter -f ./apache-jmeter-5.2.1/Dockerfile .

pack:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./server/server ./server/main.go
	#go build  -o ./server/server ./server/main.go
	#go build -o ./client/client ./client/main.go 

run:
	#docker run -it -d --name=postgres -p ${POSTGRESPORT}:5432 postgres:11.3
	#docker run -it -d --name=server -p ${REACTPORT}:${LISTENPORT}  websocket:server
	docker run -it -d --name=server -p ${REACTPORT1}:${LISTENPORT1} -p ${REACTPORT2}:${LISTENPORT2} websocket:server
	
test:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./client/client ./client/main.go 
	docker build -t websocket:client -f ./client/Dockerfile .	
	docker run -it -d --name=client1 websocket:client ./client
	#docker run -it -d --name=client2 websocket:client ./client
	#docker run -it -d --name=client3 websocket:client ./client
	#docker run -it -d --name=client4 websocket:client ./client
	#docker run -it -d --name=client5 websocket:client ./client
	#docker run -it -d -m 3G --name=jmeter1 websocket:jmeter 
	#docker run -it -d --name=jmeter1 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl
	#docker run -it -d --name=jmeter2 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl
	#docker run -it -d --name=jmeter3 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl
	#docker run -it -d --name=jmeter4 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl
	#docker run -it -d --name=jmeter5 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl
	#docker run -it -d --name=jmeter6 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl
	#docker run -it -d --name=jmeter7 websocket:jmeter ./bin/jmeter -n -t processGroup.jmx -l log.jtl

all:
	#docker rm -f `docker ps -a | grep 'client' | awk '{print $1}'`
	#docker rm -f `docker ps -a | grep 'server' | awk '{print $1}'`
	make pack
	make build
	make run
	#删除无用镜像
	#docker rmi -f  `docker images | grep '<none>' | awk '{print $3}'`