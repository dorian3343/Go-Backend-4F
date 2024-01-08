# goShop Backend

## A simple backend made to be used for a school project

### * User creation * 

---
>To create a user, send a request to '/api/createUser' that looks like this
```
{
  "Login": "your_username",
  "Password": "your_password",
  "Email": "your_email@example.com"
}

```
>And it will return something like this (PS. Save this as a cookie as you'll use it for actions)
```
{
  "Login": "john_doe",
  "Email": "john.doe@example.com",
  "ID": "some_unique_id",
  "Basket": [],
}

```
### User login

---
>To login a user,send a request to '/api/loginUser' that looks like this
```
{
  "Login": "john_doe",
  "Password": "secretpassword"
}
```
>and it will return something like this
```
{
 "Login": "john_doe",
 "Email": "john.doe@example.com",
 "ID": "123456789",
 "Basket": []
}
```
### * User deletion*

---
> To delete a user,send a request to '/api/deleteUser' that looks like this 
```

{
  "Login": "john_doe",
  "Email": "john.doe@example.com",
  "ID": "some_unique_id",
  "Basket": [],
}

```
>And it will return something like this
```
{
  "message": "User data deleted successfully"
}

```
### *Add to user basket*

---

> To add to  a user's basket,send a request to '/api/addToBasket' that looks like this 
 ```
 {
  "Login": "john_doe",
  "Email": "john.doe@example.com",
  "ID": "some_unique_id",
  "Basket": [],
  "Product_id": "cpu123"
}
 ```

### *Get product's*   

---

> To get products send whatever to the endpoint,you will get a response like this
```
[
  {
    "Id": 1,
    "Name": "Product A",
    "Price": 19.99,
  },
  {
    "Id": 2,
    "Name": "Product B",
    "Price": 29.99,
  },
  {
    "Id": 3,
    "Name": "Product C",
    "Price": 9.99,
  }
]

```# Go-Backend-4F
