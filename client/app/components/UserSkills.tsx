import { Text, View, StyleSheet, ScrollView, TouchableOpacity, TextInput } from 'react-native';
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
  const [skill, setSkill] = useState('');

  useEffect(() => {
    if (!user) {
      throw new Error('User is not set');
    }

    skillRequests.getSkills(user.token, user.id).then((topics) => {
      setSkills(topics);
    });
  }, []);

  const handleRemoveSkill = (topic: string) => {
    if (!user) {
      throw new Error('User is not set');
    }

    skillRequests.deleteSkill(user.token, topic).then(() => {
      alert('Skill removed');
      setSkills(skills.filter((t) => t !== topic));
    }).catch((error) => {
      alert('Failed to remove skill');
    });
  };

  const handleAddSkill = () => {
    if (!user) {
      throw new Error('User is not set');
    }

    skillRequests.createSkill(user.token, skill).then(() => {
      setSkills([...skills, skill]);
      setSkill('');
      alert('Skill added');
    }).catch((error) => {
      alert('Failed to add skill');
    });
  }

  return (
    <View style={styles.container}>
      {<Text style={styles.title}>Welcome, {user?.firstName}!</Text>}
      <ScrollView>
        {skills.map((topic, index) => (
          <View key={index} style={styles.skillContainer}>
            <Text>{topic}</Text>
            <TouchableOpacity onPress={() => handleRemoveSkill(topic)}>
              <Text style={styles.removeButton}>Remove</Text>
            </TouchableOpacity>
          </View>
        ))}
        <View>
          <TextInput placeholder="Enter new skill" value={skill} onChangeText={setSkill} />
          <TouchableOpacity onPress={() => handleAddSkill()}>
            <Text>Add Skill</Text>
          </TouchableOpacity>
        </View>
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
  skillContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  editButton: {
    color: 'blue',
  },
  removeButton: {
    color: 'red',
  },
});