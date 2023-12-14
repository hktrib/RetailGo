# RetailGo
RetailGo


Contributors: Colby Frey, Apramay Singh, David Nguyen, Giridhar Vadhul, Leo Asatoorian, Jonathan Michel


### Getting Started with the Backend!

```console foo@bar:~$ git clone git@github.com:hktrib/RetailGo.git```


To lessen the install burden, we'll work with postgres on a docker container via PORT# 5432

1. Ensure that you have docker desktop installed. If not download it here https://www.docker.com/products/docker-desktop/

4. Once your project has been cloned navigate to the **server/Makefile** file.


5. Makefile commands are pretty self-explanatory. 
    - ```foo@bar:RetailGo/server$ make postgres``` spins up a docker container
    - ```foo@bar:RetailGo/server$ make createdb``` creates a database -> retail_go
    - ```foo@bar:RetailGo/server$ make drodb``` deletes a database -> retail_go
    - ```console foo@bar:RetailGo/server$ make sqlc_delete``` deletes the **db/sqlc** folder.
6. run **make postgres** **make createdb** and Configure TablePlus to connect to the retail_go database via PORT# 5432 (Password: secret)
7. Run ```console foo@bar:RetailGo/server$ go run main.go``` 

##### Working With Ent
https://entgo.io/docs/getting-started

##### Docs
Go-Chi -> https://go-chi.io/


### Getting Started with the Front End!

This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

We're using [pnpm](https://pnpm.io/) as the package manager for this application. **Please** download if the following commands don't work:


Create a .env file with the following values:

    NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=
    CLERK_SECRET_KEY=
    CLERK_WEBHOOK_SECRET=
    NEXT_PUBLIC_CLERK_SIGN_IN_URL=
    NEXT_PUBLIC_CLERK_SIGN_UP_URL=
    NEXT_PUBLIC_CLERK_AFTER_SIGN_IN_URL=
    NEXT_PUBLIC_CLERK_AFTER_SIGN_UP_URL=
    EMAIL_USER=
    EMAIL_PASSWORD=
    PUBLIC_SUPABASE_URL=
    PUBLIC_SUPABASE_KEY=

Install pnpm packages:
```bash
pnpm install
```

Run the development server:

```bash
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.