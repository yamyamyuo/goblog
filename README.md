# Goblog

### API

| Endpoint					| Method | Description
|---------------------------|--------|-----------------
|**/archives**				| GET    | List all blog entries
|							| POST	 | Post a new blog entry
|**/archives/:id**			| GET	 | List a blog entry by id
|							| DELETE | Delete a blog entry by id
|**/about**					| GET	 | Get the about page

##### Usage Example
###### List all blog entries
	- GET /archives

Example response (JSON):

	[
		{
			"title":"how to write markdown?",
			"id":"how-to-write-markdown",
			"date": "2015-09-01 12:18:29.194812713",
			"text": "This is my content"
		},

		{
			"title":"how to write web app",
			"id":"how-to-write-web-app",
			"date": "2015-09-02 12:18:29.194812713",
			"text": "This is my content"
		}
	]

##### List a blog entry
	- GET /archives/:id

Example response (JSON):

	{
		"title":"how to write web app",
		"id":"how-to-write-web-app",
		"date": "2015-09-02 12:18:29.194812713",
		"text": "This is my content"
	}

##### Post a new blog entry
	- POST /archives

	{
		"title": "example post"
		"text": "this is my blog content"
	}

#### Delete a blog entry
	- DELETE /archives/:id

#### Get the about page
	- GET /about

Example response (String):

	(still in dev, need JS Markdown renderer to handle it)