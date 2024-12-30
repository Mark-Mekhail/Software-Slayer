import { Text, View, StyleSheet } from 'react-native';
import { useContext } from 'react';

import { UserContext } from '../common/UserContext';

export default function HomeScreen() {
  const userContext = useContext(UserContext);
  if (!userContext) {
    throw new Error('UserContext is not set');
  }
  const { user, setUser } = userContext;

  console.log('user:', user);

  return (
    <View style={styles.container}>
      {<Text style={styles.title}>Welcome, {user?.firstName}!</Text>}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
  },
});