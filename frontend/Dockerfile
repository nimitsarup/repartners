# Use official Node.js image as base
FROM node:alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

# Copy app source
COPY . .

# Build the app
RUN npm run build

# Production environment
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
