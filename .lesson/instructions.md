# Web API

## HTTP
HTTP communication involves two important terms and those are Client and Server.

**Client:** Client is the one who makes the HTTP request. For example Browser is the client.

**Server:** The server is the one who receives the request and sends the response. Basically server is the piece of code that is responsible for accepting the request and sending the response back. usually, the computer on which server code runs is called server. Simple HTTP communication is shown below
![http1](https://bytesofgigabytes.com/IMAGES/Networking/HTTPcommuncation/http%20communication.png)

Whenever the client or user makes a request using URL then the domain name of the website is responsible for taking requests to the server and then request is forward to a particular web page which is shown below.

![http2](https://bytesofgigabytes.com/IMAGES/Networking/HTTPcommuncation/http%20praticle%20example.png)

In the above example, www.bytesofgigabytes.com will make a request to the bytesofgigabytes.com server and then from the bytesofgigabytes.com server to the post1 web page and then the response is sent back to the client or user.

## REST API
### API
In general, an interface is a device or a system that unrelated entities use to interact. According to this definition, a remote control is an interface between you and a television set.

API is the acronym for Application Programming Interface, which is a software intermediary that allows two applications to talk to each other. Each time you use an app like Facebook, send an instant message, or check the weather on your phone, you’re using an API.

![waiter](https://i.stack.imgur.com/ChtPo.png)

### REST
REST is an acronym for REpresentational State Transfer and an architectural style for distributed hypermedia systems. Roy Fielding first presented it in 2000 in his famous dissertation.

Like other architectural styles, REST has its guiding principles and constraints. These principles must be satisfied if a service interface needs to be referred to as RESTful.

There are four components that must be considered, as follows:

#### URL Design
Consistent naming and url patterns will make for a good, easy-to-understand API. Here's an example url for product data manipulation:
```
/product
/product/123
/product/create
/product/update/123
/product/delete/123
```

in REST we can use:
```
GET /product
GET /product/123
POST /product
PATCH /product/123
DELETE /product/123
```

#### HTTP Verb
Every request made must define an http method so that the API knows what it wants. Here are four http requests that are often used:

1. GET <br>
   Used when you want to display data
2. POST <br>
   Used when you want to save data
3. PUT <br>
   Used when you want to update data
4. DELETE <br>
   Used when you want to delete data
   
By following the http standard above, do not use DELETE when you want to display data.

#### Response Code
Response Code is a standard response code used to inform a request. Here's an example:

1. 200 (Success) <br> Request success
2. 400 (Bad Request) <br>  Data used when the request does no match
3. 401 (Unauthorized) <br>  Code indicating that the request requires authentication (special access such as username and password)
4. 404 (Not Found) <br>  The requested url was not found
5. 500 <br>  (Internal Server Error)

To inform that there is a technical error from the API or server side

#### Response Format
Standardize the structure of a consistent response format every time a request is usually in the form of JSON or XML. Examples like the following
```json
{ 
  status : 200, 
  message : 'Success', 
  data: 'Welcome to rest api' 
}
```

## How FrontEnd Interact with Backend

![fe-be](https://www.researchgate.net/profile/Octavian-Dospinescu/publication/316675285/figure/fig4/AS:490411216117762@1493934553127/REST-API-integrated-with-frontend-technology-using-AJAX-and-JQuery.png)

## Lets Create Our REST API!

## Assignment
Complete Web API, so it's can get student details, update and delete student data.
Endpoint
- GET /students/:id <br>
  response :
  ```json
  {
    "type": "success",
    "message": "",
    "data": [
      {
         "id": 1, 
          "name": "Chaplin",
          "gender": "M"
      }
    ]
  }
  ```
- PATCH /students/:id <br>
  request :
   - id
   - name
   - gender
  
  response :
  ```json
  {
    "type": "success",
    "message": "successfully updated",
    "data": []
  }
  ```
- DELETE /students/:id <br>
  response :
  ```json
  {
    "type": "success",
    "message": "successfully updated",
    "data": []
  }
  ```

  _**Hint: use r.URL.Path to get dynamic ID on URL**_


## Reference
- https://www.iitk.ac.in/esc101/05Aug/tutorial/java/concepts/interface.html
- https://restfulapi.net/
- https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm
- https://www.redhat.com/en/topics/integration/whats-the-difference-between-soap-rest
- https://www.freecodecamp.org/news/rest-api-best-practices-rest-endpoint-design-examples/