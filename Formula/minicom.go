package main

// Minicom - Menu-driven communications program
// Homepage: https://packages.debian.org/sid/minicom

import (
	"fmt"
	
	"os/exec"
)

func installMinicom() {
	// Método 1: Descargar y extraer .tar.gz
	minicom_tar_url := "https://deb.debian.org/debian/pool/main/m/minicom/minicom_2.9.orig.tar.bz2"
	minicom_cmd_tar := exec.Command("curl", "-L", minicom_tar_url, "-o", "package.tar.gz")
	err := minicom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minicom_zip_url := "https://deb.debian.org/debian/pool/main/m/minicom/minicom_2.9.orig.tar.bz2"
	minicom_cmd_zip := exec.Command("curl", "-L", minicom_zip_url, "-o", "package.zip")
	err = minicom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minicom_bin_url := "https://deb.debian.org/debian/pool/main/m/minicom/minicom_2.9.orig.tar.bz2"
	minicom_cmd_bin := exec.Command("curl", "-L", minicom_bin_url, "-o", "binary.bin")
	err = minicom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minicom_src_url := "https://deb.debian.org/debian/pool/main/m/minicom/minicom_2.9.orig.tar.bz2"
	minicom_cmd_src := exec.Command("curl", "-L", minicom_src_url, "-o", "source.tar.gz")
	err = minicom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minicom_cmd_direct := exec.Command("./binary")
	err = minicom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
