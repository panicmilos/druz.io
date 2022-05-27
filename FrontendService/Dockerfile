FROM node:16.13-alpine AS build

# IMPORTANT NOTE: Docker build context should be the root of this repository!

WORKDIR /app

# Install deps
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./

# Build the app
RUN npm run build

FROM nginx:stable-alpine

# Copy build files and nginx config
COPY --from=build /app/build /usr/share/nginx/html
COPY frontend/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

# Copy env script to nginx
WORKDIR /usr/share/nginx/html
COPY frontend/env.sh ./
COPY frontend/.env ./

# Add bash
RUN apk add --no-cache bash

# Make shell script executable
RUN chmod +x env.sh

# Start Nginx server
CMD ["/bin/bash", "-c", "/usr/share/nginx/html/env.sh && nginx -g \"daemon off;\""]