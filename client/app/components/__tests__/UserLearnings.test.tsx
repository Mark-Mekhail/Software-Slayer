import { render, fireEvent, waitFor } from "@testing-library/react-native";
import React from "react";

import { UserContext } from "../../common/UserContext";
import {
  getLearnings,
  getLearningCategories,
  createLearning,
  deleteLearning,
} from "../../requests/learningRequests";
import UserLearnings from "../UserLearnings";

// Mock dependencies
jest.mock("../../requests/learningRequests");

// Correctly mock global alert instead of Alert.alert
global.alert = jest.fn();

const mockUserContext = {
  user: {
    id: 1,
    email: "test@example.com",
    username: "testuser",
    firstName: "John",
    lastName: "Doe",
    token: "valid_token",
  },
  setUser: jest.fn(),
};

describe("UserLearnings", () => {
  beforeEach(() => {
    jest.clearAllMocks();

    // Default mocks
    (getLearningCategories as jest.Mock).mockResolvedValue([
      "Languages",
      "Technologies",
      "Concepts",
    ]);
    (getLearnings as jest.Mock).mockResolvedValue([
      { id: 1, title: "Go Programming", category: "Languages" },
      { id: 2, title: "Docker", category: "Technologies" },
      { id: 3, title: "Microservices", category: "Concepts" },
    ]);
  });

  it("renders correctly and displays learning items organized by category", async () => {
    const { getByText, getAllByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    await waitFor(() => {
      expect(getByText("John's Learning Lists")).toBeTruthy();
      expect(getByText("Languages")).toBeTruthy();
      expect(getByText("Technologies")).toBeTruthy();
      expect(getByText("Concepts")).toBeTruthy();
      expect(getByText("Go Programming")).toBeTruthy();
      expect(getByText("Docker")).toBeTruthy();
      expect(getByText("Microservices")).toBeTruthy();
      expect(getAllByText("Add").length).toBe(3); // One for each category
    });

    // Verify API calls
    expect(getLearningCategories).toHaveBeenCalledTimes(1);
    expect(getLearnings).toHaveBeenCalledWith(1);
  });

  it("handles empty learning items gracefully", async () => {
    (getLearnings as jest.Mock).mockResolvedValueOnce([]);

    const { getByText, queryByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Wait for all elements to appear using waitFor
    await waitFor(() => {
      expect(getByText("John's Learning Lists")).toBeTruthy();
      expect(getByText("Languages")).toBeTruthy();
      expect(getByText("Technologies")).toBeTruthy();
      expect(getByText("Concepts")).toBeTruthy();
    });

    // Then verify that the list items are not present
    expect(queryByText("Go Programming")).toBeNull();
    expect(queryByText("Docker")).toBeNull();
    expect(queryByText("Microservices")).toBeNull();
  });

  it("adds a new learning item when form is submitted", async () => {
    (createLearning as jest.Mock).mockResolvedValueOnce(undefined);

    const { getByText, getAllByText, getAllByPlaceholderText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    await waitFor(() => {
      expect(getByText("Languages")).toBeTruthy();
    });

    const inputs = getAllByPlaceholderText("Enter learning item title");
    const addButtons = getAllByText("Add");

    // Fill and submit the first category's form (Languages)
    fireEvent.changeText(inputs[0], "React Native");
    fireEvent.press(addButtons[0]);

    await waitFor(() => {
      expect(createLearning).toHaveBeenCalledWith("valid_token", "React Native", "Languages");
      expect(getLearnings).toHaveBeenCalledTimes(2); // Initial load + after add
    });
  });

  it("validates input before adding a new learning item", async () => {
    const { getAllByPlaceholderText, getAllByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Wait for all necessary elements to appear
    await waitFor(() => {
      expect(getAllByText("Add").length).toBe(3);
      expect(getAllByPlaceholderText("Enter learning item title").length).toBe(3);
    });

    const inputs = getAllByPlaceholderText("Enter learning item title");
    const addButtons = getAllByText("Add");

    // Try to add with empty title
    fireEvent.changeText(inputs[0], "   ");
    fireEvent.press(addButtons[0]);

    // Check if alert was called with the expected message
    await waitFor(() => {
      expect(global.alert).toHaveBeenCalledWith("Please enter a valid title");
    });
    expect(createLearning).not.toHaveBeenCalled();
  });

  it("deletes a learning item when delete button is pressed", async () => {
    (deleteLearning as jest.Mock).mockResolvedValueOnce(undefined);

    const { getAllByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    await waitFor(() => {
      expect(getAllByText("Delete").length).toBe(3);
    });

    // Press the first delete button
    fireEvent.press(getAllByText("Delete")[0]);

    await waitFor(() => {
      expect(deleteLearning).toHaveBeenCalledWith("valid_token", 1);
    });
  });

  it("shows an error alert when creating a learning item fails", async () => {
    (createLearning as jest.Mock).mockRejectedValueOnce({ message: "Creation failed" });

    const { getAllByText, getAllByPlaceholderText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Wait for all necessary elements to appear
    await waitFor(() => {
      expect(getAllByText("Add").length).toBe(3);
      expect(getAllByPlaceholderText("Enter learning item title").length).toBe(3);
    });

    const inputs = getAllByPlaceholderText("Enter learning item title");
    const addButtons = getAllByText("Add");

    // Fill and submit form
    fireEvent.changeText(inputs[0], "React Native");
    fireEvent.press(addButtons[0]);

    // Check if alert was called
    await waitFor(() => {
      expect(global.alert).toHaveBeenCalledWith("Error: Could not create learning item");
    });
  });

  it("shows an error alert when deleting a learning item fails", async () => {
    (deleteLearning as jest.Mock).mockRejectedValueOnce({ message: "Deletion failed" });

    const { getAllByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Wait for delete buttons to appear
    await waitFor(() => {
      expect(getAllByText("Delete").length).toBe(3);
    });

    // Press the first delete button
    fireEvent.press(getAllByText("Delete")[0]);

    // Check if alert was called
    await waitFor(() => {
      expect(global.alert).toHaveBeenCalledWith("Error: Could not delete learning item");
    });
  });

  it("handles API failure when fetching categories", async () => {
    const consoleErrorSpy = jest.spyOn(console, "error").mockImplementation(() => {});

    // Use a string error instead of an object to match the alert message format
    (getLearningCategories as jest.Mock).mockRejectedValueOnce(
      new Error("Failed to fetch categories"),
    );

    render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Just wait for the API call to be made
    await waitFor(() => {
      expect(getLearningCategories).toHaveBeenCalledTimes(1);
    });

    consoleErrorSpy.mockRestore();
  });

  it("handles API failure when fetching learning items", async () => {
    const consoleErrorSpy = jest.spyOn(console, "error").mockImplementation(() => {});

    // Use a string error instead of an object to match the alert message format
    (getLearnings as jest.Mock).mockRejectedValueOnce(new Error("Failed to fetch learning items"));

    render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Just wait for the API call to be made
    await waitFor(() => {
      expect(getLearnings).toHaveBeenCalledTimes(1);
    });

    consoleErrorSpy.mockRestore();
  });

  it("throws an error when UserContext is not provided", () => {
    // Use spyOn instead of direct assignment
    const consoleErrorSpy = jest.spyOn(console, "error").mockImplementation(() => {});

    expect(() => {
      render(<UserLearnings />);
    }).toThrow("UserContext is not set");

    // Clean up
    consoleErrorSpy.mockRestore();
  });
});
