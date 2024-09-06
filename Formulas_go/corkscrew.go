package main

// Corkscrew - Tunnel SSH through HTTP proxies
// Homepage: https://packages.debian.org/sid/corkscrew

import (
	"fmt"
	
	"os/exec"
)

func installCorkscrew() {
	// Método 1: Descargar y extraer .tar.gz
	corkscrew_tar_url := "https://deb.debian.org/debian/pool/main/c/corkscrew/corkscrew_2.0.orig.tar.gz"
	corkscrew_cmd_tar := exec.Command("curl", "-L", corkscrew_tar_url, "-o", "package.tar.gz")
	err := corkscrew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	corkscrew_zip_url := "https://deb.debian.org/debian/pool/main/c/corkscrew/corkscrew_2.0.orig.zip"
	corkscrew_cmd_zip := exec.Command("curl", "-L", corkscrew_zip_url, "-o", "package.zip")
	err = corkscrew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	corkscrew_bin_url := "https://deb.debian.org/debian/pool/main/c/corkscrew/corkscrew_2.0.orig.bin"
	corkscrew_cmd_bin := exec.Command("curl", "-L", corkscrew_bin_url, "-o", "binary.bin")
	err = corkscrew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	corkscrew_src_url := "https://deb.debian.org/debian/pool/main/c/corkscrew/corkscrew_2.0.orig.src.tar.gz"
	corkscrew_cmd_src := exec.Command("curl", "-L", corkscrew_src_url, "-o", "source.tar.gz")
	err = corkscrew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	corkscrew_cmd_direct := exec.Command("./binary")
	err = corkscrew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
