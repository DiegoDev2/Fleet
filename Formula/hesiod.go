package main

// Hesiod - Library for the simple string lookup service built on top of DNS
// Homepage: https://github.com/achernya/hesiod

import (
	"fmt"
	
	"os/exec"
)

func installHesiod() {
	// Método 1: Descargar y extraer .tar.gz
	hesiod_tar_url := "https://github.com/achernya/hesiod/archive/refs/tags/hesiod-3.2.1.tar.gz"
	hesiod_cmd_tar := exec.Command("curl", "-L", hesiod_tar_url, "-o", "package.tar.gz")
	err := hesiod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hesiod_zip_url := "https://github.com/achernya/hesiod/archive/refs/tags/hesiod-3.2.1.zip"
	hesiod_cmd_zip := exec.Command("curl", "-L", hesiod_zip_url, "-o", "package.zip")
	err = hesiod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hesiod_bin_url := "https://github.com/achernya/hesiod/archive/refs/tags/hesiod-3.2.1.bin"
	hesiod_cmd_bin := exec.Command("curl", "-L", hesiod_bin_url, "-o", "binary.bin")
	err = hesiod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hesiod_src_url := "https://github.com/achernya/hesiod/archive/refs/tags/hesiod-3.2.1.src.tar.gz"
	hesiod_cmd_src := exec.Command("curl", "-L", hesiod_src_url, "-o", "source.tar.gz")
	err = hesiod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hesiod_cmd_direct := exec.Command("./binary")
	err = hesiod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libidn")
	exec.Command("latte", "install", "libidn").Run()
}
