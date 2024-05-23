FROM node:19-alpine3.16 AS build

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm i

COPY . .

RUN npm run build

FROM node:19-alpine3.16 AS production

ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm i --omit=dev --omit=optional

COPY --from=build /usr/src/app/dist ./dist

CMD ["node", "dist/main"]
