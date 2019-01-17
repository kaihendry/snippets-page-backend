## SnippetsPage RESTful API service  
Backend of SnippetsPage application. Echo + MongoDB.

**Project: [https://snippets.page](https://www.google.com)**\
**Frontend GitHub repository: coming soon**

## JWT
JWT authentication using HS256 algorithm and JWT is retrieved from Authorization request header.
>![POST](https://placehold.it/12/248FB1/000000?text=+) POST https://cloud.snippets.page/v1/authorization

*request*

>curl -X POST 
  -H 'Content-Type: application/json' 
  -d '{"login":"user", "password":"qwerty"}' 
  https://cloud.snippets.page/v1/authorization

*response*
```json
{
    "_id":"8c373386ca11bq001f62c31z",
    "login":"user",
    "email":"user@mail.com",
    "token":"<token>",
    "banned":false,
    "created_at":"2019-01-10T13:11:34.445Z",
    "updated_at":"2019-01-10T13:11:34.445Z"
}
```
Request a restricted resource using the token in Authorization request header.
>curl <url> -H "Authorization: Bearer _token_"

## Pagination
HTTP headers:\
X-Pagination-Total-Count: The total number of resources;\
X-Pagination-Page-Count: The number of pages;\
X-Pagination-Current-Page: The current page (1-based);\
X-Pagination-Per-Page: The number of resources in each page;\
Link: A set of navigational links allowing client to traverse the resources page by page.

## Resources
### User
![POST](https://placehold.it/12/248FB1/000000?text=+) POST https://cloud.snippets.page/v1/users <br />
![GET](https://placehold.it/12/6BBD5B/000000?text=+) GET https://cloud.snippets.page/v1/me
```json
{
  "_id": "5becbd11f23efbeiod769c40",
  "login": "user",
  "email": "email@gmail.com",
  "token": "",
  "banned": false,
  "created_at": "2018-11-15T00:27:47.302Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```
### Snippet
![GET](https://placehold.it/12/6BBD5B/000000?text=+) GET https://cloud.snippets.page/v1/snippets <br />
![GET](https://placehold.it/12/6BBD5B/000000?text=+) GET https://cloud.snippets.page/v1/me/snippets <br />
![POST](https://placehold.it/12/248FB1/000000?text=+) POST https://cloud.snippets.page/v1/me/snippets <br />
![PUT](https://placehold.it/12/DF9D43/000000?text=+) PUT https://cloud.snippets.page/v1/me/snippets/:id <br />
![DELETE](https://placehold.it/12/E27A7A/000000?text=+) DELETE https://cloud.snippets.page/v1/me/snippets/:id <br />
![GET](https://placehold.it/12/6BBD5B/000000?text=+) GET https://cloud.snippets.page/v1/me/snippets/tags


#### Query params

| Param       |      Value    | Example | Description
|-------------|:-------------:|:-------------:|:-------------:|
| sort             | created_at or -created_at | /snippets?sort=-created_at | sort collections 
| q                | string | /snippets?q=my_snippet | search snippets by title
| fields            | string | /snippets?fields=tags,title,created_at | you can specify which fields should be included in the result
| filter[tags]      | string | /snippets?filter[tags]=tag1,tag2 | filtering collections by tags
| filter[favorite]  | bool   | /snippets?filter[favorite]=true | filtering collections by favorite
| limit            | int    | /snippets?limit=50 | items per page
| page             | int    | /snippets?page=2 |  this offsets the starting point of the collection returned from the server in the results.
```json
{
  "_id": "5c393fd7f23efb0354049886",
  "user_id": "5becbd83f23efbecdd769c40",
  "favorite": true,
  "title": "Snippet title",
  "files": [],
  "public": false,
  "tags": [],
  "created_at": "2019-01-12T03:11:51.108676204+02:00",
  "updated_at": "2019-01-12T03:11:51.108676219+02:00"
}
```

## HTTP Status Codes
| Code     |      Status   | 
|----------|:-------------:|
| 200      | OK |
| 201      | OK. Created |
| 204      | No Content | 
| 400      | Bad Request | 
| 403      | Forbidden | 
| 404      | Not found | 
| 500      | Internal Server Error | 
| 503      | Service Unavailable | 


## Examples
### Get profile of current auth user
*request*

>curl -XGET -H "Content-type: application/json" -H "Authorization: Bearer _token_" 'https://cloud.snippets.page/v1/me'

*response*

```json
{
  "_id": "5becbd83f23efbecdd769c40",
  "login": "username",
  "email": "sername@gmail.com",
  "token": "",
  "banned": false,
  "created_at": "2018-11-15T00:27:47.302Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```
### Create snippet
![POST](https://placehold.it/12/248FB1/000000?text=+) POST https://cloud.snippets.page/v1/me/snippets

*request*

>curl -XPOST -H "Content-type: application/json" -H "Authorization: Bearer _token_"
 -d '{"title": "Snippet title","favorite": true,"files": [],"public": false,"tags": []}' 'https://cloud.snippets.page/v1/me/snippets'

 *response*

 ```json
 {
  "_id": "5c393ed7f23efb0354049756",
  "user_id": "5becbd83f23zfbecrw769c40",
  "favorite": true,
  "title": "Snippet title",
  "files": [],
  "public": false,
  "tags": [],
  "created_at": "2019-01-12T03:11:51.108676204+02:00",
  "updated_at": "2019-01-12T03:11:51.108676219+02:00"
}
