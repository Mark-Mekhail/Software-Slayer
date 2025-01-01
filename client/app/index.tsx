import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { UserProvider } from "./common/UserContext";
import LoginScreen from "./components/LoginScreen";
import RegisterScreen from "./components/RegisterScreen";
import UserSkills from "./components/UserSkills";

const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <UserProvider>
      <Stack.Navigator>
        <Stack.Screen name="Login" component={LoginScreen} />
        <Stack.Screen name="Register" component={RegisterScreen} />
        <Stack.Screen name="UserSkills" component={UserSkills} />
      </Stack.Navigator>
    </UserProvider>
  );
}
