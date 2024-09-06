package main

// Grok - DRY and RAD for regular expressions and then some
// Homepage: https://github.com/jordansissel/grok

import (
	"fmt"
	
	"os/exec"
)

func installGrok() {
	// Método 1: Descargar y extraer .tar.gz
	grok_tar_url := "https://github.com/jordansissel/grok/archive/refs/tags/v0.9.2.tar.gz"
	grok_cmd_tar := exec.Command("curl", "-L", grok_tar_url, "-o", "package.tar.gz")
	err := grok_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grok_zip_url := "https://github.com/jordansissel/grok/archive/refs/tags/v0.9.2.zip"
	grok_cmd_zip := exec.Command("curl", "-L", grok_zip_url, "-o", "package.zip")
	err = grok_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grok_bin_url := "https://github.com/jordansissel/grok/archive/refs/tags/v0.9.2.bin"
	grok_cmd_bin := exec.Command("curl", "-L", grok_bin_url, "-o", "binary.bin")
	err = grok_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grok_src_url := "https://github.com/jordansissel/grok/archive/refs/tags/v0.9.2.src.tar.gz"
	grok_cmd_src := exec.Command("curl", "-L", grok_src_url, "-o", "source.tar.gz")
	err = grok_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grok_cmd_direct := exec.Command("./binary")
	err = grok_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: tokyo-cabinet")
exec.Command("latte", "install", "tokyo-cabinet")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
}
