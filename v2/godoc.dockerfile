FROM 10.0.1.22:80/dev/sudo-sdk-go/doc:latest 

WORKDIR /go/sudosdk

COPY README.md /go/sudosdk/README.md
COPY pkg /go/sudosdk/pkg
COPY protobuf /go/sudosdk/protobuf
COPY go.mod /go/sudosdk/go.mod
COPY go.sum /go/sudosdk/go.sum

CMD ["pkgsite","-cache","-proxy","-gorepo=/go/go","-http=:6060","/go/sudosdk"]