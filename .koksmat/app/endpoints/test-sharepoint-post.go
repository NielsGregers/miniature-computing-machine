// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: SharePoint Test
---
*/
package endpoints

import (
	"context"

	"github.com/swaggest/usecase"

	"github.com/365admin/koksmat-test1/execution"
)

func TestSharepointPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "koksmat-test1", "50-test", "11-sharepoint.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("SharePoint Test")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Test")
	return u
}
