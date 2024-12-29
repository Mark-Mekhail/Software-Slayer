import { apiRequests } from "./apiRequests";

function createUser(
  email: string,
  firstName: string,
  lastName: string,
  username: string,
  password: string
) {
  return apiRequests.postRequest('/users', {
    email,
    firstName,
    lastName,
    username,
    password,
  });
}