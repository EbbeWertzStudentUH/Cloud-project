const { buildSchema } = require("graphql");

const generalSchema = buildSchema(`
    type User {
      id: String
      first_name: String
      last_name: String
    }
  
    type Query {
      user(id: String!): User
    }`);

const sensitiveSchema = buildSchema(`
    type User {
      id: String
      first_name: String
      last_name: String
      password_hash: String
      email: String
    }
  
    type Query {
      user(id: String, email: String): User
    }
    type Mutation {
        createUser(first_name: String!, last_name: String!, password_hash: String!, email: String!): User
    }
      `);

module.exports = { generalSchema, sensitiveSchema };
