# Software-Slayer

## Description

Software-Slayer is a full-stack application designed to manage user accounts and their learning items. The application allows users to register, log in, and manage their learning items categorized into different sections. The backend is built using Go, and the frontend is developed using React Native with Expo.

## Demo

A demo of the application can be found [here](https://youtube.com/shorts/Ndh9VpO3ayk).

## Technologies Used

### Backend

- **Go**: The backend is implemented in Go, providing RESTful APIs for user authentication and learning item management.
- **MySQL**: The application uses MySQL as the database to store user information and learning items.
- **Docker**: Docker is used to containerize the application, making it easy to deploy and manage.
- **Swagger**: Swagger is used for API documentation and testing.

### Frontend

- **React Native**: The frontend is built using React Native, allowing the application to run on both iOS and Android devices.
- **Expo**: Expo is used to streamline the development process and provide additional tools and services for building React Native applications.
- **Jest**: Jest is used for testing the frontend components.
- **ESLint**: ESLint is used to enforce code quality and style rules.
- **Prettier**: Prettier is used to format the code according to defined rules.

## Running the Application Locally

### Prerequisites

- Docker installed on your machine.
- Node.js and npm installed on your machine.
- Expo CLI installed on your machine.
- An Android emulator, iOS simulator, or a physical device with the Expo Go app installed.

### Steps

1. **Clone the repository**:

2. **Create secret files**:

   Create the following secret files in the `secrets` directory:

   - `mysql_password.txt`: Contains the MySQL user password.
   - `mysql_root_password.txt`: Contains the MySQL root password.
   - `jwt_secret.txt`: Contains the JWT secret key.

3. **Install frontend dependencies**:

   Navigate to the `client` directory and run the following command to install the dependencies:

   ```sh
   npm install
   ```

4. **Start the server application**:

   Navigate back to the root directory. Run the following command to start the server application. This will build the relevant Docker images and start the application containers. The backend API will be available at `http://localhost:8080`.

   ```sh
   ./scripts/server.sh
   ```

5. **Start the client application**:
   Run the following command to start the client application on the platform of your choice:

   ```sh
    ./scripts/client.sh <ios|android>
    ```
