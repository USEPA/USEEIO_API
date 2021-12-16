FROM public.ect.aws/docker/library/bash

RUN mkdir /opt/useeio_api && mkdir /opt/useeio_api/data
COPY awsdist/app /opt/useeio_api
WORKDIR /opt/useeio_api
RUN chmod +x /opt/useeio_api/app

EXPOSE 8080

CMD [ "./app", "-port", "8080" ]

