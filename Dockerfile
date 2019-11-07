FROM node:10-alpine
WORKDIR /var/www
COPY . /var/www
RUN npm install
EXPOSE 5000
CMD [ "node", "src/app.js" ]