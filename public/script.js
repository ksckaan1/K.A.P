//anahtar işlemi
var anahtar = 0;
$("#anahtar").click(function() {
  if (anahtar == 0) {
    anahtar = 1;
    $("#kilit").animate({ left: "30px" }, "fast");
    $("#yazialani").fadeOut(500);
    setTimeout(function() {
      $("#resimalani").fadeIn(500);
    }, 500);
  } else {
    anahtar = 0;
    $("#kilit").animate({ left: "0px" }, "fast");
    $("#resimalani").fadeOut(500);
    setTimeout(function() {
      $("#yazialani").fadeIn(500);
    }, 500);
  }
});

//LoadingBar İşlemi
$("button").click(function() {
  loadingbar();
});
$(".listekayit").click(function() {
  loadingbar();
});
function loadingbar() {
  $("#loadingBar").fadeToggle();
  $("#loadingBar").animate({ width: "800px" });
  setTimeout(function() {
    $("#loadingBar").fadeToggle();
    setTimeout(function() {
      $("#loadingBar").css("width", "0px");
    }, 500);
  }, 1500);
}

//Sayfa Yüklendiğinde Olacaklar
$(document).ready(function() {
  $("body").show();
  $("#sifreGiris").focus();
  //Listele
  window.external["invoke"]("listecek");

  //Başlangıçta çalışacaklar
  window.external["invoke"]("baslangic");

  //Başlangıçtaki typed efekti
  var options = {
    strings: [
      "                             ",
      "Kayıt Araştırma Programı",
      "Hoşgeldiniz",
      "by ksc10"
    ],
    typeSpeed: 70
  };
  var typed = new Typed("#launchtext", options);
});

//Giriş Yapılınca
function accessGranted() {
  loadingbar();
  $("#yuklenme").css("cursor", "wait");
  $("#sifreGiris").fadeOut();
  setTimeout(function() {
    $("#yuklenmeLogo").animate({ top: "-500px" });
  }, 1000);

  //Yüklenme Aşaması
  setTimeout(function() {
    $("#yuklenme").fadeOut();
  }, 1500);
  $("#aramakutusu").focus();
  var start = new Audio("./start.wav");
  ses("start");
}

//Şifre giriş işlevi
var sifreyaziliyor = 0;
$("#sifreGiris").focus(function() {
  sifreyaziliyor = 1;
});
$("#sifreGiris").blur(function() {
  sifreyaziliyor = 0;
});

//Çıkış Butonu Tıklandığında
$("#cikis").click(function() {
  window.external["invoke"]("cikis");
});

//Liste Butonuna Tıkladığında
$("#liste").click(function() {
  $("#kayitlistesi").html("");
  $("#sayfa").hide();
  $("#kayitekle").hide();
  $("#resimekle").hide();
  $("#kayitlar").fadeIn();
  window.external["invoke"]("listecek");
});

//Ekle Butonuna Tıkladığında
$("#ekle").click(function() {
  $("#sayfa").hide();
  $("#kayitlar").hide();
  $("#resimekle").hide();
  $("#kayitekle").fadeIn();
});

//Resim ekle butonuna basıltığında
$("#resimeklebtn").click(function() {
  $("#sayfa").hide();
  $("#kayitlar").hide();
  $("#kayitekle").hide();
  $("#resimekle").fadeIn();
});

//Resimekle İptale Basıldığında
$("#resimiptal").click(function() {
  $("#resimekle").hide();
  $("#resimkodu").val("");
  $("#kayitlar").fadeIn();
});

//Resimekle Kaydete Basıldığında
$("#resimkaydet").click(function() {
  var kod = $("#resimkodu").val();
  if (kod != "") {
    $("#resimekle").hide();
    var komut = "resimekle©" + kod;
    window.external["invoke"](komut);
    $("#resimkodu").val("");
    $("#kayitlar").fadeIn();
    $("#liste").click();
  } else {
    uyari("SVG Kodunu Boş Bıraktınız!");
  }
});

//Ekle İptale Basıldığında
$("#ekleiptal").click(function() {
  $("#kayitekle").hide();
  $("#ekleisim").val("");
  $("#ekleresim").val("");
  $("#eklehk").val("");
  $("#kayitlar").fadeIn();
});

//Ekle Kaydete Basıldığında
$("#eklekaydet").click(function() {
  var ad = $("#ekleisim").val();
  var pp = $("#ekleresim").val();
  var hk = $("#eklehk").val();
  var er = $("#ekleekresim").val();
  if (ad != "") {
    $("#kayitekle").hide();
    var komut = "ekle©" + ad + "©" + pp + "©" + hk + "©" + er;
    window.external["invoke"](komut);
    $("#ekleisim").val("");
    $("#ekleresim").val("");
    $("#ekleekresim").val("");
    $("#eklehk").val("");
    $("#kayitlar").fadeIn();
    $("#liste").click();
    onay("Kayıt Başarıyla Eklendi!");
  } else {
    uyari("Kayıt Adını Boş Bıraktınız!");
  }
});

//Profil Resmini Büyütme
$("#buyukresim").click(function() {
  $(this).slideToggle();
});
$("#profilresmi").click(function() {
  $("#buyukresim").slideToggle();
});

//Listeden Bir Kayda Tıkladığında
function listclick(id) {
  $("#sayfa").fadeIn();
  $("#kayitlar").hide();
  var lcid = "yazdir©" + id;
  window.external["invoke"](lcid);
}

//Bir Kaydı Profil Sayfasına Yazdırma
function yazdir(id, ad, yazi, profilresmi, galeri) {
  $("#profilid").text(id);
  $("#profilresmi").html(profilresmi);
  $("#buyukresim").html(profilresmi);
  $("#tspan10").text(ad);
  $("#yazialani").html(yazi);
  $("#duzenleisim").val(ad);
  var duzenlehk = yazi.replace(/<br\/>/g, "\n");
  duzenlehk = duzenlehk.replace(/<span class="belirgin">/g, "[");
  duzenlehk = duzenlehk.replace(/<\/span>/g, "]");
  $("#duzenlehk").val(duzenlehk);
  $("#duzenleresim").val(profilresmi);
  $("#duzenleekresim").val(galeri);
  $("#resimalani").html("");
  var glkomut = "resimler©" + galeri;
  window.external["invoke"](glkomut);
}

//Arama Kutusu İşlevi
var araniyor = 0;
$("#aramakutusu").focus(function() {
  araniyor = 1;
});
$("#aramakutusu").blur(function() {
  araniyor = 0;
});
var komutgiriliyor = 0;
$("#komutsatiri").focus(function() {
  komutgiriliyor = 1;
});
$("#komutsatiri").blur(function() {
  komutgiriliyor = 0;
  $("#komutcerceve").fadeOut()
});
// Yazı kutusu ENTER Tuşu Dinleme
window.onkeydown = function(olay) {
  if (olay.keyCode == 13 && araniyor == 1) {
    $("#aramayap").click();
  } else if (olay.keyCode == 13 && sifreyaziliyor == 1) {
    $("#girisOnay").click();
    var sifre = $("#sifreGiris").val();
    var komut = "giris©" + sifre;
    window.external["invoke"](komut);
  } else if (olay.keyCode == 222 && (araniyor == 0 && sifreyaziliyor == 0 && komutgiriliyor == 0)){
    olay.preventDefault()
    $("#komutcerceve").fadeIn();
    $("#komutsatiri").focus();
  } else if (olay.keyCode==13 && komutgiriliyor==1){
      var komutgirdi = $("#komutsatiri").val()
      var komut="komutgir©"+komutgirdi
      window.external["invoke"](komut)
      $("#komutsatiri").val("")
  }
};

//Arama Butonuna Tıklandığında
$("#aramayap").click(function() {
  var kelime = $("#aramakutusu").val();
  $("#kayitlistesi").html("");
  var komut = "ara©" + kelime;
  window.external["invoke"](komut);
});

//Arama Kutusunu Temizleme
$("#aramatemizle").click(function() {
  $("#liste").click();
  $("#aramakutusu").val("");
});

//Listeye Kayıtları Yazdırma
function listeolustur(id, ad) {
  var listeleman =
    "<div class='listekayit' onclick='listclick(" +
    id +
    ")' id='" +
    id +
    "'>" +
    ad +
    "</div>";
  $("#kayitlistesi").append(listeleman);
}

//Bir Kaydı Düzenleme
$("#duzenlebtn").click(function() {
  $("#duzenlepencere").fadeIn();
});

//Düzenlemeden Çıkmak
$("#duzenleiptal").click(function() {
  $("#duzenlepencere").fadeOut();
});

//Düzenleme Sıfırlama
$("#duzenlesifirla").click(function() {
  var id = $("#profilid").text();
  listclick(id);
});

//Kaydı Güncelleme - Değiştirme
$("#duzenlekaydet").click(function() {
  var id = $("#profilid").text();
  var ad = $("#duzenleisim").val();
  var pp = $("#duzenleresim").val();
  var hk = $("#duzenlehk").val();
  var er = $("#duzenleekresim").val();
  if (ad != "") {
    var komut = "guncelle©" + id + "©" + ad + "©" + pp + "©" + hk + "©" + er;
    window.external["invoke"](komut);
    var lcid = "yazdir©" + id;
    window.external["invoke"](lcid);
    $("#duzenlepencere").fadeOut();
    onay("Kayıt Başarıyla Güncellendi!");
  } else {
    uyari("Kayıt Adını Boş Bıraktınız!");
  }
});

//Bir Kaydı Silme
$("#silbtn").click(function() {
  var profilid = $("#profilid").text();
  var komut = "sil©" + profilid;
  window.external["invoke"](komut);
  $("#liste").click();
  onay("Kayıt Başarıyla Silindi!");
});

//Uyarı Kutucuğu
function uyari(uyari) {
  $("#uyari").text(uyari);
  $("#uyari").slideToggle();
  setTimeout(function() {
    $("#uyari").fadeOut();
  }, 3000);
  ses("uyari");
}

//Onay Kutucuğu
function onay(onay) {
  $("#onay").text(onay);
  $("#onay").slideToggle();
  setTimeout(function() {
    $("#onay").fadeOut();
  }, 3000);
  ses("onay");
}

var sesdurum = "on";
function sesAyar(durum) {
  sesdurum = durum;
}

function ses(isim) {
  if (sesdurum == "on") {
    if (isim == "start") {
      var dosya = new Audio("./start.wav");
      dosya.play();
    } else if (isim == "onay") {
      var dosya = new Audio("./onay.wav");
      dosya.play();
    } else if (isim == "uyari") {
      var dosya = new Audio("./uyari.wav");
      dosya.play();
    }
  }
}