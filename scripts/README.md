# Useful Scripts

## Glossary:

1. `generate.sh`  -- This is used to generate graphql endpoint methods from the graphql schema files in the `schema/` directory. Think of it like your proto, but this time for graphqls.
1. `mage` -- A "makefile" for go. 
1. `protogen.go` -- Executed using mage, to generate go code from the protos of each gRPC service dependency in a concise form. All outputs can be found in `services/` directory. 