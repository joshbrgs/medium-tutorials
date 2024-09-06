# Use the official Deno image as the base image
FROM denoland/deno:alpine-1.36.4

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY ./cmd/notifications/ .

# Expose the port the service will run on
EXPOSE 8091

# Set the default command to run your Deno service
CMD ["run", "--allow-net", "main.ts"]
