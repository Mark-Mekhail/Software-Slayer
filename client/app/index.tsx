import { createNativeStackNavigator } from '@react-navigation/native-stack';

import { UserProvider } from './common/UserContext';
import LoginScreen from './components/LoginScreen';
import RegisterScreen from './components/RegisterScreen';
import UserLearnings from './components/UserLearnings';

const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <UserProvider>
      <Stack.Navigator>
        <Stack.Screen name="Login" component={LoginScreen} />
        <Stack.Screen name="Register" component={RegisterScreen} />
        <Stack.Screen name="UserLearnings" component={UserLearnings} />
      </Stack.Navigator>
    </UserProvider>
  );
}
