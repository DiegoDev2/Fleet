package main

// RapidfuzzCpp - Rapid fuzzy string matching in C++ using the Levenshtein Distance
// Homepage: https://github.com/rapidfuzz/rapidfuzz-cpp

import (
	"fmt"
	
	"os/exec"
)

func installRapidfuzzCpp() {
	// Método 1: Descargar y extraer .tar.gz
	rapidfuzzcpp_tar_url := "https://github.com/rapidfuzz/rapidfuzz-cpp/archive/refs/tags/v3.0.5.tar.gz"
	rapidfuzzcpp_cmd_tar := exec.Command("curl", "-L", rapidfuzzcpp_tar_url, "-o", "package.tar.gz")
	err := rapidfuzzcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rapidfuzzcpp_zip_url := "https://github.com/rapidfuzz/rapidfuzz-cpp/archive/refs/tags/v3.0.5.zip"
	rapidfuzzcpp_cmd_zip := exec.Command("curl", "-L", rapidfuzzcpp_zip_url, "-o", "package.zip")
	err = rapidfuzzcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rapidfuzzcpp_bin_url := "https://github.com/rapidfuzz/rapidfuzz-cpp/archive/refs/tags/v3.0.5.bin"
	rapidfuzzcpp_cmd_bin := exec.Command("curl", "-L", rapidfuzzcpp_bin_url, "-o", "binary.bin")
	err = rapidfuzzcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rapidfuzzcpp_src_url := "https://github.com/rapidfuzz/rapidfuzz-cpp/archive/refs/tags/v3.0.5.src.tar.gz"
	rapidfuzzcpp_cmd_src := exec.Command("curl", "-L", rapidfuzzcpp_src_url, "-o", "source.tar.gz")
	err = rapidfuzzcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rapidfuzzcpp_cmd_direct := exec.Command("./binary")
	err = rapidfuzzcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
