import AsyncStorage from "@react-native-async-storage/async-storage";
import React, { createContext, useState, ReactNode, useEffect } from "react";

/**
 * User data structure
 */
export interface User {
  id: number;
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  token: string;
}

/**
 * User context state and methods
 */
interface UserContextType {
  user: User | null;
  setUser: (user: User | null) => void;
  isLoading: boolean;
  logout: () => void;
}

/**
 * Provider props
 */
interface UserProviderProps {
  children: ReactNode;
}

// Create the context with default undefined value
export const UserContext = createContext<UserContextType | undefined>(undefined);

// User data persistence key
const USER_STORAGE_KEY = "@SoftwareSlayer:user";

/**
 * UserProvider component for managing user authentication state
 */
export const UserProvider = ({ children }: UserProviderProps) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  // Load saved user on initial mount
  useEffect(() => {
    const loadUser = async () => {
      try {
        const savedUserData = await AsyncStorage.getItem(USER_STORAGE_KEY);
        if (savedUserData) {
          setUser(JSON.parse(savedUserData) as User);
        }
      } catch (error) {
        console.error("Failed to load user data:", error);
      } finally {
        setIsLoading(false);
      }
    };

    void loadUser();
  }, []);

  // Save user data when it changes
  useEffect(() => {
    const saveUser = async () => {
      try {
        if (user) {
          await AsyncStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user));
        } else {
          await AsyncStorage.removeItem(USER_STORAGE_KEY);
        }
      } catch (error) {
        console.error("Failed to save user data:", error);
      }
    };

    void saveUser();
  }, [user]);

  /**
   * Logs out the current user
   */
  // eslint-disable-next-line @typescript-eslint/require-await
  const logout = async () => {
    setUser(null);
  };

  // Context value with state and methods
  const value = {
    user,
    setUser,
    isLoading,
    logout,
  };

  return <UserContext.Provider value={value}>{children}</UserContext.Provider>;
};

/**
 * Custom hook for accessing the user context
 * @throws {Error} If used outside of UserProvider
 */
export const useUser = (): UserContextType => {
  const context = React.useContext(UserContext);
  if (context === undefined) {
    throw new Error("useUser must be used within a UserProvider");
  }
  return context;
};
