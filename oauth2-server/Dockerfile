FROM openjdk:8u191-jre-alpine3.9
MAINTAINER LiuFengJia

COPY ./target/oauth2-server-0.0.1-SNAPSHOT.jar /app.jar
EXPOSE 9000
ENTRYPOINT ["java","-jar","/app.jar"]