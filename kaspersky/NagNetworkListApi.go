/*
 * MIT License
 *
 * Copyright (c) [2020] [Semchenko Aleksandr]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

//	NagNetworkListApi Class Reference
//	Nagent OpenAPI to work with network lists.
//
//	List of all members.
type NagNetworkListApi service

type NetworkListParams struct {
	//Network list name
	ListName string `json:"listName"`

	//Network list item identifier
	ItemID string `json:"itemId"`

	//if true than will be get info about packed content of the file.
	//It is usefull if file is the virus and it is needed to load it.
	BNeedPacked bool `json:"bNeedPacked"`

	//start position of the chunk
	ULStartPos int64 `json:"ulStartPos"`

	//number of bytes to read
	LBytesToRead int64 `json:"lBytesToRead"`
}

//	Retrieves information about file associated with network list item.
//
//	Values for listName and itemId see in Network lists which elements are associated with
//	files Network lists which elements are associated with files.
//
//	Parameters:
//	- listName	Network list name.
//	- itemId	Network list item identifier
//	- bNeedPacked	if true than will be get info about packed content of the file.
//	It is useful if file is the virus and it is needed to load it.
//
//	Returns:
//	- pFileInfo	Contains information about queried file. contains following attributes:
//		|- KLNLST_FILE_NAME, String, name of the file
//		|- KLNLST_TOTAL_SIZE, Integer, total size of the file
//	- lResult int64, value from enum NLSTFileOperationResult.
func (nnla *NagNetworkListApi) GetListItemFileInfo(ctx context.Context, params NetworkListParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil
	}

	request, err := http.NewRequest("POST", nnla.client.Server+"/api/v1.0/NagNetworkListApi.GetListItemFileInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nnla.client.Do(ctx, request, nil)
	return raw, err
}

//	Retrieves chunk of the file associated with network list item.
//
//	Values for listName and itemId see in Network lists which elements are associated with files
//	Network lists which elements are associated with files.
//
//	Parameters:
//	- listName	Network list name
//	- itemId	Network list item identifier
//	- bNeedPacked	if true than will be get info about packed content of the file.
//	It is usefull if file is the virus and it is needed to load it.
//	- ulStartPos	start position of the chunk
//	- lBytesToRead	number of bytes to read
//
//	Return values:
//	- pChunk	binary data as array of bytes.
//	- lResult	int64, value from enum NLSTFileOperationResult
func (nnla *NagNetworkListApi) GetListItemFileChunk(ctx context.Context, params NetworkListParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil
	}

	request, err := http.NewRequest("POST", nnla.client.Server+"/api/v1.0/NagNetworkListApi.GetListItemFileChunk", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nnla.client.Do(ctx, request, nil)
	return raw, err
}
