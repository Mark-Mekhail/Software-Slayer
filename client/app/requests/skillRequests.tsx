import { apiRequests } from "./apiRequests";

/*
 * createSkill is a function that makes a POST request to create a new skill.
 * @param authToken: the user's auth token
 * @param topic: the topic of the skill
 * @returns the response data from the request
 * @throws an error if the request fails
 */
function createSkill(authToken: string, topic: string) {
  return apiRequests.postRequest(
    '/skill', 
    { Authorization: authToken }, 
    { topic }
  );
};

/*
 * updateSkill is a function that makes a PUT request to update a skill.
 * @param authToken: the user's auth token
 * @param oldTopic: the old topic of the skill
 * @param newTopic: the new topic of the skill
 * @returns the response data from the request
 * @throws an error if the request fails
 */
function updateSkill(authToken: string, oldTopic: string, newTopic: string) {
  return apiRequests.putRequest(
    '/skill', 
    { Authorization: authToken }, 
    {
      oldTopic,
      newTopic,
    }
  );
};

/*
 * deleteSkill is a function that makes a DELETE request to delete a skill.
 * @param authToken: the user's auth token
 * @param topic: the topic of the skill
 * @returns the response data from the request
 * @throws an error if the request fails
 */
function deleteSkill(authToken: string, topic: string) {
  return apiRequests.deleteRequest(
    `/skill/${topic}`,
    { Authorization: authToken }
  );
};

/*
 * getSkills is a function that makes a GET request to get all skills for a user.
 * @param authToken: the user's auth token
 * @param userId: the user's ID
 * @returns the response data from the request
 * @throws an error if the request fails
 */
function getSkills(authToken: string, userId: number): Promise<string[]> {
  return apiRequests.getRequest(`/skill/${userId}`, { Authorization: authToken });
}

export const skillRequests = {
  createSkill,
  updateSkill,
  deleteSkill,
  getSkills,
};