## Design a web service to connect to Facebook API to pull campaign data and send notifications to user when campaign budget is below a certain threshold


* campaign creation (budget threshold, duration of campaign, status, meta_data)
* check campaign status [inactive, above_threshold, below_threshold!!!] (periodically check status, set reasonable intervals)
* notify users when campaign budget falls below threshold


* notify user to star campaign if they have funds with us (Future)

* This service can be run in two ways
- ### it can be run as a single web service (like a lamda function)
This requires a data source, in the current implementation, i populate the campaigns randomly
ideally the campaign data should come from a db, service etc. the idea is, you can create a scheduler that runs this service every hour
it would take all the campaigns and run it through the budget checking process
- ### it can be run as a webserver that continuosly listens for campaigns and sends out the appropriate notification
In this case, the queue worker runs in the background everytime, as campaigns are created, they are entered
into the queue to be processed at a later time.

- to run it as a single web service run `make run-service`
- to run it as a web server, run `make run-server` _here is the [postman documentation](https://documenter.getpostman.com/view/7190909/2sA35G2MDS) for this_

### Note

i took the liberty to use the sample api response from facebook api, i also wrote a simple get request to google.com
to mimic http traffic.