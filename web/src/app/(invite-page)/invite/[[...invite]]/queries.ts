import { config } from "@/lib/hooks/config";

// Defining an asynchronous function named 'GetStoreByUUID' which takes an object as its parameter
// The object has a property 'uuid' of type string
// This function is responsible for retrieving store data based on the provided UUID
export const GetStoreByUUID = async ({
  uuid,
}: {
  uuid: string;
}) => {

  // Constructing the server URL by appending the UUID to the server URL from the 'config' object
  let serverUrl = `${config.serverURL}/store/uuid/${uuid}`;

  // Logging a message indicating the attempt to get the store by UUID
  console.log(`Attempting to get store by uuid ${uuid}`);

  try {
    // Sending a GET request to the server URL using the 'fetch' function
    let response = await fetch(serverUrl, {
      method: "GET",
    });

    // Checking if the response is not successful (status code other than 2xx)
    // If the response is not successful, return an object with an empty 'store' property and 'success' set to false
    if (!response.ok) return {
      store: {},
      success: false
    };

    // Parsing the response body as text and then parsing it as JSON
    // If the response body is empty or cannot be parsed as JSON, assign an empty object to the 'store' variable
    const store = JSON.parse(await response.text()) ?? {};

    // Logging the response status code and status text
    console.log(`response: ${response.status} + ${response.statusText}`);

    // Returning an object with the retrieved store data and 'success' set to true
    return {
      store: store,
      success: true
    }
  } catch (err) {
    // Logging an error message if an error occurs during the retrieval process
    console.log("Failed to retrieve recommendations:", err)

    // Returning an object with an empty 'store' property and 'success' set to false
    return {
      store: {},
      success: false
    }
  }
};
