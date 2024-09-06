package main

// Dropbear - Small SSH server/client for POSIX-based system
// Homepage: https://matt.ucc.asn.au/dropbear/dropbear.html

import (
	"fmt"
	
	"os/exec"
)

func installDropbear() {
	// Método 1: Descargar y extraer .tar.gz
	dropbear_tar_url := "https://matt.ucc.asn.au/dropbear/releases/dropbear-2024.85.tar.bz2"
	dropbear_cmd_tar := exec.Command("curl", "-L", dropbear_tar_url, "-o", "package.tar.gz")
	err := dropbear_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dropbear_zip_url := "https://matt.ucc.asn.au/dropbear/releases/dropbear-2024.85.tar.bz2"
	dropbear_cmd_zip := exec.Command("curl", "-L", dropbear_zip_url, "-o", "package.zip")
	err = dropbear_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dropbear_bin_url := "https://matt.ucc.asn.au/dropbear/releases/dropbear-2024.85.tar.bz2"
	dropbear_cmd_bin := exec.Command("curl", "-L", dropbear_bin_url, "-o", "binary.bin")
	err = dropbear_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dropbear_src_url := "https://matt.ucc.asn.au/dropbear/releases/dropbear-2024.85.tar.bz2"
	dropbear_cmd_src := exec.Command("curl", "-L", dropbear_src_url, "-o", "source.tar.gz")
	err = dropbear_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dropbear_cmd_direct := exec.Command("./binary")
	err = dropbear_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
