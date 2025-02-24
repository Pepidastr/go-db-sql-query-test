package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {

	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	clientID := 1
	client, err := selectClient(db, clientID)
	require.NoError(t, err)

	assert.Equal(t, client.ID, clientID)
	assert.NotEmpty(t, map[string]string{
		"FIO":      client.FIO,
		"Login":    client.Login,
		"Birthday": client.Birthday,
		"Email":    client.Email,
	})

}

func Test_SelectClient_WhenNoClient(t *testing.T) {

	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	clientID := -1
	client, err := selectClient(db, clientID)

	require.Equal(t, err, sql.ErrNoRows)

	assert.Empty(t, client)

}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	clSelect, err := selectClient(db, cl.ID)
	require.NoError(t, err)

	assert.Equal(t, cl, clSelect)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {

	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	id, err := insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Equal(t, err, sql.ErrNoRows)

}
