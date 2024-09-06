package main

// Posh - Policy-compliant ordinary shell
// Homepage: https://salsa.debian.org/clint/posh

import (
	"fmt"
	
	"os/exec"
)

func installPosh() {
	// Método 1: Descargar y extraer .tar.gz
	posh_tar_url := "https://salsa.debian.org/clint/posh/-/archive/debian/0.14.1/posh-debian-0.14.1.tar.bz2"
	posh_cmd_tar := exec.Command("curl", "-L", posh_tar_url, "-o", "package.tar.gz")
	err := posh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	posh_zip_url := "https://salsa.debian.org/clint/posh/-/archive/debian/0.14.1/posh-debian-0.14.1.tar.bz2"
	posh_cmd_zip := exec.Command("curl", "-L", posh_zip_url, "-o", "package.zip")
	err = posh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	posh_bin_url := "https://salsa.debian.org/clint/posh/-/archive/debian/0.14.1/posh-debian-0.14.1.tar.bz2"
	posh_cmd_bin := exec.Command("curl", "-L", posh_bin_url, "-o", "binary.bin")
	err = posh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	posh_src_url := "https://salsa.debian.org/clint/posh/-/archive/debian/0.14.1/posh-debian-0.14.1.tar.bz2"
	posh_cmd_src := exec.Command("curl", "-L", posh_src_url, "-o", "source.tar.gz")
	err = posh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	posh_cmd_direct := exec.Command("./binary")
	err = posh_cmd_direct.Run()
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
