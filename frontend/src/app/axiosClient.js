// Import the axios library for making HTTP requests
import axios from 'axios';

// Import the constants file that contains some utility constants
import constants from '../common/utils/constants';

// Create an axios client instance
const axiosClient = axios.create();

// Set the base URL for the axios client
axiosClient.defaults.baseURL = constants.HOST_URL;

// Set the headers for the axios client
axiosClient.defaults.headers = constants.headers;

// Set whether to share cookies to cross-site domain (change to true if needed)
axiosClient.defaults.withCredentials = false;

// Function to make a GET request
export function getRequest(URL) {
  // Make a GET request using the axios client and return the response
  return axiosClient.get(`/${URL}`).then(response => response);
}

// Function to make a POST request
export function postRequest(URL, payload) {
  // Make a POST request using the axios client and return the response
  return axiosClient.post(`/${URL}`, payload).then(response => response);
}

// Function to make a PATCH request
export function patchRequest(URL, payload) {
  // Make a PATCH request using the axios client and return the response
  return axiosClient.patch(`/${URL}`, payload).then(response => response);
}

// Function to make a DELETE request
export function deleteRequest(URL) {
  // Make a DELETE request using the axios client and return the response
  return axiosClient.delete(`/${URL}`).then(response => response);
}