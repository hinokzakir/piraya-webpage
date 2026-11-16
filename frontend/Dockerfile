
# Build the Vue app
FROM node:20-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Serve with Nginx
FROM nginx:alpine
# Copy the compiled Vue app (ensure 'dist' matches your Vue build output folder)
COPY --from=build-stage /app/dist /usr/share/nginx/html
# Copy the routing configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]