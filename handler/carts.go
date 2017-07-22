/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
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

/*
 * Revision History:
 *     Initial: 2017/07/21        Zhu Yaqiang
 */

package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"ShopApi/general"
	"ShopApi/general/errcode"
	"ShopApi/log"
	"ShopApi/models"
)

func Cartsdel(c echo.Context) error {
	var (
		err    error
		cartid models.CartsDel
	)

	if err = c.Bind(&cartid); err != nil {
		log.Logger.Error("Analysis crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.CartsService.WhetherInCart(cartid.ID, cartid.ProID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("This product doesn't exist !", err)

			return general.NewErrorWithMessage(errcode.ErrInformation, err.Error())
		}

		log.Logger.Error("Check existence exists errors", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	err = models.CartsService.CartsDelete(cartid.ID, cartid.ProID)

	if err != nil {
		log.Logger.Error("Delete product with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func AlterCartPro(c echo.Context) error {
	var (
		err     error
		cartpro models.CartPro
	)

	if err = c.Bind(&cartpro); err != nil {
		log.Logger.Error("Get crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.CartsService.AlterCartPro(cartpro.ID, cartpro.Count, cartpro.Size, cartpro.Color)
	if err == gorm.ErrRecordNotFound {
		log.Logger.Error("The product doesn't exist !", err)

		return general.NewErrorWithMessage(errcode.ErrNotFound, err.Error())
	}

	err = models.CartsService.AlterCartPro(cartpro.ID, cartpro.Count, cartpro.Size, cartpro.Color)

	if err != nil {
		log.Logger.Error("Alter product with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}
