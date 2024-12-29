import axios, { AxiosResponse } from 'axios';

const BASE_URL = 'localhost:8080';

// Generic function for GET requests with query parameters and optional JSON body
async function getRequest<T>(
  endpoint: string,
  params?: Record<string, any>
) {
  try {
    const response: AxiosResponse<T> = await axios.get(`${BASE_URL}${endpoint}`, params);
    return response.data;
  } catch (error) {
    console.error('GET request error:', error);
    throw error;
  }
};

// Generic function for POST requests
async function postRequest<T>(endpoint: string, payload: any) {
  try {
    const response: AxiosResponse<T> = await axios.post(`${BASE_URL}${endpoint}`, payload);
    return response.data;
  } catch (error) {
    console.error('POST request error:', error);
    throw error;
  }
};

export const apiRequests = {
  getRequest,
  postRequest,
};