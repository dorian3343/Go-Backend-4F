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
### Remove from basket

---

>To remove from a user's basket send a request to '/api/removeBaskets' that looks like this where 'ProductId' is the id of the product to be removed.
```
{
  "Login": "example_user",
  "Password": "example_password",
  "Email": "example@example.com",
  "ID": "123456789",
  "Basket": [
    {
      "ProductID": "product_1",
      "ProductName": "Example Product 1",
      "Price": 19.99
    },
    {
      "ProductID": "product_2",
      "ProductName": "Example Product 2",
      "Price": 29.99
    }
  ],
  "ProductId": "product_1"
}

```

### *Get product's*


---

> To get products send a GET request to '/api/getProducts',you will get a response like this
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

```

> You will get something like this
```
[
  {
    "ID": 1,
    "Name": "Product A",
    "Price": 29.99,
  },
  {
    "ID": 2,
    "Name": "Product B",
    "Price": 49.99,
  },
  {
    "ID": 3,
    "Name": "Product C",
    "Price": 19.99,
  }
]

```
