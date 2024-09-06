package main

// Tnef - Microsoft MS-TNEF attachment unpacker
// Homepage: https://github.com/verdammelt/tnef

import (
	"fmt"
	
	"os/exec"
)

func installTnef() {
	// Método 1: Descargar y extraer .tar.gz
	tnef_tar_url := "https://github.com/verdammelt/tnef/archive/refs/tags/1.4.18.tar.gz"
	tnef_cmd_tar := exec.Command("curl", "-L", tnef_tar_url, "-o", "package.tar.gz")
	err := tnef_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tnef_zip_url := "https://github.com/verdammelt/tnef/archive/refs/tags/1.4.18.zip"
	tnef_cmd_zip := exec.Command("curl", "-L", tnef_zip_url, "-o", "package.zip")
	err = tnef_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tnef_bin_url := "https://github.com/verdammelt/tnef/archive/refs/tags/1.4.18.bin"
	tnef_cmd_bin := exec.Command("curl", "-L", tnef_bin_url, "-o", "binary.bin")
	err = tnef_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tnef_src_url := "https://github.com/verdammelt/tnef/archive/refs/tags/1.4.18.src.tar.gz"
	tnef_cmd_src := exec.Command("curl", "-L", tnef_src_url, "-o", "source.tar.gz")
	err = tnef_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tnef_cmd_direct := exec.Command("./binary")
	err = tnef_cmd_direct.Run()
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
}
