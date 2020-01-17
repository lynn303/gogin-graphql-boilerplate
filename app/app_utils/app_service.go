/**
 * @Time  : 2020-01-17 10:18
 * @Author: Lynn
 * @File  : app_service
 * @Description:
 * @History:
 *  1.[2020-01-17-10:18] new created
 */
package app_utils

//响应返回参数
type Result struct {
	Success bool `json:"success"`
}

func HsResult(success bool, err error) (Result, error) {
	if !success && err != nil {
		return Result{Success: success}, err
	} else if success {
		return Result{Success: success}, nil
	}
	return Result{Success: false}, err
}
