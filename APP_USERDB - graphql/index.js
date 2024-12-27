const express = require('express');
require('dotenv').config();
const { graphqlHTTP } = require('express-graphql');
const { LISTEN_PORT } = process.env;

const { generalSchema, sensitiveSchema }  = require('./graphql_schemas.js');
const { generalResolvers, sensitiveResolvers, secret_auth}  = require('./resolvers.js');

const app = express();

app.use('/users', graphqlHTTP({
  schema: generalSchema,
  rootValue: generalResolvers,
  graphiql: true,
}));

app.use('/users-sensitive', secret_auth, graphqlHTTP({
  schema: sensitiveSchema,
  rootValue: sensitiveResolvers,
  graphiql: true,
}));

app.listen(LISTEN_PORT, () => {
  console.log(`Listening on http://localhost:${LISTEN_PORT}/users`);
  console.log(`Listening on http://localhost:${LISTEN_PORT}/auth-sensitive`);
});
