// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type City struct {
	ID      *string      `json:"ID"`
	Name    *string      `json:"name"`
	Country *string      `json:"country"`
	Coord   *Coordinates `json:"coord"`
	Weather *Weather     `json:"weather"`
}

type Clouds struct {
	All        *int `json:"all"`
	Visibility *int `json:"visibility"`
	Humidity   *int `json:"humidity"`
}

type ConfigInput struct {
	Units *Unit     `json:"units"`
	Lang  *Language `json:"lang"`
}

type Coordinates struct {
	Lon *float64 `json:"lon"`
	Lat *float64 `json:"lat"`
}

type Summary struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Icon        *string `json:"icon"`
}

type Temperature struct {
	Actual    *float64 `json:"actual"`
	FeelsLike *float64 `json:"feelsLike"`
	Min       *float64 `json:"min"`
	Max       *float64 `json:"max"`
}

type User struct {
	ID          string       `json:"ID"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Description string       `json:"description"`
	Privileged  bool         `json:"privileged"`
	Location    *Coordinates `json:"location"`
}

// Passed to createUser to create a new user
type UserInput struct {
	// The description text
	Description string `json:"description"`
	// Is the user privileged?
	Privileged *bool `json:"privileged"`
}

type Weather struct {
	Summary     *Summary     `json:"summary"`
	Temperature *Temperature `json:"temperature"`
	Wind        *Wind        `json:"wind"`
	Clouds      *Clouds      `json:"clouds"`
	Timestamp   *int         `json:"timestamp"`
}

type Wind struct {
	Speed *float64 `json:"speed"`
	Deg   *int     `json:"deg"`
}

type Language string

const (
	LanguageAf   Language = "af"
	LanguageAl   Language = "al"
	LanguageAr   Language = "ar"
	LanguageAz   Language = "az"
	LanguageBg   Language = "bg"
	LanguageCa   Language = "ca"
	LanguageCz   Language = "cz"
	LanguageDa   Language = "da"
	LanguageDe   Language = "de"
	LanguageEl   Language = "el"
	LanguageEn   Language = "en"
	LanguageEu   Language = "eu"
	LanguageFa   Language = "fa"
	LanguageFi   Language = "fi"
	LanguageFr   Language = "fr"
	LanguageGl   Language = "gl"
	LanguageHe   Language = "he"
	LanguageHi   Language = "hi"
	LanguageHr   Language = "hr"
	LanguageHu   Language = "hu"
	LanguageID   Language = "id"
	LanguageIt   Language = "it"
	LanguageJa   Language = "ja"
	LanguageKr   Language = "kr"
	LanguageLa   Language = "la"
	LanguageLt   Language = "lt"
	LanguageMk   Language = "mk"
	LanguageNo   Language = "no"
	LanguageNl   Language = "nl"
	LanguagePl   Language = "pl"
	LanguagePt   Language = "pt"
	LanguagePtBr Language = "pt_br"
	LanguageRo   Language = "ro"
	LanguageRu   Language = "ru"
	LanguageSv   Language = "sv"
	LanguageSe   Language = "se"
	LanguageSk   Language = "sk"
	LanguageSl   Language = "sl"
	LanguageSp   Language = "sp"
	LanguageEs   Language = "es"
	LanguageSr   Language = "sr"
	LanguageTh   Language = "th"
	LanguageTr   Language = "tr"
	LanguageUa   Language = "ua"
	LanguageUk   Language = "uk"
	LanguageVi   Language = "vi"
	LanguageZhCn Language = "zh_cn"
	LanguageZhTw Language = "zh_tw"
	LanguageZu   Language = "zu"
)

var AllLanguage = []Language{
	LanguageAf,
	LanguageAl,
	LanguageAr,
	LanguageAz,
	LanguageBg,
	LanguageCa,
	LanguageCz,
	LanguageDa,
	LanguageDe,
	LanguageEl,
	LanguageEn,
	LanguageEu,
	LanguageFa,
	LanguageFi,
	LanguageFr,
	LanguageGl,
	LanguageHe,
	LanguageHi,
	LanguageHr,
	LanguageHu,
	LanguageID,
	LanguageIt,
	LanguageJa,
	LanguageKr,
	LanguageLa,
	LanguageLt,
	LanguageMk,
	LanguageNo,
	LanguageNl,
	LanguagePl,
	LanguagePt,
	LanguagePtBr,
	LanguageRo,
	LanguageRu,
	LanguageSv,
	LanguageSe,
	LanguageSk,
	LanguageSl,
	LanguageSp,
	LanguageEs,
	LanguageSr,
	LanguageTh,
	LanguageTr,
	LanguageUa,
	LanguageUk,
	LanguageVi,
	LanguageZhCn,
	LanguageZhTw,
	LanguageZu,
}

func (e Language) IsValid() bool {
	switch e {
	case LanguageAf, LanguageAl, LanguageAr, LanguageAz, LanguageBg, LanguageCa, LanguageCz, LanguageDa, LanguageDe, LanguageEl, LanguageEn, LanguageEu, LanguageFa, LanguageFi, LanguageFr, LanguageGl, LanguageHe, LanguageHi, LanguageHr, LanguageHu, LanguageID, LanguageIt, LanguageJa, LanguageKr, LanguageLa, LanguageLt, LanguageMk, LanguageNo, LanguageNl, LanguagePl, LanguagePt, LanguagePtBr, LanguageRo, LanguageRu, LanguageSv, LanguageSe, LanguageSk, LanguageSl, LanguageSp, LanguageEs, LanguageSr, LanguageTh, LanguageTr, LanguageUa, LanguageUk, LanguageVi, LanguageZhCn, LanguageZhTw, LanguageZu:
		return true
	}
	return false
}

func (e Language) String() string {
	return string(e)
}

func (e *Language) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Language(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Language", str)
	}
	return nil
}

func (e Language) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleOwner   Role = "OWNER"
	RoleRegular Role = "REGULAR"
)

var AllRole = []Role{
	RoleAdmin,
	RoleOwner,
	RoleRegular,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleOwner, RoleRegular:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Unit string

const (
	UnitMetric   Unit = "metric"
	UnitImperial Unit = "imperial"
	UnitKelvin   Unit = "kelvin"
)

var AllUnit = []Unit{
	UnitMetric,
	UnitImperial,
	UnitKelvin,
}

func (e Unit) IsValid() bool {
	switch e {
	case UnitMetric, UnitImperial, UnitKelvin:
		return true
	}
	return false
}

func (e Unit) String() string {
	return string(e)
}

func (e *Unit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Unit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Unit", str)
	}
	return nil
}

func (e Unit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
