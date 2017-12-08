FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD . /
EXPOSE 8080
ENTRYPOINT ["/go-cutlass-bot"]
