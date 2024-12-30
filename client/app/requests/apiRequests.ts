import axios, { AxiosResponse } from 'axios';

const BASE_URL = 'http://localhost:8080';

// Generic function for GET requests with query parameters and optional JSON body
async function getRequest<T>(endpoint: string, headers?: any) {
  try {
    const response: AxiosResponse<T> = await axios.get(`${BASE_URL}${endpoint}`, { headers });
    return response.data;
  } catch (error) {
    console.error('GET request error:', error);
    throw error;
  }
};

// Generic function for POST requests
async function postRequest<T>(endpoint: string, headers?: any, payload?: any) {
  try {
    const response: AxiosResponse<T> = await axios.post(`${BASE_URL}${endpoint}`, payload, { headers });
    return response.data;
  } catch (error) {
    console.error('POST request error:', error);
    throw error;
  }
};

async function putRequest<T>(endpoint: string, headers?: any, payload?: any) {
  try {
    const response: AxiosResponse<T> = await axios.put(`${BASE_URL}${endpoint}`, payload, { headers });
    return response.data;
  } catch (error) {
    console.error('PUT request error:', error);
    throw error;
  }
};

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

