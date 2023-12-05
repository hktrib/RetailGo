// config.ts
import { ClassDictionary } from "clsx";
import { useFetch } from "../utils";

// http://localhost:8080/
// https://retailgo-production.up.railway.app/

const env : string = process.env.NODE_ENV

// config.tsx

let config : ClassDictionary;

if (env == "PROD") {
    config = {
        serverURL: 'https://retailgo-production.up.railway.app/'
    };
}
else {
    config = {
        serverURL: 'http://localhost:8080/'
    }; 
}


export { config };
