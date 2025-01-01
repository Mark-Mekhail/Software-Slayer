import { Text, View, StyleSheet, SectionList, TouchableOpacity, TextInput } from 'react-native';
import { useContext, useEffect, useState } from 'react';

import { UserContext } from '../common/UserContext';
import { LearningItem, getLearnings, getLearningCategories, createLearning, deleteLearning } from '../requests/learningRequests';

interface LearningSection {
  title: string;
  data: LearningItem[];
}

export default function UserLearnings() {
  const userContext = useContext(UserContext);
  if (!userContext) {
    throw new Error('UserContext is not set');
  }
  const { user } = userContext;

  const [learningCategories, setLearningCategories] = useState<string[]>([]);
  const [learnings, setLearnings] = useState<LearningSection[]>([]);
  const [newItems, setNewItems] = useState<Record<string, string>>({}); // Track new item input for each category

  useEffect(() => {
    getLearningCategories().then(setLearningCategories);
  }, []);

  useEffect(() => {
    if (!user) {
      return;
    }

    getLearnings(user.id).then((learningItems) => {
      const learningSections = learningCategories.map((category) => ({
        title: category,
        data: learningItems.filter((l) => l.category === category),
      }));
      setLearnings(learningSections);
    });
  }, [user, learningCategories]);

  const handleCreateItem = (title: string, category: string) => {
    if (!user) {
      return;
    }

    createLearning(user.token, title, category)
      .then(() => {
        // Make this more efficient by only fetching the new item
        getLearnings(user.id).then((learningItems) => {
          const learningSections = learningCategories.map((category) => ({
            title: category,
            data: learningItems.filter((l) => l.category === category),
          }));
          setLearnings(learningSections);
        });
      })
      .catch(() => {
        alert('Error: Could not create learning item');
      });
  };

  const handleDeleteItem = (id: number) => {
    if (!user) {
      return;
    }

    deleteLearning(user.token, id)
      .then(() => {
        setLearnings((prev) =>
          prev.map((section) => ({
            ...section,
            data: section.data.filter((item) => item.id !== id),
          }))
        );
      })
      .catch(() => {
        alert('Error: Could not delete learning item');
      });
  };

  const handleInputChange = (category: string, value: string) => {
    setNewItems((prev) => ({ ...prev, [category]: value }));
  };

  const handleAddClick = (category: string) => {
    const newItemTitle = newItems[category]?.trim();
    if (newItemTitle) {
      handleCreateItem(newItemTitle, category);
      setNewItems((prev) => ({ ...prev, [category]: '' })); // Clear input after adding
    } else {
      alert('Please enter a valid title');
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>{user?.firstName}'s Learning Lists</Text>
      <SectionList
        sections={learnings}
        keyExtractor={(item) => item.id.toString()}
        renderItem={({ item }) => (
          <View style={styles.listItem}>
            <Text>{item.title}</Text>
            <TouchableOpacity onPress={() => handleDeleteItem(item.id)}>
              <Text style={styles.deleteButton}>Delete</Text>
            </TouchableOpacity>
          </View>
        )}
        renderSectionHeader={({ section }) => (
          <Text style={styles.header}>{section.title}</Text>
        )}
        renderSectionFooter={({ section }) => (
          <View style={styles.footer}>
            <TextInput
              style={styles.input}
              placeholder="Enter learning item title"
              value={newItems[section.title] || ''}
              onChangeText={(text) => handleInputChange(section.title, text)}
            />
            <TouchableOpacity
              style={styles.addButton}
              onPress={() => handleAddClick(section.title)}
            >
              <Text style={styles.addButtonText}>Add</Text>
            </TouchableOpacity>
          </View>
        )}
        style={styles.list}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    flexDirection: 'column',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    margin: 20,
  },
  list: {
    width: '90%',
    display: 'flex',
    flexDirection: 'column',
  },
  header: {
    fontSize: 18,
    fontWeight: 'bold',
    marginVertical: 10,
  },
  listItem: {
    paddingVertical: 10,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  footer: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingVertical: 10,
  },
  input: {
    flex: 1,
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
    padding: 5,
    marginRight: 10,
  },
  addButton: {
    backgroundColor: '#007BFF',
    paddingHorizontal: 15,
    paddingVertical: 10,
    borderRadius: 5,
  },
  addButtonText: {
    color: '#fff',
    fontWeight: 'bold',
  },
  deleteButton: {
    color: 'red',
  },
});
