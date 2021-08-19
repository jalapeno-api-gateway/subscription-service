package arangodb

import (
	"context"
	"log"
)

func FetchLsNode(ctx context.Context, key string) LsNodeDocument {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNode FILTER d._key == \"" + key + "\" RETURN d");
	var document LsNodeDocument
	_, err := cursor.ReadDocument(ctx, &document)
	if err != nil {
		log.Fatalf("Could not fetch Node from LSNode Collection, %v", err)
	}
	return document
}

func FetchLsLink(ctx context.Context, key string) LsLinkDocument {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSLink FILTER d._key == \"" + key + "\" RETURN d");
	var document LsLinkDocument
	_, err := cursor.ReadDocument(ctx, &document)
	if err != nil {
		log.Fatalf("Could not fetch Link from LSLink Collection, %v", err)
	}
	return document
}