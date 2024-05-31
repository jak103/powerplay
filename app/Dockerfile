FROM node:20-alpine

WORKDIR /powerplay/app
RUN npm i -g @quasar/cli
COPY package*.json ./
RUN npm ci
