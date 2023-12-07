# RetailGo
RetailGo


Contributors: Colby Frey, Apramay Singh, David Nguyen, Giridhar Vadhul, Leo Asatoorian, Jonathan Michel


## Backend

#### Boot PostgreSQL / Redis / Weaviate / Stripe Database
    Requires: Docker
    
    -> run `make devpostres` or equivalient command found in makefile to start PostgreSQL service in Docker container
    
    -> run `make devcreatedb` or equivalient command found in makefile to create database
    
    -> run `make startredis` or equivalent command found in makefile to start redis service. 
    
    -> Create Weaviate Cloud Service Instance/Sandbox
    
    -> Create Clerk Account and Link Secret Key and Webhook Secret
    
    -> Create Stripe Account and provide keys. 
    
    
    -> Attach necessary secrets to env variables in a .env file.
    
        `
            DB_DRIVER=postgres
            DB_SOURCE=
            CLERK_SK=
            CLERK_WEBHOOK_SECRET=
            SERVER_ADDRESS=8080
            REDIS_HOSTNAME=
            REDIS_PORT=
            REDIS_PASSWORD=
            STRIPE_SK=
            RAILWAY_DOCKERFILE_PATH=
            STRIPE_WEBHOOK_SECRET=
            WEAVIATE_HOSTNAME=
            WEAVIATE_SK=
        `
    
#### Run Backend!

Requires: RetailGo Backend ENV variables

-> `go mod tidy`

-> `go run main.go`


#### Run Frontend!

    `
        NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=
        CLERK_SECRET_KEY=
        CLERK_WEBHOOK_SECRET=
        
        NEXT_PUBLIC_CLERK_SIGN_IN_URL=/sign-in
        NEXT_PUBLIC_CLERK_SIGN_UP_URL=/sign-up
        NEXT_PUBLIC_CLERK_AFTER_SIGN_IN_URL=/
        NEXT_PUBLIC_CLERK_AFTER_SIGN_UP_URL=/
        
        EMAIL_USER=retailgoco@gmail.com
        EMAIL_PASSWORD=RG-cse115a
        
        PUBLIC_SUPABASE_URL=
        PUBLIC_SUPABASE_KEY=
    `
Requires: RetailGo Frontend ENV variables

-> `pnpm install`
-> `pnpm run dev`

