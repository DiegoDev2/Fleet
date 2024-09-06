package main

// Unp - Unpack everything with one command
// Homepage: https://packages.debian.org/source/stable/unp

import (
	"fmt"
	
	"os/exec"
)

func installUnp() {
	// Método 1: Descargar y extraer .tar.gz
	unp_tar_url := "https://deb.debian.org/debian/pool/main/u/unp/unp_2.0~pre10.tar.xz"
	unp_cmd_tar := exec.Command("curl", "-L", unp_tar_url, "-o", "package.tar.gz")
	err := unp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unp_zip_url := "https://deb.debian.org/debian/pool/main/u/unp/unp_2.0~pre10.tar.xz"
	unp_cmd_zip := exec.Command("curl", "-L", unp_zip_url, "-o", "package.zip")
	err = unp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unp_bin_url := "https://deb.debian.org/debian/pool/main/u/unp/unp_2.0~pre10.tar.xz"
	unp_cmd_bin := exec.Command("curl", "-L", unp_bin_url, "-o", "binary.bin")
	err = unp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unp_src_url := "https://deb.debian.org/debian/pool/main/u/unp/unp_2.0~pre10.tar.xz"
	unp_cmd_src := exec.Command("curl", "-L", unp_src_url, "-o", "source.tar.gz")
	err = unp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unp_cmd_direct := exec.Command("./binary")
	err = unp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: p7zip")
exec.Command("latte", "install", "p7zip")
}
