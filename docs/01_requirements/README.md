# Short URL Web App:

- The Short URL Web App will be written in GO lang to support concurrency.
It will utilize the Postgres database.
- There will be no backend management functionalities.
- All ShortURLs will be created exclusively through APIs.

## API Endpoints:

### Group 1: Customer APIs
This group is accessible only for Admin user.

#### 1. Create Customer:
- Purpose: Used to create a customer account.
- Request data:
    - Username (unique)
    - Email
- Response data:
    - Status: success or fail
    - On success, return a system-generated password.

#### 2. Update Customer:
- Purpose: Used to disable/enable a customer's use of the API.
- Request data:
    - Customer ID
    - Status (active/inactive)
- Response data:
    - Status: success or fail

### Group 2: Short URL APIs
This group is accessible using a customer's username and password.

#### 2. Create ShortURL:
- Purpose: Used to create a ShortURL.
- Request data:
    - Long URL
- Response data:
    - Short URL

#### 3. Update ShortURL:
- Purpose: Used to update ShortURL information such as the long url or to disable/enable url tracking.
- Request data:
    - Short URL
    - New Long URL
    - Tracking status (true/false)
- Response data:
    - Status: success or fail

## Redirection :
- There will be no user interface (UI) on the frontend. When a user clicks on a ShortURL, the system will redirect them to the destination URL.
- If the short URL is disabled or not found, the system will return a 404 error. Otherwise, the system will redirect the user to the destination URL, and will keep a track of the click in the database.

Sample ShortURL:
- Format: http://x.xxx.x/yyy/zzzzz
    - x.xxx.x: domain name
    - yyy: customer prefix
    - zzzzz: short URL key

## Database:

### Customer data 
- The customer data table will store the following information:
    - Customer ID
    - Customer Prefix (3 characters, case-sensitive)
    - Username
    - Email
    - Password 
    - Status (enabled or disabled) 
    - Created At 
    - Updated At (last update time)

### Short Urls : 
- The short urls (links) table will store the following information:
    - Short URL Key (unique key for shortlink, 5-6 characters, case-sensitive)
    - Customer ID
    - Long URL
    - Status (enabled or disabled)
    - Click Count
    - First Click Date and Time (if tracking is enabled)
    - Last Click Date and Time (if tracking is enabled)
    - Created At
    - Updated At (last update time)

### Click tracking
- For click tracking, the following information will be stored in the database:
    - Click ID
    - Short URL ID
    - Click Date and Time
    - User Agent
    - IP Address

## Additional Questions:
1. How many customers do you expect?
    - Approximately 200.
2. How many redirections do you estimate in a day?
    - Between 10,000 to 20,000 redirections.
3. How do you want to deploy the service? On the cloud? Heroku maybe?
    - The service will be deployed on a local cloud.
