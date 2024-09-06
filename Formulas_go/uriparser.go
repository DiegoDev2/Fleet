package main

// Uriparser - URI parsing library (strictly RFC 3986 compliant)
// Homepage: https://uriparser.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installUriparser() {
	// Método 1: Descargar y extraer .tar.gz
	uriparser_tar_url := "https://github.com/uriparser/uriparser/releases/download/uriparser-0.9.8/uriparser-0.9.8.tar.bz2"
	uriparser_cmd_tar := exec.Command("curl", "-L", uriparser_tar_url, "-o", "package.tar.gz")
	err := uriparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uriparser_zip_url := "https://github.com/uriparser/uriparser/releases/download/uriparser-0.9.8/uriparser-0.9.8.tar.bz2"
	uriparser_cmd_zip := exec.Command("curl", "-L", uriparser_zip_url, "-o", "package.zip")
	err = uriparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uriparser_bin_url := "https://github.com/uriparser/uriparser/releases/download/uriparser-0.9.8/uriparser-0.9.8.tar.bz2"
	uriparser_cmd_bin := exec.Command("curl", "-L", uriparser_bin_url, "-o", "binary.bin")
	err = uriparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uriparser_src_url := "https://github.com/uriparser/uriparser/releases/download/uriparser-0.9.8/uriparser-0.9.8.tar.bz2"
	uriparser_cmd_src := exec.Command("curl", "-L", uriparser_src_url, "-o", "source.tar.gz")
	err = uriparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uriparser_cmd_direct := exec.Command("./binary")
	err = uriparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
