package graphql

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	URL(alias: String!, expireAfter: Time): URL
}

type Mutation {
	createURL(captchaResponse: String!, url: URLInput!, userEmail: String): URL
}

type URL {
	alias: String
	originalURL: String
	expireAt: Time
}

input URLInput {
	originalURL: String!
	customAlias: String
	expireAt: Time
}

type User {
	email: String
}

scalar Time
`
