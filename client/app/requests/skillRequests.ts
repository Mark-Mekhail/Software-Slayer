import { apiRequests } from "./apiRequests";

/*
 * createSkill is a function that makes a POST request to create a new skill.
 * @param authToken: the user's auth token
 * @param topic: the topic of the skill
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function createSkill(authToken: string, topic: string): Promise<void> {
  const response = await apiRequests.postRequest(
    '/skill', 
    { Authorization: authToken }, 
    { topic }
  );

  if (!response.ok) {
    throw new Error('Failed to add skill');
  }
};

/*
 * updateSkill is a function that makes a PUT request to update a skill.
 * @param authToken: the user's auth token
 * @param oldTopic: the old topic of the skill
 * @param newTopic: the new topic of the skill
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function updateSkill(authToken: string, oldTopic: string, newTopic: string): Promise<void> {
  const response = await apiRequests.putRequest(
    '/skill', 
    { Authorization: authToken }, 
    {
      oldTopic,
      newTopic,
    }
  );

  if (!response.ok) {
    throw new Error('Failed to update skill');
  }
};

/*
 * deleteSkill is a function that makes a DELETE request to delete a skill.
 * @param authToken: the user's auth token
 * @param topic: the topic of the skill
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function deleteSkill(authToken: string, topic: string): Promise<void> {
  const response = await apiRequests.deleteRequest(
    `/skill/${topic}`,
    { Authorization: authToken }
  );

  if (!response.ok) {
    throw new Error('Failed to remove skill');
  }
};

/*
 * getSkills is a function that makes a GET request to get all skills for a user.
 * @param authToken: the user's auth token
 * @param userId: the user's ID
 * @returns the response data from the request
 * @throws an error if the request fails
 */
async function getSkills(authToken: string, userId: number): Promise<string[]> {
  const response = await apiRequests.getRequest(`/skill/${userId}`, { Authorization: authToken });

  if (!response.ok) {
    throw new Error('Failed to get skills');
  }

  return response.json();
}

export const skillRequests = {
  createSkill,
  updateSkill,
  deleteSkill,
  getSkills,
};