package paylink

import (
	"encoding/json"
	"fmt"
	"time"
)

type PaylinkStatus string

const (
	PaylinkStatusWaiting    PaylinkStatus = "waiting"
	PaylinkStatusConfirmed  PaylinkStatus = "confirmed"
	PaylinkStatusAuthorized PaylinkStatus = "authorized"
	PaylinkStatusReserved   PaylinkStatus = "reserved"
)

// PaylinkHead represents the data which gets generated by creating
// a gateway with Payrexx
type PaylinkHead struct {
	ID        int32         `json:"id"`
	Hash      string        `json:"hash"`
	Link      string        `json:"link"`
	Invoices  []string      `json:"invoices"`
	Api       bool          `json:"api"`
	CreatedAt time.Time     `json:"created_at"`
	Status    PaylinkStatus `json:"status"`
}

func (g PaylinkHead) String() string {
	outHead, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("PaylinkHead: %s", string(outHead))
}

func (g *PaylinkHead) MarshalJSON() ([]byte, error) {
	type Alias PaylinkHead
	return json.Marshal(&struct {
		CreatedAt int64 `json:"createdAt"`
		*Alias
	}{
		CreatedAt: g.CreatedAt.Unix(),
		Alias:     (*Alias)(g),
	})
}

func (g *PaylinkHead) UnmarshalJSON(bytes []byte) error {
	type Alias PaylinkHead

	aux := &struct {
		CreatedAt int64 `json:"createdAt"`
		*Alias
	}{
		Alias: (*Alias)(g),
	}

	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	g.CreatedAt = time.Unix(aux.CreatedAt, 0)
	return nil
}
