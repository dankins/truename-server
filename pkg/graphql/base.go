package graphql

import (
	"fmt"

	"github.com/dankins/truename-server/pkg/identity"
	graphql "github.com/graph-gophers/graphql-go"
)

type schemaResolver struct {
	IdentityService *identity.IdentityService
}

type createIdentityInput struct {
	Account string  `json:"account,omitempty"`
	PubKey  *string `json:"pubKey,omitempty"`
}

func (resolver *schemaResolver) Hello() string { return "Hello, nuggets!" }

func (resolver *schemaResolver) IdentityFactoryAddress() string {
	return resolver.IdentityService.GetIdentityFactory().Address
}

func (resolver *schemaResolver) CreateFactory() string {
	address := resolver.IdentityService.DeployIdentityFactory()
	return address.String()
}

func (resolver *schemaResolver) CreateIdentity(args *struct {
	Input *createIdentityInput
}) string {
	fmt.Println("CreateIdentity", args.Input)
	return resolver.IdentityService.DeployIdentity(&identity.IdentityInput{
		Account: args.Input.Account,
	}).Address
}

func (resolver *schemaResolver) GetIdentityForAccount(args *struct {
	Account string
}) string {
	return resolver.IdentityService.GetIdentityForAccount(args.Account)
}

func (resolver *schemaResolver) GetIdentityABI() string {
	return resolver.IdentityService.GetIdentityABI()
}

func (resolver *schemaResolver) GetIdentityFactoryABI() string {
	return resolver.IdentityService.GetIdentityFactoryABI()
}

func (resolver *schemaResolver) RelayTransaction(args *struct {
	Input *identity.RelayTransactionInput
}) bool {
	return resolver.IdentityService.RelayTransaction(args.Input)
}

func BuildSchema(identityService *identity.IdentityService) *graphql.Schema {
	s := `

		input CreateIdentityInput {
			account: String!
			pubKey: String
		}

		input RelayTransactionInput {
			identityContractAddress: String!
			authorizedAccount: String!
			gas:                     String!
			messageHash:             String!
			signedHash:              String!
		}

		schema {
			query: Query
			mutation: Mutation
		}

		type Query {
			hello: String!
			identityFactoryAddress: String!
			getIdentityForAccount(account: String!): String!
			getIdentityABI: String!
			getIdentityFactoryABI: String!
		}

		type Mutation {
			createFactory: String!
			createIdentity(input: CreateIdentityInput): String!
			relayTransaction(input: RelayTransactionInput): Boolean!
		}

	`
	return graphql.MustParseSchema(s, &schemaResolver{identityService})

}
