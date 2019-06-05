FROM zeqi/micro:0.1.0

WORKDIR /go/src/vc.svc
RUN rm -rf ./*
COPY . .
RUN make build-linux-server
# CMD [ "sh", "/go/src/vc.svc/entrypoint.sh" ]
ENTRYPOINT /go/src/vc.svc/vc-svc
# ENTRYPOINT [ "/go/src/app/vc" ]
# ENTRYPOINT ['entrypoint.sh']
# COPY vc-svc /bin/vc-svc
# FROM scratch
# COPY --from=server /vc-svc /bin/vc-svc
# ENTRYPOINT ["vc-svc"]