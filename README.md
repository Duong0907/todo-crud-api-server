## **How to connect to mongodb in mongosh:**
```powershell
mongosh "mongodb+srv://cluster0.aszqp6w.mongodb.net/myFirstDatabase" --apiVersion 1 --username Duong
```

---

# API documentation

## **1. Read all todos**
**Method**
```
    GET
```

**URL**

``` 
    https://duong-todos.onrender.com/api/todos
```

**Result**

*  Success
``` json
[
    {
        "_id": "63bab04a802a07bb7a5d3594",
        "content": "Have breakfast",
        "status": "To do"
    },
    {
        "_id": "63bab0e3b4a04a63a9efcdb0",
        "content": "Go to bar",
        "status": "Done"
    }
]
```

* Failure

```json
{
    "message": "error"
}
```

## **2. Read a todo**

**Method:**
```
    GET
```
**URL**

``` 
    https://duong-todos.onrender.com/api/todo/{id}
```
*Example*
``` 
    https://duong-todos.onrender.com/api/todo/63bab04a802a07bb7a5d3594
```

**Result**

*  Success
``` json
{
    "_id": "63bab04a802a07bb7a5d3594",
    "content": "Have breakfast",
    "status": "To do"
}
```

* Failure

```json
{
    "message": "error"
}
```

## **3. Create a todo**
**Method**
```
    POST
```
**URL**

``` 
    https://duong-todos.onrender.com/api/todo/{id}
```

*Example*
``` 
    https://duong-todos.onrender.com/api/todo/63bab04a802a07bb7a5d3594
```

**Body:** Todo object to create

*Example*
```json
{
    "_id": "63bab04a802a07bb7a5d3594",
    "content": "Have breakfast",
    "status": "To do"
}
```

**Result**

*  Success
``` json
{
    "_id": "63bab04a802a07bb7a5d3594",
    "content": "Have breakfast",
    "status": "To do"
}
```

* Failure

```json
{
    "message": "error"
}
```

## **4. Update a todo**
**Method**
```
    PUT
```
**URL**

``` 
    https://duong-todos.onrender.com/api/todo/{id}
```

*Example*
``` 
    https://duong-todos.onrender.com/api/todo/63bab04a802a07bb7a5d3594
```

**Body:** Todo object to replace

*Example*
```json
{
    "_id": "63bab04a802a07bb7a5d3594",
    "content": "Have lunch",
    "status": "To do"
}
```

**Result**

*  Success
``` json
{
    "_id": "63bab04a802a07bb7a5d3594",
    "content": "Have lunch",
    "status": "To do"
}
```

* Failure

```json
{
    "message": "error"
}
```


## **5. Delete a todo**
**Method**
```
    DELETE
```
**URL**

``` 
    https://duong-todos.onrender.com/api/todo/{id}
```

*Example*
``` 
    https://duong-todos.onrender.com/api/todo/63bab04a802a07bb7a5d3594
```

**Result**

*  Success
``` json
{
    "_id": "63bab04a802a07bb7a5d3594",
    "content": "Have lunch",
    "status": "To do"
}
```

* Failure

```json
{
    "message": "error"
}
```