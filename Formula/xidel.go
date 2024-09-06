package main

// Xidel - XPath/XQuery 3.0, JSONiq interpreter to extract data from HTML/XML/JSON
// Homepage: https://www.videlibri.de/xidel.html

import (
	"fmt"
	
	"os/exec"
)

func installXidel() {
	// Método 1: Descargar y extraer .tar.gz
	xidel_tar_url := "https://github.com/benibela/xidel/releases/download/Xidel_0.9.8/xidel-0.9.8.src.tar.gz"
	xidel_cmd_tar := exec.Command("curl", "-L", xidel_tar_url, "-o", "package.tar.gz")
	err := xidel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xidel_zip_url := "https://github.com/benibela/xidel/releases/download/Xidel_0.9.8/xidel-0.9.8.src.zip"
	xidel_cmd_zip := exec.Command("curl", "-L", xidel_zip_url, "-o", "package.zip")
	err = xidel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xidel_bin_url := "https://github.com/benibela/xidel/releases/download/Xidel_0.9.8/xidel-0.9.8.src.bin"
	xidel_cmd_bin := exec.Command("curl", "-L", xidel_bin_url, "-o", "binary.bin")
	err = xidel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xidel_src_url := "https://github.com/benibela/xidel/releases/download/Xidel_0.9.8/xidel-0.9.8.src.src.tar.gz"
	xidel_cmd_src := exec.Command("curl", "-L", xidel_src_url, "-o", "source.tar.gz")
	err = xidel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xidel_cmd_direct := exec.Command("./binary")
	err = xidel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fpc")
	exec.Command("latte", "install", "fpc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
