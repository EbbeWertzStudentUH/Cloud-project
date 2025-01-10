const { buildSchema } = require("graphql");

const generalSchema = buildSchema(`
    type User {
      id: String
      first_name: String
      last_name: String
      friends: [User!]!
      friend_requests: [User!]!
    }
  
    type Query {
      user(id: String!): User
      friends(id: String!): [User!]!
      friendRequests(id: String!): [User!]!
    }
    type Mutation {
      addFriendRequest(user_id: String!, friend_id: String!): [User!]!
      addFriend(user_id: String!, friend_id: String!): [User!]!
}  
    `
  );

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
