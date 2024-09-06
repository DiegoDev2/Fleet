package main

// HttpParser - HTTP request/response parser for c
// Homepage: https://github.com/nodejs/http-parser

import (
	"fmt"
	
	"os/exec"
)

func installHttpParser() {
	// Método 1: Descargar y extraer .tar.gz
	httpparser_tar_url := "https://github.com/nodejs/http-parser/archive/refs/tags/v2.9.4.tar.gz"
	httpparser_cmd_tar := exec.Command("curl", "-L", httpparser_tar_url, "-o", "package.tar.gz")
	err := httpparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpparser_zip_url := "https://github.com/nodejs/http-parser/archive/refs/tags/v2.9.4.zip"
	httpparser_cmd_zip := exec.Command("curl", "-L", httpparser_zip_url, "-o", "package.zip")
	err = httpparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpparser_bin_url := "https://github.com/nodejs/http-parser/archive/refs/tags/v2.9.4.bin"
	httpparser_cmd_bin := exec.Command("curl", "-L", httpparser_bin_url, "-o", "binary.bin")
	err = httpparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpparser_src_url := "https://github.com/nodejs/http-parser/archive/refs/tags/v2.9.4.src.tar.gz"
	httpparser_cmd_src := exec.Command("curl", "-L", httpparser_src_url, "-o", "source.tar.gz")
	err = httpparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpparser_cmd_direct := exec.Command("./binary")
	err = httpparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
}
