# Use the official Redis image from Bitnami
FROM docker.io/bitnami/redis:7.2

# Set environment variables
ENV ALLOW_EMPTY_PASSWORD=yes
ENV REDIS_DISABLE_COMMANDS=FLUSHDB,FLUSHALL

# Expose the Redis port
EXPOSE 6379

# Define a volume for Redis data
VOLUME /bitnami/redis/data

# Note: The 'volumes' section in Docker Compose is not needed in the Dockerfile,
# as volumes are typically managed in Docker Compose or when running the container.

# Optionally, you can provide additional instructions or configurations here.
