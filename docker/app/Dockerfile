FROM ubuntu:22.04

ARG SERVICE

RUN apt update && apt install -y ca-certificates && update-ca-certificates

COPY "./$SERVICE/app-runtime" "./"

RUN ls

EXPOSE 10080 10433

CMD ["./app-runtime"]
