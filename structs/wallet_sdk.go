/*
Copyright ArxanFintech Technology Ltd. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package structs

import (
	"net/http"

	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
)

////////////////////////////////////////////////////////////////////////////////
// Wallet Client Structs

// Register wallet request structure
type RegisterWalletBody struct {
	Id        Identifier         `json:"id"`         //Optional, if empty, generated by wallet service
	Type      DidType            `json:"type"`       //Wallet Type: 1. "Organization"; 2. "Dependent"; 3. "Independent"; 4. "Asset"
	Access    string             `json:"access"`     //Register user name
	Secret    string             `json:"secret"`     //Register user passwd
	PublicKey *ed25519.PublicKey `json:"public_key"` //User public key. Optional, if empty, keypaire auto generated by wallet service
}

// Register subwallet request structure
type RegisterSubWalletBody struct {
	Id        Identifier         `json:"id"`         //main wallet id
	Type      DidType            `json:"type"`       //Wallet Type: 1."cash"; 2."fee"; 3."loan"; 4."interest"
	PublicKey *ed25519.PublicKey `json:"public_key"` //User public key. Optional, if empty, keypaire auto generated by wallet service
}

// WalletRequest common struct for transfer
type WalletRequest struct {
	Payload   string         `json:"payload"`
	Signature *SignatureBody `json:"signature"`
}

// WalletResponse common struct for wallet API response
type WalletResponse struct {
	Code           int      `json:"code"`
	Message        string   `json:"message"`
	Id             string   `json:"id"`
	Endpoint       string   `json:"endpoint"`
	KeyPair        *KeyPair `json:"key_pair"`
	Created        int64    `json:"created"`
	TransactionIds []string `json:"transaction_ids"`
}

type KeyPair struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

type WalletBalance struct {
	ColoredCoins  map[string]*ColoredCoin  `json:"colored_coins"`  //钱包中所有染色币
	DigitalAssets map[string]*DigitalAsset `json:"digital_assets"` //钱包中的数字资产
}

type (
	Identifier  string
	DidEndpoint string
	DidType     string
	DidStatus   string
)

const (
	DSValid   DidStatus = "Valid"
	DSInvalid DidStatus = "Invalid"
)

type WalletInfo struct {
	Id       Identifier                 `json:"id"`
	Type     DidType                    `json:"type"`
	Endpoint DidEndpoint                `json:"endpoint"`
	Status   DidStatus                  `json:"status"`
	Created  int64                      `json:"created"`
	Updated  int64                      `json:"updated"`
	HDS      map[Identifier]*WalletInfo `json:"hds"`
}

// IWalletClient defines the behaviors implemented by wallet sdk
type IWalletClient interface {
	Register(http.Header, *RegisterWalletBody) (*WalletResponse, error)
	RegisterSubWallet(http.Header, *RegisterSubWalletBody) (*WalletResponse, error)
	TransferCCoin(http.Header, *TransferBody, *SignatureBody) (*WalletResponse, error)
	TransferAsset(http.Header, *TransferAssetBody, *SignatureBody) (*WalletResponse, error)
	GetWalletBalance(http.Header, Identifier) (*WalletBalance, error)
	GetWalletInfo(http.Header, Identifier) (*WalletInfo, error)
}

////////////////////////////////////////////////////////////////////////////////
// POE Client Structs

// POE请求Body结构定义
type POEBody struct {
	Id         Identifier  `json:"id"`
	Name       string      `json:"name"`
	Hash       string      `json:"hash"`
	ParentId   Identifier  `json:"parent_id"`
	Owner      Identifier  `json:"owner"`
	ExpireTime int64       `json:"expire_time"`
	Metadata   interface{} `json:"metadata"`
}

// POEPayload POE查询返回结构定义
type POEPayload struct {
	Id         Identifier  `json:"id"`
	Name       string      `json:"name"`
	Hash       string      `json:"hash"`
	ParentId   Identifier  `json:"parent_id"`
	Owner      Identifier  `json:"owner"`
	ExpireTime int64       `json:"expire_time"`
	Metadata   interface{} `json:"metadata"`
	Created    int64       `json:"create_time"`
	Updated    int64       `json:"update_time"`
	Status     DidStatus   `json:"status"`
}

// OffchainMetadata offchain storage metadata
type OffchainMetadata struct {
	Filename    string `json:"filename"`
	Endpoint    string `json:"endpoint"`
	StorageType string `json:"storageType"`
	ContentHash string `json:"contentHash"`
	Size        int    `json:"size"`
}

const (
	// OffchainPOEID poe did when invoke upload file api formdata param
	OffchainPOEID = "poe_id"
	// OffchainPOEFile poe binary file when invoke upload file api formdata param
	OffchainPOEFile = "poe_file"
)

// signature const for the formdata
const (
	// SignatureCreator issuer did
	SignatureCreator = "signature.creator"
	// SignatureCreated timestamp
	SignatureCreated = "signature.created"
	// SignatureNonce nonce
	SignatureNonce = "signature.nonce"
	// SignatureSignatureValue sign value
	SignatureSignatureValue = "signature.signatureValue"
)

// IPOEClient defines the behaviors implemented by POE sdk
type IPOEClient interface {
	CreatePOE(http.Header, *POEBody, *SignatureBody) (*WalletResponse, error)
	UpdatePOE(http.Header, *POEBody, *SignatureBody) (*WalletResponse, error)
	IssuePOE(http.Header, *IssueBody, *SignatureBody) (*WalletResponse, error)
	WithdrawPOE(http.Header, *TransferBody, *SignatureBody) (*WalletResponse, error)
	QueryPOE(http.Header, Identifier) (*POEPayload, error)
}
