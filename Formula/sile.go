package main

// Sile - Modern typesetting system inspired by TeX
// Homepage: https://sile-typesetter.org

import (
	"fmt"
	
	"os/exec"
)

func installSile() {
	// Método 1: Descargar y extraer .tar.gz
	sile_tar_url := "https://github.com/sile-typesetter/sile/releases/download/v0.15.5/sile-0.15.5.tar.zst"
	sile_cmd_tar := exec.Command("curl", "-L", sile_tar_url, "-o", "package.tar.gz")
	err := sile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sile_zip_url := "https://github.com/sile-typesetter/sile/releases/download/v0.15.5/sile-0.15.5.tar.zst"
	sile_cmd_zip := exec.Command("curl", "-L", sile_zip_url, "-o", "package.zip")
	err = sile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sile_bin_url := "https://github.com/sile-typesetter/sile/releases/download/v0.15.5/sile-0.15.5.tar.zst"
	sile_cmd_bin := exec.Command("curl", "-L", sile_bin_url, "-o", "binary.bin")
	err = sile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sile_src_url := "https://github.com/sile-typesetter/sile/releases/download/v0.15.5/sile-0.15.5.tar.zst"
	sile_cmd_src := exec.Command("curl", "-L", sile_src_url, "-o", "source.tar.gz")
	err = sile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sile_cmd_direct := exec.Command("./binary")
	err = sile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: luarocks")
	exec.Command("latte", "install", "luarocks").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
}
