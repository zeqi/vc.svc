package schemas

import (
	"time"
)

type Tracking struct {
	OptTime  time.Time `json:"OptTime,omitempty" bson:"optTime,omitempty"`
	Operator string    `json:"Operator" bson:"operator"`
	Name     string    `json:"Name,omitempty" bson:"name,omitempty"`
	Action   string    `json:"Action,omitempty" bson:"action,omitempty"`
	Reason   string    `json:"Reason,omitempty" bson:"reason,omitempty"`
}
