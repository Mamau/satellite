FROM node:14-alpine

RUN apk add git
RUN npm i -g bower

COPY docker-entrypoint.sh /usr/local/bin/

RUN chmod 777 /usr/local/bin/docker-entrypoint.sh \
    && ln -s /usr/local/bin/docker-entrypoint.sh /

RUN mkdir /.config
RUN mkdir /.npm

RUN chmod 777 /.config
RUN chmod 777 /.npm

ENTRYPOINT ["docker-entrypoint.sh"]

CMD [ "bower" ]