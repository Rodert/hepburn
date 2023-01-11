FROM golang:1.18
WORKDIR /root/app/
COPY .build/ config.hcl /root/app/
CMD ["./hepburn","web"]