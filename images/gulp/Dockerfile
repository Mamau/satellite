FROM node:14-alpine

RUN apk add git
RUN npm i -g gulp

RUN chmod a+rwx -R /home/node
WORKDIR /home/node

ENTRYPOINT [ "gulp" ]