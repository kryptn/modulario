FROM golang:1.8-alpine3.6

ARG GITHUB_TOKEN

RUN apk add -U git gcc musl-dev

RUN git config --global credential.helper store \
 && echo "https://${GITHUB_TOKEN}:x-oauth-basic@github.com/" >> ~/.git-credentials \
 && git config --global url."https://github.com".insteadOf git://github.com

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 80:5000

CMD ["go-wrapper", "run"] # ["app"]