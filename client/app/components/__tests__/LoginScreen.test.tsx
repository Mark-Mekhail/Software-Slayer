import { render, fireEvent, waitFor } from "@testing-library/react-native";
import React from "react";
import { Alert } from "react-native";

import { UserContext } from "../../common/UserContext";
import { login, ApiError } from "../../requests/userRequests";
import LoginScreen from "../LoginScreen";

// Mock the UserContext
const mockSetUser = jest.fn();
const mockUserContext = {
  user: null,
  setUser: mockSetUser,
  isLoading: false,
  logout: jest.fn(),
};

// Mock the login API call
jest.mock("../../requests/userRequests", () => ({
  login: jest.fn(),
  ApiError: class ApiError extends Error {
    statusCode: number;
    constructor(message: string, statusCode: number) {
      super(message);
      this.name = "ApiError";
      this.statusCode = statusCode;
    }
  },
}));

// Mock navigation
const mockNavigation = {
  navigate: jest.fn(),
};

describe("LoginScreen", () => {
  beforeEach(() => {
    jest.clearAllMocks();
    // Fix: Use jest.spyOn instead to properly mock console.error
    jest.spyOn(console, "error").mockImplementation(() => {});
  });

  afterEach(() => {
    // Fix: Restore the mock correctly
    jest.restoreAllMocks();
  });

  it("renders correctly with all required elements", () => {
    const { getByText, getByPlaceholderText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    expect(getByText("Welcome Back")).toBeTruthy();
    expect(getByPlaceholderText("Enter your email or username")).toBeTruthy();
    expect(getByPlaceholderText("Enter your password")).toBeTruthy();
    expect(getByText("Login")).toBeTruthy();
    expect(getByText("Don't have an account? Create one")).toBeTruthy();
  });

  it("validates inputs and shows error text when fields are empty", () => {
    const { getByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const loginButton = getByText("Login");
    fireEvent.press(loginButton);

    // Should show validation error messages
    expect(getByText("Email or username is required")).toBeTruthy();
    expect(getByText("Password is required")).toBeTruthy();
  });

  it("navigates to Register screen when sign up link is pressed", () => {
    const { getByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const registerLink = getByText("Don't have an account? Create one");
    fireEvent.press(registerLink);

    expect(mockNavigation.navigate).toHaveBeenCalledWith("Register");
  });

  it("submits form with correct credentials and navigates on success", async () => {
    (login as jest.Mock).mockResolvedValueOnce({
      token: "test-token",
      user_info: {
        id: 1,
        email: "test@example.com",
        username: "testuser",
        first_name: "John",
        last_name: "Doe",
      },
    });

    const { getByPlaceholderText, getByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const usernameInput = getByPlaceholderText("Enter your email or username");
    const passwordInput = getByPlaceholderText("Enter your password");
    const loginButton = getByText("Login");

    // Fill in fields
    fireEvent.changeText(usernameInput, "testuser");
    fireEvent.changeText(passwordInput, "password123");

    // Submit form
    fireEvent.press(loginButton);

    // Check if API was called with correct params
    await waitFor(() => {
      expect(login).toHaveBeenCalledWith("testuser", "password123");
      expect(mockSetUser).toHaveBeenCalledWith({
        id: 1,
        email: "test@example.com",
        username: "testuser",
        firstName: "John",
        lastName: "Doe",
        token: "test-token",
      });
      expect(mockNavigation.navigate).toHaveBeenCalledWith("UserLearnings");
    });
  });

  it("handles login failure correctly", async () => {
    (login as jest.Mock).mockRejectedValueOnce(
      new ApiError("Incorrect email/username or password", 401),
    );

    // Mock Alert.alert to capture what is actually called
    const alertMock = jest.fn();
    jest.spyOn(Alert, "alert").mockImplementation(alertMock);

    const { getByPlaceholderText, getByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const usernameInput = getByPlaceholderText("Enter your email or username");
    const passwordInput = getByPlaceholderText("Enter your password");
    const loginButton = getByText("Login");

    // Fill in fields
    fireEvent.changeText(usernameInput, "testuser");
    fireEvent.changeText(passwordInput, "password123");

    // Submit form
    fireEvent.press(loginButton);

    await waitFor(() => {
      expect(login).toHaveBeenCalledWith("testuser", "password123");

      // Use a more flexible assertion that matches part of the message
      expect(alertMock).toHaveBeenCalledWith(
        "Login Failed",
        expect.stringContaining("Incorrect email/username or password"),
        expect.anything(),
      );

      expect(mockNavigation.navigate).not.toHaveBeenCalled();
    });
  });

  it("throws an error when UserContext is not provided", () => {
    // Mock console.error to avoid cluttering test output
    jest.spyOn(console, "error").mockImplementation(() => {});

    expect(() => {
      render(<LoginScreen navigation={mockNavigation} />);
    }).toThrow("useUser must be used within a UserProvider");

    // Restore console.error
    (console.error as jest.Mock).mockRestore();
  });
});
