package main

// Zoro - Expose local server to external network
// Homepage: https://github.com/txthinking/zoro

import (
	"fmt"
	
	"os/exec"
)

func installZoro() {
	// Método 1: Descargar y extraer .tar.gz
	zoro_tar_url := "https://github.com/txthinking/zoro/archive/refs/tags/v20240828.tar.gz"
	zoro_cmd_tar := exec.Command("curl", "-L", zoro_tar_url, "-o", "package.tar.gz")
	err := zoro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zoro_zip_url := "https://github.com/txthinking/zoro/archive/refs/tags/v20240828.zip"
	zoro_cmd_zip := exec.Command("curl", "-L", zoro_zip_url, "-o", "package.zip")
	err = zoro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zoro_bin_url := "https://github.com/txthinking/zoro/archive/refs/tags/v20240828.bin"
	zoro_cmd_bin := exec.Command("curl", "-L", zoro_bin_url, "-o", "binary.bin")
	err = zoro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zoro_src_url := "https://github.com/txthinking/zoro/archive/refs/tags/v20240828.src.tar.gz"
	zoro_cmd_src := exec.Command("curl", "-L", zoro_src_url, "-o", "source.tar.gz")
	err = zoro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zoro_cmd_direct := exec.Command("./binary")
	err = zoro_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
