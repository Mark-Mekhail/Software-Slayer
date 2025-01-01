import { apiRequests } from "./apiRequests";

/*
 * createUser is a function that makes a POST request to create a new user.
 * @param email: the user's email
 * @param firstName: the user's first name
 * @param lastName: the user's last name
 * @param username: the user's username
 * @param password: the user's password
 * @returns the response data from the request
 * @throws an error if the request fails
 */
function createUser(
  email: string,
  firstName: string,
  lastName: string,
  username: string,
  password: string
) {
  return apiRequests.postRequest('/user', null, {
    email,
    "first_name": firstName,
    "last_name": lastName,
    username,
    password,
  });
}

/*
 * login is a function that makes a POST request to log in a user.
 * @param identifier: the user's email or username
 * @param password: the user's password
 * @returns the response data from the request
 * @throws an error if the request fails
 */
function login(identifier: string, password: string) {
  return apiRequests.postRequest('/login', null, {
    identifier,
    password,
  });
}

export const userRequests = {
  createUser,
  login,
};