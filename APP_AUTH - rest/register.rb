require 'bcrypt'

def register(user_data)
  password_hash = BCrypt::Password.create(user_data['password'] + ENV['PASSWORD_HASH_PEPPER'])
  begin
    response = GRAPHQL_CLIENT.query(REGISTER_QUERY, variables: {
      first_name: user_data['first_name'],
      last_name: user_data['last_name'],
      password_hash: password_hash,
      email: user_data['email']
    })
    if response.errors.any?
      raise "GraphQL error in register, inserting user: #{response.errors.full_messages.join(', ')}"
    end
    data = response.data.to_h['createUser']
    if data == nil
      raise "No user got returned. This user probably already exists"
    end
    return data
  rescue StandardError => e
    raise "Could not register user: #{e.message}"
  end
end