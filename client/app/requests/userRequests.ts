import { apiRequests } from "./apiRequests";
import { getErrorMessageFromResponse } from "./requestUtils";

/**
 * User data interface returned from login endpoint
 */
export interface UserResponse {
  user_info: {
    id: number;
    email: string;
    username: string;
    first_name: string;
    last_name: string;
  };
  token: string;
}

/**
 * API error with message and status code
 */
export class ApiError extends Error {
  statusCode: number;

  constructor(message: string, statusCode: number) {
    super(message);
    this.name = "ApiError";
    this.statusCode = statusCode;
  }
}

/**
 * Creates a new user account
 * @param email - User's email
 * @param firstName - User's first name
 * @param lastName - User's last name
 * @param username - User's username
 * @param password - User's password
 * @throws {ApiError} If the request fails
 */
async function createUser(
  email: string,
  firstName: string,
  lastName: string,
  username: string,
  password: string,
): Promise<void> {
  const response = await apiRequests.postRequest("/user", undefined, {
    email,
    first_name: firstName,
    last_name: lastName,
    username,
    password,
  });

  if (!response.ok) {
    const errorMessage = await getErrorMessageFromResponse(response, "Failed to create user");
    throw new ApiError(errorMessage, response.status);
  }
}

/**
 * Authenticates a user and returns user info with auth token
 * @param identifier - User's email or username
 * @param password - User's password
 * @returns User information and authentication token
 * @throws {ApiError} If the request fails
 */
async function login(identifier: string, password: string): Promise<UserResponse> {
  const response = await apiRequests.postRequest("/login", undefined, {
    identifier,
    password,
  });

  if (!response.ok) {
    const errorMessage = await getErrorMessageFromResponse(response, "Failed to log in");
    throw new ApiError(errorMessage, response.status);
  }

  return response.json() as Promise<UserResponse>;
}

export { createUser, login };
