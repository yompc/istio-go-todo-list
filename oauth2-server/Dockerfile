FROM 3.6.3-jdk-8
MAINTAINER LiuFengJia
RUN ls -al
RUN mvn clean package
COPY ./target/oauth2-server-0.0.1-SNAPSHOT.jar /app.jar
EXPOSE 9000
ENTRYPOINT ["java","-jar","/app.jar"]