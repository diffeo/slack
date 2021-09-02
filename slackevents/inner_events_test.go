package slackevents

import (
	"encoding/json"
	"testing"
)

func TestAppMention(t *testing.T) {
	rawE := []byte(`
			{
				"type": "app_mention",
				"user": "U061F7AUR",
				"text": "<@U0LAN0Z89> is it everything a river should be?",
				"ts": "1515449522.000016",
				"thread_ts": "1515449522.000016",
				"channel": "C0LAN2Q65",
				"event_ts": "1515449522000016",
				"source_team": "T3MQV36V7",
				"user_team": "T3MQV36V7",
				"blah": "test"
		}
	`)
	err := json.Unmarshal(rawE, &AppMentionEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestAppUninstalled(t *testing.T) {
	rawE := []byte(`
		{
			"type": "app_uninstalled"
		}
	`)
	err := json.Unmarshal(rawE, &AppUninstalledEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestFileChangeEvent(t *testing.T) {
	rawE := []byte(`
		{
			"type": "file_change",
			"file_id": "F1234567890",
			"file": {
				"id": "F1234567890"
			}
		}
	`)

	var e FileChangeEvent
	if err := json.Unmarshal(rawE, &e); err != nil {
		t.Fatal(err)
	}
	if e.Type != "file_change" {
		t.Errorf("type should be file_change, was %s", e.Type)
	}
	if e.FileID != "F1234567890" {
		t.Errorf("file ID should be F1234567890, was %s", e.FileID)
	}
	if e.File.ID != "F1234567890" {
		t.Errorf("file.id should be F1234567890, was %s", e.File.ID)
	}
}

func TestFileDeletedEvent(t *testing.T) {
	rawE := []byte(`
		{
			"type": "file_deleted",
			"file_id": "F1234567890",
			"event_ts": "1234567890.123456"
		}
	`)

	var e FileDeletedEvent
	if err := json.Unmarshal(rawE, &e); err != nil {
		t.Fatal(err)
	}
	if e.Type != "file_deleted" {
		t.Errorf("type should be file_deleted, was %s", e.Type)
	}
	if e.FileID != "F1234567890" {
		t.Errorf("file ID should be F1234567890, was %s", e.FileID)
	}
	if e.EventTimestamp != "1234567890.123456" {
		t.Errorf("event timestamp should be 1234567890.123456, was %s", e.EventTimestamp)
	}
}

func TestFileSharedEvent(t *testing.T) {
	rawE := []byte(`
		{
			"type": "file_shared",
			"channel_id": "C1234567890",
			"file_id": "F1234567890",
			"user_id": "U11235813",
			"file": {
				"id": "F1234567890"
			},
			"event_ts": "1234567890.123456"
		}
	`)

	var e FileSharedEvent
	if err := json.Unmarshal(rawE, &e); err != nil {
		t.Fatal(err)
	}
	if e.Type != "file_shared" {
		t.Errorf("type should be file_shared, was %s", e.Type)
	}
	if e.ChannelID != "C1234567890" {
		t.Errorf("channel ID should be C1234567890, was %s", e.ChannelID)
	}
	if e.FileID != "F1234567890" {
		t.Errorf("file ID should be F1234567890, was %s", e.FileID)
	}
	if e.UserID != "U11235813" {
		t.Errorf("user ID should be U11235813, was %s", e.UserID)
	}
	if e.File.ID != "F1234567890" {
		t.Errorf("file.id should be F1234567890, was %s", e.File.ID)
	}
	if e.EventTimestamp != "1234567890.123456" {
		t.Errorf("event timestamp should be 1234567890.123456, was %s", e.EventTimestamp)
	}
}

func TestFileUnsharedEvent(t *testing.T) {
	rawE := []byte(`
		{
			"type": "file_unshared",
			"file_id": "F1234567890",
			"file": {
				"id": "F1234567890"
			}
		}
	`)

	var e FileUnsharedEvent
	if err := json.Unmarshal(rawE, &e); err != nil {
		t.Fatal(err)
	}
	if e.Type != "file_unshared" {
		t.Errorf("type should be file_shared, was %s", e.Type)
	}
	if e.FileID != "F1234567890" {
		t.Errorf("file ID should be F1234567890, was %s", e.FileID)
	}
	if e.File.ID != "F1234567890" {
		t.Errorf("file.id should be F1234567890, was %s", e.File.ID)
	}
}

func TestGridMigrationFinishedEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "grid_migration_finished",
				"enterprise_id": "EXXXXXXXX"
			}
	`)
	err := json.Unmarshal(rawE, &GridMigrationFinishedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestGridMigrationStartedEvent(t *testing.T) {
	rawE := []byte(`
			{
				"token": "XXYYZZ",
				"team_id": "TXXXXXXXX",
				"api_app_id": "AXXXXXXXXX",
				"event": {
						"type": "grid_migration_started",
						"enterprise_id": "EXXXXXXXX"
				},
				"type": "event_callback",
				"event_id": "EvXXXXXXXX",
				"event_time": 1234567890
		}
	`)
	err := json.Unmarshal(rawE, &GridMigrationStartedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestLinkSharedEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "link_shared",
				"channel": "Cxxxxxx",
				"user": "Uxxxxxxx",
				"message_ts": "123456789.9875",
				"thread_ts": "123456789.9876",
				"links":
						[
								{
										"domain": "example.com",
										"url": "https://example.com/12345"
								},
								{
										"domain": "example.com",
										"url": "https://example.com/67890"
								},
								{
										"domain": "another-example.com",
										"url": "https://yet.another-example.com/v/abcde"
								}
						]
		}
	`)
	err := json.Unmarshal(rawE, &LinkSharedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestMessageEvent(t *testing.T) {
	rawE := []byte(`
			{
				"client_msg_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
				"type": "message",
				"channel": "G024BE91L",
				"user": "U2147483697",
				"text": "Live long and prospect.",
				"ts": "1355517523.000005",
				"event_ts": "1355517523.000005",
				"channel_type": "channel",
				"source_team": "T3MQV36V7",
				"user_team": "T3MQV36V7",
				"message": {
					"text": "To infinity and beyond.",
					"edited": {
						"user": "U2147483697",
						"ts": "1355517524.000000"
					}
				},
				"previous_message": {
					"text": "Live long and prospect."
				}
		}
	`)
	err := json.Unmarshal(rawE, &MessageEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestBotMessageEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "message",
				"subtype": "bot_message",
				"ts": "1358877455.000010",
				"text": "Pushing is the answer",
				"bot_id": "BB12033",
				"username": "github",
				"icons": {}
		}
	`)
	err := json.Unmarshal(rawE, &MessageEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestThreadBroadcastEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "message",
				"subtype": "thread_broadcast",
				"channel": "G024BE91L",
				"user": "U2147483697",
				"text": "Live long and prospect.",
				"ts": "1355517523.000005",
				"event_ts": "1355517523.000005",
				"channel_type": "channel",
				"source_team": "T3MQV36V7",
				"user_team": "T3MQV36V7",
				"message": {
					"text": "To infinity and beyond.",
					"root": {
						"text": "To infinity and beyond.",
						"ts": "1355517523.000005"
					},
					"edited": {
						"user": "U2147483697",
						"ts": "1355517524.000000"
					}
				},
				"previous_message": {
					"text": "Live long and prospect."
				}
		}
	`)

	var me MessageEvent
	if err := json.Unmarshal(rawE, &me); err != nil {
		t.Error(err)
	}

	if me.Root != nil {
		t.Error("me.Root should be nil")
	}

	if me.Message.Root == nil {
		t.Fatal("me.Message.Root is nil")
	}

	if me.Message.Root.TimeStamp != "1355517523.000005" {
		t.Errorf("me.Message.Root.TimeStamp = %q, want %q", me.Root.TimeStamp, "1355517523.000005")
	}
}

func TestMemberJoinedChannelEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "member_joined_channel",
				"user": "W06GH7XHN",
				"channel": "C0698JE0H",
				"channel_type": "C",
				"team": "T024BE7LD",
				"inviter": "U123456789"
		}
	`)
	err := json.Unmarshal(rawE, &MemberJoinedChannelEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestMemberLeftChannelEvent(t *testing.T) {
	rawE := []byte(`
			{
				"type": "member_left_channel",
				"user": "W06GH7XHN",
				"channel": "C0698JE0H",
				"channel_type": "C",
				"team": "T024BE7LD"
		}
	`)
	err := json.Unmarshal(rawE, &MemberLeftChannelEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestPinAdded(t *testing.T) {
	rawE := []byte(`
			{
				"type": "pin_added",
				"user": "U061F7AUR",
				"item": {
					"type": "message",
					"channel":"C0LAN2Q65",
					"message":{
						"type":"message",
						"user":"U061F7AUR",
						"text": "<@U0LAN0Z89> is it everything a river should be?",
						"ts":"1539904112.000100",
						"pinned_to":["C0LAN2Q65"],
						"replace_original":false,
						"delete_original":false
					}
				},
				"channel_id":"C0LAN2Q65",
				"event_ts": "1515449522000016"
		}
	`)
	err := json.Unmarshal(rawE, &PinAddedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestPinRemoved(t *testing.T) {
	rawE := []byte(`
			{
				"type": "pin_removed",
				"user": "U061F7AUR",
				"item": {
					"type": "message",
					"channel":"C0LAN2Q65",
					"message":{
						"type":"message",
						"user":"U061F7AUR",
						"text": "<@U0LAN0Z89> is it everything a river should be?",
						"ts":"1539904112.000100",
						"pinned_to":["C0LAN2Q65"],
						"replace_original":false,
						"delete_original":false
					}
				},
				"channel_id":"C0LAN2Q65",
				"event_ts": "1515449522000016"
		}
	`)
	err := json.Unmarshal(rawE, &PinRemovedEvent{})
	if err != nil {
		t.Error(err)
	}
}

func TestTokensRevoked(t *testing.T) {
	rawE := []byte(`
	{
		"type": "tokens_revoked",
		"tokens": {
				"oauth": [
						"OUXXXXXXXX"
				],
				"bot": [
						"BUXXXXXXXX"
				]
		}
	}
`)
	tre := TokensRevokedEvent{}
	err := json.Unmarshal(rawE, &tre)
	if err != nil {
		t.Error(err)
	}

	if tre.Type != "tokens_revoked" {
		t.Fail()
	}

	if len(tre.Tokens.Bot) != 1 || tre.Tokens.Bot[0] != "BUXXXXXXXX" {
		t.Fail()
	}

	if len(tre.Tokens.Oauth) != 1 || tre.Tokens.Oauth[0] != "OUXXXXXXXX" {
		t.Fail()
	}
}

func TestEmojiChanged(t *testing.T) {
	var (
		ece EmojiChangedEvent
		err error
	)

	// custom emoji added event
	rawAddE := []byte(`
	{
		"type": "emoji_changed",
		"subtype": "add",
		"name": "picard_facepalm",
		"value": "https://my.slack.com/emoji/picard_facepalm/db8e287430eaa459.gif",
		"event_ts" : "1361482916.000004"
	}
`)
	ece = EmojiChangedEvent{}
	err = json.Unmarshal(rawAddE, &ece)
	if err != nil {
		t.Error(err)
	}
	if ece.Subtype != "add" {
		t.Fail()
	}
	if ece.Name != "picard_facepalm" {
		t.Fail()
	}

	// emoji removed event
	rawRemoveE := []byte(`
	{
		"type": "emoji_changed",
		"subtype": "remove",
		"names": ["picard_facepalm"],
		"event_ts" : "1361482916.000004"
	}
`)
	ece = EmojiChangedEvent{}
	err = json.Unmarshal(rawRemoveE, &ece)
	if err != nil {
		t.Error(err)
	}
	if ece.Subtype != "remove" {
		t.Fail()
	}
	if len(ece.Names) != 1 {
		t.Fail()
	}
	if ece.Names[0] != "picard_facepalm" {
		t.Fail()
	}

	// custom emoji rename event
	rawRenameE := []byte(`
	{
		"type": "emoji_changed",
		"subtype": "rename",
		"old_name": "grin",
		"new_name": "cheese-grin",
		"value": "https://my.slack.com/emoji/picard_facepalm/db8e287430eaa459.gif",
		"event_ts" : "1361482916.000004"
	}
`)
	ece = EmojiChangedEvent{}
	err = json.Unmarshal(rawRenameE, &ece)
	if err != nil {
		t.Error(err)
	}
	if ece.Subtype != "rename" {
		t.Fail()
	}
	if ece.OldName != "grin" {
		t.Fail()
	}
	if ece.NewName != "cheese-grin" {
		t.Fail()
	}
}
