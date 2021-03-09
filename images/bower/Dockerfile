FROM node:14-alpine

RUN apk add git
RUN npm i -g bower

RUN mkdir /.config
RUN chmod a+rwx -R /home/node
RUN chmod a+rwx -R /.config
WORKDIR /home/node

ENTRYPOINT [ "bower" ]