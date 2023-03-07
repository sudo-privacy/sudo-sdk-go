FROM 10.0.1.22:80/dev/sudo-sdk-go/doc:latest AS build

WORKDIR /go/sudosdk

COPY . .
RUN sh scripts/sync_protobuf.sh

FROM 10.0.1.22:80/dev/sudo-sdk-go/doc:latest AS final
WORKDIR /go/sudosdk
COPY --from=build /go/sudosdk/README.md /go/sudosdk/README.md
COPY --from=build /go/sudosdk/pkg /go/sudosdk/pkg
COPY --from=build /go/sudosdk/protobuf /go/sudosdk/protobuf
COPY --from=build /go/sudosdk/go.mod /go/sudosdk/go.mod
COPY --from=build /go/sudosdk/go.sum /go/sudosdk/go.sum

CMD ["pkgsite","-cache","-proxy","-gorepo=/go/go","-http=:6060","/go/sudosdk"]