# description
Include a specification of some kind describing the RESTful API. Specification of choice would be OpenApi 3.x. OpenApi 3.x provides extensive support for describing RESTful APIs. There are many tools and frameworks that support auto generating client packages.

# Notes and thoughts
The api.yaml is incomplete. I would review the api.yaml with the product team and the various stakeholders to ask any questions I might have about the different properties.

### Example questions for stakeholders
    - How long can a description be?
    - How should the sort work. What order can they sort the services on.
    - How exactly should the search capabilities work. Does the search only apply to the name property. Should we search within the description etc.

# Testing
My environment has not been cooperative with me for testing. I am working off a VM that I have brought up as I don't have a personal environment currently. Using a more inplace development process and strategy agreed upon by multiple engineers would support faster iteration over development. Some benefits of having an integrated development process are other developers checking running changes to the applications locally in order to evaluate the components being changed. I unfortunately did not have the time to dedicate time to adding many tests due to environment issues and I figured that it would be different than what is currently in place for the team.

# Authorization
Generally authorization is applied within the context of the given operation being performed. There are easy authorizations to check what the authenticated principal has access to including scope and role validations following the OAuth2 and OIDC protocols. There are also more indepth context related authorizations like returning a subset of the resources in a collection based on the users id, group, or organization.