# Taylor Swift Songs API

This is a simple Go project, a REST API for managing Taylor Swift songs, with CRUD operations.

## Prerequisites

Make sure you have Go installed on your computer. You can download it from [golang.org](https://golang.org/).

You'll also need an SQL database to establish the necessary connections. In the project directory, create a `.env` file based on the example provided in `.env.example` and fill in the information for your PostgreSQL database.

Example `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database_name
```

Ensure you create the corresponding database in PostgreSQL.

## Installation and Execution

1. Clone the repository:

    ```bash
    git clone https://github.com/ana-flav/learning-golang.git
    cd your-project
    ```

2. Download Go dependencies:

    ```bash
    go mod download
    ```

3. Run the project (make sure you are inside the cmd directory):

    ```bash
    go run main.go
    ```

The API will be available at [http://localhost:8000](http://localhost:8000).

## Example Usage with Postman

Let's provide an example of how to use Postman to send a POST request and add a song using your API endpoint.

1. Open Postman and create a new request.

2. Select the POST method.

3. Enter the URL of your API to add a song, for example: `http://localhost:8000/add-song`.

4. Go to the "Body" tab, select "raw", and choose "JSON (application/json)" from the menu.

5. In the request body, enter the details of the song you want to add. For example:

    ```json
    {
        "Name": "mirrorball",
        "LinkSong": "https://open.spotify.com/intl-pt/track/0ZNU020wNYvgW84iljPkPP?si=20728167f34247c4"
    }
    ```

6. Click the "Send" button to send the request.

This will send a POST request to the `http://localhost:8000/add-song` endpoint with the song details provided in the request body. Remember to adapt the details as needed to match the expected fields in your application.
