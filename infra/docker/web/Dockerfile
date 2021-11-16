FROM node:14.18.1-alpine

ENV CHOKIDAR_USEPOLLING=true
ENV HOST 0.0.0.0

WORKDIR /src

RUN apk add --no-cache \
        alpine-sdk \
        git

RUN apk update && \
    yarn global add create-nuxt-app

RUN yarn install
CMD ["yarn", "dev"]

EXPOSE 3000