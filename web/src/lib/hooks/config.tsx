// config.ts
import { ClassDictionary } from "clsx";
import * as dotenv from "dotenv";

dotenv.config({ path: "../../../.env.local" });

// http://localhost:8080/
// https://retailgo-production.up.railway.app/

let env: string = process.env.NODE_ENV;

let config: ClassDictionary;

// env = "production";
// console.log(env)

if (env === "development") {
  config = {
    serverURL: "http://localhost:8080",
    recServerURL: "http://recommendation-server-production.up.railway.app"
  };
} else if (env === "production") {
  config = {
    serverURL: "https://retailgo-production.up.railway.app",
    recServerURL: "http://recommendation-server-production.up.railway.app"
  };
}

// console.log(serverURL);

export { config };
