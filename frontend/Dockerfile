FROM node:16-alpine AS development
ENV NODE_ENV development
# Add a work directory
WORKDIR /templatefrontend
# Cache and Install dependencies
COPY package.json .
RUN npm install
# Copy app files
COPY . .
# Expose port
EXPOSE 3000
# Start the app
CMD [ "npm", "start" ]

FROM node:16-alpine AS builder
ENV NODE_ENV production
# Add a work directory
WORKDIR /templatefrontend
# Cache and Install dependencies
COPY package.json .
RUN npm install --production
# Copy app files
COPY . .
# Build the app
RUN npm run build

# Bundle static assets with nginx
FROM nginx:1.23.0-alpine as production
ENV NODE_ENV production
# Copy built assets from builder
COPY --from=builder /templatefrontend/build /usr/share/nginx/html
# Add your nginx.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf
# Expose port
EXPOSE 80
# Start nginx
CMD ["nginx", "-g", "daemon off;"]