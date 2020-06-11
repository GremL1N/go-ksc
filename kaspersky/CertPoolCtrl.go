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
	"fmt"
	"net/http"
)

//	CertPoolCtrl Class Reference
//
//	Interface to manage the pool of certificates used by the
//	Kaspersky Security Center Server.
//
//	Public Member Functions
type CertPoolCtrl service

//	Returns information about certificate from server's certificates pool.
//
//	Parameters:
//	- nVServerId (int64) Virtual server id (-1 for current, 0 for main server)
//	- nFunction (int64) Certificate function (see "KLCERTP.CertificateFunction enum values")
//
//	Returns:
//	- (params) If certificate present then it returns params
//	with "CPublic" (paramBinary) field only.
func (cpc *CertPoolCtrl) GetCertificateInfo(ctx context.Context, nVServerId, nFunction int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nVServerId": %d, "nFunction" : %d }`, nVServerId, nFunction))
	request, err := http.NewRequest("POST", cpc.client.Server+"/api/v1.0/CertPoolCtrl.GetCertificateInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := cpc.client.Do(ctx, request, nil)
	return raw, err
}

//	Sets certificate of a given function for a given virtual server.
//
//	Parameters:
//	- params interface{}
//	|- nVServerId	(int64) Virtual server id (-1 for current, 0 for main server)
//	|- nFunction	(int64) Certificate function (see "KLCERTP.CertificateFunction enum values")
//	|- pCertData	Params with certificate data (see Common format for certificate params).
//	If pCertData is empty or NULL, then certificate will be deleted.
func (cpc *CertPoolCtrl) SetCertificate(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", cpc.client.Server+"/api/v1.0/CertPoolCtrl.SetCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := cpc.client.Do(ctx, request, nil)
	return raw, err
}
