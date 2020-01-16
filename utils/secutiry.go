/**
 * @Time  : 2020-01-16 18:27
 * @Author: Lynn
 * @File  : secutiry
 * @Description:
 * @History:
 *  1.[2020-01-16-18:27] new created
 */
package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func MD5(password, salt string) string {
	p := md5.New()
	io.WriteString(p, password)
	pwd := fmt.Sprintf("%x\n", p.Sum(nil))
	buf := bytes.NewBufferString("")
	io.WriteString(buf, pwd)
	io.WriteString(buf, salt)
	_md5 := md5.New()
	io.WriteString(_md5, buf.String())
	return hex.EncodeToString(_md5.Sum(nil))
}
