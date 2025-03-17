import { Platform } from "react-native";

/**
 * Base URL for API requests, adjusted for platform-specific needs
 * - Android emulator needs 10.0.2.2 to access localhost on host machine
 * - iOS and web can use localhost directly
 */
const BASE_URL = Platform.OS === "android" ? "http://10.0.2.2:8080" : "http://localhost:8080";

type Payload = Record<string, string | number | boolean | null>;

/**
 * Default headers for all requests
 */
const DEFAULT_HEADERS = {
  "Content-Type": "application/json",
  Accept: "application/json",
};

/**
 * Makes a GET request to the specified endpoint
 * @param endpoint - API endpoint path
 * @param headers - Optional custom headers
 * @returns Promise with the fetch Response
 */
async function getRequest(endpoint: string, headers?: Record<string, string>): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    headers: { ...DEFAULT_HEADERS, ...headers },
  });
}

/**
 * Makes a POST request to the specified endpoint
 * @param endpoint - API endpoint path
 * @param headers - Optional custom headers
 * @param payload - Optional request body data
 * @returns Promise with the fetch Response
 */
async function postRequest(
  endpoint: string,
  headers?: Record<string, string>,
  payload?: Payload,
): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    method: "POST",
    headers: { ...DEFAULT_HEADERS, ...headers },
    body: JSON.stringify(payload),
  });
}

/**
 * Makes a PUT request to the specified endpoint
 * @param endpoint - API endpoint path
 * @param headers - Optional custom headers
 * @param payload - Optional request body data
 * @returns Promise with the fetch Response
 */
async function putRequest(
  endpoint: string,
  headers?: Record<string, string>,
  payload?: Payload,
): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    method: "PUT",
    headers: { ...DEFAULT_HEADERS, ...headers },
    body: JSON.stringify(payload),
  });
}

/**
 * Makes a DELETE request to the specified endpoint
 * @param endpoint - API endpoint path
 * @param headers - Optional custom headers
 * @returns Promise with the fetch Response
 */
async function deleteRequest(
  endpoint: string,
  headers?: Record<string, string>,
): Promise<Response> {
  return await fetch(`${BASE_URL}${endpoint}`, {
    method: "DELETE",
    headers: { ...DEFAULT_HEADERS, ...headers },
  });
}

export const apiRequests = {
  getRequest,
  postRequest,
  putRequest,
  deleteRequest,
};
