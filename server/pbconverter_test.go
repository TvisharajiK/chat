package main

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pbx "github.com/tinode/chat/pbx"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestPbServCtrlSerialize_ValidInput_001(t *testing.T) {
	ctrl := &MsgServerCtrl{
		Id:     "test-id",
		Topic:  "test-topic",
		Params: map[string]any{"key": "value"},
		Code:   200,
		Text:   "OK",
	}

	result := pbServCtrlSerialize(ctrl)

	require.NotNil(t, result)
	require.NotNil(t, result.Ctrl)
	assert.Equal(t, ctrl.Id, result.Ctrl.Id)
	assert.Equal(t, ctrl.Topic, result.Ctrl.Topic)
	assert.Equal(t, int32(ctrl.Code), result.Ctrl.Code)
	assert.Equal(t, ctrl.Text, result.Ctrl.Text)
	assert.NotNil(t, result.Ctrl.Params)
	assert.Equal(t, []byte(`"value"`), result.Ctrl.Params["key"])
}

// Test generated using Keploy
func TestPbServDataSerialize_ValidInput_002(t *testing.T) {
	timestamp := time.Now()
	data := &MsgServerData{
		Topic:     "test-topic",
		From:      "test-user",
		Timestamp: timestamp,
		DeletedAt: &timestamp,
		SeqId:     123,
		Head:      map[string]any{"header-key": "header-value"},
		Content:   "test-content",
	}

	result := pbServDataSerialize(data)

	require.NotNil(t, result)
	require.NotNil(t, result.Data)
	assert.Equal(t, data.Topic, result.Data.Topic)
	assert.Equal(t, data.From, result.Data.FromUserId)
	assert.Equal(t, timeToInt64(&data.Timestamp), result.Data.Timestamp)
	assert.Equal(t, timeToInt64(data.DeletedAt), result.Data.DeletedAt)
	assert.Equal(t, int32(data.SeqId), result.Data.SeqId)
	assert.NotNil(t, result.Data.Head)
	assert.Equal(t, []byte(`"header-value"`), result.Data.Head["header-key"])
	assert.Equal(t, []byte(`"test-content"`), result.Data.Content)
}

// Test generated using Keploy
func TestPbServPresSerialize_ValidPresence_003(t *testing.T) {
	pres := &MsgServerPres{
		Topic:     "test-topic",
		Src:       "test-src",
		What:      "on",
		UserAgent: "test-agent",
		SeqId:     123,
		DelId:     456,
		DelSeq:    []MsgDelRange{{LowId: 1, HiId: 10}},
		AcsTarget: "target-user",
		AcsActor:  "actor-user",
		Acs:       &MsgAccessMode{Want: "RW", Given: "R"},
	}

	result := pbServPresSerialize(pres)

	require.NotNil(t, result)
	require.NotNil(t, result.Pres)
	assert.Equal(t, pres.Topic, result.Pres.Topic)
	assert.Equal(t, pres.Src, result.Pres.Src)
	assert.Equal(t, pbx.ServerPres_ON, result.Pres.What)
	assert.Equal(t, pres.UserAgent, result.Pres.UserAgent)
	assert.Equal(t, int32(pres.SeqId), result.Pres.SeqId)
	assert.Equal(t, int32(pres.DelId), result.Pres.DelId)
	assert.NotNil(t, result.Pres.DelSeq)
	assert.Equal(t, int32(pres.DelSeq[0].LowId), result.Pres.DelSeq[0].Low)
	assert.Equal(t, int32(pres.DelSeq[0].HiId), result.Pres.DelSeq[0].Hi)
	assert.Equal(t, pres.AcsTarget, result.Pres.TargetUserId)
	assert.Equal(t, pres.AcsActor, result.Pres.ActorUserId)
	assert.NotNil(t, result.Pres.Acs)
	assert.Equal(t, pres.Acs.Want, result.Pres.Acs.Want)
	assert.Equal(t, pres.Acs.Given, result.Pres.Acs.Given)
}

// Test generated using Keploy
func TestPbServSerialize_Ctrl_589(t *testing.T) {
	ctrl := &MsgServerCtrl{Id: "ctrl1", Code: 200, Text: "OK"}
	msg := &ServerComMessage{Ctrl: ctrl}

	result := pbServSerialize(msg)

	require.NotNil(t, result)
	assert.NotNil(t, result.GetCtrl())
	assert.Nil(t, result.GetData())
	assert.Nil(t, result.GetPres())
	assert.Nil(t, result.GetInfo())
	assert.Nil(t, result.GetMeta())
	assert.Equal(t, ctrl.Id, result.GetCtrl().Id)
	assert.Equal(t, int32(ctrl.Code), result.GetCtrl().Code)
}

// Test generated using Keploy
func TestPbServSerialize_Data_912(t *testing.T) {
	now := time.Now()
	data := &MsgServerData{Topic: "topic1", From: "user1", Timestamp: now, SeqId: 1}
	msg := &ServerComMessage{Data: data}

	result := pbServSerialize(msg)

	require.NotNil(t, result)
	assert.Nil(t, result.GetCtrl())
	assert.NotNil(t, result.GetData())
	assert.Nil(t, result.GetPres())
	assert.Nil(t, result.GetInfo())
	assert.Nil(t, result.GetMeta())
	assert.Equal(t, data.Topic, result.GetData().Topic)
	assert.Equal(t, data.From, result.GetData().FromUserId)
}

// Test generated using Keploy
func TestPbServSerialize_Pres_333(t *testing.T) {
	pres := &MsgServerPres{Topic: "topic2", Src: "user2", What: "on"}
	msg := &ServerComMessage{Pres: pres}

	result := pbServSerialize(msg)

	require.NotNil(t, result)
	assert.Nil(t, result.GetCtrl())
	assert.Nil(t, result.GetData())
	assert.NotNil(t, result.GetPres())
	assert.Nil(t, result.GetInfo())
	assert.Nil(t, result.GetMeta())
	assert.Equal(t, pres.Topic, result.GetPres().Topic)
	assert.Equal(t, pbx.ServerPres_ON, result.GetPres().What)
}

// Test generated using Keploy
func TestPbServSerialize_Info_747(t *testing.T) {
	info := &MsgServerInfo{Topic: "topic3", From: "user3", What: "read", SeqId: 5}
	msg := &ServerComMessage{Info: info}

	result := pbServSerialize(msg)

	require.NotNil(t, result)
	assert.Nil(t, result.GetCtrl())
	assert.Nil(t, result.GetData())
	assert.Nil(t, result.GetPres())
	assert.NotNil(t, result.GetInfo())
	assert.Nil(t, result.GetMeta())
	assert.Equal(t, info.Topic, result.GetInfo().Topic)
	assert.Equal(t, pbx.InfoNote_READ, result.GetInfo().What)
}

// Test generated using Keploy
func TestPbServSerialize_Meta_109(t *testing.T) {
	meta := &MsgServerMeta{Id: "meta1", Topic: "topic4"}
	msg := &ServerComMessage{Meta: meta}

	result := pbServSerialize(msg)

	require.NotNil(t, result)
	assert.Nil(t, result.GetCtrl())
	assert.Nil(t, result.GetData())
	assert.Nil(t, result.GetPres())
	assert.Nil(t, result.GetInfo())
	assert.NotNil(t, result.GetMeta())
	assert.Equal(t, meta.Id, result.GetMeta().Id)
	assert.Equal(t, meta.Topic, result.GetMeta().Topic)
}

// Test generated using Keploy
func TestPbServDeserialize_Ctrl_661(t *testing.T) {
	params := map[string][]byte{"key1": []byte(`"value1"`), "key2": []byte(`123`)}
	ctrl := &pbx.ServerCtrl{
		Id:     "ctrl-deser",
		Topic:  "topic-deser",
		Code:   404,
		Text:   "Not Found",
		Params: params,
	}
	pkt := &pbx.ServerMsg{Message: &pbx.ServerMsg_Ctrl{Ctrl: ctrl}}

	msg := pbServDeserialize(pkt)

	require.NotNil(t, msg)
	require.NotNil(t, msg.Ctrl)
	assert.Nil(t, msg.Data)
	assert.Nil(t, msg.Pres)
	assert.Nil(t, msg.Info)
	assert.Nil(t, msg.Meta)

	assert.Equal(t, ctrl.Id, msg.Ctrl.Id)
	assert.Equal(t, ctrl.Topic, msg.Ctrl.Topic)
	assert.Equal(t, int(ctrl.Code), msg.Ctrl.Code)
	assert.Equal(t, ctrl.Text, msg.Ctrl.Text)
	require.NotNil(t, msg.Ctrl.Params)
	require.IsType(t, map[string]any{}, msg.Ctrl.Params)
	p, ok := msg.Ctrl.Params.(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "value1", p["key1"])
	// Note: JSON unmarshals numbers into float64 by default
	assert.Equal(t, float64(123), p["key2"])
}

// Test generated using Keploy
func TestPbServDeserialize_Pres_AllCases_173(t *testing.T) {
	testCases := []struct {
		name      string
		whatPbx   pbx.ServerPres_What
		whatStr   string
		hasAcs    bool
		hasDelSeq bool
	}{
		{"on", pbx.ServerPres_ON, "on", true, true},
		{"off", pbx.ServerPres_OFF, "off", false, false},
		{"ua", pbx.ServerPres_UA, "ua", true, false},
		{"upd", pbx.ServerPres_UPD, "upd", false, true},
		{"gone", pbx.ServerPres_GONE, "gone", true, true},
		{"acs", pbx.ServerPres_ACS, "acs", true, false},
		{"term", pbx.ServerPres_TERM, "term", false, false},
		{"msg", pbx.ServerPres_MSG, "msg", false, true},
		{"read", pbx.ServerPres_READ, "read", true, true},
		{"recv", pbx.ServerPres_RECV, "recv", false, false},
		{"del", pbx.ServerPres_DEL, "del", true, false},
		{"tags", pbx.ServerPres_TAGS, "tags", false, true},
		{"default", pbx.ServerPres_What(99), "", true, true}, // Unknown value
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pres := &pbx.ServerPres{
				Topic:        "pres-topic",
				Src:          "pres-src",
				What:         tc.whatPbx,
				UserAgent:    "pres-agent",
				SeqId:        1,
				DelId:        2,
				TargetUserId: "target",
				ActorUserId:  "actor",
			}
			if tc.hasAcs {
				pres.Acs = &pbx.AccessMode{Want: "W", Given: "N"}
			}
			if tc.hasDelSeq {
				pres.DelSeq = []*pbx.SeqRange{{Low: 5, Hi: 10}}
			}

			pkt := &pbx.ServerMsg{Message: &pbx.ServerMsg_Pres{Pres: pres}}
			msg := pbServDeserialize(pkt)

			require.NotNil(t, msg)
			require.NotNil(t, msg.Pres)
			assert.Equal(t, pres.Topic, msg.Pres.Topic)
			assert.Equal(t, pres.Src, msg.Pres.Src)
			assert.Equal(t, tc.whatStr, msg.Pres.What)
			assert.Equal(t, pres.UserAgent, msg.Pres.UserAgent)
			assert.Equal(t, int(pres.SeqId), msg.Pres.SeqId)
			assert.Equal(t, int(pres.DelId), msg.Pres.DelId)
			assert.Equal(t, pres.TargetUserId, msg.Pres.AcsTarget)
			assert.Equal(t, pres.ActorUserId, msg.Pres.AcsActor)

			if tc.hasAcs {
				require.NotNil(t, msg.Pres.Acs)
				assert.Equal(t, pres.Acs.Want, msg.Pres.Acs.Want)
				assert.Equal(t, pres.Acs.Given, msg.Pres.Acs.Given)
			} else {
				assert.Nil(t, msg.Pres.Acs)
			}

			if tc.hasDelSeq {
				require.NotNil(t, msg.Pres.DelSeq)
				require.Len(t, msg.Pres.DelSeq, 1)
				assert.Equal(t, int(pres.DelSeq[0].Low), msg.Pres.DelSeq[0].LowId)
				assert.Equal(t, int(pres.DelSeq[0].Hi), msg.Pres.DelSeq[0].HiId)
			} else {
				assert.Nil(t, msg.Pres.DelSeq)
			}
		})
	}
}

// Test generated using Keploy
func TestPbServDeserialize_Meta_025(t *testing.T) {
	now := time.Now()
	ts := now.UnixNano() / int64(time.Millisecond)
	desc := &pbx.TopicDesc{
		CreatedAt:         ts - 100000,
		UpdatedAt:         ts - 50000,
		TouchedAt:         ts - 10000,
		State:             "ok",
		Online:            true,
		IsChan:            true,
		Defacs:            &pbx.DefaultAcsMode{Auth: "RW", Anon: "R"},
		Acs:               &pbx.AccessMode{Want: "J", Given: "P"},
		SeqId:             100,
		ReadId:            90,
		RecvId:            95,
		DelId:             5,
		Public:            []byte(`{"desc":"public desc"}`),
		Trusted:           []byte(`{"desc":"trusted desc"}`),
		Private:           []byte(`{"desc":"private desc"}`),
		LastSeenTime:      ts - 20000,
		LastSeenUserAgent: "meta-agent",
	}
	sub := []*pbx.TopicSub{
		{
			UpdatedAt:         ts - 40000,
			DeletedAt:         0,
			Online:            false,
			Acs:               &pbx.AccessMode{Want: "S", Given: "O"},
			ReadId:            80,
			RecvId:            85,
			Public:            []byte(`{"sub":"public sub"}`),
			Trusted:           []byte(`{"sub":"trusted sub"}`),
			Private:           []byte(`{"sub":"private sub"}`),
			UserId:            "sub-user",
			Topic:             "meta-topic",
			TouchedAt:         ts - 30000,
			SeqId:             99,
			DelId:             4,
			LastSeenTime:      ts - 15000,
			LastSeenUserAgent: "sub-agent",
		},
	}
	del := &pbx.DelValues{DelId: 5, DelSeq: []*pbx.SeqRange{{Low: 1, Hi: 3}}}
	cred := []*pbx.ServerCred{{Method: "tel", Value: "+123", Done: true}}
	meta := &pbx.ServerMeta{
		Id:    "meta-deser",
		Topic: "meta-topic",
		Desc:  desc,
		Sub:   sub,
		Del:   del,
		Tags:  []string{"meta-tag1", "meta-tag2"},
		Cred:  cred,
	}
	pkt := &pbx.ServerMsg{Message: &pbx.ServerMsg_Meta{Meta: meta}}

	msg := pbServDeserialize(pkt)

	require.NotNil(t, msg)
	require.NotNil(t, msg.Meta)

	// Assert Meta fields
	assert.Equal(t, meta.Id, msg.Meta.Id)
	assert.Equal(t, meta.Topic, msg.Meta.Topic)
	assert.Equal(t, meta.Tags, msg.Meta.Tags)

	// Assert Desc
	require.NotNil(t, msg.Meta.Desc)
	assert.WithinDuration(t, int64ToTime(desc.CreatedAt).UTC(), *msg.Meta.Desc.CreatedAt, time.Millisecond)
	assert.WithinDuration(t, int64ToTime(desc.UpdatedAt).UTC(), *msg.Meta.Desc.UpdatedAt, time.Millisecond)
	assert.WithinDuration(t, int64ToTime(desc.TouchedAt).UTC(), *msg.Meta.Desc.TouchedAt, time.Millisecond)
	assert.Equal(t, desc.State, msg.Meta.Desc.State)
	assert.Equal(t, desc.Online, msg.Meta.Desc.Online)
	assert.Equal(t, desc.IsChan, msg.Meta.Desc.IsChan)
	require.NotNil(t, msg.Meta.Desc.DefaultAcs)
	assert.Equal(t, desc.Defacs.Auth, msg.Meta.Desc.DefaultAcs.Auth)
	assert.Equal(t, desc.Defacs.Anon, msg.Meta.Desc.DefaultAcs.Anon)
	require.NotNil(t, msg.Meta.Desc.Acs)
	assert.Equal(t, desc.Acs.Want, msg.Meta.Desc.Acs.Want)
	assert.Equal(t, desc.Acs.Given, msg.Meta.Desc.Acs.Given)
	assert.Equal(t, int(desc.SeqId), msg.Meta.Desc.SeqId)
	assert.Equal(t, int(desc.ReadId), msg.Meta.Desc.ReadSeqId)
	assert.Equal(t, int(desc.RecvId), msg.Meta.Desc.RecvSeqId)
	assert.Equal(t, int(desc.DelId), msg.Meta.Desc.DelId)
	assert.Equal(t, map[string]interface{}{"desc": "public desc"}, msg.Meta.Desc.Public)
	assert.Equal(t, map[string]interface{}{"desc": "trusted desc"}, msg.Meta.Desc.Trusted)
	assert.Equal(t, map[string]interface{}{"desc": "private desc"}, msg.Meta.Desc.Private)
	require.NotNil(t, msg.Meta.Desc.LastSeen)
	assert.WithinDuration(t, int64ToTime(desc.LastSeenTime).UTC(), *msg.Meta.Desc.LastSeen.When, time.Millisecond)
	assert.Equal(t, desc.LastSeenUserAgent, msg.Meta.Desc.LastSeen.UserAgent)

	// Assert Sub
	require.Len(t, msg.Meta.Sub, 1)
	msgSub := msg.Meta.Sub[0]
	pbSub := sub[0]
	assert.WithinDuration(t, int64ToTime(pbSub.UpdatedAt).UTC(), *msgSub.UpdatedAt, time.Millisecond)
	assert.Nil(t, msgSub.DeletedAt)
	assert.Equal(t, pbSub.Online, msgSub.Online)
	require.NotNil(t, msgSub.Acs) // Acs is value type in MsgTopicSub
	assert.Equal(t, pbSub.Acs.Want, msgSub.Acs.Want)
	assert.Equal(t, pbSub.Acs.Given, msgSub.Acs.Given)
	assert.Equal(t, int(pbSub.ReadId), msgSub.ReadSeqId)
	assert.Equal(t, int(pbSub.RecvId), msgSub.RecvSeqId)
	assert.Equal(t, map[string]interface{}{"sub": "public sub"}, msgSub.Public)
	assert.Equal(t, map[string]interface{}{"sub": "trusted sub"}, msgSub.Trusted)
	assert.Equal(t, map[string]interface{}{"sub": "private sub"}, msgSub.Private)
	assert.Equal(t, pbSub.UserId, msgSub.User)
	assert.Equal(t, pbSub.Topic, msgSub.Topic)
	assert.WithinDuration(t, int64ToTime(pbSub.TouchedAt).UTC(), *msgSub.TouchedAt, time.Millisecond)
	assert.Equal(t, int(pbSub.SeqId), msgSub.SeqId)
	assert.Equal(t, int(pbSub.DelId), msgSub.DelId)
	require.NotNil(t, msgSub.LastSeen)
	assert.WithinDuration(t, int64ToTime(pbSub.LastSeenTime).UTC(), *msgSub.LastSeen.When, time.Millisecond)
	assert.Equal(t, pbSub.LastSeenUserAgent, msgSub.LastSeen.UserAgent)

	// Assert Del
	require.NotNil(t, msg.Meta.Del)
	assert.Equal(t, int(del.DelId), msg.Meta.Del.DelId)
	require.Len(t, msg.Meta.Del.DelSeq, 1)
	assert.Equal(t, int(del.DelSeq[0].Low), msg.Meta.Del.DelSeq[0].LowId)
	assert.Equal(t, int(del.DelSeq[0].Hi), msg.Meta.Del.DelSeq[0].HiId)

	// Assert Cred
	require.Len(t, msg.Meta.Cred, 1)
	assert.Equal(t, cred[0].Method, msg.Meta.Cred[0].Method)
	assert.Equal(t, cred[0].Value, msg.Meta.Cred[0].Value)
	assert.Equal(t, cred[0].Done, msg.Meta.Cred[0].Done)
}

// Test generated using Keploy
func TestPbCliSerialize_Hi_711(t *testing.T) {
	hi := &MsgClientHi{
		Id:         "hi-1",
		UserAgent:  "TinodeTest/1.0",
		Version:    "0.16",
		DeviceID:   "dev123",
		Platform:   "test",
		Lang:       "en-US",
		Background: false,
	}
	msg := &ClientComMessage{Hi: hi}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetHi())
	assert.Equal(t, hi.Id, result.GetHi().Id)
	assert.Equal(t, hi.UserAgent, result.GetHi().UserAgent)
	assert.Equal(t, hi.Version, result.GetHi().Ver)
	assert.Equal(t, hi.DeviceID, result.GetHi().DeviceId)
	assert.Equal(t, hi.Platform, result.GetHi().Platform)
	assert.Equal(t, hi.Lang, result.GetHi().Lang)
	assert.Equal(t, hi.Background, result.GetHi().Background)
	assert.Nil(t, result.Extra)
}

// Test generated using Keploy
func TestPbCliSerialize_Acc_AuthLevels_258(t *testing.T) {
	testCases := []struct {
		name      string
		authLevel string
		expected  pbx.AuthLevel
	}{
		{"none", "none", pbx.AuthLevel_NONE},
		{"NONE", "NONE", pbx.AuthLevel_NONE},
		{"empty", "", pbx.AuthLevel_NONE},
		{"anon", "anon", pbx.AuthLevel_ANON},
		{"ANON", "ANON", pbx.AuthLevel_ANON},
		{"auth", "auth", pbx.AuthLevel_AUTH},
		{"AUTH", "AUTH", pbx.AuthLevel_AUTH},
		{"root", "root", pbx.AuthLevel_NONE}, // root not supported here
		{"ROOT", "ROOT", pbx.AuthLevel_NONE},
		{"invalid", "invalid", pbx.AuthLevel_NONE}, // default case
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			acc := &MsgClientAcc{
				Id:        "acc-1",
				User:      "usrTest",
				AuthLevel: tc.authLevel,
				Scheme:    "basic",
				Secret:    []byte("password"),
				Tags:      []string{"t1", "t2"},
				Desc:      &MsgSetDesc{Public: "Public Desc"},
				Cred:      []MsgCredClient{{Method: "email", Value: "e@example.com"}},
			}
			msg := &ClientComMessage{Acc: acc}

			result := pbCliSerialize(msg)
			require.NotNil(t, result)
			require.NotNil(t, result.GetAcc())
			assert.Equal(t, tc.expected, result.GetAcc().AuthLevel)
			assert.Equal(t, acc.Id, result.GetAcc().Id)
			assert.Equal(t, acc.User, result.GetAcc().UserId)
			assert.Equal(t, acc.Scheme, result.GetAcc().Scheme)
			assert.Equal(t, acc.Secret, result.GetAcc().Secret)
			assert.Equal(t, acc.Tags, result.GetAcc().Tags)
			require.NotNil(t, result.GetAcc().Desc)
			assert.Equal(t, `"`+acc.Desc.Public.(string)+`"`, string(result.GetAcc().Desc.Public))
			require.Len(t, result.GetAcc().Cred, 1)
			assert.Equal(t, acc.Cred[0].Method, result.GetAcc().Cred[0].Method)
			assert.Equal(t, acc.Cred[0].Value, result.GetAcc().Cred[0].Value)
		})
	}
}

// Test generated using Keploy
func TestPbCliSerialize_Login_843(t *testing.T) {
	login := &MsgClientLogin{
		Id:     "login-1",
		Scheme: "basic",
		Secret: []byte("secret"),
		Cred:   []MsgCredClient{{Method: "code", Response: "1234"}},
	}
	msg := &ClientComMessage{Login: login}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetLogin())
	assert.Equal(t, login.Id, result.GetLogin().Id)
	assert.Equal(t, login.Scheme, result.GetLogin().Scheme)
	assert.Equal(t, login.Secret, result.GetLogin().Secret)
	require.Len(t, result.GetLogin().Cred, 1)
	assert.Equal(t, login.Cred[0].Method, result.GetLogin().Cred[0].Method)
	assert.Equal(t, login.Cred[0].Response, result.GetLogin().Cred[0].Response)
}

// Test generated using Keploy
func TestPbCliSerialize_Sub_599(t *testing.T) {
	sub := &MsgClientSub{
		Id:    "sub-1",
		Topic: "grpTest",
		Set: &MsgSetQuery{
			Desc: &MsgSetDesc{Private: "private info"},
			Sub:  &MsgSetSub{User: "usrABC", Mode: "RW"},
		},
		Get: &MsgGetQuery{
			What: "data sub desc",
			Data: &MsgGetOpts{Limit: 10},
			Sub:  &MsgGetOpts{Limit: 5},
			Desc: &MsgGetOpts{},
		},
	}
	msg := &ClientComMessage{Sub: sub}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetSub())
	assert.Equal(t, sub.Id, result.GetSub().Id)
	assert.Equal(t, sub.Topic, result.GetSub().Topic)

	require.NotNil(t, result.GetSub().GetQuery)
	assert.Equal(t, sub.Get.What, result.GetSub().GetQuery.What)
	require.NotNil(t, result.GetSub().GetQuery.Data)
	assert.Equal(t, int32(sub.Get.Data.Limit), result.GetSub().GetQuery.Data.Limit)
	require.NotNil(t, result.GetSub().GetQuery.Sub)
	assert.Equal(t, int32(sub.Get.Sub.Limit), result.GetSub().GetQuery.Sub.Limit)
	require.NotNil(t, result.GetSub().GetQuery.Desc) // Should not be nil even if empty

	require.NotNil(t, result.GetSub().SetQuery)
	require.NotNil(t, result.GetSub().SetQuery.Desc)
	assert.Equal(t, `"`+sub.Set.Desc.Private.(string)+`"`, string(result.GetSub().SetQuery.Desc.Private))
	require.NotNil(t, result.GetSub().SetQuery.Sub)
	assert.Equal(t, sub.Set.Sub.User, result.GetSub().SetQuery.Sub.UserId)
	assert.Equal(t, sub.Set.Sub.Mode, result.GetSub().SetQuery.Sub.Mode)
}

// Test generated using Keploy
func TestPbCliSerialize_Leave_123(t *testing.T) {
	leave := &MsgClientLeave{
		Id:    "leave-1",
		Topic: "grpLeave",
		Unsub: true,
	}
	msg := &ClientComMessage{Leave: leave}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetLeave())
	assert.Equal(t, leave.Id, result.GetLeave().Id)
	assert.Equal(t, leave.Topic, result.GetLeave().Topic)
	assert.Equal(t, leave.Unsub, result.GetLeave().Unsub)
}

// Test generated using Keploy
func TestPbCliSerialize_Pub_456(t *testing.T) {
	pub := &MsgClientPub{
		Id:      "pub-1",
		Topic:   "grpPub",
		NoEcho:  true,
		Head:    map[string]any{"mime": "text/plain"},
		Content: "Hello world",
	}
	msg := &ClientComMessage{Pub: pub}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetPub())
	assert.Equal(t, pub.Id, result.GetPub().Id)
	assert.Equal(t, pub.Topic, result.GetPub().Topic)
	assert.Equal(t, pub.NoEcho, result.GetPub().NoEcho)
	require.NotNil(t, result.GetPub().Head)
	assert.Equal(t, []byte(`"text/plain"`), result.GetPub().Head["mime"])
	assert.Equal(t, []byte(`"Hello world"`), result.GetPub().Content)
}

// Test generated using Keploy
func TestPbCliSerialize_Get_789(t *testing.T) {
	get := &MsgClientGet{
		Id:    "get-1",
		Topic: "grpGet",
		MsgGetQuery: MsgGetQuery{
			What: "data",
			Data: &MsgGetOpts{SinceId: 5, Limit: 2},
		},
	}
	msg := &ClientComMessage{Get: get}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetGet())
	assert.Equal(t, get.Id, result.GetGet().Id)
	assert.Equal(t, get.Topic, result.GetGet().Topic)
	require.NotNil(t, result.GetGet().Query)
	assert.Equal(t, get.What, result.GetGet().Query.What)
	require.NotNil(t, result.GetGet().Query.Data)
	assert.Equal(t, int32(get.Data.SinceId), result.GetGet().Query.Data.SinceId)
	assert.Equal(t, int32(get.Data.Limit), result.GetGet().Query.Data.Limit)
	assert.Nil(t, result.GetGet().Query.Desc)
	assert.Nil(t, result.GetGet().Query.Sub)
}

// Test generated using Keploy
func TestPbCliSerialize_Set_987(t *testing.T) {
	set := &MsgClientSet{
		Id:    "set-1",
		Topic: "grpSet",
		MsgSetQuery: MsgSetQuery{
			Tags: []string{"settag1", "settag2"},
		},
	}
	msg := &ClientComMessage{Set: set}

	result := pbCliSerialize(msg)
	require.NotNil(t, result)
	require.NotNil(t, result.GetSet())
	assert.Equal(t, set.Id, result.GetSet().Id)
	assert.Equal(t, set.Topic, result.GetSet().Topic)
	require.NotNil(t, result.GetSet().Query)
	assert.Equal(t, set.Tags, result.GetSet().Query.Tags)
	assert.Nil(t, result.GetSet().Query.Desc)
	assert.Nil(t, result.GetSet().Query.Sub)
	assert.Nil(t, result.GetSet().Query.Cred)
}

// Test generated using Keploy
func TestPbCliSerialize_Del_WhatCases_654(t *testing.T) {
	testCases := []struct {
		what     string
		expected pbx.ClientDel_What
	}{
		{"msg", pbx.ClientDel_MSG},
		{"topic", pbx.ClientDel_TOPIC},
		{"sub", pbx.ClientDel_SUB},
		{"user", pbx.ClientDel_USER},
		{"cred", pbx.ClientDel_CRED},
		{"invalid", pbx.ClientDel_X0}, // default
	}

	for _, tc := range testCases {
		t.Run(tc.what, func(t *testing.T) {
			del := &MsgClientDel{
				Id:     "del-1",
				Topic:  "grpDel",
				What:   tc.what,
				DelSeq: []MsgDelRange{{LowId: 1}},
				User:   "usrDel",
				Cred:   &MsgCredClient{Method: "email"},
				Hard:   true,
			}
			msg := &ClientComMessage{Del: del}

			result := pbCliSerialize(msg)
			require.NotNil(t, result)
			require.NotNil(t, result.GetDel())
			assert.Equal(t, del.Id, result.GetDel().Id)
			assert.Equal(t, del.Topic, result.GetDel().Topic)
			assert.Equal(t, tc.expected, result.GetDel().What)
			require.Len(t, result.GetDel().DelSeq, 1)
			assert.Equal(t, int32(del.DelSeq[0].LowId), result.GetDel().DelSeq[0].Low)
			assert.Equal(t, del.User, result.GetDel().UserId)
			require.NotNil(t, result.GetDel().Cred)
			assert.Equal(t, del.Cred.Method, result.GetDel().Cred.Method)
			assert.Equal(t, del.Hard, result.GetDel().Hard)
		})
	}
}

// Test generated using Keploy
func TestPbCliSerialize_Empty_852(t *testing.T) {
	msg := &ClientComMessage{} // No message type set
	result := pbCliSerialize(msg)
	assert.Nil(t, result)
}

// Test generated using Keploy
func TestPbCliDeserialize_Hi_159(t *testing.T) {
	hi := &pbx.ClientHi{
		Id:         "hi-deser",
		UserAgent:  "TestAgent/2.0",
		Ver:        "0.17",
		DeviceId:   "dev456",
		Platform:   "srv",
		Lang:       "fr-FR",
		Background: true,
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Hi{Hi: hi}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Hi)
	assert.Equal(t, hi.Id, msg.Hi.Id)
	assert.Equal(t, hi.UserAgent, msg.Hi.UserAgent)
	assert.Equal(t, hi.Ver, msg.Hi.Version)
	assert.Equal(t, hi.DeviceId, msg.Hi.DeviceID)
	assert.Equal(t, hi.Platform, msg.Hi.Platform)
	assert.Equal(t, hi.Lang, msg.Hi.Lang)
	assert.Equal(t, hi.Background, msg.Hi.Background)
	assert.Nil(t, msg.Extra)
}

// Test generated using Keploy
func TestPbCliDeserialize_Acc_753(t *testing.T) {
	acc := &pbx.ClientAcc{
		Id:        "acc-deser",
		UserId:    "usrDeser",
		State:     "suspended",
		TmpScheme: "code",
		TmpSecret: []byte("12345"),
		AuthLevel: pbx.AuthLevel_AUTH,
		Scheme:    "basic",
		Secret:    []byte("newpass"),
		Login:     true,
		Tags:      []string{"tagA", "tagB"},
		Desc:      &pbx.SetDesc{Public: []byte(`{"name":"Test User"}`)},
		Cred:      []*pbx.ClientCred{{Method: "tel", Value: "+123"}},
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Acc{Acc: acc}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Acc)
	assert.Equal(t, acc.Id, msg.Acc.Id)
	assert.Equal(t, acc.UserId, msg.Acc.User)
	assert.Equal(t, acc.State, msg.Acc.State)
	assert.Equal(t, acc.TmpScheme, msg.Acc.TmpScheme)
	assert.Equal(t, acc.TmpSecret, msg.Acc.TmpSecret)
	assert.Equal(t, "AUTH", msg.Acc.AuthLevel)
	assert.Equal(t, acc.Scheme, msg.Acc.Scheme)
	assert.Equal(t, acc.Secret, msg.Acc.Secret)
	assert.Equal(t, acc.Login, msg.Acc.Login)
	assert.Equal(t, acc.Tags, msg.Acc.Tags)
	require.NotNil(t, msg.Acc.Desc)
	assert.Equal(t, map[string]interface{}{"name": "Test User"}, msg.Acc.Desc.Public)
	require.Len(t, msg.Acc.Cred, 1)
	assert.Equal(t, acc.Cred[0].Method, msg.Acc.Cred[0].Method)
	assert.Equal(t, acc.Cred[0].Value, msg.Acc.Cred[0].Value)
}

// Test generated using Keploy
func TestPbCliDeserialize_Login_951(t *testing.T) {
	login := &pbx.ClientLogin{
		Id:     "login-deser",
		Scheme: "token",
		Secret: []byte("abcdef"),
		Cred:   []*pbx.ClientCred{{Method: "email", Value: "a@b.c"}},
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Login{Login: login}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Login)
	assert.Equal(t, login.Id, msg.Login.Id)
	assert.Equal(t, login.Scheme, msg.Login.Scheme)
	assert.Equal(t, login.Secret, msg.Login.Secret)
	require.Len(t, msg.Login.Cred, 1)
	assert.Equal(t, login.Cred[0].Method, msg.Login.Cred[0].Method)
	assert.Equal(t, login.Cred[0].Value, msg.Login.Cred[0].Value)
}

// Test generated using Keploy
func TestPbCliDeserialize_Sub_NilQueries_357(t *testing.T) {
	sub := &pbx.ClientSub{
		Id:    "sub-deser",
		Topic: "grpDeser",
		// GetQuery and SetQuery are nil
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Sub{Sub: sub}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Sub)
	assert.Equal(t, sub.Id, msg.Sub.Id)
	assert.Equal(t, sub.Topic, msg.Sub.Topic)
	assert.Nil(t, msg.Sub.Get) // Should be nil when pb is nil
	assert.Nil(t, msg.Sub.Set) // Should be nil when pb is nil
}

// Test generated using Keploy
func TestPbCliDeserialize_Leave_864(t *testing.T) {
	leave := &pbx.ClientLeave{
		Id:    "leave-deser",
		Topic: "grpLeaveDeser",
		Unsub: false,
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Leave{Leave: leave}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Leave)
	assert.Equal(t, leave.Id, msg.Leave.Id)
	assert.Equal(t, leave.Topic, msg.Leave.Topic)
	assert.Equal(t, leave.Unsub, msg.Leave.Unsub)
}

// Test generated using Keploy
func TestPbCliDeserialize_Pub_246(t *testing.T) {
	pub := &pbx.ClientPub{
		Id:      "pub-deser",
		Topic:   "grpPubDeser",
		NoEcho:  false,
		Head:    map[string][]byte{"reply": []byte(`"123"`)},
		Content: []byte(`{"text":"response"}`),
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Pub{Pub: pub}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Pub)
	assert.Equal(t, pub.Id, msg.Pub.Id)
	assert.Equal(t, pub.Topic, msg.Pub.Topic)
	assert.Equal(t, pub.NoEcho, msg.Pub.NoEcho)
	require.NotNil(t, msg.Pub.Head)
	assert.Equal(t, "123", msg.Pub.Head["reply"])
	require.NotNil(t, msg.Pub.Content)
	assert.Equal(t, map[string]interface{}{"text": "response"}, msg.Pub.Content)
}

// Test generated using Keploy
func TestPbCliDeserialize_Get_WithQuery_135(t *testing.T) {
	get := &pbx.ClientGet{
		Id:    "get-deser",
		Topic: "grpGetDeser",
		Query: &pbx.GetQuery{
			What: "sub data",
			Sub:  &pbx.GetOpts{Limit: 1},
			Data: &pbx.GetOpts{BeforeId: 10},
		},
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Get{Get: get}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Get)
	assert.Equal(t, get.Id, msg.Get.Id)
	assert.Equal(t, get.Topic, msg.Get.Topic)
	assert.Equal(t, get.Query.What, msg.Get.What)
	require.NotNil(t, msg.Get.Sub)
	assert.Equal(t, int(get.Query.Sub.Limit), msg.Get.Sub.Limit)
	require.NotNil(t, msg.Get.Data)
	assert.Equal(t, int(get.Query.Data.BeforeId), msg.Get.Data.BeforeId)
	assert.Nil(t, msg.Get.Desc) // Was nil in source
}

// Test generated using Keploy
func TestPbCliDeserialize_Set_WithQuery_248(t *testing.T) {
	set := &pbx.ClientSet{
		Id:    "set-deser",
		Topic: "grpSetDeser",
		Query: &pbx.SetQuery{
			Desc: &pbx.SetDesc{Public: []byte(`"Public Desc"`)},
			Sub:  &pbx.SetSub{UserId: "usrSet", Mode: "RWP"},
			Tags: []string{"newtag"},
		},
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Set{Set: set}}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Set)
	assert.Equal(t, set.Id, msg.Set.Id)
	assert.Equal(t, set.Topic, msg.Set.Topic)

	require.NotNil(t, msg.Set.Desc)
	assert.Equal(t, "Public Desc", msg.Set.Desc.Public)
	require.NotNil(t, msg.Set.Sub)
	assert.Equal(t, set.Query.Sub.UserId, msg.Set.Sub.User)
	assert.Equal(t, set.Query.Sub.Mode, msg.Set.Sub.Mode)
	assert.Equal(t, set.Query.Tags, msg.Set.Tags)
	assert.Nil(t, msg.Set.Cred) // Was nil in source
}

// Test generated using Keploy
func TestPbCliDeserialize_Del_WhatCases_963(t *testing.T) {
	testCases := []struct {
		whatPbx  pbx.ClientDel_What
		expected string
	}{
		{pbx.ClientDel_MSG, "msg"},
		{pbx.ClientDel_TOPIC, "topic"},
		{pbx.ClientDel_SUB, "sub"},
		{pbx.ClientDel_USER, "user"},
		{pbx.ClientDel_CRED, "cred"},
		{pbx.ClientDel_What(99), ""}, // Default
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			del := &pbx.ClientDel{
				Id:     "del-deser",
				Topic:  "grpDelDeser",
				What:   tc.whatPbx,
				DelSeq: []*pbx.SeqRange{{Low: 10}},
				UserId: "usrDelDeser",
				Cred:   &pbx.ClientCred{Method: "sms"},
				Hard:   false,
			}
			pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Del{Del: del}}

			msg := pbCliDeserialize(pkt)
			require.NotNil(t, msg)
			require.NotNil(t, msg.Del)
			assert.Equal(t, del.Id, msg.Del.Id)
			assert.Equal(t, del.Topic, msg.Del.Topic)
			assert.Equal(t, tc.expected, msg.Del.What)
			require.Len(t, msg.Del.DelSeq, 1)
			assert.Equal(t, int(del.DelSeq[0].Low), msg.Del.DelSeq[0].LowId)
			assert.Equal(t, del.UserId, msg.Del.User)
			require.NotNil(t, msg.Del.Cred)
			assert.Equal(t, del.Cred.Method, msg.Del.Cred.Method)
			assert.Equal(t, del.Hard, msg.Del.Hard)
		})
	}
}

// Test generated using Keploy
func TestPbCliDeserialize_WithExtra_257(t *testing.T) {
	hi := &pbx.ClientHi{Id: "hi-extra-deser"}
	extra := &pbx.ClientExtra{
		Attachments: []string{"attA", "attB"},
		OnBehalfOf:  "usrOther",
		AuthLevel:   pbx.AuthLevel_ANON,
	}
	pkt := &pbx.ClientMsg{Message: &pbx.ClientMsg_Hi{Hi: hi}, Extra: extra}

	msg := pbCliDeserialize(pkt)
	require.NotNil(t, msg)
	require.NotNil(t, msg.Hi)
	assert.Equal(t, hi.Id, msg.Hi.Id)

	require.NotNil(t, msg.Extra)
	assert.Equal(t, extra.Attachments, msg.Extra.Attachments)
	assert.Equal(t, extra.OnBehalfOf, msg.Extra.AsUser)
	assert.Equal(t, "ANON", msg.Extra.AuthLevel)
}

// Test generated using Keploy
func TestBytesToInterface_ValidEmptyInvalid_702(t *testing.T) {
	// Valid
	validJSON := []byte(`{"key":"value"}`)
	var expected any = map[string]interface{}{"key": "value"}
	assert.Equal(t, expected, bytesToInterface(validJSON))

	// Empty bytes
	assert.Nil(t, bytesToInterface([]byte{}))

	// Null JSON
	nullJSON := []byte(`null`)
	assert.Nil(t, bytesToInterface(nullJSON)) // JSON null becomes Go nil

	// Invalid JSON - should log warning and return nil
	invalidJSON := []byte(`invalid json`)
	assert.Nil(t, bytesToInterface(invalidJSON))
	// Note: We cannot easily assert the log message here without more complex setup.
	// We rely on the function returning nil as documented.
}

// Test generated using Keploy
func TestPbGetQueryDeserialize_NilAndOptions_146(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbGetQueryDeserialize(nil))

	// Test with only 'what'
	inWhat := &pbx.GetQuery{What: "data"}
	outWhat := pbGetQueryDeserialize(inWhat)
	require.NotNil(t, outWhat)
	assert.Equal(t, "data", outWhat.What)
	assert.Nil(t, outWhat.Desc)
	assert.Nil(t, outWhat.Sub)
	assert.Nil(t, outWhat.Data)

	// Test with Desc options (ignore User/Topic as they are not deserialized)
	nowMillis := time.Now().UnixNano() / int64(time.Millisecond)
	inDesc := &pbx.GetQuery{
		What: "desc",
		Desc: &pbx.GetOpts{IfModifiedSince: nowMillis, Limit: 15},
	}
	outDesc := pbGetQueryDeserialize(inDesc)
	require.NotNil(t, outDesc)
	require.NotNil(t, outDesc.Desc)
	require.NotNil(t, outDesc.Desc.IfModifiedSince)
	assert.True(t, int64ToTime(inDesc.Desc.IfModifiedSince).Equal(*outDesc.Desc.IfModifiedSince))
	assert.Equal(t, int(inDesc.Desc.Limit), outDesc.Desc.Limit)
	assert.Nil(t, outDesc.Sub)
	assert.Nil(t, outDesc.Data)

	// Test with Sub options (ignore User/Topic)
	inSub := &pbx.GetQuery{
		What: "sub",
		Sub:  &pbx.GetOpts{IfModifiedSince: nowMillis - 10000, Limit: 7},
	}
	outSub := pbGetQueryDeserialize(inSub)
	require.NotNil(t, outSub)
	require.NotNil(t, outSub.Sub)
	require.NotNil(t, outSub.Sub.IfModifiedSince)
	assert.True(t, int64ToTime(inSub.Sub.IfModifiedSince).Equal(*outSub.Sub.IfModifiedSince))
	assert.Equal(t, int(inSub.Sub.Limit), outSub.Sub.Limit)
	assert.Nil(t, outSub.Desc)
	assert.Nil(t, outSub.Data)

	// Test with Data options
	inData := &pbx.GetQuery{
		What: "data",
		Data: &pbx.GetOpts{BeforeId: 200, SinceId: 150, Limit: 30},
	}
	outData := pbGetQueryDeserialize(inData)
	require.NotNil(t, outData)
	require.NotNil(t, outData.Data)
	assert.Equal(t, int(inData.Data.BeforeId), outData.Data.BeforeId)
	assert.Equal(t, int(inData.Data.SinceId), outData.Data.SinceId)
	assert.Equal(t, int(inData.Data.Limit), outData.Data.Limit)
	assert.Nil(t, outData.Desc)
	assert.Nil(t, outData.Sub)
}

// Test generated using Keploy
func TestPbSetDescDeserialize_NilAndEmpty_368(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbSetDescDeserialize(nil))

	// Test empty input (no fields set in pb)
	inEmpty := &pbx.SetDesc{}
	assert.Nil(t, pbSetDescDeserialize(inEmpty), "Should return nil when no fields are set")

	// Test with some fields
	inPartial := &pbx.SetDesc{
		Public:  []byte(`"some public data"`),
		Trusted: []byte(`{"valid":true}`),
	}
	outPartial := pbSetDescDeserialize(inPartial)
	require.NotNil(t, outPartial)
	assert.Equal(t, "some public data", outPartial.Public)
	assert.Equal(t, map[string]interface{}{"valid": true}, outPartial.Trusted)
	assert.Nil(t, outPartial.DefaultAcs)
	assert.Nil(t, outPartial.Private)

	// Test with all fields
	inFull := &pbx.SetDesc{
		DefaultAcs: &pbx.DefaultAcsMode{Auth: "RW", Anon: "R"},
		Public:     []byte(`123`),
		Trusted:    []byte(`["a","b"]`),
		Private:    []byte(`true`),
	}
	outFull := pbSetDescDeserialize(inFull)
	require.NotNil(t, outFull)
	require.NotNil(t, outFull.DefaultAcs)
	assert.Equal(t, inFull.DefaultAcs.Auth, outFull.DefaultAcs.Auth)
	assert.Equal(t, inFull.DefaultAcs.Anon, outFull.DefaultAcs.Anon)
	assert.Equal(t, float64(123), outFull.Public) // JSON number -> float64
	assert.Equal(t, []interface{}{"a", "b"}, outFull.Trusted)
	assert.Equal(t, true, outFull.Private)
}

// Test generated using Keploy
func TestPbInfoNoteWhatSerialize_AllCases_601(t *testing.T) {
	testCases := []struct {
		in  string
		out pbx.InfoNote
	}{
		{"kp", pbx.InfoNote_KP},
		{"read", pbx.InfoNote_READ},
		{"recv", pbx.InfoNote_RECV},
		{"call", pbx.InfoNote_CALL},
		{"unknown", pbx.InfoNote_X1}, // Default case
		{"", pbx.InfoNote_X1},        // Default case
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.out, pbInfoNoteWhatSerialize(tc.in))
		})
	}
}

// Test generated using Keploy
func TestPbInfoNoteWhatDeserialize_AllCases_712(t *testing.T) {
	testCases := []struct {
		in  pbx.InfoNote
		out string
	}{
		{pbx.InfoNote_KP, "kp"},
		{pbx.InfoNote_READ, "read"},
		{pbx.InfoNote_RECV, "recv"},
		{pbx.InfoNote_CALL, "call"},
		{pbx.InfoNote_X1, ""},  // Default case
		{pbx.InfoNote(99), ""}, // Unknown value
	}

	for _, tc := range testCases {
		t.Run(tc.out, func(t *testing.T) {
			assert.Equal(t, tc.out, pbInfoNoteWhatDeserialize(tc.in))
		})
	}
}

// Test generated using Keploy
func TestPbCallEventSerialize_AllCases_823(t *testing.T) {
	testCases := []struct {
		in  string
		out pbx.CallEvent
	}{
		{"accept", pbx.CallEvent_ACCEPT},
		{"answer", pbx.CallEvent_ANSWER},
		{"hang-up", pbx.CallEvent_HANG_UP},
		{"ice-candidate", pbx.CallEvent_ICE_CANDIDATE},
		{"invite", pbx.CallEvent_INVITE},
		{"offer", pbx.CallEvent_OFFER},
		{"ringing", pbx.CallEvent_RINGING},
		{"", pbx.CallEvent_X2},        // Empty maps to X2
		{"unknown", pbx.CallEvent_X2}, // Default case
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.out, pbCallEventSerialize(tc.in))
		})
	}
}

// Test generated using Keploy
func TestPbCallEventDeserialize_AllCases_934(t *testing.T) {
	testCases := []struct {
		in  pbx.CallEvent
		out string
	}{
		{pbx.CallEvent_ACCEPT, "accept"},
		{pbx.CallEvent_ANSWER, "answer"},
		{pbx.CallEvent_HANG_UP, "hang-up"},
		{pbx.CallEvent_ICE_CANDIDATE, "ice-candidate"},
		{pbx.CallEvent_INVITE, "invite"},
		{pbx.CallEvent_OFFER, "offer"},
		{pbx.CallEvent_RINGING, "ringing"},
		{pbx.CallEvent_X2, ""},  // Default case
		{pbx.CallEvent(99), ""}, // Unknown value
	}

	for _, tc := range testCases {
		t.Run(tc.out, func(t *testing.T) {
			assert.Equal(t, tc.out, pbCallEventDeserialize(tc.in))
		})
	}
}

// Test generated using Keploy
func TestPbDefaultAcsSerialize_NilNonNil_267(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbDefaultAcsSerialize(nil))

	// Test non-nil input
	in := &MsgDefaultAcsMode{Auth: "RW", Anon: "R"}
	out := pbDefaultAcsSerialize(in)
	require.NotNil(t, out)
	assert.Equal(t, in.Auth, out.Auth)
	assert.Equal(t, in.Anon, out.Anon)
}

// Test generated using Keploy
func TestPbDefaultAcsDeserialize_NilEmptyNonEmpty_378(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbDefaultAcsDeserialize(nil))

	// Test empty fields input (should return nil)
	inEmpty := &pbx.DefaultAcsMode{Auth: "", Anon: ""}
	assert.Nil(t, pbDefaultAcsDeserialize(inEmpty))

	// Test non-empty input
	inNonEmpty := &pbx.DefaultAcsMode{Auth: "J", Anon: "P"}
	outNonEmpty := pbDefaultAcsDeserialize(inNonEmpty)
	require.NotNil(t, outNonEmpty)
	assert.Equal(t, inNonEmpty.Auth, outNonEmpty.Auth)
	assert.Equal(t, inNonEmpty.Anon, outNonEmpty.Anon)

	// Test partial input
	inPartial := &pbx.DefaultAcsMode{Auth: "R"}
	outPartial := pbDefaultAcsDeserialize(inPartial)
	require.NotNil(t, outPartial)
	assert.Equal(t, inPartial.Auth, outPartial.Auth)
	assert.Empty(t, outPartial.Anon)
}

// Test generated using Keploy
func TestPbTopicDescDeserialize_NilAndLastSeen_590(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbTopicDescDeserialize(nil))

	nowMillis := time.Now().UnixNano() / int64(time.Millisecond)

	// Test without LastSeen
	inNoLastSeen := &pbx.TopicDesc{
		CreatedAt: nowMillis - 10000, UpdatedAt: nowMillis - 5000, TouchedAt: nowMillis,
		State: "suspended", Online: false, IsChan: true,
		Defacs: &pbx.DefaultAcsMode{Anon: "N"}, Acs: &pbx.AccessMode{Want: "W"},
		SeqId: 10, ReadId: 9, RecvId: 8, DelId: 7,
		Public: []byte(`{"p":1}`), Trusted: []byte(`{"t":2}`), Private: []byte(`{"r":3}`),
		// LastSeenTime is 0
	}
	outNoLastSeen := pbTopicDescDeserialize(inNoLastSeen)
	require.NotNil(t, outNoLastSeen)
	require.NotNil(t, outNoLastSeen.CreatedAt)
	assert.True(t, int64ToTime(inNoLastSeen.CreatedAt).Equal(*outNoLastSeen.CreatedAt))
	require.NotNil(t, outNoLastSeen.UpdatedAt)
	assert.True(t, int64ToTime(inNoLastSeen.UpdatedAt).Equal(*outNoLastSeen.UpdatedAt))
	require.NotNil(t, outNoLastSeen.TouchedAt)
	assert.True(t, int64ToTime(inNoLastSeen.TouchedAt).Equal(*outNoLastSeen.TouchedAt))
	assert.Equal(t, inNoLastSeen.State, outNoLastSeen.State)
	assert.Equal(t, inNoLastSeen.Online, outNoLastSeen.Online)
	assert.Equal(t, inNoLastSeen.IsChan, outNoLastSeen.IsChan)
	require.NotNil(t, outNoLastSeen.DefaultAcs)
	assert.Equal(t, inNoLastSeen.Defacs.Anon, outNoLastSeen.DefaultAcs.Anon)
	require.NotNil(t, outNoLastSeen.Acs)
	assert.Equal(t, inNoLastSeen.Acs.Want, outNoLastSeen.Acs.Want)
	assert.Equal(t, int(inNoLastSeen.SeqId), outNoLastSeen.SeqId)
	assert.Equal(t, int(inNoLastSeen.ReadId), outNoLastSeen.ReadSeqId)
	assert.Equal(t, int(inNoLastSeen.RecvId), outNoLastSeen.RecvSeqId)
	assert.Equal(t, int(inNoLastSeen.DelId), outNoLastSeen.DelId)
	assert.Equal(t, map[string]interface{}{"p": float64(1)}, outNoLastSeen.Public)
	assert.Equal(t, map[string]interface{}{"t": float64(2)}, outNoLastSeen.Trusted)
	assert.Equal(t, map[string]interface{}{"r": float64(3)}, outNoLastSeen.Private)
	assert.Nil(t, outNoLastSeen.LastSeen) // Should be nil

	// Test with LastSeen
	inWithLastSeen := &pbx.TopicDesc{
		LastSeenTime:      nowMillis - 1000,
		LastSeenUserAgent: "agent2",
	}
	outWithLastSeen := pbTopicDescDeserialize(inWithLastSeen)
	require.NotNil(t, outWithLastSeen)
	require.NotNil(t, outWithLastSeen.LastSeen)
	require.NotNil(t, outWithLastSeen.LastSeen.When)
	assert.True(t, int64ToTime(inWithLastSeen.LastSeenTime).Equal(*outWithLastSeen.LastSeen.When))
	assert.Equal(t, inWithLastSeen.LastSeenUserAgent, outWithLastSeen.LastSeen.UserAgent)
}

// Test generated using Keploy
func TestPbTopicSerializeToDesc_NilNonNil_601(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbTopicSerializeToDesc(nil))

	// Test non-nil input
	now := time.Now()
	topic := &Topic{
		created:    now.Add(-time.Hour),
		updated:    now.Add(-time.Minute),
		accessAuth: types.ModeRead | types.ModeWrite, // RW
		accessAnon: types.ModeRead,                   // R
		lastID:     50,
		delID:      5,
		public:     map[string]string{"name": "Public Topic"},
		trusted:    map[string]string{"verified": "yes"},
		// private is not serialized here
	}
	out := pbTopicSerializeToDesc(topic)
	require.NotNil(t, out)
	assert.Equal(t, timeToInt64(&topic.created), out.CreatedAt)
	assert.Equal(t, timeToInt64(&topic.updated), out.UpdatedAt)
	require.NotNil(t, out.Defacs)
	assert.Equal(t, topic.accessAuth.String(), out.Defacs.Auth)
	assert.Equal(t, topic.accessAnon.String(), out.Defacs.Anon)
	assert.Equal(t, int32(topic.lastID), out.SeqId)
	assert.Equal(t, int32(topic.delID), out.DelId)
	assert.Equal(t, `{"name":"Public Topic"}`, string(out.Public))
	assert.Equal(t, `{"verified":"yes"}`, string(out.Trusted))
	// Ensure fields not present in Topic struct are zero/nil
	assert.Zero(t, out.TouchedAt)
	assert.Nil(t, out.Acs)
	assert.Zero(t, out.ReadId)
	assert.Zero(t, out.RecvId)
	assert.Nil(t, out.Private)
}

// Test generated using Keploy
func TestPbTopicSubSliceSerialize_EmptyNonEmpty_712(t *testing.T) {
	// Test nil slice
	assert.Nil(t, pbTopicSubSliceSerialize(nil))

	// Test empty slice
	assert.Nil(t, pbTopicSubSliceSerialize([]MsgTopicSub{}))

	// Test non-empty slice
	subs := []MsgTopicSub{
		{User: "usr1", Topic: "topic1", Acs: MsgAccessMode{Want: "R"}},
		{User: "usr2", Topic: "topic1", Acs: MsgAccessMode{Given: "W"}},
	}
	out := pbTopicSubSliceSerialize(subs)
	require.NotNil(t, out)
	require.Len(t, out, 2)
	assert.Equal(t, subs[0].User, out[0].UserId)
	assert.Equal(t, subs[0].Topic, out[0].Topic)
	require.NotNil(t, out[0].Acs)
	assert.Equal(t, subs[0].Acs.Want, out[0].Acs.Want)
	assert.Equal(t, subs[1].User, out[1].UserId)
	assert.Equal(t, subs[1].Topic, out[1].Topic)
	require.NotNil(t, out[1].Acs)
	assert.Equal(t, subs[1].Acs.Given, out[1].Acs.Given)
}

// Test generated using Keploy
func TestPbTopicSubSliceDeserialize_EmptyNilNonNil_934(t *testing.T) {
	// Test nil slice
	assert.Nil(t, pbTopicSubSliceDeserialize(nil))

	// Test empty slice
	assert.Nil(t, pbTopicSubSliceDeserialize([]*pbx.TopicSub{}))

	nowMillis := time.Now().UnixNano() / int64(time.Millisecond)

	// Test with data, nil Acs, no LastSeen
	in1 := []*pbx.TopicSub{
		{
			UpdatedAt: nowMillis, DeletedAt: 0, Online: true,
			Acs:    nil, // Test nil Acs
			ReadId: 1, RecvId: 2, SeqId: 3, DelId: 4,
			Public: []byte(`"p"`), Trusted: []byte(`"t"`), Private: []byte(`"r"`),
			UserId: "usr1", Topic: "topic1", TouchedAt: nowMillis - 1000,
			// LastSeenTime is 0
		},
		{ // With Acs and LastSeen
			Acs:               &pbx.AccessMode{Want: "W", Given: "G"},
			LastSeenTime:      nowMillis - 2000,
			LastSeenUserAgent: "agent2",
			UserId:            "usr2", Topic: "topic1",
		},
	}
	out1 := pbTopicSubSliceDeserialize(in1)
	require.NotNil(t, out1)
	require.Len(t, out1, 2)

	// Check sub 1
	sub1 := out1[0]
	pbSub1 := in1[0]
	require.NotNil(t, sub1.UpdatedAt)
	assert.True(t, int64ToTime(pbSub1.UpdatedAt).Equal(*sub1.UpdatedAt))
	assert.Nil(t, sub1.DeletedAt)
	assert.Equal(t, pbSub1.Online, sub1.Online)
	assert.Empty(t, sub1.Acs.Want) // Acs should be zero value because input was nil
	assert.Empty(t, sub1.Acs.Given)
	assert.Equal(t, int(pbSub1.ReadId), sub1.ReadSeqId)
	assert.Equal(t, int(pbSub1.RecvId), sub1.RecvSeqId)
	assert.Equal(t, int(pbSub1.SeqId), sub1.SeqId)
	assert.Equal(t, int(pbSub1.DelId), sub1.DelId)
	assert.Equal(t, "p", sub1.Public)
	assert.Equal(t, "t", sub1.Trusted)
	assert.Equal(t, "r", sub1.Private)
	assert.Equal(t, pbSub1.UserId, sub1.User)
	assert.Equal(t, pbSub1.Topic, sub1.Topic)
	require.NotNil(t, sub1.TouchedAt)
	assert.True(t, int64ToTime(pbSub1.TouchedAt).Equal(*sub1.TouchedAt))
	assert.Nil(t, sub1.LastSeen)

	// Check sub 2
	sub2 := out1[1]
	pbSub2 := in1[1]
	require.NotNil(t, sub2.Acs)
	assert.Equal(t, pbSub2.Acs.Want, sub2.Acs.Want)
	assert.Equal(t, pbSub2.Acs.Given, sub2.Acs.Given)
	require.NotNil(t, sub2.LastSeen)
	require.NotNil(t, sub2.LastSeen.When)
	assert.True(t, int64ToTime(pbSub2.LastSeenTime).Equal(*sub2.LastSeen.When))
	assert.Equal(t, pbSub2.LastSeenUserAgent, sub2.LastSeen.UserAgent)
	assert.Equal(t, pbSub2.UserId, sub2.User)
	assert.Equal(t, pbSub2.Topic, sub2.Topic)
}

// Test generated using Keploy
func TestPbSubSliceDeserialize_ToTypesSubscription_045(t *testing.T) {
	// Test nil slice
	assert.Nil(t, pbSubSliceDeserialize(nil))

	// Test empty slice
	assert.Nil(t, pbSubSliceDeserialize([]*pbx.TopicSub{}))

	nowMillis := time.Now().UnixNano() / int64(time.Millisecond)
	nowTime := int64ToTime(nowMillis)
	require.NotNil(t, nowTime)

	// Test with data, nil Acs, no LastSeen
	in1 := []*pbx.TopicSub{
		{
			UpdatedAt: nowMillis, DeletedAt: 0,
			Acs:    nil, // Test nil Acs
			UserId: "usr1", Topic: "topic1", DelId: 4,
			Public: []byte(`"p"`), Trusted: []byte(`"t"`), Private: []byte(`"r"`),
			// LastSeenTime is 0
		},
		{ // With Acs and LastSeen
			UpdatedAt:         nowMillis,
			Acs:               &pbx.AccessMode{Want: "W", Given: "G"},
			LastSeenTime:      nowMillis - 2000,
			LastSeenUserAgent: "agent2",
			UserId:            "usr2", Topic: "topic1",
		},
	}
	out1 := pbSubSliceDeserialize(in1)
	require.NotNil(t, out1)
	require.Len(t, out1, 2)

	// Check sub 1
	sub1 := out1[0]
	pbSub1 := in1[0]
	assert.True(t, nowTime.Equal(sub1.UpdatedAt))
	assert.Nil(t, sub1.DeletedAt)
	assert.Equal(t, types.AccessMode(0), sub1.ModeGiven) // Should be zero due to nil Acs
	assert.Equal(t, types.AccessMode(0), sub1.ModeWant)
	assert.Equal(t, pbSub1.UserId, sub1.User)
	assert.Equal(t, pbSub1.Topic, sub1.Topic)
	assert.Equal(t, int(pbSub1.DelId), sub1.DelId)
	assert.Equal(t, "p", sub1.GetPublic())
	assert.Equal(t, "t", sub1.GetTrusted())
	assert.Equal(t, "r", sub1.Private)
	assert.Nil(t, sub1.GetLastSeen()) // No LastSeen time provided

	// Check sub 2
	sub2 := out1[1]
	pbSub2 := in1[1]
	assert.True(t, nowTime.Equal(sub2.UpdatedAt))
	var expectedGiven, expectedWant types.AccessMode
	expectedGiven.UnmarshalText([]byte(pbSub2.Acs.Given))
	expectedWant.UnmarshalText([]byte(pbSub2.Acs.Want))
	assert.Equal(t, expectedGiven, sub2.ModeGiven)
	assert.Equal(t, expectedWant, sub2.ModeWant)
	require.NotNil(t, sub2.GetLastSeen())
	expectedLastSeen := int64ToTime(pbSub2.LastSeenTime)
	assert.True(t, expectedLastSeen.Equal(*sub2.GetLastSeen()))
	assert.Equal(t, pbSub2.LastSeenUserAgent, sub2.GetUserAgent())
	assert.Equal(t, pbSub2.UserId, sub2.User)
	assert.Equal(t, pbSub2.Topic, sub2.Topic)
}

// Test generated using Keploy
func TestPbDelValuesSerialize_NilNonNil_378(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbDelValuesSerialize(nil))

	// Test non-nil input
	in := &MsgDelValues{
		DelId:  12,
		DelSeq: []MsgDelRange{{LowId: 1, HiId: 2}, {LowId: 5}},
	}
	out := pbDelValuesSerialize(in)
	require.NotNil(t, out)
	assert.Equal(t, int32(in.DelId), out.DelId)
	require.Len(t, out.DelSeq, 2)
	assert.Equal(t, int32(in.DelSeq[0].LowId), out.DelSeq[0].Low)
	assert.Equal(t, int32(in.DelSeq[0].HiId), out.DelSeq[0].Hi)
	assert.Equal(t, int32(in.DelSeq[1].LowId), out.DelSeq[1].Low)
	assert.Equal(t, int32(0), out.DelSeq[1].Hi)
}

// Test generated using Keploy
func TestPbDelValuesDeserialize_NilNonNil_489(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbDelValuesDeserialize(nil))

	// Test non-nil input
	in := &pbx.DelValues{
		DelId:  21,
		DelSeq: []*pbx.SeqRange{{Low: 3, Hi: 4}, {Low: 8}},
	}
	out := pbDelValuesDeserialize(in)
	require.NotNil(t, out)
	assert.Equal(t, int(in.DelId), out.DelId)
	require.Len(t, out.DelSeq, 2)
	assert.Equal(t, int(in.DelSeq[0].Low), out.DelSeq[0].LowId)
	assert.Equal(t, int(in.DelSeq[0].Hi), out.DelSeq[0].HiId)
	assert.Equal(t, int(in.DelSeq[1].Low), out.DelSeq[1].LowId)
	assert.Equal(t, int(0), out.DelSeq[1].HiId)
}

// Test generated using Keploy
func TestPbClientCredsSerialize_NilEmptyNonEmpty_611(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbClientCredsSerialize(nil))

	// Test empty slice
	assert.NotNil(t, pbClientCredsSerialize([]MsgCredClient{}))
	assert.Len(t, pbClientCredsSerialize([]MsgCredClient{}), 0)

	// Test non-empty slice
	in := []MsgCredClient{
		{Method: "m1", Value: "v1"},
		{Method: "m2", Response: "r2"},
	}
	out := pbClientCredsSerialize(in)
	require.NotNil(t, out)
	require.Len(t, out, 2)
	assert.Equal(t, "m1", out[0].Method)
	assert.Equal(t, "v1", out[0].Value)
	assert.Equal(t, "m2", out[1].Method)
	assert.Equal(t, "r2", out[1].Response)
}

// Test generated using Keploy
func TestPbClientCredDeserialize_NilNonNil_722(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbClientCredDeserialize(nil))

	// Test non-nil input
	in := &pbx.ClientCred{
		Method:   "tel",
		Value:    "+123",
		Response: "code123",
		Params:   map[string][]byte{"country": []byte(`"US"`)},
	}
	out := pbClientCredDeserialize(in)
	require.NotNil(t, out)
	assert.Equal(t, in.Method, out.Method)
	assert.Equal(t, in.Value, out.Value)
	assert.Equal(t, in.Response, out.Response)
	require.NotNil(t, out.Params)
	assert.Equal(t, map[string]any{"country": "US"}, out.Params)
}

// Test generated using Keploy
func TestPbClientCredsDeserialize_NilEmptyNonEmpty_833(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbClientCredsDeserialize(nil))

	// Test empty slice
	assert.NotNil(t, pbClientCredsDeserialize([]*pbx.ClientCred{}))
	assert.Len(t, pbClientCredsDeserialize([]*pbx.ClientCred{}), 0)

	// Test non-empty slice
	in := []*pbx.ClientCred{
		{Method: "cm1", Value: "cv1"},
		{Method: "cm2", Response: "cr2"},
	}
	out := pbClientCredsDeserialize(in)
	require.NotNil(t, out)
	require.Len(t, out, 2)
	assert.Equal(t, "cm1", out[0].Method)
	assert.Equal(t, "cv1", out[0].Value)
	assert.Equal(t, "cm2", out[1].Method)
	assert.Equal(t, "cr2", out[1].Response)
}

// Test generated using Keploy
func TestPbServerCredsSerialize_NilEmptyNonEmpty_944(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbServerCredsSerialize(nil))

	// Test empty slice
	assert.NotNil(t, pbServerCredsSerialize([]*MsgCredServer{}))
	assert.Len(t, pbServerCredsSerialize([]*MsgCredServer{}), 0)

	// Test non-empty slice
	in := []*MsgCredServer{
		{Method: "sm1", Value: "sv1"}, // Done field does not exist in MsgCredServer
		{Method: "sm2", Value: "sv2"},
	}
	out := pbServerCredsSerialize(in)
	require.NotNil(t, out)
	require.Len(t, out, 2)
	assert.Equal(t, "sm1", out[0].Method)
	assert.Equal(t, "sv1", out[0].Value)
	assert.False(t, out[0].Done) // Done field is not serialized from MsgCredServer
	assert.Equal(t, "sm2", out[1].Method)
	assert.Equal(t, "sv2", out[1].Value)
	assert.False(t, out[1].Done)
}

// Test generated using Keploy
func TestPbServerCredsDeserialize_NilEmptyNonEmpty_055(t *testing.T) {
	// Test nil input
	assert.Nil(t, pbServerCredsDeserialize(nil))

	// Test empty slice
	assert.NotNil(t, pbServerCredsDeserialize([]*pbx.ServerCred{}))
	assert.Len(t, pbServerCredsDeserialize([]*pbx.ServerCred{}), 0)

	// Test non-empty slice
	in := []*pbx.ServerCred{
		{Method: "sm1d", Value: "sv1d", Done: true},
		{Method: "sm2d", Value: "sv2d", Done: false},
	}
	out := pbServerCredsDeserialize(in)
	require.NotNil(t, out)
	require.Len(t, out, 2)
	assert.Equal(t, "sm1d", out[0].Method)
	assert.Equal(t, "sv1d", out[0].Value)
	assert.True(t, out[0].Done)
	assert.Equal(t, "sm2d", out[1].Method)
	assert.Equal(t, "sv2d", out[1].Value)
	assert.False(t, out[1].Done)
}
