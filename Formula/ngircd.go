package main

// Ngircd - Lightweight Internet Relay Chat server
// Homepage: https://ngircd.barton.de/

import (
	"fmt"
	
	"os/exec"
)

func installNgircd() {
	// Método 1: Descargar y extraer .tar.gz
	ngircd_tar_url := "https://ngircd.barton.de/pub/ngircd/ngircd-27.tar.xz"
	ngircd_cmd_tar := exec.Command("curl", "-L", ngircd_tar_url, "-o", "package.tar.gz")
	err := ngircd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ngircd_zip_url := "https://ngircd.barton.de/pub/ngircd/ngircd-27.tar.xz"
	ngircd_cmd_zip := exec.Command("curl", "-L", ngircd_zip_url, "-o", "package.zip")
	err = ngircd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ngircd_bin_url := "https://ngircd.barton.de/pub/ngircd/ngircd-27.tar.xz"
	ngircd_cmd_bin := exec.Command("curl", "-L", ngircd_bin_url, "-o", "binary.bin")
	err = ngircd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ngircd_src_url := "https://ngircd.barton.de/pub/ngircd/ngircd-27.tar.xz"
	ngircd_cmd_src := exec.Command("curl", "-L", ngircd_src_url, "-o", "source.tar.gz")
	err = ngircd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ngircd_cmd_direct := exec.Command("./binary")
	err = ngircd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libident")
	exec.Command("latte", "install", "libident").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
