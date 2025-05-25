# Software-Slayer

## Description

Software-Slayer is a full-stack learning management application designed to help developers track and organize their learning journey. Users can create accounts, authenticate securely, and manage personalized learning lists categorized by programming languages, technologies, concepts, and projects. The application features a robust Go backend with comprehensive testing and a modern React Native frontend with cross-platform support.

## Key Features

### Security & Authentication
- **JWT-based Authentication**: Secure token-based authentication system
- **Password Encryption**: Bcrypt hashing for secure password storage
- **Token Authorization**: Protected endpoints with middleware validation
- **Secure Secret Management**: File-based secret storage for production security

### Cross-Platform Frontend
- **React Native with Expo**: Native iOS and Android app development
- **TypeScript**: Type-safe development with comprehensive type definitions
- **Responsive Design**: Optimized for both mobile platforms
- **Context-based State Management**: Efficient user state management

### Robust Backend Architecture
- **Go REST API**: High-performance backend with clean architecture
- **MySQL Database**: Reliable data persistence with proper schema design
- **Swagger Documentation**: Auto-generated API documentation at `/swagger/`
- **Docker Containerization**: Easy deployment and development setup

### Learning Management
- **Categorized Learning Lists**: Organize learning items by Languages, Technologies, Concepts, Projects, and Other
- **CRUD Operations**: Full create, read, update, delete functionality
- **User-specific Data**: Secure user isolation and data privacy
- **Real-time Updates**: Immediate UI updates after operations

### Comprehensive Testing
- **Backend Testing**: Unit tests, integration tests, and fuzz testing
- **Frontend Testing**: Component tests with React Testing Library
- **Mock Services**: Comprehensive mocking for isolated testing
- **95%+ Test Coverage**: Extensive test coverage across all layers
- **Continuous Integration**: Automated testing with GitHub Actions

## Technologies Used

### Backend
- **Go 1.21+**: Modern Go with latest features
- **MySQL 8.0**: Robust relational database
- **JWT (golang-jwt)**: Secure authentication tokens
- **Bcrypt**: Industry-standard password hashing
- **Docker & Docker Compose**: Containerization and orchestration
- **Swagger/OpenAPI**: API documentation and testing
- **Testify & SQLMock**: Advanced testing frameworks

### Frontend
- **React Native**: Cross-platform mobile development
- **Expo SDK**: Streamlined development and deployment
- **TypeScript**: Type-safe JavaScript development
- **React Navigation**: Native navigation patterns
- **Jest & React Testing Library**: Comprehensive testing suite
- **ESLint & Prettier**: Code quality and formatting

### DevOps & Tooling
- **Docker Compose**: Multi-service orchestration
- **Shell Scripts**: Automated development workflows
- **Environment Variables**: Secure configuration management
- **Git**: Version control with feature branching
- **GitHub Actions**: CI/CD for automated testing

## Demo

A demo of the application can be found [here](https://youtube.com/shorts/Ndh9VpO3ayk).

## Running the Application Locally

### Prerequisites

- **Docker & Docker Compose**: For containerized services
- **Node.js 18+** and **npm**: For frontend development
- **Expo CLI**: For React Native development
- **Mobile Development**: Android emulator, iOS simulator, or physical device with Expo Go

### Quick Start

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Mark-Mekhail/Software-Slayer.git
   cd Software-Slayer
   ```

2. **Create secret files**:
   Create the following files in the `secrets/` directory:
   ```bash
   mkdir -p secrets
   echo "mysql_password" > secrets/mysql_password.txt
   echo "mysql_root_password" > secrets/mysql_root_password.txt
   echo "jwt_secret_key" > secrets/jwt_secret.txt
   ```

3. **Install frontend dependencies**:
   ```bash
   cd client
   npm install
   cd ..
   ```

4. **Start the backend services**:
   ```bash
   ./scripts/server.sh
   ```
   Backend API will be available at `http://localhost:8080`
   Swagger documentation at `http://localhost:8080/swagger/`

5. **Start the mobile application**:
   ```bash
   ./scripts/client.sh ios     # For iOS
   ./scripts/client.sh android # For Android
   ```

## Testing

### Backend Testing
```bash
cd server/app/src/go
go test ./...                    # Run all tests
go test -v ./...                 # Verbose output
go test -cover ./...             # Coverage report
go test -fuzz=FuzzTest ./...     # Fuzz testing
```

### Frontend Testing
```bash
cd client
npm test                         # Run test suite
npm run test:watch               # Watch mode
npm run test:coverage            # Coverage report
```

## API Documentation

The API is fully documented using Swagger/OpenAPI specification. Once the server is running, visit:
- **Swagger UI**: `http://localhost:8080/swagger/`
- **Interactive API Testing**: Available through Swagger interface

### Key Endpoints
- `POST /user` - User registration
- `POST /login` - User authentication
- `GET /user?current=true` - Get current user info
- `POST /learning` - Create learning item
- `GET /learning/{user_id}` - Get user's learning items
- `DELETE /learning/{id}` - Delete learning item
- `GET /learning/categories` - Get available categories

## Architecture Highlights

### Security Architecture
- **JWT Token Management**: Secure token generation and validation
- **Password Security**: Bcrypt hashing with salt rounds
- **Authorization Middleware**: Protected endpoint access control
- **Input Validation**: Comprehensive request validation and sanitization

### Database Design
- **Normalized Schema**: Efficient relational database design
- **User Isolation**: Secure data separation between users

### Testing Strategy
- **Unit Testing**: Individual component and function testing
- **Integration Testing**: End-to-end API testing with database
- **Fuzz Testing**: Security and edge case validation
- **Component Testing**: React Native component isolation testing
- **Mock Services**: Comprehensive service mocking for reliable tests

### Performance Optimizations
- **Context Timeouts**: Request timeout management
- **Graceful Shutdown**: Clean service termination
- **Connection Pooling**: Database connection optimization
- **Efficient Queries**: Optimized database operations

## Development

### Code Quality
- **TypeScript**: Type safety across frontend
- **ESLint & Prettier**: Automated code formatting and linting
- **Go Modules**: Dependency management

### Environment Configuration
- **Docker Environment**: Isolated development environment
- **Environment Variables**: Secure configuration management
- **Secret Management**: File-based secrets for production security
- **Cross-platform Support**: Works on macOS, Linux, and Windows
