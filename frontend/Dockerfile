FROM node:lts-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm clean-install
COPY . .
RUN npm run build

FROM build
COPY --from=build /app/.output /app/.output
EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]

