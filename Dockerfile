FROM golang:latest AS base
RUN mkdir /app
ADD . /app/
WORKDIR /app

FROM base AS artifacts
RUN chmod +x /app/wait-for-it.sh
RUN make

CMD ["/app/data-migrator"]

EXPOSE 8080
CMD ["/app/apiserver"]
