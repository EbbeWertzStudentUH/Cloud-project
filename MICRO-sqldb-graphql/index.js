const express = require('express');
require('dotenv').config();
const { graphqlHTTP } = require('express-graphql');
const { LISTEN_PORT } = process.env;

import { generalSchema, sensitiveSchema } from './graphql_schemas.js';
import { generalResolvers, sensitiveResolvers, secret_auth} from './resolvers.js';

const app = express();

app.use('/users', graphqlHTTP({
  schema: generalSchema,
  rootValue: generalResolvers,
  graphiql: true,
}));

app.use('/auth-info', secret_auth, graphqlHTTP({
  schema: sensitiveSchema,
  rootValue: sensitiveResolvers,
  graphiql: true,
}));

app.listen(LISTEN_PORT, () => {
  console.log(`Listening on http://localhost:${LISTEN_PORT}/users`);
  console.log(`Listening on http://localhost:${LISTEN_PORT}/auth-info`);
});
