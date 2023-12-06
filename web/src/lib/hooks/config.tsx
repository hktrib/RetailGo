// config.ts
import { ClassDictionary } from "clsx";
import * as dotenv from "dotenv";

dotenv.config({ path: "../../../.env.local" });

// http://localhost:8080/
// https://retailgo-production.up.railway.app/

const env: string = process.env.NODE_ENV;

let config: ClassDictionary;

if (env === "development") {
  config = {
    serverURL: "http://localhost:8080/",
  };
} else if (env === "production") {
  config = {
    serverURL: "https://retailgo-production.up.railway.app/",
  };
}

export { config };
