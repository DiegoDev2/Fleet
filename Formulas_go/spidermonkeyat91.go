package main

// SpidermonkeyAT91 - JavaScript-C Engine
// Homepage: https://spidermonkey.dev

import (
	"fmt"
	
	"os/exec"
)

func installSpidermonkeyAT91() {
	// Método 1: Descargar y extraer .tar.gz
	spidermonkeyat91_tar_url := "https://archive.mozilla.org/pub/firefox/releases/91.13.0esr/source/firefox-91.13.0esr.source.tar.xz"
	spidermonkeyat91_cmd_tar := exec.Command("curl", "-L", spidermonkeyat91_tar_url, "-o", "package.tar.gz")
	err := spidermonkeyat91_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spidermonkeyat91_zip_url := "https://archive.mozilla.org/pub/firefox/releases/91.13.0esr/source/firefox-91.13.0esr.source.tar.xz"
	spidermonkeyat91_cmd_zip := exec.Command("curl", "-L", spidermonkeyat91_zip_url, "-o", "package.zip")
	err = spidermonkeyat91_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spidermonkeyat91_bin_url := "https://archive.mozilla.org/pub/firefox/releases/91.13.0esr/source/firefox-91.13.0esr.source.tar.xz"
	spidermonkeyat91_cmd_bin := exec.Command("curl", "-L", spidermonkeyat91_bin_url, "-o", "binary.bin")
	err = spidermonkeyat91_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spidermonkeyat91_src_url := "https://archive.mozilla.org/pub/firefox/releases/91.13.0esr/source/firefox-91.13.0esr.source.tar.xz"
	spidermonkeyat91_cmd_src := exec.Command("curl", "-L", spidermonkeyat91_src_url, "-o", "source.tar.gz")
	err = spidermonkeyat91_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spidermonkeyat91_cmd_direct := exec.Command("./binary")
	err = spidermonkeyat91_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf@2.13")
exec.Command("latte", "install", "autoconf@2.13")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.9")
exec.Command("latte", "install", "python@3.9")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: nspr")
exec.Command("latte", "install", "nspr")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
