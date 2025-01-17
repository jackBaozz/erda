// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

import (
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/providers/component-protocol/components/list"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/common"
)

func (l *List) GenAppKvInfo(item apistructs.ApplicationDTO) (kvs []list.KvInfo) {
	var isPublic = "privateApp"
	var publicIcon = "private"
	if item.IsPublic {
		isPublic = "publicApp"
		publicIcon = "public"
	}
	updated := common.UpdatedTime(l.sdk.Ctx, item.UpdatedAt)
	runtimeUrlQuery, err := common.GenerateUrlQueryParams(map[string]interface{}{
		"app": []string{item.Name},
	})
	if err != nil {
		logrus.Errorf("run time url query encode failed, error: %v", err)
		panic(err)
	}
	kvs = []list.KvInfo{
		{
			Icon:  publicIcon,
			Value: l.sdk.I18n(isPublic),
			Tip:   l.sdk.I18n("publicProperty"),
		},
		{
			Icon:  "list-numbers",
			Tip:   l.sdk.I18n("runtime"),
			Value: strconv.Itoa(int(item.Stats.CountRuntimes)),
			Operations: map[cptype.OperationKey]cptype.Operation{
				list.OpItemClickGoto{}.OpKey(): cputil.NewOpBuilder().
					WithSkipRender(true).
					WithServerDataPtr(list.OpItemClickGotoServerData{
						OpItemBasicServerData: list.OpItemBasicServerData{
							Params: map[string]interface{}{
								"projectId": item.ProjectID,
								"env":       "dev",
							},
							Target: "projectDeployEnv",
							Query: map[string]interface{}{
								"advanceFilter__urlQuery": runtimeUrlQuery,
							},
						},
					}).
					Build(),
			},
		},
		{
			Icon:  "time",
			Tip:   l.sdk.I18n("updatedAt") + ": " + item.UpdatedAt.Format("2006-01-02 15:04:05"),
			Value: updated,
		},
		{
			Icon:  "category-management",
			Tip:   l.sdk.I18n("appType"),
			Value: l.sdk.I18n("appMode" + item.Mode),
		},
	}
	return
}
