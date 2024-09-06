package main

// Qjson - Map JSON to QVariant objects
// Homepage: https://qjson.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installQjson() {
	// Método 1: Descargar y extraer .tar.gz
	qjson_tar_url := "https://github.com/flavio/qjson/archive/refs/tags/0.9.0.tar.gz"
	qjson_cmd_tar := exec.Command("curl", "-L", qjson_tar_url, "-o", "package.tar.gz")
	err := qjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qjson_zip_url := "https://github.com/flavio/qjson/archive/refs/tags/0.9.0.zip"
	qjson_cmd_zip := exec.Command("curl", "-L", qjson_zip_url, "-o", "package.zip")
	err = qjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qjson_bin_url := "https://github.com/flavio/qjson/archive/refs/tags/0.9.0.bin"
	qjson_cmd_bin := exec.Command("curl", "-L", qjson_bin_url, "-o", "binary.bin")
	err = qjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qjson_src_url := "https://github.com/flavio/qjson/archive/refs/tags/0.9.0.src.tar.gz"
	qjson_cmd_src := exec.Command("curl", "-L", qjson_src_url, "-o", "source.tar.gz")
	err = qjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qjson_cmd_direct := exec.Command("./binary")
	err = qjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
}
