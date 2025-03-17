import React, { useState } from "react";
import {
  Text,
  TextInput,
  TouchableOpacity,
  View,
  Alert,
  ActivityIndicator,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
} from "react-native";

import { useUser } from "../common/UserContext";
import { login, ApiError } from "../requests/userRequests";
import { formStyles } from "./shared/FormStyles";

interface LoginScreenProps {
  navigation: {
    navigate: (screen: string) => void;
  };
}

/**
 * LoginScreen component handles user authentication
 */
export default function LoginScreen({ navigation }: LoginScreenProps) {
  const { setUser } = useUser();
  const [identifier, setIdentifier] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [inputErrors, setInputErrors] = useState<{
    identifier?: string;
    password?: string;
  }>({});

  /**
   * Validates the login form inputs
   * @returns true if validation passes, false otherwise
   */
  const validateInputs = (): boolean => {
    const errors: { identifier?: string; password?: string } = {};
    let isValid = true;

    // Validate identifier
    if (!identifier.trim()) {
      errors.identifier = "Email or username is required";
      isValid = false;
    }

    // Validate password
    if (!password) {
      errors.password = "Password is required";
      isValid = false;
    }

    setInputErrors(errors);
    return isValid;
  };

  /**
   * Handles the login form submission
   */
  const handleLogin = async (): Promise<void> => {
    if (!validateInputs()) {
      return;
    }

    try {
      setIsLoading(true);
      const res = await login(identifier, password);

      setUser({
        id: res.user_info.id,
        email: res.user_info.email,
        username: res.user_info.username,
        firstName: res.user_info.first_name,
        lastName: res.user_info.last_name,
        token: res.token,
      });

      navigation.navigate("UserLearnings");
    } catch (error) {
      let errorMessage = "Incorrect email/username or password. Please try again.";
      if (error instanceof ApiError) {
        // Use specific error message if available
        errorMessage = error.message;
      }

      Alert.alert("Login Failed", errorMessage, [{ text: "OK" }]);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
    >
      <ScrollView contentContainerStyle={formStyles.scrollContainer}>
        <View style={formStyles.container}>
          <Text style={formStyles.title} testID="login-title">
            Welcome
          </Text>
          <Text style={formStyles.subtitle}>Sign in to continue to Software Slayer</Text>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Email or Username</Text>
            <TextInput
              style={[formStyles.input, inputErrors.identifier ? formStyles.inputError : null]}
              placeholder="Enter your email or username"
              value={identifier}
              onChangeText={(text) => {
                setIdentifier(text);
                if (inputErrors.identifier) {
                  setInputErrors((prev) => ({ ...prev, identifier: undefined }));
                }
              }}
              keyboardType="email-address"
              autoCapitalize="none"
              autoComplete="email"
              textContentType="emailAddress"
              accessibilityLabel="Email or username input"
              editable={!isLoading}
            />
            {inputErrors.identifier ? (
              <Text style={formStyles.errorText}>{inputErrors.identifier}</Text>
            ) : null}
          </View>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Password</Text>
            <TextInput
              style={[formStyles.input, inputErrors.password ? formStyles.inputError : null]}
              placeholder="Enter your password"
              value={password}
              onChangeText={(text) => {
                setPassword(text);
                if (inputErrors.password) {
                  setInputErrors((prev) => ({ ...prev, password: undefined }));
                }
              }}
              secureTextEntry
              autoComplete="password"
              textContentType="password"
              accessibilityLabel="Password input"
              editable={!isLoading}
            />
            {inputErrors.password ? (
              <Text style={formStyles.errorText}>{inputErrors.password}</Text>
            ) : null}
          </View>

          <TouchableOpacity
            style={[formStyles.button, isLoading ? formStyles.buttonDisabled : null]}
            onPress={() => void handleLogin()}
            disabled={isLoading}
            testID="login-button"
            accessibilityLabel="Login button"
            accessibilityRole="button"
          >
            {isLoading ? (
              <ActivityIndicator size="small" color="#fff" />
            ) : (
              <Text style={formStyles.buttonText}>Login</Text>
            )}
          </TouchableOpacity>

          <TouchableOpacity
            style={formStyles.linkButton}
            onPress={() => navigation.navigate("Register")}
            disabled={isLoading}
            accessibilityLabel="Create account button"
            accessibilityRole="button"
          >
            <Text style={formStyles.linkText}>Don&apos;t have an account? Create one</Text>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}
