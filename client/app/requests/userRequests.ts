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
async function createUser(
  email: string,
  firstName: string,
  lastName: string,
  username: string,
  password: string,
): Promise<void> {
  const response = await apiRequests.postRequest("/user", null, {
    email,
    first_name: firstName,
    last_name: lastName,
    username,
    password,
  });

  if (!response.ok) {
    throw new Error("Failed to create user");
  }
}

/*
 * login is a function that makes a POST request to log in a user.
 * @param identifier: the user's email or username
 * @param password: the user's password
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function login(identifier: string, password: string): Promise<object> {
  const response = await apiRequests.postRequest("/login", null, {
    identifier,
    password,
  });

  if (!response.ok) {
    throw new Error("Failed to log in");
  }

  return response.json();
}

export { createUser, login };
