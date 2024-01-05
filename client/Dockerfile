FROM node:18-alpine

WORKDIR /app

COPY /public /app/public
COPY /src /app/src
COPY . .

RUN npm install

CMD [ "npm", "run", "dev" ]