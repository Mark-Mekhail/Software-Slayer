import React, { useEffect, useState, useCallback } from "react";
import {
  Text,
  View,
  StyleSheet,
  SectionList,
  TouchableOpacity,
  TextInput,
  ActivityIndicator,
  Alert,
  RefreshControl,
  KeyboardAvoidingView,
  Platform,
} from "react-native";

import { useUser } from "../common/UserContext";
import {
  LearningItem,
  getLearnings,
  getLearningCategories,
  createLearning,
  deleteLearning,
} from "../requests/learningRequests";

/**
 * Section data structure for grouped learning items
 */
interface LearningSection {
  title: string;
  data: LearningItem[];
}

/**
 * UserLearnings component displays and manages a user's learning items
 * organized by categories
 */
export default function UserLearnings() {
  const { user, logout } = useUser();
  const [learningCategories, setLearningCategories] = useState<string[]>([]);
  const [learnings, setLearnings] = useState<LearningSection[]>([]);
  const [newItems, setNewItems] = useState<Record<string, string>>({});
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [isRefreshing, setIsRefreshing] = useState<boolean>(false);
  const [isSubmitting, setIsSubmitting] = useState<Record<string, boolean>>({});
  const [error, setError] = useState<string | null>(null);

  /**
   * Fetches learning categories from the server
   */
  const fetchCategories = useCallback(async () => {
    try {
      const categories = await getLearningCategories();
      setLearningCategories(categories);
      setError(null);
    } catch (err) {
      console.error("Failed to fetch categories:", err);
      setError("Failed to fetch categories. Please try again later.");
    }
  }, []);

  /**
   * Fetches user's learning items from the server
   */
  const fetchLearningItems = useCallback(async () => {
    if (!user || learningCategories.length === 0) {
      return;
    }

    try {
      const learningItems = await getLearnings(user.id);

      // Group items by category
      const learningSections = learningCategories.map((category) => ({
        title: category,
        data: learningItems.filter((item) => item.category === category),
      }));

      setLearnings(learningSections);
      setError(null);
    } catch (err) {
      console.error("Failed to fetch learning items:", err);
      setError("Failed to fetch learning items. Please try again later.");
    } finally {
      setIsLoading(false);
      setIsRefreshing(false);
    }
  }, [user, learningCategories]);

  // Initial data loading
  useEffect(() => {
    void fetchCategories();
  }, [fetchCategories]);

  useEffect(() => {
    void fetchLearningItems();
  }, [fetchLearningItems]);

  /**
   * Handles manual refresh of the learning items list
   */
  const handleRefresh = useCallback(() => {
    setIsRefreshing(true);
    void fetchLearningItems();
  }, [fetchLearningItems]);

  /**
   * Creates a new learning item
   * @param title - The title of the learning item
   * @param category - The category of the learning item
   */
  const handleCreateItem = async (title: string, category: string) => {
    if (!user) {
      return;
    }

    try {
      setIsSubmitting((prev) => ({ ...prev, [category]: true }));
      await createLearning(user.token, title, category);

      // Refresh the learning items
      await fetchLearningItems();

      // Clear input field after successful creation
      setNewItems((prev) => ({ ...prev, [category]: "" }));
    } catch (err) {
      console.error("Error creating learning item:", err);
      Alert.alert("Error", "Could not create learning item. Please try again.", [{ text: "OK" }]);
    } finally {
      setIsSubmitting((prev) => ({ ...prev, [category]: false }));
    }
  };

  /**
   * Deletes a learning item
   * @param id - The ID of the learning item to delete
   */
  // eslint-disable-next-line @typescript-eslint/require-await
  const handleDeleteItem = async (id: number) => {
    if (!user) {
      return;
    }

    // Confirm deletion
    Alert.alert("Confirm Deletion", "Are you sure you want to delete this learning item?", [
      { text: "Cancel", style: "cancel" },
      {
        text: "Delete",
        style: "destructive",
        onPress: () => {
          void (async () => {
            try {
              await deleteLearning(user.token, id);

              // Update local state to reflect the deletion
              setLearnings((prev) =>
                prev.map((section) => ({
                  ...section,
                  data: section.data.filter((item) => item.id !== id),
                })),
              );
            } catch (err) {
              console.error("Error deleting learning item:", err);
              Alert.alert("Error", "Could not delete learning item. Please try again.", [
                { text: "OK" },
              ]);
            }
          })();
        },
      },
    ]);
  };

  /**
   * Handles input change for new learning item
   * @param category - The category for which input is changing
   * @param value - The new input value
   */
  const handleInputChange = (category: string, value: string) => {
    setNewItems((prev) => ({ ...prev, [category]: value }));
  };

  /**
   * Handles the add button click for a new learning item
   * @param category - The category for which to add an item
   */
  const handleAddClick = (category: string) => {
    const newItemTitle = newItems[category]?.trim();
    if (newItemTitle) {
      void handleCreateItem(newItemTitle, category);
    } else {
      Alert.alert("Input Error", "Please enter a valid title", [{ text: "OK" }]);
    }
  };

  /**
   * Handles user logout with confirmation
   */
  const handleLogout = () => {
    Alert.alert("Confirm Logout", "Are you sure you want to log out?", [
      { text: "Cancel", style: "cancel" },
      { text: "Log Out", style: "destructive", onPress: logout },
    ]);
  };

  // Render error state
  if (error) {
    return (
      <View style={styles.centeredContainer}>
        <Text style={styles.errorText}>{error}</Text>
        <TouchableOpacity style={styles.retryButton} onPress={() => void fetchLearningItems()}>
          <Text style={styles.retryButtonText}>Retry</Text>
        </TouchableOpacity>
      </View>
    );
  }

  // Render loading state
  if (isLoading && learnings.length === 0) {
    return (
      <View style={styles.centeredContainer}>
        <ActivityIndicator size="large" color="#007BFF" />
        <Text style={styles.loadingText}>Loading learning items...</Text>
      </View>
    );
  }

  return (
    <KeyboardAvoidingView
      style={styles.container}
      behavior={Platform.OS === "ios" ? "padding" : undefined}
      keyboardVerticalOffset={100}
    >
      <View style={styles.header}>
        <Text style={styles.title}>{user?.firstName}&apos;s Learning Lists</Text>
        <TouchableOpacity style={styles.logoutButton} onPress={handleLogout}>
          <Text style={styles.logoutText}>Logout</Text>
        </TouchableOpacity>
      </View>

      <SectionList
        sections={learnings}
        keyExtractor={(item) => item.id.toString()}
        renderItem={({ item }) => (
          <View style={styles.listItem}>
            <Text style={styles.itemTitle}>{item.title}</Text>
            <TouchableOpacity
              onPress={() => void handleDeleteItem(item.id)}
              hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
              accessibilityLabel={`Delete ${item.title}`}
              accessibilityRole="button"
            >
              <Text style={styles.deleteButton}>Delete</Text>
            </TouchableOpacity>
          </View>
        )}
        renderSectionHeader={({ section }) => (
          <Text style={styles.sectionHeader}>{section.title}</Text>
        )}
        renderSectionFooter={({ section }) => (
          <View style={styles.footer}>
            <TextInput
              style={styles.input}
              placeholder={`Add a new ${section.title} item`}
              value={newItems[section.title] || ""}
              onChangeText={(text) => handleInputChange(section.title, text)}
              editable={!isSubmitting[section.title]}
              returnKeyType="done"
              onSubmitEditing={() => handleAddClick(section.title)}
            />
            {isSubmitting[section.title] ? (
              <ActivityIndicator size="small" color="#007BFF" style={styles.addButtonSpinner} />
            ) : (
              <TouchableOpacity
                style={styles.addButton}
                onPress={() => handleAddClick(section.title)}
                disabled={!newItems[section.title]?.trim()}
              >
                <Text
                  style={[
                    styles.addButtonText,
                    !newItems[section.title]?.trim() && styles.addButtonTextDisabled,
                  ]}
                >
                  Add
                </Text>
              </TouchableOpacity>
            )}
          </View>
        )}
        stickySectionHeadersEnabled
        style={styles.list}
        refreshControl={<RefreshControl refreshing={isRefreshing} onRefresh={handleRefresh} />}
        ListEmptyComponent={
          <Text style={styles.emptyText}>
            No learning items yet. Add some using the fields below!
          </Text>
        }
      />
    </KeyboardAvoidingView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    width: "100%",
    alignItems: "center",
  },
  centeredContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    padding: 20,
  },
  header: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    width: "90%",
    paddingVertical: 15,
    marginTop: 10,
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
  },
  logoutButton: {
    padding: 8,
  },
  logoutText: {
    color: "#FF3B30",
    fontWeight: "600",
  },
  list: {
    width: "90%",
    flexDirection: "column",
  },
  sectionHeader: {
    fontSize: 18,
    fontWeight: "bold",
    marginVertical: 10,
    backgroundColor: "#f8f8f8",
    paddingVertical: 8,
    paddingHorizontal: 15,
    borderRadius: 5,
  },
  listItem: {
    paddingVertical: 12,
    paddingHorizontal: 15,
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    borderBottomWidth: 1,
    borderBottomColor: "#f0f0f0",
    backgroundColor: "#fff",
  },
  itemTitle: {
    fontSize: 16,
  },
  footer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    paddingVertical: 10,
    paddingHorizontal: 15,
    marginBottom: 20,
  },
  input: {
    flex: 3,
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 5,
    padding: 10,
    marginRight: 10,
    fontSize: 16,
  },
  addButton: {
    flex: 1,
    paddingVertical: 10,
    borderRadius: 5,
  },
  addButtonText: {
    color: "#007BFF",
    fontWeight: "bold",
    textAlign: "right",
  },
  addButtonTextDisabled: {
    color: "#80bdff",
  },
  addButtonSpinner: {
    flex: 1,
  },
  deleteButton: {
    color: "red",
    paddingVertical: 5,
    paddingHorizontal: 10,
    fontWeight: "500",
  },
  errorText: {
    color: "red",
    fontSize: 16,
    textAlign: "center",
    marginBottom: 20,
  },
  loadingText: {
    marginTop: 10,
    color: "#555",
    fontSize: 16,
  },
  retryButton: {
    marginTop: 15,
    paddingVertical: 10,
    paddingHorizontal: 20,
    backgroundColor: "#007BFF",
    borderRadius: 5,
  },
  retryButtonText: {
    color: "#fff",
    fontWeight: "bold",
  },
  emptyText: {
    textAlign: "center",
    color: "#777",
    marginVertical: 20,
    fontStyle: "italic",
  },
});
