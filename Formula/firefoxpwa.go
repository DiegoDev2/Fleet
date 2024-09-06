package main

// Firefoxpwa - Tool to install, manage and use Progressive Web Apps in Mozilla Firefox
// Homepage: https://pwasforfirefox.filips.si/

import (
	"fmt"
	
	"os/exec"
)

func installFirefoxpwa() {
	// Método 1: Descargar y extraer .tar.gz
	firefoxpwa_tar_url := "https://github.com/filips123/PWAsForFirefox/archive/refs/tags/v2.12.1.tar.gz"
	firefoxpwa_cmd_tar := exec.Command("curl", "-L", firefoxpwa_tar_url, "-o", "package.tar.gz")
	err := firefoxpwa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	firefoxpwa_zip_url := "https://github.com/filips123/PWAsForFirefox/archive/refs/tags/v2.12.1.zip"
	firefoxpwa_cmd_zip := exec.Command("curl", "-L", firefoxpwa_zip_url, "-o", "package.zip")
	err = firefoxpwa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	firefoxpwa_bin_url := "https://github.com/filips123/PWAsForFirefox/archive/refs/tags/v2.12.1.bin"
	firefoxpwa_cmd_bin := exec.Command("curl", "-L", firefoxpwa_bin_url, "-o", "binary.bin")
	err = firefoxpwa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	firefoxpwa_src_url := "https://github.com/filips123/PWAsForFirefox/archive/refs/tags/v2.12.1.src.tar.gz"
	firefoxpwa_cmd_src := exec.Command("curl", "-L", firefoxpwa_src_url, "-o", "source.tar.gz")
	err = firefoxpwa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	firefoxpwa_cmd_direct := exec.Command("./binary")
	err = firefoxpwa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bzip2")
	exec.Command("latte", "install", "bzip2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
