package main

// UniAlgo - Unicode Algorithms Implementation for C/C++
// Homepage: https://github.com/uni-algo/uni-algo

import (
	"fmt"
	
	"os/exec"
)

func installUniAlgo() {
	// Método 1: Descargar y extraer .tar.gz
	unialgo_tar_url := "https://github.com/uni-algo/uni-algo/archive/refs/tags/v1.2.0.tar.gz"
	unialgo_cmd_tar := exec.Command("curl", "-L", unialgo_tar_url, "-o", "package.tar.gz")
	err := unialgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unialgo_zip_url := "https://github.com/uni-algo/uni-algo/archive/refs/tags/v1.2.0.zip"
	unialgo_cmd_zip := exec.Command("curl", "-L", unialgo_zip_url, "-o", "package.zip")
	err = unialgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unialgo_bin_url := "https://github.com/uni-algo/uni-algo/archive/refs/tags/v1.2.0.bin"
	unialgo_cmd_bin := exec.Command("curl", "-L", unialgo_bin_url, "-o", "binary.bin")
	err = unialgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unialgo_src_url := "https://github.com/uni-algo/uni-algo/archive/refs/tags/v1.2.0.src.tar.gz"
	unialgo_cmd_src := exec.Command("curl", "-L", unialgo_src_url, "-o", "source.tar.gz")
	err = unialgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unialgo_cmd_direct := exec.Command("./binary")
	err = unialgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
