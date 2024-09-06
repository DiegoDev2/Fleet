package main

// Cracklib - LibCrack password checking library
// Homepage: https://github.com/cracklib/cracklib

import (
	"fmt"
	
	"os/exec"
)

func installCracklib() {
	// Método 1: Descargar y extraer .tar.gz
	cracklib_tar_url := "https://github.com/cracklib/cracklib/releases/download/v2.10.2/cracklib-2.10.2.tar.bz2"
	cracklib_cmd_tar := exec.Command("curl", "-L", cracklib_tar_url, "-o", "package.tar.gz")
	err := cracklib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cracklib_zip_url := "https://github.com/cracklib/cracklib/releases/download/v2.10.2/cracklib-2.10.2.tar.bz2"
	cracklib_cmd_zip := exec.Command("curl", "-L", cracklib_zip_url, "-o", "package.zip")
	err = cracklib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cracklib_bin_url := "https://github.com/cracklib/cracklib/releases/download/v2.10.2/cracklib-2.10.2.tar.bz2"
	cracklib_cmd_bin := exec.Command("curl", "-L", cracklib_bin_url, "-o", "binary.bin")
	err = cracklib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cracklib_src_url := "https://github.com/cracklib/cracklib/releases/download/v2.10.2/cracklib-2.10.2.tar.bz2"
	cracklib_cmd_src := exec.Command("curl", "-L", cracklib_src_url, "-o", "source.tar.gz")
	err = cracklib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cracklib_cmd_direct := exec.Command("./binary")
	err = cracklib_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
