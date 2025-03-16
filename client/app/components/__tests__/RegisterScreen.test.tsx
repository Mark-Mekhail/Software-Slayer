import { render, fireEvent, waitFor } from "@testing-library/react-native";
import React from "react";
import { Alert } from "react-native";

import { createUser } from "../../requests/userRequests";
import RegisterScreen from "../RegisterScreen";

// Mock dependencies
jest.mock("../../requests/userRequests");
jest.spyOn(Alert, "alert").mockImplementation(() => {});

// Mock navigation
const mockNavigation = {
  navigate: jest.fn(),
};

describe("RegisterScreen", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("renders correctly with all form fields", () => {
    const { getByText, getByPlaceholderText } = render(
      <RegisterScreen navigation={mockNavigation} />,
    );

    expect(getByText("Create Account")).toBeTruthy();
    expect(getByPlaceholderText("Email")).toBeTruthy();
    expect(getByPlaceholderText("Username")).toBeTruthy();
    expect(getByPlaceholderText("First Name")).toBeTruthy();
    expect(getByPlaceholderText("Last Name")).toBeTruthy();
    expect(getByPlaceholderText("Password")).toBeTruthy();
    expect(getByText("Register")).toBeTruthy();
    expect(getByText("Already have an account? Log in")).toBeTruthy();
  });

  it("updates state when text inputs change", () => {
    const { getByPlaceholderText } = render(<RegisterScreen navigation={mockNavigation} />);

    const emailInput = getByPlaceholderText("Email");
    const usernameInput = getByPlaceholderText("Username");
    const firstNameInput = getByPlaceholderText("First Name");
    const lastNameInput = getByPlaceholderText("Last Name");
    const passwordInput = getByPlaceholderText("Password");

    fireEvent.changeText(emailInput, "test@example.com");
    fireEvent.changeText(usernameInput, "testuser");
    fireEvent.changeText(firstNameInput, "John");
    fireEvent.changeText(lastNameInput, "Doe");
    fireEvent.changeText(passwordInput, "password123");

    expect(emailInput.props.value).toBe("test@example.com");
    expect(usernameInput.props.value).toBe("testuser");
    expect(firstNameInput.props.value).toBe("John");
    expect(lastNameInput.props.value).toBe("Doe");
    expect(passwordInput.props.value).toBe("password123");
  });

  it("validates all fields are filled before submission", () => {
    const { getByText, getByPlaceholderText } = render(
      <RegisterScreen navigation={mockNavigation} />,
    );

    // Fill only some fields
    fireEvent.changeText(getByPlaceholderText("Email"), "test@example.com");
    fireEvent.changeText(getByPlaceholderText("Username"), "testuser");
    // Leave First Name, Last Name and Password empty

    // Try to register
    const registerButton = getByText("Register");
    fireEvent.press(registerButton);

    // Should show an error alert
    expect(Alert.alert).toHaveBeenCalledWith("Error", "Please fill out all fields.");
    expect(createUser).not.toHaveBeenCalled();
  });

  it("calls createUser API and navigates to Login on successful registration", async () => {
    (createUser as jest.Mock).mockResolvedValueOnce(undefined);

    const { getByText, getByPlaceholderText } = render(
      <RegisterScreen navigation={mockNavigation} />,
    );

    // Fill all fields
    fireEvent.changeText(getByPlaceholderText("Email"), "test@example.com");
    fireEvent.changeText(getByPlaceholderText("Username"), "testuser");
    fireEvent.changeText(getByPlaceholderText("First Name"), "John");
    fireEvent.changeText(getByPlaceholderText("Last Name"), "Doe");
    fireEvent.changeText(getByPlaceholderText("Password"), "password123");

    // Submit form
    const registerButton = getByText("Register");
    fireEvent.press(registerButton);

    await waitFor(() => {
      expect(createUser).toHaveBeenCalledWith(
        "test@example.com",
        "John",
        "Doe",
        "testuser",
        "password123",
      );
      expect(mockNavigation.navigate).toHaveBeenCalledWith("Login");
    });
  });

  it("handles registration failure correctly", async () => {
    (createUser as jest.Mock).mockRejectedValueOnce(new Error("Registration failed"));

    const { getByText, getByPlaceholderText } = render(
      <RegisterScreen navigation={mockNavigation} />,
    );

    // Fill all fields
    fireEvent.changeText(getByPlaceholderText("Email"), "test@example.com");
    fireEvent.changeText(getByPlaceholderText("Username"), "testuser");
    fireEvent.changeText(getByPlaceholderText("First Name"), "John");
    fireEvent.changeText(getByPlaceholderText("Last Name"), "Doe");
    fireEvent.changeText(getByPlaceholderText("Password"), "password123");

    // Submit form
    const registerButton = getByText("Register");
    fireEvent.press(registerButton);

    await waitFor(() => {
      expect(createUser).toHaveBeenCalled();
      expect(Alert.alert).toHaveBeenCalledWith("Error", "An error occurred. Please try again.");
      expect(mockNavigation.navigate).not.toHaveBeenCalled();
    });
  });

  it('navigates to Login screen when "Log in" link is clicked', () => {
    const { getByText } = render(<RegisterScreen navigation={mockNavigation} />);

    const loginLink = getByText("Already have an account? Log in");
    fireEvent.press(loginLink);

    expect(mockNavigation.navigate).toHaveBeenCalledWith("Login");
  });
});
