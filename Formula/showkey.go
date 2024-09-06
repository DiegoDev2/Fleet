package main

// Showkey - Simple keystroke visualizer
// Homepage: http://www.catb.org/~esr/showkey/

import (
	"fmt"
	
	"os/exec"
)

func installShowkey() {
	// Método 1: Descargar y extraer .tar.gz
	showkey_tar_url := "http://www.catb.org/~esr/showkey/showkey-1.9.tar.gz"
	showkey_cmd_tar := exec.Command("curl", "-L", showkey_tar_url, "-o", "package.tar.gz")
	err := showkey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	showkey_zip_url := "http://www.catb.org/~esr/showkey/showkey-1.9.zip"
	showkey_cmd_zip := exec.Command("curl", "-L", showkey_zip_url, "-o", "package.zip")
	err = showkey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	showkey_bin_url := "http://www.catb.org/~esr/showkey/showkey-1.9.bin"
	showkey_cmd_bin := exec.Command("curl", "-L", showkey_bin_url, "-o", "binary.bin")
	err = showkey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	showkey_src_url := "http://www.catb.org/~esr/showkey/showkey-1.9.src.tar.gz"
	showkey_cmd_src := exec.Command("curl", "-L", showkey_src_url, "-o", "source.tar.gz")
	err = showkey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	showkey_cmd_direct := exec.Command("./binary")
	err = showkey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
}
