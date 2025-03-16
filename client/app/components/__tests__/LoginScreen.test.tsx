import { render, fireEvent, waitFor } from "@testing-library/react-native";
import React from "react";
import { Alert } from "react-native";

import { UserContext } from "../../common/UserContext";
import { login } from "../../requests/userRequests";
import LoginScreen from "../LoginScreen";

// Mock dependencies
jest.mock("../../requests/userRequests");
jest.spyOn(Alert, "alert").mockImplementation(() => {});

// Mock navigation
const mockNavigation = {
  navigate: jest.fn(),
};

// Mock user context
const mockSetUser = jest.fn();
const mockUserContext = {
  user: null,
  setUser: mockSetUser,
};

describe("LoginScreen", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("renders correctly with all elements", () => {
    const { getByText, getByPlaceholderText, getByTestId } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    // Check if all elements are rendered
    expect(getByTestId("login-title")).toBeTruthy();
    expect(getByPlaceholderText("Enter your email or username")).toBeTruthy();
    expect(getByPlaceholderText("Enter your password")).toBeTruthy();
    expect(getByTestId("login-button")).toBeTruthy();
    expect(getByText("Don't have an account? Create one")).toBeTruthy();
  });

  it("updates state when text inputs change", () => {
    const { getByPlaceholderText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const identifierInput = getByPlaceholderText("Enter your email or username");
    const passwordInput = getByPlaceholderText("Enter your password");

    fireEvent.changeText(identifierInput, "testuser");
    fireEvent.changeText(passwordInput, "password123");

    expect(identifierInput.props.value).toBe("testuser");
    expect(passwordInput.props.value).toBe("password123");
  });

  it("validates inputs and shows alert when fields are empty", () => {
    const { getByTestId } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const loginButton = getByTestId("login-button");
    fireEvent.press(loginButton);

    expect(Alert.alert).toHaveBeenCalledWith(
      "Error",
      "Please enter both email/username and password.",
    );
    expect(login).not.toHaveBeenCalled();
  });

  it("calls login API and navigates on successful login", async () => {
    const mockResponse = {
      token: "test-token",
      user_info: {
        id: 1,
        email: "test@example.com",
        username: "testuser",
        first_name: "John",
        last_name: "Doe",
      },
    };

    (login as jest.Mock).mockResolvedValueOnce(mockResponse);

    const { getByTestId, getByPlaceholderText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const identifierInput = getByPlaceholderText("Enter your email or username");
    const passwordInput = getByPlaceholderText("Enter your password");
    const loginButton = getByTestId("login-button");

    fireEvent.changeText(identifierInput, "testuser");
    fireEvent.changeText(passwordInput, "password123");
    fireEvent.press(loginButton);

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
    (login as jest.Mock).mockRejectedValueOnce(new Error("Login failed"));

    const { getByTestId, getByPlaceholderText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const identifierInput = getByPlaceholderText("Enter your email or username");
    const passwordInput = getByPlaceholderText("Enter your password");
    const loginButton = getByTestId("login-button");

    fireEvent.changeText(identifierInput, "testuser");
    fireEvent.changeText(passwordInput, "password123");
    fireEvent.press(loginButton);

    await waitFor(() => {
      expect(login).toHaveBeenCalledWith("testuser", "password123");
      expect(Alert.alert).toHaveBeenCalledWith("Error", "An error occurred. Please try again.");
      expect(mockNavigation.navigate).not.toHaveBeenCalled();
    });
  });

  it('navigates to Register screen when "Create one" link is clicked', () => {
    const { getByText } = render(
      <UserContext.Provider value={mockUserContext}>
        <LoginScreen navigation={mockNavigation} />
      </UserContext.Provider>,
    );

    const registerLink = getByText("Don't have an account? Create one");
    fireEvent.press(registerLink);

    expect(mockNavigation.navigate).toHaveBeenCalledWith("Register");
  });

  it("throws an error when UserContext is not provided", () => {
    // Spy on console.error to suppress error messages in test output
    jest.spyOn(console, "error").mockImplementation(() => {});

    expect(() => {
      render(<LoginScreen navigation={mockNavigation} />);
    }).toThrow("UserContext is not set");

    // Restore console.error
    (console.error as jest.Mock).mockRestore();
  });
});
