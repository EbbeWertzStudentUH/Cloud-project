FROM node:18-alpine AS base

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install


FROM base AS build

COPY . .

RUN npm run build


FROM node:18-alpine AS runtime

WORKDIR /app
COPY --from=build /app/build /app/build
COPY --from=build /app/node_modules /app/node_modules
COPY --from=build /app/package.json /app/package.json

EXPOSE 3000

COPY docker_entrypoint.sh ./docker_entrypoint.sh
ENTRYPOINT ["/app/docker_entrypoint.sh"]
