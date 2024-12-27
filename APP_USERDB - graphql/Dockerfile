# Image: nodejs server
FROM node:18-alpine

WORKDIR /app

# dingen die niet moeten staan in docker ignore
COPY . .

# dependencies
RUN npm install --force

# Node server runt op 3000
EXPOSE 3002

ENTRYPOINT ["/app/docker_entrypoint.sh"]