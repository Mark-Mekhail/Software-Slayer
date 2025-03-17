import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { StatusBar } from "expo-status-bar";
import React from "react";
import { ActivityIndicator, View, Text } from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";

import { UserProvider, useUser } from "./common/UserContext";
import LoginScreen from "./components/LoginScreen";
import RegisterScreen from "./components/RegisterScreen";
import UserLearnings from "./components/UserLearnings";

// Define the stack navigator parameter list
type RootStackParamList = {
  Login: undefined;
  Register: undefined;
  UserLearnings: undefined;
};

// Create a stack navigator
const Stack = createNativeStackNavigator<RootStackParamList>();

/**
 * AuthNavigator handles navigation between authentication screens
 */
const AuthNavigator = () => (
  <Stack.Navigator
    initialRouteName="Login"
    screenOptions={{
      headerShown: false,
      contentStyle: { backgroundColor: "#f5f5f5" },
      animation: "slide_from_right",
    }}
  >
    <Stack.Screen name="Login" component={LoginScreen} />
    <Stack.Screen name="Register" component={RegisterScreen} />
  </Stack.Navigator>
);

/**
 * AppNavigator handles navigation for authenticated users
 */
const AppNavigator = () => (
  <Stack.Navigator
    screenOptions={{
      headerShown: false,
      contentStyle: { backgroundColor: "#f5f5f5" },
    }}
  >
    <Stack.Screen name="UserLearnings" component={UserLearnings} />
  </Stack.Navigator>
);

/**
 * Main navigation container that checks authentication status
 */
const RootNavigator = () => {
  const { user, isLoading } = useUser();

  // Show loading screen while checking auth status
  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" color="#007BFF" />
        <Text style={{ marginTop: 20, fontSize: 16 }}>Loading...</Text>
      </View>
    );
  }

  // Return navigation structure without NavigationContainer
  return user ? <AppNavigator /> : <AuthNavigator />;
};

/**
 * Root component that provides the user context and navigation
 */
export default function App() {
  return (
    <SafeAreaProvider>
      <StatusBar style="auto" />
      <UserProvider>
        <RootNavigator />
      </UserProvider>
    </SafeAreaProvider>
  );
}
