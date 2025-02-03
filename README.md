# Software-Slayer

## Description

Software-Slayer is a full-stack application designed to manage user accounts and their learning items. The application allows users to register, log in, and manage their learning items categorized into different sections. The backend is built using Go, and the frontend is developed using React Native with Expo.

## Technologies Used

### Backend

- **Go**: The backend is implemented in Go, providing RESTful APIs for user authentication and learning item management.
- **MySQL**: The application uses MySQL as the database to store user information and learning items.
- **Docker**: Docker is used to containerize the application, making it easy to deploy and manage.
- **Swagger**: Swagger is used for API documentation and testing.

### Frontend

- **React Native**: The frontend is built using React Native, allowing the application to run on both iOS and Android devices.
- **Expo**: Expo is used to streamline the development process and provide additional tools and services for building React Native applications.

## Running the Application Locally

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Node.js and npm installed on your machine.

### Steps

1. **Clone the repository**:

   ```sh
   git clone https://github.com/your-username/software-slayer.git
   cd software-slayer
   ```

2. **Create secret files**:

   Create the following secret files in the `secrets` directory:

   - `mysql_password.txt`: Contains the MySQL user password.
   - `mysql_root_password.txt`: Contains the MySQL root password.
   - `jwt_secret.txt`: Contains the JWT secret key.

3. **Start the backend services**:

   Navigate to the root directory and run the following command to start the backend services using Docker Compose:

   ```sh
   docker-compose up
   ```

   This will start the MySQL database and the Go server.

4. **Install frontend dependencies**:

   Navigate to the `client` directory and run the following command to install the dependencies:

   ```sh
   npm install
   ```

5. **Start the frontend application**:

   Run the following command to start the Expo development server:

   ```sh
   npx expo start
   ```

   Follow the instructions in the terminal to open the application on an Android emulator, iOS simulator, or a physical device using the Expo Go app.

## Conclusion

You have successfully set up and run the Software-Slayer application locally. You can now explore the features and functionality of the application. If you encounter any issues or have any questions, feel free to open an issue on the GitHub repository.
