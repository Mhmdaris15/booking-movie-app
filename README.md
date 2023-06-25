# Booking Movie App - Server Side

This is the server-side application for the Booking Movie App written in Golang and built with the Gin Gonic framework. The application provides APIs for managing movie bookings and theater schedules.

## Features

- Authentication: Users can register, log in, and manage their authentication tokens.
- Movie Management: Admin users can add, update, and delete movie details such as title, description, genre, and duration.
- Theater Management: Admin users can manage theaters, including adding, updating, and deleting theater information.
- Schedule Management: Admin users can create and manage movie schedules for different theaters.
- Booking Management: Users can search for available movie schedules and book seats for a specific movie and theater.

## Requirements

- Go 1.16 or higher
- Gin Gonic framework

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/Mhmdaris15/booking-movie-app.git
   ```

2. Change into the project repository

`cd booking-movie-app`

3. Install The Dependencies

`go mod download`

4. Build The Project

`go build .\cmd\server\main.go`

5. Run The Application

`.\main.exe`

6. The Server should now be running on `http://localhost:3000`

## Configuration

The server can be configured by modifying the `config.go` file. You can change the server port, database connection details, and other settings as needed.

## API Documentation

For detailed information on the available API endpoints, please refer to the routes.go

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## Licence

This Projcet is licensed under the MIT License.
