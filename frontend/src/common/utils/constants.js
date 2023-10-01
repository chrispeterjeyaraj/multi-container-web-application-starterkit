/**
 * A constant object containing various configuration values.
 * @typedef {Object} Constants
 * @property {string} HOST_URL - The host URL for the API.
 * @property {Object} headers - The headers for API requests.
 * @property {string} headers.Content-Type - The content type for API requests.
 * @property {string} headers.Accept - The accept header for API requests.
 */

/**
 * The constant object containing various configuration values.
 * @type {Constants}
 */
const constants = {
  /**
   * The host URL for the API.
   * Note: This should be defined in the .env file. And the url we get from the backend container app
   * @type {string}
   */
  HOST_URL: process.env.REACT_APP_HOST_URL,
  /**
   * The headers for API requests.
   * @type {Object}
   */
  headers: {
    /**
     * The content type for API requests.
     * @type {string}
     */
    'Content-Type': 'application/json',
    /**
     * The accept header for API requests.
     * @type {string}
     */
    Accept: 'application/json',
  },
};

export default constants;