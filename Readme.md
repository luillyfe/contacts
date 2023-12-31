# Contacts API

This repository contains a simple API that allows users to create, read, update, and delete contacts.

## Features

- Create a new contact
- Get a specific contact by ID
- Update a specific contact by ID
- Delete a specific contact by ID

## Installation

To install the API, you will need to have Go installed. You can then install the dependencies by running the following command:

`go get .`

## Usage

To use the API, you will need to make HTTP requests to the following endpoints:

- `/contacts`: Get a list of all contacts
- `/contacts/:id`: Get a specific contact by ID
- `/contacts`: Create a new contact
- `/contacts/:id`: Update a specific contact by ID
- `/contacts/:id`: Delete a specific contact by ID

The requests should be made in JSON format. The response will also be in JSON format.

## Security

It has been added a restriction access to just acept request that includes an API key. You must be able to create yours, from the [APIs & Services](https://console.cloud.google.com/apis/dashboard) section in Google Cloud Console. When generating the API key, make sure to limit it to just the API in development (Contacts API fro this case). [See building secure APIs in Google Cloud](https://medium.com/@luillyfe/securing-cloud-endpoints-with-api-keys-d16c00f77e5c).

## Documentation

For more documentation, please see the following:

- [README file](README.md)
- [Tutorial](tutorial.md)
- [Blog post](https://luillyfe.medium.com/building-rest-apis-in-google-cloud-6498aea274ea)

## Turorial

To deploy this to Cloud the API to Cloud Endpoints use,
`gcloud endpoints services deploy swagger.yaml`

To the deploy the backend API use,
`gcloud app deploy`

## Contributing

If you would like to contribute to the API, please fork the repository and submit a pull request.

## License

The API is licensed under the MIT License.

```

```
