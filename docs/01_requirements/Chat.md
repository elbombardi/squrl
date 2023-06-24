1. Create ShortURL endpoint will receive a request with a long URL and will return a short URL?
- Yes

2.UpdateLink endpoint will receive an id of a short url [or a short URL](Created with endpoint: Create ShortURL), and will change the long URL for redirect?
- id of short url created from A --> Yes but this will allow update of LongURL and whether this shortURL should track click or not



3, There are endpoint for Disable customer, Do you need only to disable customer account, or also you need an endpoint to enable customer account?
- Yes, also need to enable customer account



4. The shortURL will have option to turn on/off logging of click -> So this will be an endpoint, for example:
mywebsite.com/api/enableLogging/true/\[shortURL] -> will enable the logging
mywebsite.com/api/enableLogging/false/\[shortURL] -> will disable the logging
- No, we use the same one as #2 just pass variable like tracking: true or false



5. Do you have any preference for the API constraints? I propose to use REST API...
- Yes, REST API



And, is not clear what information to store in DB. From what I see I shall log in the DB:
1. Date and time of request -- >Yes
2. ShortURL --> Yes
What other information is mandatory to store in DB logs?
For Click tracking
- First click date and time, Last click date and time, count total number of times the link has been click, Last user agent
table "click_logs" would be like below
id
linkid
click_date = YYYY-MM-DD
click_time = HH🇲🇲ss
click_month = store YYYYMM
user_agent
ip



table "links" would be
id = link_id
customer_id
customer_prefix = each customer has own prefix (set by system) to prevent possible shortlink duplication (3 chars case sensitive)
key = unique key for shortlink (5-6 chars case sensitive)
longURL
tracking = 1 or 0
created_at
updated_at = last update time



For shortURL
- log of changes to the linkid or shorturl like
table "link_logs" would be like below and you just fill this everytime the 2.UpdateLink endpoint is called
id
linkid
longURL
track
ip
created_date


[update] table "links" would be
id = link_id
customer_id
customer_prefix = each customer has own prefix (set by system) to prevent possible shortlink duplication (3 chars case sensitive)
key = unique key for shortlink (5-6 chars case sensitive)
longURL
tracking = 1 or 0
count = total click count
first_clickdatetime = first time the link was clicked (if tracking is enabled)
last_clickdatetime = latest time the link was clicked (if tracking is enabled)
created_at
updated_at = last update time

---
1. How many customers do you expect? ==> About 200
2. How many redirections do you estimate in a day? ==> 10-20K
3. How do you want to deploy the service? On the cloud? Heroku maybe? ==> Local cloud
4. What's the max length of the short URL do you want? (domain name not included) ==> 6 should be enough. The link should be case-sensitive

    