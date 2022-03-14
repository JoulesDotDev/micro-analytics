FROM alpine
ADD analytics /analytics
ENTRYPOINT [ "/analytics" ]
