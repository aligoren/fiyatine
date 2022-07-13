/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/aligoren/fiyatine/internal/models"
	"github.com/aligoren/fiyatine/internal/render"
	"github.com/aligoren/fiyatine/internal/services"
	"github.com/spf13/cobra"
)

// n11Cmd represents the n11 command
var n11Cmd = &cobra.Command{
	Use:   "n11",
	Short: "n11 üzerinde arama yap",
	Long:  "Spesifik olarak n11 üzerinde arama yapmanıza olanak sağlar",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("lütfen aranacak ürünün adını giriniz")
		}

		productName := strings.Join(args, " ")

		n11 := services.N11Service{
			SearchParams: models.ProductSearchModel{
				ProductName: productName,
			},
		}

		service := services.BaseService{ProductService: n11}

		products := service.Search()

		if len(products) == 0 {
			log.Println("n11 sitesinde aradığınız kriterlere uygun ürün bulunamadı")
		}

		headers := []string{"Satıcı", "Ürün Adı", "Fiyat", "Url"}
		rows := [][]string{}

		for _, product := range products {
			rows = append(rows, []string{"N11", product.Title, product.Price, product.Url})
		}

		renderer := render.TableRenderer{
			Headers:        headers,
			Rows:           rows,
			AutoWrapText:   false,
			RowLine:        true,
			AutoMergeCells: false,
		}

		renderer.RenderOutput()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(n11Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// n11Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// n11Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}