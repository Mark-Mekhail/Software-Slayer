/* eslint-env jest */
jest.mock("@react-native-async-storage/async-storage", () => ({
  setItem: jest.fn(() => Promise.resolve()),
  getItem: jest.fn(() => Promise.resolve(null)),
  removeItem: jest.fn(() => Promise.resolve()),
  clear: jest.fn(() => Promise.resolve()),
  getAllKeys: jest.fn(() => Promise.resolve([])),
  multiGet: jest.fn(() => Promise.resolve([])),
  multiSet: jest.fn(() => Promise.resolve()),
  multiRemove: jest.fn(() => Promise.resolve()),
}));

// Mock React Native's Alert module with a more robust implementation
jest.mock("react-native/Libraries/Alert/Alert", () => {
  return {
    alert: jest.fn(),
  };
});

// Silence the warning about the native module being unavailable in tests
jest.mock("react-native/Libraries/LogBox/LogBox", () => ({
  ignoreLogs: jest.fn(),
}));

// Silence React Native and other warnings/errors during tests
const originalConsoleError = console.error;
const originalConsoleWarn = console.warn;

// Silent all console.error calls during tests unless they match certain patterns
console.error = (...args) => {
  // Allow specific errors to show through if needed
  const shouldShow = false;

  // Only show errors if explicitly requested
  if (shouldShow) {
    originalConsoleError(...args);
  }
};

// Silent React Navigation warnings
console.warn = (...args) => {
  // Allow specific warnings to show through if needed
  const shouldShow = false;

  // Only show warnings if explicitly requested
  if (shouldShow) {
    originalConsoleWarn(...args);
  }
};
