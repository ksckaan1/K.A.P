package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zserge/webview"
)

func kontrol(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Yakala ...
func Yakala(w webview.WebView, data string) {
	veri := strings.Split(data, "©")
	switch veri[0] {
	case "cikis": //Çıkış İşlemi
		w.Terminate()
	case "giris":
		vt, err := sql.Open("sqlite3", "./veritabanı.db")
		if err != nil {
			fmt.Println(err)
		}
		tablo, _ := vt.Query("SELECT * FROM ayarlar")
		var ayar, deger string
		for tablo.Next() {
			aktarma := tablo.Scan(&ayar, &deger)
			if ayar == "şifre" {
				şifre := md5.New()
				io.WriteString(şifre, veri[1])
				if hex.EncodeToString(şifre.Sum(nil)) == deger { //Burada şifre kontrolü yapıyor
					w.Eval("accessGranted()")
					w.Eval("onay('Başarıyla Giriş Yapıldı!');")
				} else {
					w.Eval("uyari('Yanlış Şifre Girdiniz!');")
				}
			}
			if aktarma == nil {

			}
		}
		tablo.Close()
		vt.Close()

	case "yazdir": //Profil Sayfasını Yazdırma
		id := veri[1]
		if id == "1" {

		}
		vt, err := sql.Open("sqlite3", "./veritabanı.db")
		if err != nil {
			fmt.Println(err)
		}
		tablo, _ := vt.Query("SELECT * FROM kisiler")
		var veriid int
		var veriad, verires, verihk, verigl string
		for tablo.Next() {
			aktarma := tablo.Scan(&veriid, &veriad, &verires, &verihk, &verigl)

			denk, _ := strconv.Atoi(id)
			if aktarma == nil && veriid == denk {
				fmt.Println("#"+strconv.Itoa(veriid), "id'li profile bakma: "+veriad)
				icerik := "yazdir('" + strconv.Itoa(veriid) + "','" + strings.ToUpperSpecial(unicode.TurkishCase, veriad) + "',`" + verihk + "`,`" + verires + "`,'" + verigl + "');"
				w.Eval(icerik)
			} else {

			}
		}
		tablo.Close()
		vt.Close()
	case "listecek": //Kayıt Listesini Oluşturma
		var kayitsayisi = 0
		fmt.Println("Kayıtlar listeleniyor...")
		vt, err := sql.Open("sqlite3", "./veritabanı.db")
		if err != nil {
			fmt.Println(err)
		}
		tablo, _ := vt.Query("SELECT * FROM kisiler")
		var veriid int
		var veriad, verires, verihk, verigl string
		for tablo.Next() {
			aktarma := tablo.Scan(&veriid, &veriad, &verires, &verihk, &verigl)

			if aktarma == nil {
				icerik := "listeolustur('" + strconv.Itoa(veriid) + "',`" + strings.ToUpperSpecial(unicode.TurkishCase, veriad) + "`);"
				w.Eval(icerik)
				kayitsayisi++
			} else {

			}
		}
		tablo.Close()
		vt.Close()
		fmt.Println("Toplam kayıt sayısı:", kayitsayisi)
	case "guncelle":
		var id = veri[1]
		var ad = veri[2]
		var pp = veri[3]
		var hk = veri[4]
		var gl = veri[5]
		vt, _ := sql.Open("sqlite3", "./veritabanı.db")
		işlem, _ := vt.Prepare("update kisiler set ad=?, resim=?, hakkinda=?, galeri=? where id=" + id)
		//Güncellenecek kısmı belirtiyoruz
		hk = strings.ReplaceAll(hk, "\n", "<br/>")
		hk = strings.ReplaceAll(hk, "[", `<span class="belirgin">`)
		hk = strings.ReplaceAll(hk, "]", `</span>`)
		veri, _ := işlem.Exec(strings.ToLowerSpecial(unicode.TurkishCase, ad), pp, hk, gl)
		//Değişiklik ve Değiştirilen verinin id'si
		değişiklik, _ := veri.RowsAffected()
		fmt.Println(değişiklik, "kayıt güncellendi. ID'si:", id, ".Kayıt ismi:", ad)
		vt.Close() //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
	case "sil":
		silid := veri[1]
		vt, _ := sql.Open("sqlite3", "./veritabanı.db")
		işlem, _ := vt.Prepare("delete from kisiler where id=" + silid)
		//id numarasına göre sileceğiz
		veri, _ := işlem.Exec()              //Silinecek kişinin id'si
		değişiklik, _ := veri.RowsAffected() //Silinen kişinin id'sini aldık
		fmt.Println(değişiklik, "kayıt silindi. ID'si:", silid)
		vt.Close()
	case "ekle":
		ad := veri[1]
		pp := veri[2]
		hk := veri[3]
		gl := veri[4]
		vt, _ := sql.Open("sqlite3", "./veritabanı.db")
		işlem, _ := vt.Prepare("INSERT INTO kisiler(ad, resim, hakkinda, galeri) values(?, ?, ?, ?)")
		//Hangi bölüme eklenecekse yukarıda orayı belirtiyoruz
		hk = strings.ReplaceAll(hk, "\n", "<br/>")
		hk = strings.ReplaceAll(hk, "[", `<span class="belirgin">`)
		hk = strings.ReplaceAll(hk, "]", `</span>`)
		veri, _ := işlem.Exec(ad, pp, hk, gl) //Eklenecek değer
		id, _ := veri.LastInsertId()          //Son girişin id numarısını aldık
		fmt.Println("Eklenen Kaydın ID'si:", id, ".Kayıt ismi:", ad)
		vt.Close() //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
	case "resimekle":
		svgkod := veri[1]
		vt, _ := sql.Open("sqlite3", "./veritabanı.db")
		işlem, _ := vt.Prepare("INSERT INTO resimler(svg) values(?)")
		//Hangi bölüme eklenecekse yukarıda orayı belirtiyoruz
		veri, _ := işlem.Exec(svgkod) //Eklenecek değer
		id, _ := veri.LastInsertId()  //Son girişin id numarısını aldık
		fmt.Println("Eklenen Resmin ID'si:", id)
		vt.Close() //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
		numara := strconv.Itoa(int(id))
		bildirim := `onay("Eklenen resmin id'si: ` + numara + `")`
		w.Eval(bildirim)
	case "ara":
		var kelime = strings.ToLowerSpecial(unicode.TurkishCase, veri[1])
		fmt.Println("Aranan kelime:", kelime)
		vt, err := sql.Open("sqlite3", "./veritabanı.db")
		if err != nil {
			fmt.Println(err)
		}
		tablo, _ := vt.Query("SELECT * FROM kisiler")
		var id int
		var ad, pp, hk string
		for tablo.Next() {
			aktarma := tablo.Scan(&id, &ad, &pp, &hk)
			numara := strconv.Itoa(id)
			if aktarma == nil && (strings.Contains(numara, kelime) || strings.Contains(strings.ToLowerSpecial(unicode.TurkishCase, ad), kelime) || strings.Contains(strings.ToLowerSpecial(unicode.TurkishCase, hk), kelime)) {
				icerik := "listeolustur('" + strconv.Itoa(id) + "',`" + strings.ToUpperSpecial(unicode.TurkishCase, ad) + "`);"
				w.Eval(icerik)
			} else {

			}
		}
		tablo.Close()
		vt.Close()
	case "baslangic":
		var showver = "$('.versiyon').text('versiyon " + Versiyon + "')"
		w.Eval(showver)
		vt, err := sql.Open("sqlite3", "./veritabanı.db")
		if err != nil {
			fmt.Println(err)
		}
		tablo, _ := vt.Query("SELECT * FROM ayarlar")
		var ayar, deger string
		for tablo.Next() {
			aktarma := tablo.Scan(&ayar, &deger)
			if ayar == "ses" {
				var komut = "sesAyar('" + deger + "')"
				w.Eval(komut)
			}
			if aktarma == nil {

			}
		}
		tablo.Close()
		vt.Close()
	case "resimler":
		etiketres := strings.Split(veri[1], ",")
		if veri[1] == "" {
			var boskomut = "$('#resimalani').append('Hiç resim bulunmamaktadır!')"
			w.Eval(boskomut)
		}
		vt, err := sql.Open("sqlite3", "./veritabanı.db")
		if err != nil {
			fmt.Println(err)
		}
		tablo, _ := vt.Query("SELECT * FROM resimler")
		var resid int
		var resvg string
		for tablo.Next() {
			aktarma := tablo.Scan(&resid, &resvg)

			for i, d := range etiketres {
				if i == 0 {
				}
				if d == strconv.Itoa(resid) {
					var komut = "$('#resimalani').append(`" + resvg + "`)"
					w.Eval(komut)
				}

			}

			if aktarma == nil {

			}
		}
		tablo.Close()
		vt.Close()
	case "komutgir":
		komut := strings.Split(veri[1], " ")
		if komut[0] == "ses" && (komut[1] == "on" || komut[1] == "off") {
			vt, err := sql.Open("sqlite3", "./veritabanı.db")
			kontrol(err)
			işlem, err := vt.Prepare("update ayarlar set deger=? where ayar='ses'")
			kontrol(err)
			veri, err := işlem.Exec(komut[1])
			kontrol(err)
			değişiklik, err := veri.RowsAffected()
			kontrol(err)
			fmt.Println(değişiklik)
			var komut = "sesAyar('" + komut[1] + "')"
			w.Eval(komut)
			var çıktı string = "onay(\"Ses durumu değiştirildi\")"
			w.Eval(çıktı)
			vt.Close()
		} else if komut[0] == "sifre" {
			şifre := md5.New()
			io.WriteString(şifre, komut[1])
			vt, err := sql.Open("sqlite3", "./veritabanı.db")
			kontrol(err)
			işlem, err := vt.Prepare("update ayarlar set deger=? where ayar='şifre'")
			kontrol(err)
			veri, err := işlem.Exec(hex.EncodeToString(şifre.Sum(nil)))
			kontrol(err)
			değişiklik, err := veri.RowsAffected()
			kontrol(err)
			fmt.Println(değişiklik)
			var çıktı string = "onay(\"Şifre değiştirildi\")"
			w.Eval(çıktı)
			vt.Close()
		} else {
			w.Eval("uyari(\"Geçersiz komut girdiniz!\")")
		}
	}
}
