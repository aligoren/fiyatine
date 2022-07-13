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

// trendyolCmd represents the trendyol command
var trendyolCmd = &cobra.Command{
	Use:   "trendyol",
	Short: "Trendyol üzerinde arama yap",
	Long:  "Spesifik olarak trendol üzerinde arama yapmanıza olanak sağlar",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("lütfen aranacak ürünün adını giriniz")
		}

		productName := strings.Join(args, " ")

		trendyol := services.TrendyolService{
			SearchParams: models.ProductSearchModel{
				ProductName: productName,
			},
		}

		service := services.BaseService{ProductService: trendyol}

		products := service.Search()

		if len(products) == 0 {
			log.Println("Trendyol sitesinde aradığınız kriterlere uygun ürün bulunamadı")
		}

		headers := []string{"Satıcı", "Ürün Adı", "Fiyat", "Url"}
		rows := [][]string{}

		for _, product := range products {
			rows = append(rows, []string{"Trendyol", product.Title, product.Price, product.Url})
		}

		render.RenderOutput(headers, rows)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(trendyolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trendyolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trendyolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
