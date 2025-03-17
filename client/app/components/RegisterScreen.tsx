import React, { useState } from "react";
import {
  Text,
  TextInput,
  TouchableOpacity,
  View,
  Alert,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  ActivityIndicator,
} from "react-native";

import { createUser, ApiError } from "../requests/userRequests";
import { formStyles } from "./shared/FormStyles";

interface RegisterScreenProps {
  navigation: {
    navigate: (screen: string) => void;
  };
}

/**
 * RegisterScreen component handles new user registration
 */
export default function RegisterScreen({ navigation }: RegisterScreenProps) {
  const [email, setEmail] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [firstName, setFirstName] = useState<string>("");
  const [lastName, setLastName] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [confirmPassword, setConfirmPassword] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [inputErrors, setInputErrors] = useState<{
    email?: string;
    username?: string;
    firstName?: string;
    lastName?: string;
    password?: string;
    confirmPassword?: string;
  }>({});

  /**
   * Validates the registration form inputs
   * @returns true if validation passes, false otherwise
   */
  const validateInputs = (): boolean => {
    const errors: Record<string, string> = {};
    let isValid = true;

    // Email validation
    if (!email.trim()) {
      errors.email = "Email is required";
      isValid = false;
    }

    // Username validation
    if (!username.trim()) {
      errors.username = "Username is required";
      isValid = false;
    }

    // Name validation
    if (!firstName.trim()) {
      errors.firstName = "First name is required";
      isValid = false;
    }

    if (!lastName.trim()) {
      errors.lastName = "Last name is required";
      isValid = false;
    }

    // Password validation
    if (!password) {
      errors.password = "Password is required";
      isValid = false;
    }

    // Confirm password
    if (password !== confirmPassword) {
      errors.confirmPassword = "Passwords don't match";
      isValid = false;
    }

    setInputErrors(errors);
    return isValid;
  };

  /**
   * Handles user registration
   */
  const handleRegister = async () => {
    if (!validateInputs()) {
      return;
    }

    try {
      setIsLoading(true);
      await createUser(email, firstName, lastName, username, password);

      Alert.alert("Registration Successful", "Your account has been created. Please log in.", [
        { text: "OK", onPress: () => navigation.navigate("Login") },
      ]);
    } catch (error) {
      console.error("Registration error:", error);

      let errorMessage = "An error occurred during registration. Please try again.";
      if (error instanceof ApiError) {
        errorMessage = error.message;
      }

      Alert.alert("Registration Failed", errorMessage, [{ text: "OK" }]);
    } finally {
      setIsLoading(false);
    }
  };

  /**
   * Handles input change and clears associated error
   */
  const handleInputChange = (
    field: keyof typeof inputErrors,
    value: string,
    setter: (value: string) => void,
  ) => {
    setter(value);
    if (inputErrors[field]) {
      setInputErrors((prev) => ({ ...prev, [field]: undefined }));
    }
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
    >
      <ScrollView contentContainerStyle={formStyles.scrollContainer}>
        <View style={formStyles.container}>
          <Text style={formStyles.title}>Create Account</Text>
          <Text style={formStyles.subtitle}>Join Software Slayer today</Text>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Email</Text>
            <TextInput
              style={[formStyles.input, inputErrors.email ? formStyles.inputError : null]}
              placeholder="Enter your email"
              value={email}
              onChangeText={(text) => handleInputChange("email", text, setEmail)}
              keyboardType="email-address"
              autoCapitalize="none"
              autoComplete="email"
              textContentType="emailAddress"
              editable={!isLoading}
            />
            {inputErrors.email && <Text style={formStyles.errorText}>{inputErrors.email}</Text>}
          </View>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Username</Text>
            <TextInput
              style={[formStyles.input, inputErrors.username ? formStyles.inputError : null]}
              placeholder="Choose a username"
              value={username}
              onChangeText={(text) => handleInputChange("username", text, setUsername)}
              autoCapitalize="none"
              editable={!isLoading}
            />
            {inputErrors.username && (
              <Text style={formStyles.errorText}>{inputErrors.username}</Text>
            )}
          </View>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>First Name</Text>
            <TextInput
              style={[formStyles.input, inputErrors.firstName ? formStyles.inputError : null]}
              placeholder="Enter your first name"
              value={firstName}
              onChangeText={(text) => handleInputChange("firstName", text, setFirstName)}
              autoComplete="name-given"
              textContentType="givenName"
              editable={!isLoading}
            />
            {inputErrors.firstName && (
              <Text style={formStyles.errorText}>{inputErrors.firstName}</Text>
            )}
          </View>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Last Name</Text>
            <TextInput
              style={[formStyles.input, inputErrors.lastName ? formStyles.inputError : null]}
              placeholder="Enter your last name"
              value={lastName}
              onChangeText={(text) => handleInputChange("lastName", text, setLastName)}
              autoComplete="name-family"
              textContentType="familyName"
              editable={!isLoading}
            />
            {inputErrors.lastName && (
              <Text style={formStyles.errorText}>{inputErrors.lastName}</Text>
            )}
          </View>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Password</Text>
            <TextInput
              style={[formStyles.input, inputErrors.password ? formStyles.inputError : null]}
              placeholder="Create a password"
              value={password}
              onChangeText={(text) => handleInputChange("password", text, setPassword)}
              secureTextEntry
              textContentType="newPassword"
              editable={!isLoading}
            />
            {inputErrors.password && (
              <Text style={formStyles.errorText}>{inputErrors.password}</Text>
            )}
          </View>

          <View style={formStyles.formGroup}>
            <Text style={formStyles.label}>Confirm Password</Text>
            <TextInput
              style={[formStyles.input, inputErrors.confirmPassword ? formStyles.inputError : null]}
              placeholder="Confirm your password"
              value={confirmPassword}
              onChangeText={(text) =>
                handleInputChange("confirmPassword", text, setConfirmPassword)
              }
              secureTextEntry
              textContentType="newPassword"
              editable={!isLoading}
            />
            {inputErrors.confirmPassword && (
              <Text style={formStyles.errorText}>{inputErrors.confirmPassword}</Text>
            )}
          </View>

          <TouchableOpacity
            style={[formStyles.button, isLoading ? formStyles.buttonDisabled : null]}
            onPress={() => void handleRegister()}
            disabled={isLoading}
            accessibilityLabel="Register button"
            accessibilityRole="button"
            testID="register-button"
          >
            {isLoading ? (
              <ActivityIndicator size="small" color="#fff" />
            ) : (
              <Text style={formStyles.buttonText}>Register</Text>
            )}
          </TouchableOpacity>

          <TouchableOpacity
            style={formStyles.linkButton}
            onPress={() => navigation.navigate("Login")}
            disabled={isLoading}
            accessibilityLabel="Login link"
            accessibilityRole="button"
          >
            <Text style={formStyles.linkText}>Already have an account? Log in</Text>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}
