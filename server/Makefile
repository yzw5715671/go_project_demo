default: test

NAME := demo-server

server:
	go version
	go build -o $(NAME) .; \
	mkdir -p bin; \
	mv -f $(NAME) ./bin

test: server
	cp ./conf/logtest.xml ./conf/log.xml
	cp ./conf/test.ini ./conf/conf.ini

production: server
	cp ./conf/logproduction.xml ./conf/log.xml
	cp ./conf/production.ini ./conf/conf.ini

start:
	./bin/$(NAME)
