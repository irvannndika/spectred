package protowire

import (
	"github.com/pkg/errors"
	"github.com/spectre-project/spectred/app/appmessage"
	"github.com/spectre-project/spectred/domain/consensus/model/externalapi"
)

func (x *SpectredMessage_RequestIBDChainBlockLocator) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SpectredMessage_RequestIBDChainBlockLocator is nil")
	}
	return x.RequestIBDChainBlockLocator.toAppMessage()
}

func (x *RequestIBDChainBlockLocatorMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestIBDChainBlockLocatorMessage is nil")
	}
	var err error
	var highHash, lowHash *externalapi.DomainHash
	if x.HighHash != nil {
		highHash, err = x.HighHash.toDomain()
		if err != nil {
			return nil, err
		}
	}
	if x.LowHash != nil {
		lowHash, err = x.LowHash.toDomain()
		if err != nil {
			return nil, err
		}
	}
	return &appmessage.MsgRequestIBDChainBlockLocator{
		HighHash: highHash,
		LowHash:  lowHash,
	}, nil

}

func (x *SpectredMessage_RequestIBDChainBlockLocator) fromAppMessage(msgGetBlockLocator *appmessage.MsgRequestIBDChainBlockLocator) error {
	var highHash, lowHash *Hash
	if msgGetBlockLocator.HighHash != nil {
		highHash = domainHashToProto(msgGetBlockLocator.HighHash)
	}
	if msgGetBlockLocator.LowHash != nil {
		lowHash = domainHashToProto(msgGetBlockLocator.LowHash)
	}
	x.RequestIBDChainBlockLocator = &RequestIBDChainBlockLocatorMessage{
		HighHash: highHash,
		LowHash:  lowHash,
	}

	return nil
}