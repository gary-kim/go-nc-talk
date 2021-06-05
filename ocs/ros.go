// Copyright (c) 2020 Gary Kim <gary@garykim.dev>, All Rights Reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// File describes Nextcloud's Rich Object Strings (https://github.com/nextcloud/server/issues/1706)

package ocs

import "encoding/json"

// RichObjectString describes the content of placeholders in TalkRoomMessageData
type RichObjectString struct {
	Type RichObjectStringType `json:"type"`
	ID   string               `json:"id"`
	Name string               `json:"name"`
	Path string               `json:"path"`
	Link string               `json:"link"`
}

// RichObjectStringType describes what a rich object string is describing
type RichObjectStringType string

const (
	// ROSTypeUser describes a rich object string that is a user
	ROSTypeUser RichObjectStringType = "user"
	// ROSTypeGroup describes a rich object string that is a group
	ROSTypeGroup RichObjectStringType = "group"
	// ROSTypeFile describes a rich object string that is a file
	ROSTypeFile RichObjectStringType = "file"
)

// RichObjectStringDefinitions represents a Nextcloud rich object
//
// See https://github.com/nextcloud/server/blob/master/lib/public/RichObjectStrings/Definitions.php
// for details on supported rich objects.
type RichObjectStringDefinitions interface {
	ToJSON() ([]byte, error)
	ObjectType() string
}

type ROSAddressBook struct {
	// ID is the id used to identify the addressbook on the instance
	ID string `json:"id"`
	// Name is the display name of the addressbook which should be used in the visual representation
	Name string `json:"name"`
}

func (r *ROSAddressBook) ObjectType() string {
	return "addressbook"
}

func (r *ROSAddressBook) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSAddressBookContact struct {
	// ID is the id used to identify the contact on the instance
	ID string `json:"id"`
	// Name is the display name of the contact which should be used in the visual representation
	Name string `json:"name"`
}

func (r *ROSAddressBookContact) ObjectType() string {
	return "addressbook-contact"
}

func (r *ROSAddressBookContact) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSAnnouncement struct {
	// ID is the id used to identify the announcement on the instance
	ID string `json:"id"`
	// Name is the announcement subject which should be used in the visual representation
	Name string `json:"name"`
	// Link is the full URL to the file
	Link string `json:"link"`
}

func (r *ROSAnnouncement) ObjectType() string {
	return "announcement"
}

func (r *ROSAnnouncement) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSApp struct {
	// ID is the app id
	ID string `json:"id"`
	// Name is the name of the app which should be used in the visual representation
	Name string `json:"name"`
}

func (r *ROSApp) ObjectType() string {
	return "app"
}

func (r *ROSApp) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSCalendar struct {
	// ID is the id used to identify the calendar on the instance
	ID string `json:"id"`
	// Name is the display name of the calendar which should be used in the visual representation
	Name string `json:"name"`
}

func (r *ROSCalendar) ObjectType() string {
	return "calendar"
}

func (r *ROSCalendar) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSCalendarEvent struct {
	// ID is the id used to identify the event on the instance
	ID string `json:"id"`
	// Name is the display name of the event which should be used in the visual representation
	Name string `json:"name"`
	// Link is a link to the page displaying the calendar
	Link string `json:"link"`
}

func (r *ROSCalendarEvent) ObjectType() string {
	return "calendar-event"
}

func (r *ROSCalendarEvent) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSCall struct {
	// ID is the id used to identify the call on the instance
	ID string `json:"id"`
	// Name is the display name of the call which should be used in the visual representation
	Name string `json:"name"`
	// CallType is the type of call: "one2one", "group", or "public"
	CallType string `json:"call-type"`
	// Link is the link to the conversation
	Link string `json:"link"`
}

func (r *ROSCall) ObjectType() string {
	return "call"
}

func (r *ROSCall) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSCircle struct {
	// ID is the id used to identify the circle on the instance
	ID string `json:"id"`
	// Name is the display name of the circle which should be used in the visual representation
	Name string `json:"name"`
	// Link is the full URL to the circle
	Link string `json:"link"`
}

func (r *ROSCircle) ObjectType() string {
	return "circle"
}

func (r *ROSCircle) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSDeckBoard struct {
	// ID is the id used to identify the board on the instance
	ID string `json:"id"`
	// Name is the display name of the deck board
	Name string `json:"name"`
	// Link is the full URL to the board
	Link string `json:"link"`
}

func (r *ROSDeckBoard) ObjectType() string {
	return "deck-board"
}

func (r *ROSDeckBoard) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSDeckCard struct {
	// ID is the id used to identify the card on the instance
	ID string `json:"id"`
	// Name is the title of the deck card
	Name string `json:"name"`
	// BoardName is the display name of the board which contains the card
	BoardName string `json:"boardname"`
	// StackName is the display name of the stack which contains the card in the board
	StackName string `json:"stackname"`
	// Link is the full URL to the card directly
	Link string `json:"link"`
}

func (r *ROSDeckBoard) ObjectType() string {
	return "deck-card"
}

func (r *ROSDeckCard) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSEmail struct {
	// ID is the mail address used to identify the event on the instance
	ID string `json:"id"`
	// Name is the display name of a matching contact or the email (fallback) which should be used in the visual representation
	Name string `json:"name"`
}

func (r *ROSEmail) ObjectType() string {
	return "email"
}

func (r *ROSEmail) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSFile struct {
	// ID is the id used to identify the file on the instance
	ID string `json:"id"`
	// Name is the name which should be used in the visual representation
	Name string `json:"name"`
	// Size is the file size in bytes
	Size string `json:"size"`
	// Path is the full path of the file for the user, should not start with a slash
	Path string `json:"path"`
	// Link is the full URL to the file
	Link string `json:"link"`
	// MimeType is the mimetype of the file/folder to allow clients to show a placeholder
	MimeType string `json:"mimetype"`
	// PreviewAvailable indicates whether or not a preview is available. If "no" the mimetype icon should be used. Can also be set to "yes"
	PreviewAvailable string `json:"preview-available"`
}

func (r *ROSFile) ObjectType() string {
	return "file"
}

func (r *ROSFile) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSFormsForm struct {
	// ID is the form-hash of the form
	ID string `json:"id"`
	// Name is the title of the form
	Name string `json:"name"`
	// Link is the full URL to the board
	Link string `json:"link"`
}

func (r *ROSFormsForm) ObjectType() string {
	return "forms-form"
}

func (r *ROSFormsForm) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSGuest struct {
	// ID is the id used to identify the guest user
	ID string `json:"id"`
	// Name is the potential displayname of the guest user
	Name string `json:"name"`
}

func (r *ROSGuest) ObjectType() string {
	return "guest"
}

func (r *ROSGuest) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSHighlight struct {
	// ID is the id used to identify the highlighted object on the instance
	ID string `json:"id"`
	// Name is the string that should be highlighted
	Name string `json:"name"`
	// Link is the full URL that should be opened when clicking the highlighted text
	Link string `json:"link"`
}

func (r *ROSHighlight) ObjectType() string {
	return "highlight"
}

func (r *ROSHighlight) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSGeoLocation struct {
	// ID is the geo URI (https://en.wikipedia.org/wiki/Geo_URI_scheme) to identify the location
	ID string `json:"id"`
	// Name is a description of the location
	Name string `json:"name"`
	// Latitude is the latitude of the location which MUST be the same as in the id
	Latitude string `json:"latitude"`
	// Longitude is the longitude of the location which MUST be the same as in the id
	Longitude string `json:"longitude"`
}

func (r *ROSGeoLocation) ObjectType() string {
	return "geo-location"
}

func (r *ROSGeoLocation) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSOpenGraph struct {
	// ID is the id used to identify the open graphs data on the instance
	ID string `json:"id"`
	// Name is the open graph title of the website
	Name string `json:"name"`
	// Description is the open graph description from the website
	Description string `json:"description"`
	// Thumb is the full URL of the open graph thumbnail
	Thumb string `json:"thumb"`
	// Website is the name of the described website
	Website string `json:"website"`
	// Link is the full link to the website
	Link string `json:"link"`
}

func (r *ROSOpenGraph) ObjectType() string {
	return "open-graph"
}

func (r *ROSOpenGraph) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSPendingFederatedShare struct {
	// ID is used to identify the federated share on the instance
	ID string `json:"id"`
	// Name is the name of the shared item which should be used in the visual representation
	Name string `json:"name"`
}

func (r *ROSPendingFederatedShare) ObjectType() string {
	return "pending-federated-share"
}

func (r *ROSPendingFederatedShare) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSSystemTag struct {
	// ID is the id used to identify the systemtag on the instance
	ID string `json:"id"`
	// Name is the display name of the systemtag which should be used in the visual representation
	Name string `json:"name"`
	// Visibility indicates whether the user can see the system tag. Should be "1" for yes and "0" for no
	Visibility string `json:"visibility"`
	// Assignable is if the user can assign the systemtag
	Assignable string `json:"assignable"`
}

func (r *ROSSystemTag) ObjectType() string {
	return "systemtag"
}

func (r *ROSSystemTag) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSTalkAttachment struct {
	// ID is the id used to identify the attachment on the instance
	ID string `json:"id"`
	// Name is the name of the attachment
	Name string `json:"name"`
	// Conversation is the token of the conversation
	Conversation string `json:"conversation"`
	// MimeType is the mimetype of the file/folder to allow client to show a placeholder
	MimeType string `json:"mimetype"`
	// PreviewAvailable is whether or not a preview is available. If "no", the mimetype icon should be used
	PreviewAvailable string `json:"preview-available"`
}

func (r *ROSTalkAttachment) ObjectType() string {
	return "talk-attachment"
}

func (r *ROSTalkAttachment) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSUser struct {
	// ID is the id used to identify the user on the instance
	ID string `json:"id"`
	// Name is the display name of the user which should be used in the visual representation
	Name string `json:"name"`
	// Server is the URL of the instance the user lives on
	Server string `json:"server"`
}

func (r *ROSUser) ObjectType() string {
	return "user"
}

func (r *ROSUser) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

type ROSUserGroup struct {
	// ID is the id used to identify the group on the instance
	ID string `json:"id"`
	// Name is the display name of the group which should be used in the visual representation
	Name string `json:"name"`
}

