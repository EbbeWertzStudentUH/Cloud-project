require 'graphql/client'
require 'graphql/client/http'

graphql_http = GraphQL::Client::HTTP.new("#{ENV['DB_GRAPHQL_URL']}/users-sensitive") do
  def headers(context)
    { "Authorization" => "Bearer #{ENV['DB_GRAPHQL_API_AUTH_SECRET']}" }
  end
end

schema = GraphQL::Client.load_schema(graphql_http)
GRAPHQL_CLIENT = GraphQL::Client.new(schema: schema, execute: graphql_http)

