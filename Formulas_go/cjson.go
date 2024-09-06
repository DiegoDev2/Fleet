package main

// Cjson - Ultralightweight JSON parser in ANSI C
// Homepage: https://github.com/DaveGamble/cJSON

import (
	"fmt"
	
	"os/exec"
)

func installCjson() {
	// Método 1: Descargar y extraer .tar.gz
	cjson_tar_url := "https://github.com/DaveGamble/cJSON/archive/refs/tags/v1.7.18.tar.gz"
	cjson_cmd_tar := exec.Command("curl", "-L", cjson_tar_url, "-o", "package.tar.gz")
	err := cjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cjson_zip_url := "https://github.com/DaveGamble/cJSON/archive/refs/tags/v1.7.18.zip"
	cjson_cmd_zip := exec.Command("curl", "-L", cjson_zip_url, "-o", "package.zip")
	err = cjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cjson_bin_url := "https://github.com/DaveGamble/cJSON/archive/refs/tags/v1.7.18.bin"
	cjson_cmd_bin := exec.Command("curl", "-L", cjson_bin_url, "-o", "binary.bin")
	err = cjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cjson_src_url := "https://github.com/DaveGamble/cJSON/archive/refs/tags/v1.7.18.src.tar.gz"
	cjson_cmd_src := exec.Command("curl", "-L", cjson_src_url, "-o", "source.tar.gz")
	err = cjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cjson_cmd_direct := exec.Command("./binary")
	err = cjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
