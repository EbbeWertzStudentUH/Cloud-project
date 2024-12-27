REGISTER_QUERY = GRAPHQL_CLIENT.parse <<-GRAPHQL
  mutation($first_name: String!, $last_name: String!, $password_hash: String!, $email: String!) {
    createUser(
      first_name: $first_name,
      last_name: $last_name,
      password_hash: $password_hash,
      email: $email
    ) {
      id
      first_name
      last_name
      email
    }
  }
GRAPHQL

QUERY_HASH_AND_ID_BY_EMAIL = GRAPHQL_CLIENT.parse <<-GRAPHQL
query($email: String!) {
  user(email: $email) {
    id
    password_hash
  }
}
GRAPHQL