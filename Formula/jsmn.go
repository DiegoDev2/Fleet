package main

// Jsmn - World fastest JSON parser/tokenizer
// Homepage: https://zserge.com/jsmn/

import (
	"fmt"
	
	"os/exec"
)

func installJsmn() {
	// Método 1: Descargar y extraer .tar.gz
	jsmn_tar_url := "https://github.com/zserge/jsmn/archive/refs/tags/v1.1.0.tar.gz"
	jsmn_cmd_tar := exec.Command("curl", "-L", jsmn_tar_url, "-o", "package.tar.gz")
	err := jsmn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsmn_zip_url := "https://github.com/zserge/jsmn/archive/refs/tags/v1.1.0.zip"
	jsmn_cmd_zip := exec.Command("curl", "-L", jsmn_zip_url, "-o", "package.zip")
	err = jsmn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsmn_bin_url := "https://github.com/zserge/jsmn/archive/refs/tags/v1.1.0.bin"
	jsmn_cmd_bin := exec.Command("curl", "-L", jsmn_bin_url, "-o", "binary.bin")
	err = jsmn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsmn_src_url := "https://github.com/zserge/jsmn/archive/refs/tags/v1.1.0.src.tar.gz"
	jsmn_cmd_src := exec.Command("curl", "-L", jsmn_src_url, "-o", "source.tar.gz")
	err = jsmn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsmn_cmd_direct := exec.Command("./binary")
	err = jsmn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
