import { Text, View, StyleSheet, ScrollView } from 'react-native';
import { useContext, useEffect, useState } from 'react';

import { UserContext } from '../common/UserContext';
import { skillRequests } from '../requests/skillRequests';

export default function UserSkills() {
  const userContext = useContext(UserContext);
  if (!userContext) {
    throw new Error('UserContext is not set');
  }
  const { user, setUser } = userContext;

  const [skills, setSkills] = useState<string[]>([]);

  useEffect(() => {
    if (!user) {
      throw new Error('User is not set');
    }

    skillRequests.getSkills(user.token, user.id).then((topics) => {
      setSkills(topics);
    });
  }, []);

  return (
    <View style={styles.container}>
      {<Text style={styles.title}>Welcome, {user?.firstName}!</Text>}
      <ScrollView>
        {skills.map((skill, index) => (
          <Text key={index}>{skill}</Text>
        ))}
      </ScrollView>
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