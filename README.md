# api-gateway

**Orchestrating Microservices for Seamless Integration**

The `api-gateway` serves as a central orchestrator, seamlessly integrating key microservices:

- **Product Service:** Manages product-related operations.
- **Auth Service:** Handles user authentication and authorization.
- **Order Service:** Orchestrates order-related activities.

## Getting Started

### 1. Set Up Environment Variables

Create a `.env` file with the following configurations:

```bash
# Configuration for Product Service
PRODUCT_SERVICE_HOST=http://localhost:8080
LIST_PRODUCT_URI=/v1/products
CREATE_PRODUCT_URI=/v1/products
GET_PRODUCT_URI=/v1/products/:id
UPDATE_PRODUCT_URI=/v1/products/:id
DISABLE_PRODUCT_URI=/v1/products/:id/disable
ENABLE_PRODUCT_URI=/v1/products/:id/enable
INCREASE_BOOKED_QUOTA=/v1/products/increase-booked-quota
DECREASE_BOOKED_QUOTA=/v1/products/decrease-booked-quota

# Configuration for Auth Service
AUTH_SERVICE_HOST=http://localhost:8081
LOGIN_URI=/v1/login
LOGOUT_URI=/v1/logout
AUTHENTICATE_URI=/v1/authenticate

# Configuration for Order Service
ORDER_SERVICE_HOST=http://localhost:8082
ADD_TO_CART_URI=/v1/carts
CREATE_ORDER_URI=/v1/orders
CANCEL_ORDER_URI=/v1/orders/:id/cancel
LIST_ORDERS_URI=/v1/orders
```

### 2. Run the Application

Launch the application using the following command:

```bash
go run main.go
```

### 3. Access the Server

The server will be accessible at [http://localhost:8000](http://localhost:8000).

## Contributing

We welcome contributions! Feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
