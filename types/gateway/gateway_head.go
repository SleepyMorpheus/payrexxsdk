package gateway

import (
	"encoding/json"
	"fmt"
	"time"
)

type Status string

const (
	StatusWaiting    Status = "waiting"
	StatusConfirmed  Status = "confirmed"
	StatusAuthorized Status = "authorized"
	StatusReserved   Status = "reserved"
)

// GatewayHead represents the data which gets generated by creating
// a gateway with Payrexx
type GatewayHead struct {
	ID        int32     `json:"id"`
	Status    Status    `json:"status"`
	Hash      string    `json:"hash"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

func (g GatewayHead) String() string {
	outHead, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("GatewayHead: %s", string(outHead))
}

func (g *GatewayHead) MarshalJSON() ([]byte, error) {
	type Alias GatewayHead
	return json.Marshal(&struct {
		CreatedAt int64 `json:"createdAt"`
		*Alias
	}{
		CreatedAt: g.CreatedAt.Unix(),
		Alias:     (*Alias)(g),
	})
}

func (g *GatewayHead) UnmarshalJSON(bytes []byte) error {
	type Alias GatewayHead

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
