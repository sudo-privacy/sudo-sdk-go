FROM 10.0.1.22:80/dev/sudo-sdk-go/doc:latest 

WORKDIR /go/sudosdk

COPY v2 /go/sudosdk/v2
COPY README.md /go/sudosdk/v2/README.md
COPY LICENSE /go/sudosdk/v2/LICENSE

CMD ["pkgsite","-cache","-proxy","-gorepo=/go/go","-http=:6060","/go/sudosdk/v2"]