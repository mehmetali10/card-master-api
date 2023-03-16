# Card Master API
Card Master API is a RESTful API built using Golang and PostgreSQL that allows CRUD (Create, Read, Update, Delete) operations for card information. This API uses CORS (Cross-Origin Resource Sharing) policy to enable cross-domain requests.
<br />

## Installation
To run Card Master API locally, you need to have the following tools installed on your machine:

<li>Docker</li>
<li>Docker Compose</li>
<li>Go</li>
<br />
Once you have installed these tools, follow these steps:

<ol>
<li>Clone this repository to your local machine:</li>


<code>git clone https://github.com/mehmetali10/card-master-api.git</code>

<li>Navigate to the project directory:</li>

<code>cd card-master-api/src</code>

<li>Run the following command to start the Docker containers:</li>

<code>docker-compose up</code>
<br /><br />
This will start the PostgreSQL database and the Golang API server.
<br />

<li>Navigate to http://localhost:8000 to access the API endpoints.</li>
</ol>
<br />

## Usage
Card Master API provides the following endpoints:
<br />
<li><code>GET /cards</code> : Get a list of all cards</li>
<li><code>GET /cards/{id}</code> : Get a card by ID</li>
<li><code>POST /cards</code> : Create a new card</li>
<li><code>PUT /cards</code> : Update an existing card</li>
<li><code>DELETE /cards/{id}</code> : Delete a card by ID</li>
<br /><br />
To use the API, you can use a tool like curl or a REST client like Postman.
<br /><br />
For example, to get a list of all cards, you can run the following command:
<br /><br />
<code>curl http://localhost:8080/cards</code>
<br /><br />
This will return a JSON response containing a list of all cards.
<br /><br />

## Configuration
Card Master API can be configured using environment variables. The following environment variables are supported:
<br /><br />

<strong>DB_HOST</strong> : The hostname of the PostgreSQL database (default: <strong>localhost</strong>)<br />
<strong>DB_PORT</strong> : The port number of the PostgreSQL database (default: <strong>5439</strong>)<br />
<strong>DB_USER</strong> : The username to use to connect to the PostgreSQL database (default: <strong>postgres</strong>)<br />
<strong>DB_PASSWORD</strong> : The password to use to connect to the PostgreSQL database (default: <strong>unbroken</strong>)<br />
<strong>DB_NAME</strong> : The name of the PostgreSQL database (default: <strong>card_master_db</strong>)<br /><br />
You can set these environment variables either in your shell or by modifying the docker-compose.yml file.<br /><br />

## Contributing
If you find a bug or have an idea for a new feature, feel free to open an issue or submit a pull request. We welcome contributions from the community!
<br /><br />

## License
Card Master API is licensed under the MIT License. See the LICENSE file for more information.
