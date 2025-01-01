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
    const response = await fetch(`${BASE_URL}${endpoint}`, { headers });
    if (!response.ok) {
      throw new Error('GET request failed');
    }
    return await response.json();
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
    const response = await fetch(`${BASE_URL}${endpoint}`, {
      method: 'POST',
      headers,
      body: JSON.stringify(payload),
    });
    if (!response.ok) {
      throw new Error('POST request failed');
    }

    return await response.json();
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
    const response = await fetch(`${BASE_URL}${endpoint}`, {
      method: 'PUT',
      headers,
      body: JSON.stringify(payload),
    });
    if (!response.ok) {
      throw new Error('PUT request failed');
    }
    return await response.json();
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
    const response = await fetch(`${BASE_URL}${endpoint}`, {
      method: 'DELETE',
      headers,
    });
    if (!response.ok) {
      throw new Error('DELETE request failed');
    }
    return await response.json();
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