FROM node:20-alpine

WORKDIR /powerplay/site
RUN npm i -g @quasar/cli
COPY package*.json ./
RUN npm ci