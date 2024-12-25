require 'jwt'

def generate_jwt(user_id)
  data = {
    user_id: user_id,
    exp: Time.now.to_i + 3600 # 1 uur
  }
  secret_key = ENV['JWT_SECRET']
  JWT.encode(data, secret_key, 'HS256')
end

def login(user_data)
  begin
    response = GRAPHQL_CLIENT.query(QUERY_HASH_AND_ID_BY_EMAIL, variables: { email: user_data['email'] })
    if response.errors.any?
      raise "GraphQL error in login, quering user data: #{response.errors.full_messages.join(', ')}"
    end
    hash_and_id = response.data.to_h['user']
    if !hash_and_id.nil? && BCrypt::Password.new(hash_and_id['password_hash']) == (user_data['password'] + ENV['PASSWORD_HASH_PEPPER'])
      return generate_jwt(hash_and_id['id'])
    else
      raise "invlaid email or password"
    end

  rescue StandardError => e
    raise "Could not login user: #{e.message}"
  end
end