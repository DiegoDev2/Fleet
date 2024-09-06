package main

// Mailcatcher - Catches mail and serves it through a dream
// Homepage: https://mailcatcher.me

import (
	"fmt"
	
	"os/exec"
)

func installMailcatcher() {
	// Método 1: Descargar y extraer .tar.gz
	mailcatcher_tar_url := "https://github.com/sj26/mailcatcher/archive/refs/tags/v0.10.0.tar.gz"
	mailcatcher_cmd_tar := exec.Command("curl", "-L", mailcatcher_tar_url, "-o", "package.tar.gz")
	err := mailcatcher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mailcatcher_zip_url := "https://github.com/sj26/mailcatcher/archive/refs/tags/v0.10.0.zip"
	mailcatcher_cmd_zip := exec.Command("curl", "-L", mailcatcher_zip_url, "-o", "package.zip")
	err = mailcatcher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mailcatcher_bin_url := "https://github.com/sj26/mailcatcher/archive/refs/tags/v0.10.0.bin"
	mailcatcher_cmd_bin := exec.Command("curl", "-L", mailcatcher_bin_url, "-o", "binary.bin")
	err = mailcatcher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mailcatcher_src_url := "https://github.com/sj26/mailcatcher/archive/refs/tags/v0.10.0.src.tar.gz"
	mailcatcher_cmd_src := exec.Command("curl", "-L", mailcatcher_src_url, "-o", "source.tar.gz")
	err = mailcatcher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mailcatcher_cmd_direct := exec.Command("./binary")
	err = mailcatcher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libedit")
	exec.Command("latte", "install", "libedit").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
