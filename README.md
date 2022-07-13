## Fiyati Ne

Bu proje bir CLI projesidir. Bu proje ile birlikte Amazon, Trendyol, Hepsiburada vb. sitelerden ürün arayıp tek seferde listeleme ya da site bazlı (Örn: Amazon) listeleme yapılabilir.

Provider sayısı daha da genişletilebilir.

## NOT

**Ekonominin çok iyi olması, şahlanması nedeniyle fiyatları ucuzdan pahalıya olacak şekilde sıraladım. Durduk yere dış güçler master'a pushlasın istemem**


## Kullanım

Kullanımı şu şekillerde olabilir;
### Amazon

`go run . amazon Ütü Masası`
### Hepsiburada

`go run . hepsiburada Ütü Masası`

### Trendyol

`go run . trendyol Ütü Masası`

### Tümü

Bu şekilde arama yapacaksanız da tamamında arama yapar.

`go run . tumu Ütü Masası`

## Örnek Çıktı

![/images/screenshot.png](/images/screenshot.png)