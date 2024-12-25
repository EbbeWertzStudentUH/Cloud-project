require 'dotenv'
Dotenv.load('./.env')
require 'sinatra'
require './graphql_client.rb'
require './queries.rb'
require './register.rb'
# require 'jwt'
# require 'net/http'
# require 'json'


set :port, ENV['LISTEN_PORT']

post '/register' do
  content_type :json
  data = JSON.parse(request.body.read)
  begin
    user = register(data)
    return {
      message: "new user succesfully registered",
      user: user }.to_json
  rescue StandardError => e
    puts "error: #{e}"
    halt 500, { message: e.message }.to_json
  end
end
