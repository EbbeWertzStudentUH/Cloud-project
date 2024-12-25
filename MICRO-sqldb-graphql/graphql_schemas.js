const { buildSchema } = require("graphql");

const generalSchema = buildSchema(`
    type User {
      id: String
      first_name: String
      last_name: String
    }
  
    type Query {
      getUser(id: String!): User
    }`);

const sensitiveSchema = buildSchema(`
    type User {
      id: String
      first_name: String
      last_name: String
      password_hash: String
    }
  
    type Query {
      getAuthInfo(id: String!): User
    }`);

module.exports = { generalSchema, sensitiveSchema };
