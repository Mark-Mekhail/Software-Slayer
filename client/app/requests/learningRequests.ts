import { apiRequests } from "./apiRequests";
import { getErrorMessageFromResponse } from "./requestUtils";
import { ApiError } from "./userRequests";

/**
 * Learning item data structure
 */
export interface LearningItem {
  id: number;
  title: string;
  category: string;
}

/**
 * Creates a new learning item
 * @param authToken - User's authentication token
 * @param title - Title of the learning item
 * @param category - Category of the learning item
 * @throws {ApiError} If the request fails
 */
async function createLearning(authToken: string, title: string, category: string): Promise<void> {
  const response = await apiRequests.postRequest(
    "/learning",
    { Authorization: authToken },
    { title, category },
  );

  if (!response.ok) {
    const errorMessage = await getErrorMessageFromResponse(response, "Failed to add learning item");
    throw new ApiError(errorMessage, response.status);
  }
}

/**
 * Deletes a learning item
 * @param authToken - User's authentication token
 * @param id - ID of the learning item to delete
 * @throws {ApiError} If the request fails
 */
async function deleteLearning(authToken: string, id: number): Promise<void> {
  const response = await apiRequests.deleteRequest(`/learning/${id}`, { Authorization: authToken });

  if (!response.ok) {
    const errorMessage = await getErrorMessageFromResponse(
      response,
      "Failed to remove learning item",
    );
    throw new ApiError(errorMessage, response.status);
  }
}

/**
 * Retrieves learning items for a specific user
 * @param userId - ID of the user
 * @returns Array of learning items
 * @throws {ApiError} If the request fails
 */
async function getLearnings(userId: number): Promise<LearningItem[]> {
  const response = await apiRequests.getRequest(`/learning/${userId}`);

  if (!response.ok) {
    const errorMessage = await getErrorMessageFromResponse(
      response,
      "Failed to get learning items",
    );
    throw new ApiError(errorMessage, response.status);
  }

  return response.json() as Promise<LearningItem[]>;
}

/**
 * Retrieves available learning categories
 * @returns Array of category names
 * @throws {ApiError} If the request fails
 */
async function getLearningCategories(): Promise<string[]> {
  const response = await apiRequests.getRequest("/learning/categories");

  if (!response.ok) {
    const errorMessage = await getErrorMessageFromResponse(
      response,
      "Failed to get learning categories",
    );
    throw new ApiError(errorMessage, response.status);
  }

  return response.json() as Promise<string[]>;
}

export { createLearning, deleteLearning, getLearnings, getLearningCategories };
