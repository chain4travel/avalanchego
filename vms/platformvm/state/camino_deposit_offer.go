// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks"
)

const interestRateDenominator uint64 = 1_000_000

var bigInterestRateDenominator = (&big.Int{}).SetUint64(interestRateDenominator)

type DepositOffer struct {
	id ids.ID

	GradualUnlock         bool   `serialize:"true"`
	InterestRateNominator uint64 `serialize:"true"`
	Start                 uint64 `serialize:"true"`
	End                   uint64 `serialize:"true"`
	MinAmount             uint64 `serialize:"true"`
	Duration              uint64 `serialize:"true"`
}

func (o *DepositOffer) ID() ids.ID {
	return o.id
}

func (o *DepositOffer) setID() error {
	bytes, err := blocks.GenesisCodec.Marshal(blocks.Version, o)
	if err != nil {
		return err
	}
	o.id = hashing.ComputeHash256Array(bytes)
	return nil
}

func (o DepositOffer) StartTime() time.Time {
	return time.Unix(int64(o.Start), 0)
}

func (o DepositOffer) EndTime() time.Time {
	return time.Unix(int64(o.End), 0)
}

func (o DepositOffer) InterestRateFloat64() float64 {
	return float64(o.InterestRateNominator) / float64(interestRateDenominator)
}

func (cs *caminoState) AddDepositOffer(offer *DepositOffer) {
	offer.setID()
	cs.modifiedDepositOffers[offer.id] = offer
}

func (cs *caminoState) GetDepositOffer(offerID ids.ID) (*DepositOffer, error) {
	// Try to get from modified state
	offer, ok := cs.modifiedDepositOffers[offerID]
	// Try to get it from state
	if !ok {
		if offer, ok = cs.depositOffers[offerID]; !ok {
			return nil, database.ErrNotFound
		}
	}
	return offer, nil
}

func (cs *caminoState) GetAllDepositOffers() ([]*DepositOffer, error) {
	offers := make([]*DepositOffer, len(cs.modifiedDepositOffers))

	for _, offer := range cs.modifiedDepositOffers {
		offers = append(offers, offer)
	}

	for _, offer := range cs.depositOffers {
		offers = append(offers, offer)
	}

	return offers, nil
}

func (cs *caminoState) loadDepositOffers() error {
	depositOffersIt := cs.depositOffersList.NewIterator()
	defer depositOffersIt.Release()
	for depositOffersIt.Next() {
		depositOfferIDBytes := depositOffersIt.Key()
		depositOfferID, err := ids.ToID(depositOfferIDBytes)
		if err != nil {
			return err
		}

		depositOfferBytes := depositOffersIt.Value()
		depositOffer := &DepositOffer{
			id: depositOfferID,
		}
		if _, err := blocks.GenesisCodec.Unmarshal(depositOfferBytes, depositOffer); err != nil {
			return err
		}

		cs.depositOffers[depositOfferID] = depositOffer
	}

	return nil
}

func (cs *caminoState) writeDepositOffers() error {
	for offerID, offer := range cs.modifiedDepositOffers {
		delete(cs.modifiedDepositOffers, offerID)

		offerBytes, err := blocks.GenesisCodec.Marshal(blocks.Version, offer)
		if err != nil {
			return fmt.Errorf("failed to serialize deposit offer: %w", err)
		}

		if err := cs.depositOffersList.Put(offerID[:], offerBytes); err != nil {
			return err
		}
	}
	return nil
}
