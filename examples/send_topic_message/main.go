/*
Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ldnvnbl/hms-push-serverdemo-go/examples/common"
	"github.com/ldnvnbl/hms-push-serverdemo-go/push/constant"
	"github.com/ldnvnbl/hms-push-serverdemo-go/push/model"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic! Error is %s\n", err)
		}
	}()

	sendTopicMessage()
}

func sendTopicMessage() {
	msgRequest, err := getTopicMsgRequest()
	if err != nil {
		fmt.Printf("Failed to get message request! Error is %s\n", err.Error())
		return
	}

	client := common.GetPushClient()

	resp, err := client.SendMessage(context.Background(), msgRequest)
	if err != nil {
		fmt.Printf("Failed to send message! Error is %s\n", err.Error())
		return
	}

	if resp.Code != constant.Success {
		fmt.Printf("Failed to send message! Response is %+v\n", resp)
		return
	}

	fmt.Printf("Succeed to send message! Response is %+v\n", resp)
}

func getTopicMsgRequest() (*model.MessageRequest, error) {
	msgRequest := model.NewNotificationMsgRequest()
	msgRequest.Message.Topic = common.TargetTopic
	msgRequest.Message.Android = model.GetDefaultAndroid()
	msgRequest.Message.Data = ""
	msgRequest.Message.Android.Notification = model.GetDefaultAndroidNotification()

	b, err := json.Marshal(msgRequest)
	if err != nil {
		fmt.Printf("Failed to marshal the default message! Error is %s\n", err.Error())
		return nil, err
	}

	fmt.Printf("Default message is %s\n", string(b))
	return msgRequest, nil
}
