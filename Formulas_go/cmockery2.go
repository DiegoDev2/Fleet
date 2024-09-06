package main

// Cmockery2 - Reviving cmockery unit test framework from Google
// Homepage: https://github.com/lpabon/cmockery2

import (
	"fmt"
	
	"os/exec"
)

func installCmockery2() {
	// Método 1: Descargar y extraer .tar.gz
	cmockery2_tar_url := "https://github.com/lpabon/cmockery2/archive/refs/tags/v1.3.9.tar.gz"
	cmockery2_cmd_tar := exec.Command("curl", "-L", cmockery2_tar_url, "-o", "package.tar.gz")
	err := cmockery2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmockery2_zip_url := "https://github.com/lpabon/cmockery2/archive/refs/tags/v1.3.9.zip"
	cmockery2_cmd_zip := exec.Command("curl", "-L", cmockery2_zip_url, "-o", "package.zip")
	err = cmockery2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmockery2_bin_url := "https://github.com/lpabon/cmockery2/archive/refs/tags/v1.3.9.bin"
	cmockery2_cmd_bin := exec.Command("curl", "-L", cmockery2_bin_url, "-o", "binary.bin")
	err = cmockery2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmockery2_src_url := "https://github.com/lpabon/cmockery2/archive/refs/tags/v1.3.9.src.tar.gz"
	cmockery2_cmd_src := exec.Command("curl", "-L", cmockery2_src_url, "-o", "source.tar.gz")
	err = cmockery2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmockery2_cmd_direct := exec.Command("./binary")
	err = cmockery2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
