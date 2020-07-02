FROM maven:3.6.3-jdk-8
MAINTAINER LiuFengJia
WORKDIR /go-todo-list
RUN git clone https://github.com/yompc/go-todo-list.git
RUN ls -al
RUN /go-todo-list/oauth2-server/ mvn clean package
COPY /go-todo-list/oauth2-server/target/oauth2-server-0.0.1-SNAPSHOT.jar /app.jar
EXPOSE 9000
ENTRYPOINT ["java","-jar","/app.jar"]