FROM mcr.microsoft.com/java/maven:8u252-zulu-debian10
MAINTAINER LiuFengJia
WORKDIR /oauth2-server
RUN mvn clean package
COPY ./target/oauth2-server-0.0.1-SNAPSHOT.jar /app.jar
EXPOSE 9000
ENTRYPOINT ["java","-jar","/app.jar"]