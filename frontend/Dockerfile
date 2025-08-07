FROM node:alpine as base
WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .

FROM base as dev
CMD ["npm", "run", "dev"]

FROM base as build
RUN npm run build

FROM nginx:stable-alpine as prod
COPY --from=build /app/ /usr/share/nginx/html