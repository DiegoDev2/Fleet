package main

// Groonga - Fulltext search engine and column store
// Homepage: https://groonga.org/

import (
	"fmt"
	
	"os/exec"
)

func installGroonga() {
	// Método 1: Descargar y extraer .tar.gz
	groonga_tar_url := "https://github.com/groonga/groonga/releases/download/v14.0.7/groonga-14.0.7.tar.gz"
	groonga_cmd_tar := exec.Command("curl", "-L", groonga_tar_url, "-o", "package.tar.gz")
	err := groonga_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	groonga_zip_url := "https://github.com/groonga/groonga/releases/download/v14.0.7/groonga-14.0.7.zip"
	groonga_cmd_zip := exec.Command("curl", "-L", groonga_zip_url, "-o", "package.zip")
	err = groonga_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	groonga_bin_url := "https://github.com/groonga/groonga/releases/download/v14.0.7/groonga-14.0.7.bin"
	groonga_cmd_bin := exec.Command("curl", "-L", groonga_bin_url, "-o", "binary.bin")
	err = groonga_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	groonga_src_url := "https://github.com/groonga/groonga/releases/download/v14.0.7/groonga-14.0.7.src.tar.gz"
	groonga_cmd_src := exec.Command("curl", "-L", groonga_src_url, "-o", "source.tar.gz")
	err = groonga_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	groonga_cmd_direct := exec.Command("./binary")
	err = groonga_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: mecab")
	exec.Command("latte", "install", "mecab").Run()
	fmt.Println("Instalando dependencia: mecab-ipadic")
	exec.Command("latte", "install", "mecab-ipadic").Run()
	fmt.Println("Instalando dependencia: msgpack")
	exec.Command("latte", "install", "msgpack").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
