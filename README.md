# go-openapi-generator-example

Personally I perfer this: https://github.com/deepmap/oapi-codegen.
But I found this at my job and decided to build some kind of example using this: https://github.com/OpenAPITools/openapi-generator

and maybe add something else.

## Generating the code

Be sure that you have installed docker and got pre-built docker image with openapi generator. You can get it there:
https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#table-of-contents

Use this command:
`docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/openapi/openapi.yaml \
    -g go-server \
    -o /local/go-openapi-check \
    --openapi-generator-ignore-list main.go,README.md,go.mod,internal \
    -c /local/openapi/generate-config.yaml`
