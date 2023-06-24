# Design and Architecture

## Architecture 
- There are 2 main components in this system
    - API Server
    - Redirection Server
- API Server is the main server that handle all the API requests from the client
- Redirection Server is the server that handle all the redirection requests from the client
- The separation between the two components has two benifits : 
    - The servers can be scaled independently from each other.
    - Security: The API server is not exposed public users, it's only used by authenticated users. The redirection server is exposed to public users, it doesn't require authentication, but it doesn't modify customer and short url data, it only reads data, and inserts tracking data.
- This architecture can letter be deployed as microservices, where each component is a microservice.

<p align="center"><img src="images/ArchitectureDiagrams.drawio.png"/></p>

## API Server
- The API server is the main server that handle all the API requests from the admin and customers.
- The API server is a REST API server, it's built using Golang & Swagger.
- The API server is stateless, it doesn't store any data, it only reads and writes data to the database.
- The API server is secured using API keys, each customer has an API key stored in the database, and the admin has a specific API key stored as an environement variable.

Here's a brief description of the API endpoints:
1. Create Customer:
   - Endpoint: `/api/admin/customer/`
   - HTTP Method: POST
   - Request JSON Structure:
     ```json
     {
       "username": "string",
       "email": "string"
     }
     ```
   - Response JSON Structure (on success):
     ```json
     {
       "status": "string",
       "api_key": "string",
     }
     ```

2. Update Customer:
   - Endpoint: `/api/admin/customer/`
   - HTTP Method: PUT
   - Request JSON Structure:
     ```json
     {
       "username": "string",
       "status": "string"
     }
     ```
   - Response JSON Structure (on success):
     ```json
     {
       "status": "string"
     }
     ```

3. Create ShortURL:
   - Endpoint: `/api/customer/short-url/create`
   - HTTP Method: POST
   - Request JSON Structure:
     ```json
     {
       "long_url": "string"
     }
     ```
   - Response JSON Structure (on success):
     ```json
     {
       "short_url": "string"
     }
     ```

4. Update ShortURL:
   - Endpoint: `/api/customer/short-url/update`
   - HTTP Method: PUT
   - Request JSON Structure:
     ```json
     {
       "short_url": "string",
       "new_long_url": "string",
       "tracking_status": "boolean"
     }
     ```
   - Response JSON Structure (on success):
     ```json
     {
       "status": "string"
     }
     ```

## Redirection Server

## Database Design

## Code Structure
