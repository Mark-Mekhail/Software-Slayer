import { apiRequests } from "./apiRequests";

function createUser(
  email: string,
  firstName: string,
  lastName: string,
  username: string,
  password: string
) {
  return apiRequests.postRequest('/user', {
    email,
    "first_name": firstName,
    "last_name": lastName,
    username,
    password,
  });
}

function login(identifier: string, password: string) {
  return apiRequests.postRequest('/login', {
    identifier,
    password,
  });
}

export const userRequests = {
  createUser,
  login,
};