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
	"sort"
	"strings"
	"sync"

	"github.com/aligoren/fiyatine/internal/models"
	"github.com/aligoren/fiyatine/internal/render"
	"github.com/aligoren/fiyatine/internal/services"
	"github.com/spf13/cobra"
)

// tumuCmd represents the tumu command
var tumuCmd = &cobra.Command{
	Use:   "tumu",
	Short: "Tüm satıcılar üzerinde arama yap",
	Long:  "Tüm satıcılar üzerinde arama yapmanıza olanak sağlar",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("lütfen aranacak ürünün adını giriniz")
		}

		productName := strings.Join(args, " ")

		searchParams := models.ProductSearchModel{
			ProductName: productName,
		}

		baseServices := []services.BaseService{
			{
				ProductService: services.AmazonService{SearchParams: searchParams},
			},
			{
				ProductService: services.TrendyolService{SearchParams: searchParams},
			},
			{
				ProductService: services.HepsiburadaService{SearchParams: searchParams},
			},
			{
				ProductService: services.N11Service{SearchParams: searchParams},
			},
		}

		var products []models.ResponseModel

		var wg sync.WaitGroup
		var mu sync.Mutex
		for _, service := range baseServices {
			wg.Add(1)
			go func(baseService services.BaseService) {
				defer wg.Done()
				mu.Lock()
				products = append(products, baseService.Search()...)
				mu.Unlock()
			}(service)

		}
		wg.Wait()

		if len(products) == 0 {
			log.Println("Tüm aramalar içerisinde aradığınız kriterlere uygun ürün bulunamadı")
		}

		sort.Slice(products, func(i, j int) bool {
			return products[i].PriceField < products[j].PriceField
		})

		headers := []string{"Satıcı", "Ürün Adı", "Fiyat", "Url"}
		rows := [][]string{}

		for _, product := range products {
			rows = append(rows, []string{product.Vendor, product.Title, product.Price, product.Url})
		}

		renderer := render.TableRenderer{
			Headers:        headers,
			Rows:           rows,
			AutoWrapText:   true,
			RowLine:        true,
			AutoMergeCells: false,
		}

		renderer.RenderOutput()

		return nil

	},
}

func init() {
	rootCmd.AddCommand(tumuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tumuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tumuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
