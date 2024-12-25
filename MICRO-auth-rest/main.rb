require 'dotenv'
Dotenv.load('./.env')
require 'sinatra'
require './graphql_client.rb'
require './queries.rb'
require './register.rb'
require './login_and_auth.rb'

set :port, ENV['LISTEN_PORT']

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

post '/logout' do
end

get '/verify_token' do
end

post '/refresh_token' do
end