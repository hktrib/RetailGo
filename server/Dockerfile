# Use the official Redis image from Bitnami
FROM docker.io/bitnami/redis:7.2

# Set environment variables
ENV ALLOW_EMPTY_PASSWORD=yes
ENV REDIS_DISABLE_COMMANDS=FLUSHDB,FLUSHALL

# Expose the Redis port
EXPOSE 6379

# Define a volume for Redis data
VOLUME /bitnami/redis/data