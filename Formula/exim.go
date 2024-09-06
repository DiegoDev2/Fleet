package main

// Exim - Complete replacement for sendmail
// Homepage: https://exim.org

import (
	"fmt"
	
	"os/exec"
)

func installExim() {
	// Método 1: Descargar y extraer .tar.gz
	exim_tar_url := "https://ftp.exim.org/pub/exim/exim4/exim-4.98.tar.xz"
	exim_cmd_tar := exec.Command("curl", "-L", exim_tar_url, "-o", "package.tar.gz")
	err := exim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exim_zip_url := "https://ftp.exim.org/pub/exim/exim4/exim-4.98.tar.xz"
	exim_cmd_zip := exec.Command("curl", "-L", exim_zip_url, "-o", "package.zip")
	err = exim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exim_bin_url := "https://ftp.exim.org/pub/exim/exim4/exim-4.98.tar.xz"
	exim_cmd_bin := exec.Command("curl", "-L", exim_bin_url, "-o", "binary.bin")
	err = exim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exim_src_url := "https://ftp.exim.org/pub/exim/exim4/exim-4.98.tar.xz"
	exim_cmd_src := exec.Command("curl", "-L", exim_src_url, "-o", "source.tar.gz")
	err = exim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exim_cmd_direct := exec.Command("./binary")
	err = exim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
