FROM golang:1.8-alpine3.6

ARG GITHUB_TOKEN

WORKDIR /go/src/app
COPY . .

RUN apk add -U --virtual .deps gcc musl-dev \
 && git config --global credential.helper store \
 && echo "https://${GITHUB_TOKEN}:x-oauth-basic@github.com/" >> ~/.git-credentials \
 && git config --global url."https://github.com".insteadOf git://github.com \
 && go-wrapper download \
 && go-wrapper install \
 && apk del .deps \
 && rm ~/.git-credentials

EXPOSE 80:5000

CMD ["go-wrapper", "run"] # ["app"]
