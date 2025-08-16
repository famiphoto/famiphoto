FROM node:20-bookworm-slim as local
WORKDIR /var/www/famiphoto_web
COPY ../../web .
RUN rm -rf node_modules package-lock.json
RUN npm i
EXPOSE 3000

FROM node:20-bookworm-slim as builder
WORKDIR /var/www/famiphoto_web
COPY ../../web .
RUN npm ci
RUN npm run build

FROM node:20-bookworm-slim
WORKDIR /var/www/famiphoto_web
ADD .env .
COPY --from=builder /var/www/famiphoto_web/build ./
COPY --from=builder /var/www/famiphoto_web/server.js ./server.js
COPY --from=builder /var/www/famiphoto_web/package.json ./package.json
EXPOSE 3000
CMD ["npm", "run", "start"]
