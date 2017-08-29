FROM gitlab.priv.als:5050/docker/go-project:master

WORKDIR /app

COPY . /app

RUN go get -u github.com/aws/aws-sdk-go

RUN go build -v \
    && wget -O /usr/local/share/ca-certificates/Altus_CA.crt http://hauler.priv.als/share/Utilidades/Altus_CA.crt \
    && update-ca-certificates

CMD ["/app/app"]

EXPOSE 9090
