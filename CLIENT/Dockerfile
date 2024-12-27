# Image: nodejs server
FROM node:18-alpine

WORKDIR /app

# dingen die niet moeten staan in docker ignore
COPY . .

# dependencies
RUN npm install

# build pure js voor Node
RUN npm run build

# Node server runt op 3000
EXPOSE 3000

ENTRYPOINT ["/app/docker_entrypoint.sh"]