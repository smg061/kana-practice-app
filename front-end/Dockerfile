FROM node:alpine
COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i
RUN npm run build
EXPOSE 3000
CMD ["npm", "run", "start"]
