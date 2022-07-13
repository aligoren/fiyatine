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
	"log"

	"github.com/aligoren/fiyatine/internal/models"
	"github.com/aligoren/fiyatine/internal/render"
	"github.com/aligoren/fiyatine/internal/services"
	"github.com/spf13/cobra"
)

// hepsiburadaCmd represents the hepsiburada command
var hepsiburadaCmd = &cobra.Command{
	Use:   "hepsiburada",
	Short: "Hepsiburada üzerinde arama yap",
	Long:  `Spesifik olarak hepsiburada üzerinde arama yapmanıza olanak sağlar`,
	Run: func(cmd *cobra.Command, args []string) {

		hb := services.HepsiburadaService{
			SearchParams: models.ProductSearchModel{
				ProductName: "Ütü Masası",
			},
		}

		service := services.BaseService{ProductService: hb}

		products := service.Search()

		if len(products) == 0 {
			log.Println("Hepsiburada sitesinde aradığınız kriterlere uygun ürün bulunamadı")
		}

		headers := []string{"Satıcı", "ID", "Ürün Adı", "Fiyat"}
		rows := [][]string{}

		for _, product := range products {
			rows = append(rows, []string{"Hepsiburada", product.ID, product.Title, product.Price})
		}

		render.RenderOutput(headers, rows)
	},
}

func init() {
	rootCmd.AddCommand(hepsiburadaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hepsiburadaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hepsiburadaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
