package types

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
)

type ItemStatus string
type ItemPriority string

const (
	// ItemStatus
	ItemStatus_InProgress ItemStatus = "IN_PROGRESS"
	ItemStatus_Completed  ItemStatus = "COMPLETED"

	// ItemPriority
	ItemPriority_High   ItemPriority = "HIGH"
	ItemPriority_Medium ItemPriority = "MEDIUM"
	ItemPriority_Low    ItemPriority = "LOW"
	ItemPriority_None   ItemPriority = "NONE"
)

// ItemStatus
func (ItemStatus) Values() []string {
	var variants []string
	for _, s := range []ItemStatus{ItemStatus_InProgress, ItemStatus_Completed} {
		variants = append(variants, string(s))
	}
	return variants
}

func (i ItemStatus) Value() (driver.Value, error) {
	return string(i), nil
}

func (i *ItemStatus) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("ItemStatus: expected a value")
	}
	switch src := src.(type) {
	case string:
		*i = ItemStatus(src)
	case ItemStatus:
		*i = src
	default:
		return fmt.Errorf("ItemStatus: unexpected type, %T", src)
	}
	return nil
}

func (i *ItemStatus) UnmarshalGQL(v interface{}) error {
	return i.Scan(v)
}

func (i ItemStatus) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(i)))
}

// ItemPriority
func (ItemPriority) Values() []string {
	var variants []string
	for _, s := range []ItemPriority{ItemPriority_High, ItemPriority_Medium, ItemPriority_Low, ItemPriority_None} {
		variants = append(variants, string(s))
	}
	return variants
}

func (i ItemPriority) Value() (driver.Value, error) {
	return string(i), nil
}

func (i *ItemPriority) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("ItemPriority: expected a value")
	}
	switch src := src.(type) {
	case string:
		*i = ItemPriority(src)
	case ItemPriority:
		*i = src
	default:
		return fmt.Errorf("ItemPriority: unexpected type, %T", src)
	}
	return nil
}

func (i *ItemPriority) UnmarshalGQL(v interface{}) error {
	return i.Scan(v)
}

func (i ItemPriority) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(i)))
}
