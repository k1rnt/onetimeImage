FROM node:14.18.1-alpine as build-stage
WORKDIR /client
COPY ./client .
RUN yarn install

ENV CHOKIDAR_USEPOLLING=true
EXPOSE 3000
CMD [ "yarn" "dev" ]