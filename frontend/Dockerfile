FROM node:21-alpine as development

WORKDIR /code

COPY package.json .
COPY package-lock.json .

RUN npm ci

COPY . .

ENV CI=true

CMD [ "npm", "run", "dev" ]

FROM development AS build

RUN npm run build

FROM nginx:1.25-alpine

COPY --chown=nginx:nginx ../nginx/nginx.conf /etc/nginx/nginx.conf
