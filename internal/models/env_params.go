package models

type Config struct {
	ADR_AUTH                  interface{}
	PROTOCOL_WITH_DOMAIN_AUTH interface{}

	ADR_CONTROL_USERS                  interface{}
	PROTOCOL_WITH_DOMAIN_CONTROL_USERS interface{}

	ADR_NEWS                  interface{}
	PROTOCOL_WITH_DOMAIN_NEWS interface{}

	POSTGRES_HOST        interface{}
	POSTGRES_PORT        interface{}
	POSTGRES_USER        interface{}
	POSTGRES_BD_NAME     interface{}
	POSTGRES_PASSWORD    interface{}
	POSTGRES_TABLE_USERS interface{}

	POSTGRES_TABLE_NEWS interface{}

	HASH_SALT interface{}
}
