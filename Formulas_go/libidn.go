package main

// Libidn - International domain name library
// Homepage: https://www.gnu.org/software/libidn/

import (
	"fmt"
	
	"os/exec"
)

func installLibidn() {
	// Método 1: Descargar y extraer .tar.gz
	libidn_tar_url := "https://ftp.gnu.org/gnu/libidn/libidn-1.42.tar.gz"
	libidn_cmd_tar := exec.Command("curl", "-L", libidn_tar_url, "-o", "package.tar.gz")
	err := libidn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libidn_zip_url := "https://ftp.gnu.org/gnu/libidn/libidn-1.42.zip"
	libidn_cmd_zip := exec.Command("curl", "-L", libidn_zip_url, "-o", "package.zip")
	err = libidn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libidn_bin_url := "https://ftp.gnu.org/gnu/libidn/libidn-1.42.bin"
	libidn_cmd_bin := exec.Command("curl", "-L", libidn_bin_url, "-o", "binary.bin")
	err = libidn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libidn_src_url := "https://ftp.gnu.org/gnu/libidn/libidn-1.42.src.tar.gz"
	libidn_cmd_src := exec.Command("curl", "-L", libidn_src_url, "-o", "source.tar.gz")
	err = libidn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libidn_cmd_direct := exec.Command("./binary")
	err = libidn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
