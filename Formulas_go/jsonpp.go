package main

// Jsonpp - Command-line JSON pretty-printer
// Homepage: https://jmhodges.github.io/jsonpp/

import (
	"fmt"
	
	"os/exec"
)

func installJsonpp() {
	// Método 1: Descargar y extraer .tar.gz
	jsonpp_tar_url := "https://github.com/jmhodges/jsonpp/archive/refs/tags/1.3.0.tar.gz"
	jsonpp_cmd_tar := exec.Command("curl", "-L", jsonpp_tar_url, "-o", "package.tar.gz")
	err := jsonpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonpp_zip_url := "https://github.com/jmhodges/jsonpp/archive/refs/tags/1.3.0.zip"
	jsonpp_cmd_zip := exec.Command("curl", "-L", jsonpp_zip_url, "-o", "package.zip")
	err = jsonpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonpp_bin_url := "https://github.com/jmhodges/jsonpp/archive/refs/tags/1.3.0.bin"
	jsonpp_cmd_bin := exec.Command("curl", "-L", jsonpp_bin_url, "-o", "binary.bin")
	err = jsonpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonpp_src_url := "https://github.com/jmhodges/jsonpp/archive/refs/tags/1.3.0.src.tar.gz"
	jsonpp_cmd_src := exec.Command("curl", "-L", jsonpp_src_url, "-o", "source.tar.gz")
	err = jsonpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonpp_cmd_direct := exec.Command("./binary")
	err = jsonpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
