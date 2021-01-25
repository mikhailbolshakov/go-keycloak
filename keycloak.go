package main

import "github.com/Nerzal/gocloak/v7"

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

func newKeycloak() *keycloak {
	return &keycloak{
		gocloak:      gocloak.NewClient("http://localhost:8086"),
		clientId:     "my-go-service",
		clientSecret: "abfa2984-9125-486b-b360-03386ad13e08",
		realm:        "medium",
	}
}
