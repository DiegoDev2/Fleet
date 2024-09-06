package main

// Spidermonkey - JavaScript-C Engine
// Homepage: https://spidermonkey.dev

import (
	"fmt"
	
	"os/exec"
)

func installSpidermonkey() {
	// Método 1: Descargar y extraer .tar.gz
	spidermonkey_tar_url := "https://archive.mozilla.org/pub/firefox/releases/128.2.0esr/source/firefox-128.2.0esr.source.tar.xz"
	spidermonkey_cmd_tar := exec.Command("curl", "-L", spidermonkey_tar_url, "-o", "package.tar.gz")
	err := spidermonkey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spidermonkey_zip_url := "https://archive.mozilla.org/pub/firefox/releases/128.2.0esr/source/firefox-128.2.0esr.source.tar.xz"
	spidermonkey_cmd_zip := exec.Command("curl", "-L", spidermonkey_zip_url, "-o", "package.zip")
	err = spidermonkey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spidermonkey_bin_url := "https://archive.mozilla.org/pub/firefox/releases/128.2.0esr/source/firefox-128.2.0esr.source.tar.xz"
	spidermonkey_cmd_bin := exec.Command("curl", "-L", spidermonkey_bin_url, "-o", "binary.bin")
	err = spidermonkey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spidermonkey_src_url := "https://archive.mozilla.org/pub/firefox/releases/128.2.0esr/source/firefox-128.2.0esr.source.tar.xz"
	spidermonkey_cmd_src := exec.Command("curl", "-L", spidermonkey_src_url, "-o", "source.tar.gz")
	err = spidermonkey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spidermonkey_cmd_direct := exec.Command("./binary")
	err = spidermonkey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cbindgen")
exec.Command("latte", "install", "cbindgen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: nspr")
exec.Command("latte", "install", "nspr")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
