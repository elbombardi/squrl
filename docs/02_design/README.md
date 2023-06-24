# Design and Architecture

## Architecture 
- There are 2 main components in this system
    - API Server
    - Redirection Server
- API Server is the main server that handle all the API requests from the client
- Redirection Server is the server that handle all the redirection requests from the client
- The separation between the two components has two benifits : 
    - The servers can be scaled independently from each other.
    - Security: The API server is not exposed public users, it's only used by authenticated users. The redirection server is exposed to public users, it doesn't require authentication, but it doesn't modify the database, it only reads from it.
- This architecture can letter be deployed as microservices, where each component is a microservice.

<p align="center"><img src="images/ArchitectureDiagrams.drawio.png"/></p>

## API Design

## Redirection Server

## Database Design

## Code Structure
