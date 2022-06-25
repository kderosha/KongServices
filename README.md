# KongServices
An application providing an implementation for the take home.

# If I had the time
### Authentication
I would have added a piece of middleware for the service that provides authentication via whatever the requirements would define. Most notably I'm very proficient in OAuth2 and OIDC integrations and prefer authentication to be handled by an IdP as in house authentication is expensive to maintain. 

### Code generation
I would have flushed out the api.yaml further to fully describe the RESTful API and all the capabilities through the OpenApi specification. There are multiple code generation tools that allow developers to generate server stubs based on the specification defined. This requires an agreed upon specification between developers and product stakeholders based on the requirements give.

OpenApi also provides the ability to generate client packages. I normally wrap the client packages in an SDK that is then available for integration with common dependency management tools.