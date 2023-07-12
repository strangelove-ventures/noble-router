package keeper

import (
	"encoding/binary"
)

// TODO copy pasted from github.com/strangelove-ventures/noble-cctp, change to reference that

type BurnMessage struct {
	Version       uint32
	BurnToken     []byte
	MintRecipient []byte
	Amount        uint64
	MessageSender []byte
}

type Message struct {
	Version           uint32
	SourceDomainBytes []byte
	SourceDomain      uint32
	DestinationDomain uint32
	NonceBytes        []byte
	Nonce             uint64
	Sender            []byte
	Recipient         []byte
	DestinationCaller []byte
	MessageBody       []byte
}

const (
	// Indices of each field in message
	VersionIndex           = 0
	SourceDomainIndex      = 4
	DestinationDomainIndex = 8
	NonceIndex             = 12
	SenderIndex            = 20
	RecipientIndex         = 52
	DestinationCallerIndex = 84
	MessageBodyIndex       = 116

	// Indices of each field in BurnMessage
	BurnMsgVersionIndex = 0
	VersionLen          = 4
	BurnTokenIndex      = 4
	BurnTokenLen        = 32
	MintRecipientIndex  = 36
	MintRecipientLen    = 32
	AmountIndex         = 68
	AmountLen           = 32
	MsgSenderIndex      = 100
	MsgSenderLen        = 32
	// 4 byte version + 32 bytes burnToken + 32 bytes mintRecipient + 32 bytes amount + 32 bytes messageSender
	BurnMessageLen = 132

	NobleMessageVersion = 0
	MessageBodyVersion  = 0
	NobleDomainId       = 4
	Bytes32Len          = 32
)

func decodeMessage(msg []byte) Message {
	message := Message{
		Version:           binary.BigEndian.Uint32(msg[VersionIndex:SourceDomainIndex]),
		SourceDomainBytes: msg[SourceDomainIndex:DestinationDomainIndex],
		SourceDomain:      binary.BigEndian.Uint32(msg[SourceDomainIndex:DestinationDomainIndex]),
		DestinationDomain: binary.BigEndian.Uint32(msg[DestinationDomainIndex:NonceIndex]),
		NonceBytes:        msg[NonceIndex:SenderIndex],
		Nonce:             binary.BigEndian.Uint64(msg[NonceIndex:SenderIndex]),
		Sender:            msg[SenderIndex:RecipientIndex],
		Recipient:         msg[RecipientIndex:DestinationCallerIndex],
		DestinationCaller: msg[DestinationCallerIndex:MessageBodyIndex],
		MessageBody:       msg[MessageBodyIndex:],
	}

	return message
}

func decodeBurnMessage(msg []byte) BurnMessage {
	message := BurnMessage{
		Version:       binary.BigEndian.Uint32(msg[BurnMsgVersionIndex:BurnTokenIndex]),
		BurnToken:     msg[BurnTokenIndex:MintRecipientIndex],
		MintRecipient: msg[MintRecipientIndex:AmountIndex],
		Amount:        binary.BigEndian.Uint64(msg[AmountIndex:MsgSenderIndex]),
		MessageSender: msg[MsgSenderIndex:BurnMessageLen],
	}

	return message
}

func parseBurnMessageIntoBytes(msg BurnMessage) []byte {
	result := make([]byte, BurnMessageLen)

	versionBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(versionBytes, msg.Version)

	amountBytes := make([]byte, Bytes32Len)
	binary.LittleEndian.PutUint64(amountBytes, msg.Amount)

	copyBytes(BurnMsgVersionIndex, BurnTokenIndex, versionBytes, &result)
	copyBytes(BurnTokenIndex, MintRecipientIndex, msg.BurnToken, &result)
	copyBytes(MintRecipientIndex, AmountIndex, msg.MintRecipient, &result)
	copyBytes(AmountIndex, MsgSenderIndex, amountBytes, &result)

	return result
}

func copyBytes(start int, end int, copyFrom []byte, copyInto *[]byte) {
	for i := start; i < end; i++ {
		(*copyInto)[i] = copyFrom[i]
	}
}
