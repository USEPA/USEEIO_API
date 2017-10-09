FROM scratch

COPY app /opt/smm-tool
WORKDIR /opt/smm-tool

EXPOSE 5000

CMD [ "./app" ]
