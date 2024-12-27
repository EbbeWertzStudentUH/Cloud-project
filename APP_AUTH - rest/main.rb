require 'dotenv'
Dotenv.load('./.env')
require 'sinatra'
require './graphql_client.rb'
require './queries.rb'
require './register.rb'
require './login_and_auth.rb'

set :port, ENV['LISTEN_PORT']
set :bind, '0.0.0.0'

post '/register' do
  content_type :json
  data = JSON.parse(request.body.read)
  begin
    user = register(data)
    return { message: "new user succesfully registered", user: user }.to_json
  rescue StandardError => e
    halt 500, { message: e.message }.to_json
  end
end

post '/login' do
  data = JSON.parse(request.body.read)
  begin
    token = login(data)
    return { message: "user succesfully logged in", token: token }.to_json
  rescue StandardError => e
    halt 500, { message: e.message }.to_json
  end
end

get '/verify_token' do
  token = request.env['HTTP_AUTHORIZATION']&.split(' ')&.last  # auth header is "Bearer " en dan token
  if token.nil?
    halt 500, { message: "token should be in auth header" }.to_json
  end

  begin
    user_id = decode_jwt(token)
    return { message: "token is valid", valid: true, user_id: user_id }.to_json
  rescue StandardError => e
    halt 500, { message: e.message, valid: false }.to_json
  end
end