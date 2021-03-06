// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package slackbot

import "github.com/nlopes/slack"

// LoadChannelIDs loads channel ids for the Slack client
func LoadChannelIDs(api slack.Client) (map[string]string, error) {
	channels, err := api.GetChannels(true)
	if err != nil {
		return nil, err
	}
	channelIDs := make(map[string]string)
	for _, c := range channels {
		channelIDs[c.Name] = c.ID
	}
	return channelIDs, nil
}
