FROM alpine:3.6

COPY modulario .

RUN cp /vol/modulario . \
 && chmod +x modulario

EXPOSE 80:5000

CMD ["modulario"]