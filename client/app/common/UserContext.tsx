import { createContext, useState, ReactNode } from "react";

interface User {
  id: number;
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  token: string;
}

// UserContextType is an object that contains the user and a function to set the user.
interface UserContextType {
  user: User | null;
  setUser: (user: User | null) => void;
}

// UserContext is a context that provides the user and a function to set the user.
export const UserContext = createContext<UserContextType | undefined>(undefined);

// UserProvider is a component that provides the user and a function to set the user to its children.
export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);

  return <UserContext.Provider value={{ user, setUser }}>{children}</UserContext.Provider>;
};
