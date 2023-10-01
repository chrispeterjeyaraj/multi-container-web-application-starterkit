# Multi-Container Web Application with Docker

This repository contains a multi-container web application built using Docker, with separate containers for the server, UI, and MongoDB for data storage. The server-side code is written in Golang, and the user interface is developed using React. This setup allows you to easily deploy and manage the various components of your web application as isolated containers.

## Features

- **Server Container (Golang):** The server-side code is written in Golang, providing a fast and efficient backend for your application. It handles data processing and communication with the MongoDB container.

- **UI Container (React):** The user interface is built using React, a popular JavaScript library for creating dynamic and responsive web applications. The UI container serves as the frontend of your application.

- **MongoDB Container:** MongoDB is used as the database for storing and retrieving data. The MongoDB container ensures data persistence and can be easily configured to suit your application's needs.

- **Docker Compose:** The `docker-compose.yml` file is included to orchestrate and manage the deployment of these containers. Docker Compose simplifies the process of starting and connecting multiple containers.

## Getting Started

To get started with this multi-container web application, follow these steps:

1. **Clone the Repository:** Clone this repository to your local machine using `git clone`.

2. **Set Up Environment Variables:** Create a `.env` file in the project root directory and configure environment variables such as database connection settings, server configurations, and any secret keys. For initial setup, i have .env file already. But you can go ahead and make changes as per your requirement.

3. **Build and Start Containers:** Use Docker Compose to build and start the containers defined in the `docker-compose.yml` file. Run the following command:

   ```bash
   docker-compose up --build
   ```

   This command will start the server, UI, and MongoDB containers.

4. **Access the Application:** Once the containers are up and running, you can access the web application by opening your browser and navigating to `http://localhost:3000`.

5. **Customize and Extend:** Customize the server and UI code to suit your project's requirements. You can modify the Golang server code in the `server/` directory and the React UI code in the `ui/` directory.

6. **Scaling and Deployment:** Docker Compose makes it easy to scale your application and deploy it to various environments. You can configure additional settings for production, staging, or development environments as needed.

## Contributing

We welcome contributions to this project. If you find a bug, have a feature request, or want to contribute code, please open an issue or submit a pull request. Check out our [Contribution Guidelines](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Thank you for choosing our multi-container web application template. I hope this setup simplifies the development and deployment of your server and UI components, along with MongoDB as your data store. If you have any questions or need assistance, feel free to reach out. Happy coding!


[Tech Stack](#Tech) |
[Pre-requisites](#prerequisites) |
[Basic commands](#Basic-commands-to-run-the-containers) |
[Database diagram](#Database-diagram)

## Tech
- Docker for creating containers for UI, Server and DB
- React for user interface and visualization
- Go for backend APIs
- Mongo for DB

## Prerequisites
- Docker (LTS)
- NodeJS (LTS)
- MongoDB (LTS)
- Go Compiler (1.19.2)

## Basic commands to run the application

```sh
sudo docker compose up [--build]
api1 -> http://localhost:4000 [non cors]
api2 -> http://localhost:4545 [cors] [use with web application]
mongo express -> http://localhost:8081
UI -> http://localhost:3000
```

## CI/CD Pipeline

- GO CI for API code -> Action Name -> [![Go](https://github.com/chrispeterjeyaraj/multi-container-web-application-boilerplate/actions/workflows/go.yml/badge.svg)](https://github.com/chrispeterjeyaraj/multi-container-web-application-boilerplate/actions/workflows/go.yml)

- Node CI for UI Code -> [![Node.js CI](https://github.com/chrispeterjeyaraj/multi-container-web-application-boilerplate/actions/workflows/node.js.yml/badge.svg)](https://github.com/chrispeterjeyaraj/multi-container-web-application-boilerplate/actions/workflows/node.js.yml)

- Docker CI -> [![Exam Center Docker CI](https://github.com/chrispeterjeyaraj/multi-container-web-application-boilerplate/actions/workflows/docker-image.yml/badge.svg)](https://github.com/chrispeterjeyaraj/multi-container-web-application-boilerplate/actions/workflows/docker-image.yml)