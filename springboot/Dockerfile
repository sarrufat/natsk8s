FROM maven:3-openjdk-8 AS BUILD

COPY ./ ./

RUN mvn package

FROM openjdk:8-jdk-alpine
ARG JAR_FILE=target/*.jar
COPY --from=BUILD ${JAR_FILE} app.jar
ENTRYPOINT ["java","-jar","/app.jar"]

