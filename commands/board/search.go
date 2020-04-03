// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package board

import (
	"context"

	"github.com/arduino/arduino-cli/commands"
	rpc "github.com/arduino/arduino-cli/rpc/commands"
	"github.com/pkg/errors"
)

// Search FIXMEDOC
func Search(ctx context.Context, req *rpc.BoardSearchReq) (r *rpc.BoardSearchResp, e error) {
	pm := commands.GetPackageManager(req.GetInstance().GetId())
	if pm == nil {
		return nil, errors.New("invalid instance")
	}

	boards := pm.Registry.SearchBoards(req.GetQuery())
	resBoards := []*rpc.SearchedBoard{}
	for _, board := range boards {
		extURL := ""
		if board.ExternalPlatformURL != nil {
			extURL = board.ExternalPlatformURL.String()
		}
		resBoards = append(resBoards, &rpc.SearchedBoard{
			Name:                board.Name,
			Alias:               board.Alias,
			Fqbn:                board.FQBN.String(),
			ExternalPlatformUrl: extURL,
		})
	}
	res := &rpc.BoardSearchResp{
		Boards: resBoards,
	}
	return res, nil
}
