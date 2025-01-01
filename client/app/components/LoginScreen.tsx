import { useState, useContext } from 'react';
import { StyleSheet, Text, TextInput, TouchableOpacity, View, Alert } from 'react-native';

import { userRequests } from '../requests/userRequests';
import { UserContext } from '../common/UserContext';

interface LoginScreenProps {
  navigation: {
    navigate: (screen: string) => void;
  }
};

export default function LoginScreen({ navigation }: LoginScreenProps) {
  const userContext = useContext(UserContext);
  if (!userContext) {
    throw new Error('UserContext is not set');
  }

  const { setUser } = userContext;

  const [identifier, setIdentifier] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  const handleLogin = (): void => {
    if (!identifier || !password) {
      Alert.alert('Error', 'Please enter both email/username and password.');
      return;
    }

    userRequests.login(identifier, password)
      .then((res: any) => {
        setUser( {
          id: res.user_info.id,
          email: res.user_info.email,
          username: res.user_info.username,
          firstName: res.user_info.first_name,
          lastName: res.user_info.last_name,
          token: res.token,
        });
        navigation.navigate('UserSkills');
      })
      .catch(() => {
        Alert.alert('Error', 'An error occurred. Please try again.');
      });
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Login</Text>

      <TextInput
        style={styles.input}
        placeholder="Enter your email or username"
        value={identifier}
        onChangeText={setIdentifier}
        keyboardType="email-address"
        autoCapitalize="none"
      />

      <TextInput
        style={styles.input}
        placeholder="Enter your password"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />

      <TouchableOpacity style={styles.button} onPress={handleLogin}>
        <Text style={styles.buttonText}>Login</Text>
      </TouchableOpacity>

      <TouchableOpacity
        style={styles.linkButton}
        onPress={() =>
          navigation.navigate('Register')
        }
      >
        <Text style={styles.linkText}>Don't have an account? Create one</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#f5f5f5',
    padding: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  input: {
    width: '100%',
    padding: 15,
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
    marginBottom: 15,
    backgroundColor: '#fff',
  },
  button: {
    width: '100%',
    padding: 15,
    backgroundColor: '#007BFF',
    borderRadius: 5,
    alignItems: 'center',
  },
  buttonText: {
    color: '#fff',
    fontWeight: 'bold',
    fontSize: 16,
  },
  linkButton: {
    marginTop: 15,
  },
  linkText: {
    color: '#007BFF',
    textDecorationLine: 'underline',
  },
});