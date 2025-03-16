import { Platform } from "react-native";

const BASE_URL = Platform.OS === "android" ? "http://10.0.2.2:8080" : "http://localhost:8080";

/*
 * getRequest is a generic function that makes a GET request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @returns the response from the request
 */
async function getRequest(endpoint: string, headers?: any): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, { headers });
}

/*
 * postRequest is a generic function that makes a POST request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @param payload: optional payload to include in the request
 * @returns the response from the request
 */
async function postRequest(endpoint: string, headers?: any, payload?: any): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    method: "POST",
    headers: {
      ...headers,
    },
    body: JSON.stringify(payload),
  });
}

/*
 * putRequest is a generic function that makes a PUT request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @param payload: optional payload to include in the request
 * @returns the response from the request
 */
async function putRequest(endpoint: string, headers?: any, payload?: any): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    method: "PUT",
    headers,
    body: JSON.stringify(payload),
  });
}

/*
 * deleteRequest is a generic function that makes a DELETE request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @returns the response from the request
 */
async function deleteRequest(endpoint: string, headers?: any): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    method: "DELETE",
    headers,
  });
}

export const apiRequests = {
  getRequest,
  postRequest,
  putRequest,
  deleteRequest,
};
