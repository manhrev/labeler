FROM golang:1.22.3-alpine
# RUN apk add git
WORKDIR /app
COPY ./ /app/
WORKDIR /app/cmd
RUN go build -o /cmd
COPY wait-for-it.sh /wait-for-it.sh
COPY /docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /wait-for-it.sh /docker-entrypoint.sh
RUN apk add --no-cache bash
EXPOSE ${LISTEN_PORT}

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/cmd"]