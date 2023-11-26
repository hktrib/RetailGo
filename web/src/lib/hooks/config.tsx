// config.ts
import { ClassDictionary } from "clsx";
import { useFetch } from "../utils";

// http://localhost:8080/
// https://retailgo-production.up.railway.app/

const Environment : string = "TEST"

// config.tsx

let config : ClassDictionary;

if (Environment == "PRO") {
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
