/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//	FileCategorizer2 Class Reference
//
//	Interface for working with FileCategorizer subsystem..
//	There are 3 types of categories: simple, autoupdate and silverimage.
//	Simple category can be created by user manually.
//	Autoupdate category is working on server side and calculating hashes of files from choosen directory.
//	SilverImage category type accumulates hashes of files from choosen hosts.
//	List of all members.
type FileCategorizer2 service

//TODO AddExpressions

//	Cancel file metadata operations.
//
//	Method canceles operation (GetFileMetadata, GetFilesMetadata, GetFilesMetadataFromMSI) initialized using current connection.
func (fc *FileCategorizer2) CancelFileMetadataOperations(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2."+
		"CancelFileMetadataOperations", nil)

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Cancel file upload for file categorizer subsystem.
//
//	This methode canceles file upload.
//	Call FileCategorizer2.InitFileUpload to start new upload.
func (fc *FileCategorizer2) CancelFileUpload(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2."+
		"CancelFileUpload", nil)

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//CategoryParams struct <- TODO Correct this... working but, need all fields
type CategoryParams struct {
	PCategory *PCategory `json:"pCategory,omitempty"`
}

type PCategory struct {
	CategoryType                *int64       `json:"CategoryType,omitempty"`
	CustomCategoryCipCompatible *bool        `json:"CustomCategoryCipCompatible,omitempty"`
	Md5WithoutSha256Exists      *bool        `json:"Md5WithoutSha256Exists,omitempty"`
	Exclusions                  []Exclusions `json:"exclusions"`
	FromMaster                  *bool        `json:"fromMaster,omitempty"`
	Inclusions                  []Inclusion  `json:"inclusions"`
	Name                        *string      `json:"name,omitempty"`
	Descr                       *string      `json:"descr,omitempty"`
	Version                     *int64       `json:"version,omitempty"`
}

type Exclusions struct {
	//TODO Body fields
}

type Inclusion struct {
	Type           *string         `json:"type,omitempty"`
	InclusionValue *InclusionValue `json:"value,omitempty"`
}

type InclusionValue struct {
	ExType *int64  `json:"ex_type,omitempty"`
	Str    *string `json:"str,omitempty"`
	StrOp  *int64  `json:"str_op,omitempty"`
	//TODO Body fields
}

//	Create category (simple, autoupdate or silverimage)
//
//	Parameters:
//	- pCategory	(params) Category body (see Custom category format)
//
//	Returns:
//	- (int64) Category id
//
//	Exceptions:
//	- KLSTD.STDE_OBJ_EXISTS	- name or UUID is not unique
func (fc *FileCategorizer2) CreateCategory(ctx context.Context, params CategoryParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.CreateCategory", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := fc.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Delete category.
//
//	Parameters:
//	- nCategoryId (int) Category id
//
//	Exceptions:
//	- KLSTD.STDE_NOTFOUND	- category not found
//	- KLSTD.STDE_NOACCESS
func (fc *FileCategorizer2) DeleteCategory(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DeleteCategory",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//TODO DeleteExpression
//TODO DoStaticAnalysisAsync
//TODO DoStaticAnalysisAsync2
//TODO DoTestStaticAnalysisAsync
//TODO DoTestStaticAnalysisAsync2
//TODO FinishStaticAnalysis

//	Force process of automatic update (for autoupdate and silverimage)
//
//	Parameters:
//	- nCategoryId (int64) Category id
//
//	Exceptions:
//	KLSTD.STDE_NOTPERM	- wrong category type
func (fc *FileCategorizer2) ForceCategoryUpdate(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.ForceCategoryUpdate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns modification counter. It increments on every category change.
//
//	Returns:
//	- (int64) Modification counter
func (fc *FileCategorizer2) GetCategoriesModificationCounter(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2."+
		"GetCategoriesModificationCounter", nil)

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Get category by id.
//
//	Parameters:
//	 - nCategoryId	(int64) Category id
//
//	Return:
//	- pCategory	(params) Category body (see Custom category format)
//
//	Exceptions:
//	- KLSTD.STDE_NOTFOUND	- category not found
//	- KLSTD.STDE_BADFORMAT	- format of category is wrong
func (fc *FileCategorizer2) GetCategory(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetCategory",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Get category by uuid.
//
//	Parameters:
//	- pCategoryUUID	(string) Category uuid
//
//	Return:
//		- pCategory	(params) Category body (see Custom category format)
//
//	Exceptions:
//	- KLSTD.STDE_NOTFOUND	- category not found
func (fc *FileCategorizer2) GetCategoryByUUID(ctx context.Context, pCategoryUUID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pCategoryUUID": "%s"}`, pCategoryUUID))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetCategoryByUUID",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Get file metadata.
//
//	Parameters:
//	- ulFlag (int64) Requested meta information. Use like a mask of flags. See File metadata flags.
//
//	Return:
//	- wstrAsyncId (string) Id of async operation.
//
//	To get result use AsyncActionStateChecker.CheckActionState.
//	It returns params with requested attributes.
//	See list of attributes File metadata flags.
//
//	See also:
//	AsyncActionStateChecker
func (fc *FileCategorizer2) GetFileMetadata(ctx context.Context, ulFlag int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"ulFlag": %d}`, ulFlag))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetFileMetadata",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Get files metadata from zip-archive.
//
//	Parameters:
//	- ulFlag (int64) Requested meta information. Use like a mask of flags. See File metadata flags.
//
//	Return:
//	- wstrAsyncId (string) Id of async operation.
//
//	To get action status use AsyncActionStateChecker.CheckActionState.
//	When action is not finished and lStateCode equals 2 then task in progress
//	and pStateData may contain attribute "Progress" (int).
//	When action is successfully finished it returns pStateData with an array "FilesMetadata".
//	Each element is a params with requested attributes. See list of attributes File metadata flags.
//
//	See also:
//	AsyncActionStateChecker
func (fc *FileCategorizer2) GetFilesMetadata(ctx context.Context, ulFlag int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"ulFlag": %d}`, ulFlag))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetFilesMetadata",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Get files metadata from MSI.
//
//	Parameters:
//	- ulFlag (int64) Requested meta information. Use like a mask of flags. See File metadata flags.
//
//	Return:
//	- wstrAsyncId	(string) Id of async operation.
//	To get action status use AsyncActionStateChecker.CheckActionState.
//	When action is not finished and lStateCode equals 2 then task in progress
//	and pStateData may contain attribute "Progress" (int).
//	When action is successfully finished it returns pStateData with an array "FilesMetadata".
//	Each element is a params with requested attributes. See list of attributes File metadata flags.
//
//	See also:
//	AsyncActionStateChecker
func (fc *FileCategorizer2) GetFilesMetadataFromMSI(ctx context.Context, ulFlag int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"ulFlag": %d}`, ulFlag))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetFilesMetadataFromMSI",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//RefPolicies struct
type RefPolicies struct {
	PPolicies *PPolicies `json:"pPolicies,omitempty"`
}

type PPolicies struct {
	RefPolicies []RefPolicy `json:"RefPolicies"`
}

type RefPolicy struct {
	Type           *string         `json:"type,omitempty"`
	RefPolicyValue *RefPolicyValue `json:"value,omitempty"`
}

type RefPolicyValue struct {
	FromMaster  *bool   `json:"FromMaster,omitempty"`
	GroupID     *int64  `json:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty"`
	PolID       *int64  `json:"PolId,omitempty"`
	PolName     *string `json:"PolName,omitempty"`
	VServerID   *int64  `json:"VServerId,omitempty"`
	VServerName *string `json:"VServerName,omitempty"`
}

//	Returns array of policies with references to specified category.
//
//	Parameters:
//	- nCatId (int64) Category id
//
//	Return:
//	- pPolicies	(params) See Policies array
//
//	Exceptions:
//	- KLSTD.STDE_NOACCESS	- access denied
func (fc *FileCategorizer2) GetRefPolicies(ctx context.Context, nCatId int64) (*RefPolicies, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCatId": %d}`, nCatId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetRefPolicies",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	refPolicies := new(RefPolicies)
	raw, err := fc.client.Do(ctx, request, &refPolicies)
	return refPolicies, raw, err
}

//	Returns serialized category body for plugin.
//
//	Warning:
//	Deprecated for using in OpenAPI. Use FileCategorizer2.GetSerializedCategoryBody2 instead.
//
//	Parameters:
//	- nCategoryId (int) Category id
//
//	Return:
//	 - pCategory (params) Category serialized body
func (fc *FileCategorizer2) GetSerializedCategoryBody(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetSerializedCategoryBody",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns serialized category body for plugin.
//
//	Parameters:
//	- nCategoryId (int64) Category id
//
//	Return:
//	- pCategory	(params) Category serialized body
//
//	See also:
//	See Serialized category format
func (fc *FileCategorizer2) GetSerializedCategoryBody2(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetSerializedCategoryBody2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns categories synchronization id.
//
//	Returns:
//	- (int64) Synchronization id
//
//	See also:
//	See GroupSync
func (fc *FileCategorizer2) GetSyncId(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2."+
		"GetSyncId", nil)

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//TODO InitFileUpload
//TODO UpdateCategory
//TODO UpdateExpressions