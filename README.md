# trello-cards-creator-oc
Web service that interfaces with trello with certain rules in the middle to create cards in an easier way.
 
## How to make it run?

- Run the following command: 
```  
make setup  
```  
This will install the required dependencies, generate a copy of the environment variables file and build the project.

- Then, it is necessary to configure the following environment variables:

|Name|Description|Default value|
|---|---|---|---|
|TRELLO_BASE_URL| Base URL of Trello REST API|https://api.trello.com|
|TRELLO_API_KEY| API Key for the power up created on [https://trello.com/power-ups/](https://trello.com/power-ups/)|-|
|TRELLO_USER_TOKEN| Generated based on the API Key|-|
|TRELLO_BOARD_ID| ID of the board in Trello, used to search for the members participating on that board (for Bug card random assignment)|-|
|TRELLO_GENERAL_LIST_ID| ID of the list in which bug and task cards will be created|-|
|TRELLO_ISSUES_LIST_ID| ID of the list in which issue cards will be created|Could be the same as TRELLO_GENERAL_LIST_ID|
|TRELLO_BUG_LABEL_ID| ID of the label used to tag bug cards|-|
|TRELLO_MAINTENANCE_LABEL_ID| ID of the label used to tag task cards with the category of Maintenance|-|
|TRELLO_RESEARCH_LABEL_ID| ID of the label used to tag task cards with the category of Research|-|
|TRELLO_TEST_LABEL_ID| ID of the label used to tag task cards with the category of Test|-|
|SERVER_PORT|Number of the port in which the web service will run|3030|

- Once variables are in place (you can use [direnv](https://direnv.net/)), you can run the service using the following command:
```
make run
```

## How to test it?

Although the exercise description specified that the endpoint has to be the root path, I changed it a little bit, to add the card type as the last URL segment. This will  help to keep the payload more lean and containing only the information used to populate the Trello card.

The following examples can be used to create Issues, Bugs and Tasks on Trello. Additionally you can use the Insomnia REST suite in the root folder of this repository.

```
curl --request POST \
  --url http://localhost:3030/cards/issue \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Update enum with new tax types",
	"description": "The enum that calculates taxes for payrol is currently outdated. Just needs to get the codes updated with the new codes."
}'
```

```
curl --request POST \
  --url http://localhost:3030/cards/bug \
  --header 'Content-Type: application/json' \
  --data '{
	"description": "NaN shown on UI when calculating salary for an employee with less than a month"
}'
```

```
curl --request POST \
  --url http://localhost:3030/cards/task \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Enable feature flag for automatic loan discountz",
	"category": "test"
}'
```
