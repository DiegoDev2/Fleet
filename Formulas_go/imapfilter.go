package main

// Imapfilter - IMAP message processor/filter
// Homepage: https://github.com/lefcha/imapfilter/

import (
	"fmt"
	
	"os/exec"
)

func installImapfilter() {
	// Método 1: Descargar y extraer .tar.gz
	imapfilter_tar_url := "https://github.com/lefcha/imapfilter/archive/refs/tags/v2.8.2.tar.gz"
	imapfilter_cmd_tar := exec.Command("curl", "-L", imapfilter_tar_url, "-o", "package.tar.gz")
	err := imapfilter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imapfilter_zip_url := "https://github.com/lefcha/imapfilter/archive/refs/tags/v2.8.2.zip"
	imapfilter_cmd_zip := exec.Command("curl", "-L", imapfilter_zip_url, "-o", "package.zip")
	err = imapfilter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imapfilter_bin_url := "https://github.com/lefcha/imapfilter/archive/refs/tags/v2.8.2.bin"
	imapfilter_cmd_bin := exec.Command("curl", "-L", imapfilter_bin_url, "-o", "binary.bin")
	err = imapfilter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imapfilter_src_url := "https://github.com/lefcha/imapfilter/archive/refs/tags/v2.8.2.src.tar.gz"
	imapfilter_cmd_src := exec.Command("curl", "-L", imapfilter_src_url, "-o", "source.tar.gz")
	err = imapfilter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imapfilter_cmd_direct := exec.Command("./binary")
	err = imapfilter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
