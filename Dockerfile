FROM alpine:latest

RUN mkdir /app

COPY friendManagementApp /app

CMD [ "/app/friendManagementApp"]