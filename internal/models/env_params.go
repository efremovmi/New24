package models

type Config struct {
	ADDR_AUTH                 interface{}
	PROTOCOL_WITH_DOMAIN_AUTH interface{}

	ADDR_CONTROL_USERS                 interface{}
	PROTOCOL_WITH_DOMAIN_CONTROL_USERS interface{}

	POSTGRES_HOST        interface{}
	POSTGRES_PORT        interface{}
	POSTGRES_USER        interface{}
	POSTGRES_BD_NAME     interface{}
	POSTGRES_PASSWORD    interface{}
	POSTGRES_TABLE_USERS interface{}

	HASH_SALT interface{}
}
