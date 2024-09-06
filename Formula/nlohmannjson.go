package main

// NlohmannJson - JSON for modern C++
// Homepage: https://github.com/nlohmann/json

import (
	"fmt"
	
	"os/exec"
)

func installNlohmannJson() {
	// Método 1: Descargar y extraer .tar.gz
	nlohmannjson_tar_url := "https://github.com/nlohmann/json/archive/refs/tags/v3.11.3.tar.gz"
	nlohmannjson_cmd_tar := exec.Command("curl", "-L", nlohmannjson_tar_url, "-o", "package.tar.gz")
	err := nlohmannjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nlohmannjson_zip_url := "https://github.com/nlohmann/json/archive/refs/tags/v3.11.3.zip"
	nlohmannjson_cmd_zip := exec.Command("curl", "-L", nlohmannjson_zip_url, "-o", "package.zip")
	err = nlohmannjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nlohmannjson_bin_url := "https://github.com/nlohmann/json/archive/refs/tags/v3.11.3.bin"
	nlohmannjson_cmd_bin := exec.Command("curl", "-L", nlohmannjson_bin_url, "-o", "binary.bin")
	err = nlohmannjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nlohmannjson_src_url := "https://github.com/nlohmann/json/archive/refs/tags/v3.11.3.src.tar.gz"
	nlohmannjson_cmd_src := exec.Command("curl", "-L", nlohmannjson_src_url, "-o", "source.tar.gz")
	err = nlohmannjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nlohmannjson_cmd_direct := exec.Command("./binary")
	err = nlohmannjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
