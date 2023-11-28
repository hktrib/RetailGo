# Setting environment variables for RetailGo for the current user
# Created by Colby Frey

param (
    [string]$env = "pro"
)

# Setting environment variables for RetailGo for the current user
# Created by Colby Frey

# Set DB_DRIVER
[System.Environment]::SetEnvironmentVariable("DB_DRIVER", "postgres", [System.EnvironmentVariableTarget]::User)
Write-Host "DB_DRIVER has been set for the current user."

# Set DB_SOURCE based on environment argument
if ($env -eq "local") {
    # Local environment - PostgreSQL running in Docker
    $dbSource = "postgresql://root:secret@localhost:5432/retail_go?sslmode=disable"
} else {
    # Production environment
    $dbSource = "postgresql://postgres:76ashcuCoOhkEhgb@db.zvevvgcnviqxagbysekg.supabase.co:5432/postgres"
}
[System.Environment]::SetEnvironmentVariable("DB_SOURCE", $dbSource, [System.EnvironmentVariableTarget]::User)
Write-Host "DB_SOURCE has been set for the current user to $dbSource."

# Set other environment variables
[System.Environment]::SetEnvironmentVariable("CLERK_SK", "sk_test_wCmeudOz44ArIXVFbzTjFOOqhbPquW94kdazRMmjfQ", [System.EnvironmentVariableTarget]::User)
Write-Host "CLERK_SK has been set for the current user."

[System.Environment]::SetEnvironmentVariable("REDIS_HOSTNAME", "viaduct.proxy.rlwy.net", [System.EnvironmentVariableTarget]::User)
Write-Host "REDIS_HOSTNAME has been set for the current user."

[System.Environment]::SetEnvironmentVariable("SERVER_ADDRESS", "8080", [System.EnvironmentVariableTarget]::User)
Write-Host "SERVER_ADDRESS has been set for the current user."

[System.Environment]::SetEnvironmentVariable("REDIS_ADDRESS", "redis://default:HBCmEOKMFFGN6525oJombkA6IfnfKaHn@viaduct.proxy.rlwy.net:38806", [System.EnvironmentVariableTarget]::User)
Write-Host "REDIS_ADDRESS has been set for the current user."

[System.Environment]::SetEnvironmentVariable("REDIS_PORT", "38806", [System.EnvironmentVariableTarget]::User)
Write-Host "REDIS_PORT has been set for the current user."

[System.Environment]::SetEnvironmentVariable("REDIS_PASSWORD", "HBCmEOKMFFGN6525oJombkA6IfnfKaHn", [System.EnvironmentVariableTarget]::User)
Write-Host "REDIS_PASSWORD has been set for the current user."

[System.Environment]::SetEnvironmentVariable("STRIPE_SK", "sk_test_51ODz7pHWQUATs9zV4fWYLtRag0GwwLPticrlOe5FqicEWwdnWUlsZkRh90o1YOkt3qsOduJQNSbbUJupkm4i9xLm00hcffWjDm", [System.EnvironmentVariableTarget]::User)
Write-Host "STRIPE_SK has been set for the current user."

# Output a message
Write-Host "Environment variables have been set for the current user with the $environment environment."
