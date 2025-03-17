import { Alert } from "react-native";

// Mock Alert implementation that makes it easier to test in both formats
export const mockAlert = () => {
  const originalAlert = Alert.alert;

  // Mock Alert.alert to also call global.alert with a simplified message
  // This helps tests that expect global.alert to be called
  jest.spyOn(Alert, "alert").mockImplementation((title, message, buttons) => {
    // Call global.alert with combined message for tests that expect this format
    if (global.alert && typeof global.alert === "function") {
      global.alert(`${title}: ${message}`);
    }

    // For tests that check Alert.alert directly
    return originalAlert(title, message, buttons);
  });

  return () => {
    jest.restoreAllMocks();
  };
};

// Helper to set up global.alert if tests need it
export const setupGlobalAlert = () => {
  if (!global.alert) {
    global.alert = jest.fn();
  }
  return () => {
    delete global.alert;
  };
};
