package circleci

import "time"

// String returns a pointer to the given string.
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the given bool.
func Bool(v bool) *bool {
	return &v
}

// Time returns a pointer to the given time.Time
func Time(v time.Time) *time.Time {
	return &v
}

// ReportingWindow returns a pointer to the given reportingWindow.
func ReportingWindow(v reportingWindow) *reportingWindow {
	return &v
}

// OwnerType returs a pointer to the given ownerType.
func OwnerType(v ownerType) *ownerType {
	return &v
}

// CheckoutKeyType return a pointer to the given checkoutKeyType
func CheckoutKeyType(v checkoutKeyType) *checkoutKeyType {
	return &v
}
