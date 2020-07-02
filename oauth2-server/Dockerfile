FROM maven:3.6.3-jdk-8
MAINTAINER LiuFengJia
WORKDIR /root
RUN git clone https://github.com/yompc/go-todo-list.git
RUN ls -al
WORKDIR /root/go-todo-list/oauth2-server
RUN mvn clean package
COPY ./target/oauth2-server-0.0.1-SNAPSHOT.jar /app.jar
EXPOSE 9000
ENTRYPOINT ["java","-jar","/app.jar"]