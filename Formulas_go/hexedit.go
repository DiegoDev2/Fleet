package main

// Hexedit - View and edit files in hexadecimal or ASCII
// Homepage: https://rigaux.org/hexedit.html

import (
	"fmt"
	
	"os/exec"
)

func installHexedit() {
	// Método 1: Descargar y extraer .tar.gz
	hexedit_tar_url := "https://github.com/pixel/hexedit/archive/refs/tags/1.6.tar.gz"
	hexedit_cmd_tar := exec.Command("curl", "-L", hexedit_tar_url, "-o", "package.tar.gz")
	err := hexedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hexedit_zip_url := "https://github.com/pixel/hexedit/archive/refs/tags/1.6.zip"
	hexedit_cmd_zip := exec.Command("curl", "-L", hexedit_zip_url, "-o", "package.zip")
	err = hexedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hexedit_bin_url := "https://github.com/pixel/hexedit/archive/refs/tags/1.6.bin"
	hexedit_cmd_bin := exec.Command("curl", "-L", hexedit_bin_url, "-o", "binary.bin")
	err = hexedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hexedit_src_url := "https://github.com/pixel/hexedit/archive/refs/tags/1.6.src.tar.gz"
	hexedit_cmd_src := exec.Command("curl", "-L", hexedit_src_url, "-o", "source.tar.gz")
	err = hexedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hexedit_cmd_direct := exec.Command("./binary")
	err = hexedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
