import { apiRequests } from "./apiRequests";

interface LearningItem {
  id: number;
  title: string;
  category: string;
}

/*
 * createLearning is a function that makes a POST request to create a new learning item.
 * @param authToken: the user's auth token
 * @param title: the title of the learning item
 * @param category: the category of the learning item
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function createLearning(authToken: string, title: string, category: string): Promise<void> {
  const response = await apiRequests.postRequest(
    "/learning",
    { Authorization: authToken },
    {
      title,
      category,
    },
  );

  if (!response.ok) {
    throw new Error("Failed to add learning item");
  }
}

/*
 * deleteLearning is a function that makes a DELETE request to delete a learning item.
 * @param authToken: the user's auth token
 * @param id: the ID of the learning item
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function deleteLearning(authToken: string, id: number): Promise<void> {
  const response = await apiRequests.deleteRequest(`/learning/${id}`, { Authorization: authToken });

  if (!response.ok) {
    throw new Error("Failed to remove learning item");
  }
}

/*
 * getLearnings is a function that makes a GET request to get all learning items for a user.
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function getLearnings(userId: number): Promise<LearningItem[]> {
  const response = await apiRequests.getRequest(`/learning/${userId}`);

  if (!response.ok) {
    throw new Error("Failed to get learning items");
  }

  return response.json();
}

async function getLearningCategories(): Promise<string[]> {
  const response = await apiRequests.getRequest("/learning/categories");

  if (!response.ok) {
    throw new Error("Failed to get learning categories");
  }

  return response.json();
}

export { LearningItem, createLearning, deleteLearning, getLearnings, getLearningCategories };
