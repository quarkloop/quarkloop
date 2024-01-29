package main

import (
	"context"
	"errors"
	"io"
	"log"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var schema = `
definition org {
    relation admin: user
    relation member: user

    permission read = admin + member
    permission create = admin
    permission update = admin
    permission delete = admin
    permission * = read + create + update + delete
}
`

func main() {
	// alice := &v1.SubjectReference{Object: &v1.ObjectReference{
	// 	ObjectType: "user",
	// 	ObjectId:   "alice",
	// }}

	// devs := &v1.ObjectReference{
	// 	ObjectType: "group",
	// 	ObjectId:   "devs",
	// }

	client, err := authzed.NewClient(
		"localhost:50051",
		grpcutil.WithInsecureBearerToken("my_passphrase_key"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	wResp, err := client.WriteSchema(context.Background(), &v1.WriteSchemaRequest{
		Schema: schema,
	})
	if err != nil {
		//log.Fatalf("failed to write schema: %s", err)
	}
	wResp.GetWrittenAt()

	// rResp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
	// 	Resource:   devs,
	// 	Permission: "member",
	// 	Subject:    alice,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to check permission: %s", err)
	// }

	// if rResp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
	// 	log.Println("allowed!")
	// }

	resp, err := client.LookupSubjects(context.Background(), &v1.LookupSubjectsRequest{
		Consistency: &v1.Consistency{
			Requirement: &v1.Consistency_FullyConsistent{
				FullyConsistent: true,
			},
		},
		Resource: &v1.ObjectReference{
			ObjectType: "workspace",
			ObjectId:   "it_department",
		},
		Permission:        "create",
		SubjectObjectType: "user",
	})
	if err != nil {
		log.Fatalf("failed to LookupResources: %s", err)
	}

	found := []*v1.LookupSubjectsResponse{}
	for {
		resp, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalf("failed to recv: %s", err)
		}

		found = append(found, resp)
	}

	log.Println("LookupResources", found)
}
