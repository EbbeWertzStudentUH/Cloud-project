# app.rb
require 'sinatra'
require 'json'

set :bind, '0.0.0.0'

get '/' do
  'Hello, world!'
end

get '/api/greet' do
  content_type :json
  { message: 'Hello from the Ruby API!' }.to_json
end

post '/api/echo' do
  content_type :json
  request_data = JSON.parse(request.body.read)
  { received: request_data }.to_json
end
