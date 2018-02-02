# antifreeze-back-end
Repository for the back-end of Freeze-B-Gone

## Development

This section contains instructions for setting up and using the development environment.

### Environment

- Node.js 8.9.4
- NPM 5.6.0
- Serverless framework 1.26 (global install)
- Go 1.9.3
- Go Dep 0.4.1
- JRE 6.x or newer
- [DynamoDB Local](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html)
- Postman
- Git

### Set Up

1. Install the required tools
2. Clone the repository
3. Use Dep to install Go dependencies
4. Add a `.env` file to the top level directory with the following contents (for development purposes):

```
PORT=3000
SECRET=thisisoursecret
SALT=10

PROD_DB_URI=pending
DEV_DB_URI=pending
TEST_DB_URI=pending

USERNAME=test@ttu.edu
PASSWORD=test
```

### Test Changes

Instructions pending...