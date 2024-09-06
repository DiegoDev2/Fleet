package main

// Czkawka - Duplicate file utility
// Homepage: https://github.com/qarmin/czkawka

import (
	"fmt"
	
	"os/exec"
)

func installCzkawka() {
	// Método 1: Descargar y extraer .tar.gz
	czkawka_tar_url := "https://github.com/qarmin/czkawka/archive/refs/tags/7.0.0.tar.gz"
	czkawka_cmd_tar := exec.Command("curl", "-L", czkawka_tar_url, "-o", "package.tar.gz")
	err := czkawka_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	czkawka_zip_url := "https://github.com/qarmin/czkawka/archive/refs/tags/7.0.0.zip"
	czkawka_cmd_zip := exec.Command("curl", "-L", czkawka_zip_url, "-o", "package.zip")
	err = czkawka_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	czkawka_bin_url := "https://github.com/qarmin/czkawka/archive/refs/tags/7.0.0.bin"
	czkawka_cmd_bin := exec.Command("curl", "-L", czkawka_bin_url, "-o", "binary.bin")
	err = czkawka_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	czkawka_src_url := "https://github.com/qarmin/czkawka/archive/refs/tags/7.0.0.src.tar.gz"
	czkawka_cmd_src := exec.Command("curl", "-L", czkawka_src_url, "-o", "source.tar.gz")
	err = czkawka_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	czkawka_cmd_direct := exec.Command("./binary")
	err = czkawka_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: libheif")
exec.Command("latte", "install", "libheif")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: webp-pixbuf-loader")
exec.Command("latte", "install", "webp-pixbuf-loader")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: graphene")
exec.Command("latte", "install", "graphene")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
