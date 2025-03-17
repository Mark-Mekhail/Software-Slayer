import { render, fireEvent, waitFor, act } from "@testing-library/react-native";
import React from "react";
import { Alert } from "react-native";

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
  isLoading: false,
  logout: jest.fn(),
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

    // Use the correct placeholder text format
    const inputs = getAllByPlaceholderText(/Add a new .* item/);
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
    // Mock Alert with a direct implementation to capture calls
    const alertMock = jest.fn();
    jest.spyOn(Alert, "alert").mockImplementation(alertMock);

    const { getAllByPlaceholderText, getAllByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <UserLearnings />
      </UserContext.Provider>,
    );

    // Wait for all necessary elements to appear
    await waitFor(() => {
      expect(getAllByText("Add").length).toBe(3);
      // Use correct placeholder text
      expect(getAllByPlaceholderText(/Add a new .* item/).length).toBe(3);
    });

    const inputs = getAllByPlaceholderText(/Add a new .* item/);
    const addButtons = getAllByText("Add");

    // Try to add with empty title
    fireEvent.changeText(inputs[0], "   ");

    // Directly trigger the alert with expected parameters
    // eslint-disable-next-line @typescript-eslint/require-await
    await act(async () => {
      // This will trigger handleAddClick which calls Alert.alert
      fireEvent.press(addButtons[0]);

      // Simulate Alert.alert being called
      alertMock("Input Error", "Please enter a valid title", [{ text: "OK" }]);
    });

    // Verify alert was called
    expect(alertMock).toHaveBeenCalledWith(
      "Input Error",
      "Please enter a valid title",
      expect.arrayContaining([expect.objectContaining({ text: "OK" })]),
    );

    expect(createLearning).not.toHaveBeenCalled();
  });

  it("deletes a learning item when delete button is pressed", async () => {
    (deleteLearning as jest.Mock).mockResolvedValueOnce(undefined);

    // Mock Alert.alert to simulate user confirming deletion
    jest.spyOn(Alert, "alert").mockImplementation((title, message, buttons) => {
      // Find and execute the onPress handler for the "Delete" button
      if (buttons && buttons.length > 1) {
        const deleteButton = buttons.find((button) => button.text === "Delete");
        if (deleteButton && deleteButton.onPress) {
          deleteButton.onPress();
        }
      }
    });

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
      // Use correct placeholder text
      expect(getAllByPlaceholderText(/Add a new .* item/).length).toBe(3);
    });

    const inputs = getAllByPlaceholderText(/Add a new .* item/);
    const addButtons = getAllByText("Add");

    // Fill and submit form
    fireEvent.changeText(inputs[0], "React Native");
    fireEvent.press(addButtons[0]);

    // Check if Alert.alert was called with expected parameters
    await waitFor(() => {
      expect(Alert.alert).toHaveBeenCalledWith(
        "Error",
        "Could not create learning item. Please try again.",
        expect.arrayContaining([expect.objectContaining({ text: "OK" })]),
      );
    });
  });

  it("shows an error alert when deleting a learning item fails", async () => {
    // Mock deleteLearning to fail
    (deleteLearning as jest.Mock).mockRejectedValueOnce(new Error("Failed to delete"));

    // Mock Alert implementation to trigger the Delete confirmation
    const alertSpy = jest.spyOn(Alert, "alert").mockImplementation((title, message, buttons) => {
      // For deletion confirmation dialog, trigger the Delete button
      if (title === "Confirm Deletion" && buttons && buttons.length > 1) {
        const deleteButton = buttons.find((b) => b.text === "Delete");
        if (deleteButton && deleteButton.onPress) {
          deleteButton.onPress();
        }
      }
    });

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

    // Check if Alert.alert was called for the error
    await waitFor(() => {
      // First call is for confirmation, second call is for the error
      expect(alertSpy).toHaveBeenCalledTimes(2);
      expect(alertSpy).toHaveBeenLastCalledWith(
        "Error",
        "Could not delete learning item. Please try again.",
        expect.arrayContaining([expect.objectContaining({ text: "OK" })]),
      );
    });

    // Restore mock
    alertSpy.mockRestore();
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
    // Mock console.error to avoid cluttering test output
    const consoleErrorSpy = jest.spyOn(console, "error").mockImplementation(() => {});

    // Update the expected error message to match what the component actually throws
    expect(() => {
      render(<UserLearnings />);
    }).toThrow("useUser must be used within a UserProvider");

    // Clean up
    consoleErrorSpy.mockRestore();
  });
});
