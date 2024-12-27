require 'graphql/client'
require 'graphql/client/http'

graphql_http = GraphQL::Client::HTTP.new("#{ENV['DB_GRAPHQL_URL']}/users-sensitive") do
  def headers(context)
    { "Authorization" => "Bearer #{ENV['DB_GRAPHQL_SENSITIVE_ENDPOINT_API_SECRET']}" }
  end
end

schema = nil

while true
  begin
    schema = GraphQL::Client.load_schema(graphql_http)
    break
  rescue
    puts "could not connect to graphql query, trying again in 3s"
  end
  sleep 3
end

GRAPHQL_CLIENT = GraphQL::Client.new(schema: schema, execute: graphql_http)

