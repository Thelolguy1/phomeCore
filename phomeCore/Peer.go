// This file is for handling Peers, including pairing and storage of existing paired devices.
package phomeCore

import (
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

type BundleType int64

const (
	LocationReport BundleType = iota
	InitialPairing
) // This may be extended for personal use.

type JSONBundle struct { // JSON Bundles are used for pairing and general purpose.
	UUID   string
	PubKey string
	BundleType BundleType // Don't rely on this, since bundles may be mislabeled.
	Data   any // This is empty during initial pairing.
}

// This function generates the initial pairing JSON from a JSONBundle.
// It is recommended to convert the string output to base64 for pairing.
func (newPairingJSON *JSONBundle) GeneratePairingJSON () string {
	newPairingJSON.BundleType = InitialPairing
	jsonStr, err := json.Marshal(newPairingJSON)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonStr)
}

// This function unmarshals a pairing JSON string into a JSONBundle
func (newPairingJSON *JSONBundle) DecodePairingJSON (pairingJSONstr string){
	err := json.Unmarshal([]byte(pairingJSONstr), &newPairingJSON)
	if err != nil {
		log.Fatal(err)
	}
}

// Note: GenCerts in Encryption.go also generates the localhost UUID.
func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func EncodeB64(in string) string { // Output can be used in a external program, like a QR generator.
	return base64.URLEncoding.EncodeToString([]byte(in))
}

func DecodeB64(in string) string {
	data, err := base64.URLEncoding.DecodeString(in)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}