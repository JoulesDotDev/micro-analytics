FROM alpine
ADD micro-analytics /micro-analytics
ENTRYPOINT [ "/micro-analytics" ]
