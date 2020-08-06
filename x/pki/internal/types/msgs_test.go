//nolint:testpackage
package types

import (
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
	MsgProposeAddX509RootCert
*/

func TestNewMsgProposeAddX509RootCert(t *testing.T) {
	var msg = NewMsgProposeAddX509RootCert(testconstants.RootCertPem, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgProposeAddX509RootCert
	}{
		{true, NewMsgProposeAddX509RootCert(
			testconstants.RootCertPem, testconstants.Signer)},
		{false, NewMsgProposeAddX509RootCert(
			"", testconstants.Signer)},
		{false, NewMsgProposeAddX509RootCert(
			testconstants.RootCertPem, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgProposeAddX509RootCertGetSignBytes(t *testing.T) {
	var msg = NewMsgProposeAddX509RootCert(testconstants.StubCert, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/ProposeAddX509RootCert","value":{"cert":` +
		`"pem certificate","signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgApproveAddX509RootCert(
*/

func TestNewMsgApproveAddX509RootCert(t *testing.T) {
	var msg = NewMsgApproveAddX509RootCert(testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_add_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveAddX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgApproveAddX509RootCert
	}{
		{true, NewMsgApproveAddX509RootCert(
			testconstants.LeafSubject, testconstants.LeafSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			"", testconstants.LeafSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			testconstants.LeafSubject, "", testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			testconstants.LeafSubject, testconstants.LeafSubjectKeyID, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgApproveAddX509RootCertGetSignBytes(t *testing.T) {
	var msg = NewMsgApproveAddX509RootCert(testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID, testconstants.Signer)

	expected := `{"type":"pki/ApproveAddX509RootCert","value":{"signer":` +
		`"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=dsr-corporation.com","subject_key_id":` +
		`"8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgAddX509Cert
*/

func TestNewMsgAddX509Cert(t *testing.T) {
	var msg = NewMsgAddX509Cert(testconstants.LeafCertPem, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_x509_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgAddX509Cert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddX509Cert
	}{
		{true, NewMsgAddX509Cert(
			testconstants.LeafCertPem, testconstants.Signer)},
		{false, NewMsgAddX509Cert(
			"", testconstants.Signer)},
		{false, NewMsgAddX509Cert(
			testconstants.LeafCertPem, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgAddX509CertGetSignBytes(t *testing.T) {
	var msg = NewMsgAddX509Cert(testconstants.StubCert, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/AddX509Cert","value":{"cert":"pem certificate",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgRevokeX509Cert
*/

func TestNewMsgRevokeX509Cert(t *testing.T) {
	var msg = NewMsgRevokeX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID, testconstants.Signer)

	require.Equal(t, RouterKey, msg.Route())
	require.Equal(t, "revoke_x509_cert", msg.Type())
	require.Equal(t, []sdk.AccAddress{testconstants.Signer}, msg.GetSigners())
}

func TestValidateMsgRevokeX509Cert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgRevokeX509Cert
	}{
		{true, NewMsgRevokeX509Cert(testconstants.LeafIssuer, testconstants.LeafSerialNumber, testconstants.Signer)},
		{false, NewMsgRevokeX509Cert("", testconstants.LeafSerialNumber, testconstants.Signer)},
		{false, NewMsgRevokeX509Cert(testconstants.LeafIssuer, "", testconstants.Signer)},
		{false, NewMsgRevokeX509Cert(testconstants.LeafIssuer, testconstants.LeafSerialNumber, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgRevokeX509CertGetSignBytes(t *testing.T) {
	var msg = NewMsgRevokeX509Cert(testconstants.LeafIssuer, testconstants.LeafSerialNumber, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/RevokeX509Cert","value":{` +
		`"issuer":"CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US",` +
		`"serial_number":"393904870890265262371394210372104514174397",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}
