FROM node:18-alpine

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./

RUN npm install

COPY frontend/ ./

# 開発用のReactサーバーを起動
CMD ["npm", "run", "dev"]