FROM node:lts-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm clean-install
COPY . .
RUN npm run build


# production stage
FROM nginx:stable-alpine as production
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /app
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

