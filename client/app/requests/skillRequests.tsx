import { apiRequests } from "./apiRequests";

function createSkill(authToken: string, topic: string) {
  return apiRequests.postRequest(
    '/skill', 
    { Authorization: authToken }, 
    { topic }
  );
};

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

function deleteSkill(authToken: string, topic: string) {
  return apiRequests.deleteRequest(
    `/skill/${topic}`,
    { Authorization: authToken }
  );
};

function getSkills(authToken: string, userId: number): Promise<string[]> {
  return apiRequests.getRequest(`/skill/${userId}`, { Authorization: authToken });
}

export const skillRequests = {
  createSkill,
  updateSkill,
  deleteSkill,
  getSkills,
};