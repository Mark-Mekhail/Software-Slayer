import axios, { AxiosResponse } from 'axios';

// TODO: Update BASE_URL to the correct URL for the backend API server
const BASE_URL = 'http://localhost:8080';

/*
 * getRequest is a generic function that makes a GET request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function getRequest<T>(endpoint: string, headers?: any): Promise<T> {
  try {
    const response: AxiosResponse<T> = await axios.get(`${BASE_URL}${endpoint}`, { headers });
    return response.data;
  } catch (error) {
    console.error('GET request error:', error);
    throw error;
  }
};

/*
 * postRequest is a generic function that makes a POST request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @param payload: optional payload to include in the request
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function postRequest<T>(endpoint: string, headers?: any, payload?: any): Promise<T> {
  try {
    const response: AxiosResponse<T> = await axios.post(`${BASE_URL}${endpoint}`, payload, { headers });
    return response.data;
  } catch (error) {
    console.error('POST request error:', error);
    throw error;
  }
};

/*
 * putRequest is a generic function that makes a PUT request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @param payload: optional payload to include in the request
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function putRequest<T>(endpoint: string, headers?: any, payload?: any): Promise<T> {
  try {
    const response: AxiosResponse<T> = await axios.put(`${BASE_URL}${endpoint}`, payload, { headers });
    return response.data;
  } catch (error) {
    console.error('PUT request error:', error);
    throw error;
  }
};

/*
 * deleteRequest is a generic function that makes a DELETE request to the specified endpoint.
 * @param endpoint: the endpoint to make the request to
 * @param headers: optional headers to include in the request
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function deleteRequest<T>(endpoint: string, headers?: any) {
  try {
    const response: AxiosResponse<T> = await axios.delete(`${BASE_URL}${endpoint}`, { headers });
    return response.data;
  } catch (error) {
    console.error('DELETE request error:', error);
    throw error;
  }
};

export const apiRequests = {
  getRequest,
  postRequest,
  putRequest,
  deleteRequest,
};

