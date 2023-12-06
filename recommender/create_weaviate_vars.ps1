# Set WEAVIATE_HOSTNAME
[System.Environment]::SetEnvironmentVariable("WEAVIATE_HOSTNAME", "https://retailgo-recengine-eb6uzggu.weaviate.network", [System.EnvironmentVariableTarget]::User)
Write-Host "WEAVIATE_HOSTNAME has been set for the current user."

# Set WEAVIATE_SK
[System.Environment]::SetEnvironmentVariable("WEAVIATE_SK", "isYZjIQAxvMOhFkTUt0bI5xqETVAHGHqO6fU", [System.EnvironmentVariableTarget]::User)
Write-Host "WEAVIATE_SK has been set for the current user."
