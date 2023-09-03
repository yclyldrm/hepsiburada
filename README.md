# Hepsiburada Case Study

This project is developed as an application for an e-commerce website. It processes data from Excel files and performs read and write operations to a MySQL database for products, campaigns, and orders.


## Run Locally

Clone the project

```bash
  git clone git@github.com:yclyldrm/hepsiburada.git
```
or
```bash
  git clone https://github.com/yclyldrm/hepsiburada.git
``` 

Go to the project directory

```bash
  cd hepsiburada
```

Start the server

```bash
  make run-app
```


## Running Tests

To run tests, run the following command

```bash
  make run-test
```


## API Reference

#### Sample data

```http
  GET /sample
```
#### response
```http
{
    "message": "Sample data implemented successfully",
    "output": [
        "Product created; code=P1; price=100; stock=1000",
        "Campaign created; name C1, product P1, duration 5, limit 20, target sales count 100",
        "Product P1 info; price=100, stock=1000",
        "01:00",
        "Product P1 info; price=90, stock=1000",
        "02:00",
        "Product P1 info; price=80, stock=1000",
        "03:00",
        "Product P1 info; price=80, stock=1000",
        "04:00",
        "Product P1 info; price=80, stock=1000",
        "06:00",
        "Product P1 info; price=100, stock=1000",
        "Campaign C1 info; Status false, Target Sales 100, Total Sales 0, Turnover 100, Average Item Price 0"
    ],
    "status": "success"
}
```

#### Custom data

```http
  POST /custom
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `custom_file`      | `form_data` | **Required**|
