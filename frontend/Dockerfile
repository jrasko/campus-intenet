FROM node:lts-alpine3.23 as build
WORKDIR /app
COPY package*.json ./
RUN npm clean-install
COPY . .
RUN npm run build

FROM nginx:stable-alpine
COPY --from=build /app/.output /usr/share/nginx/html
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]
