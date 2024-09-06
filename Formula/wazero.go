package main

// Wazero - Zero dependency WebAssembly runtime
// Homepage: https://wazero.io

import (
	"fmt"
	
	"os/exec"
)

func installWazero() {
	// Método 1: Descargar y extraer .tar.gz
	wazero_tar_url := "https://github.com/tetratelabs/wazero/archive/refs/tags/v1.8.0.tar.gz"
	wazero_cmd_tar := exec.Command("curl", "-L", wazero_tar_url, "-o", "package.tar.gz")
	err := wazero_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wazero_zip_url := "https://github.com/tetratelabs/wazero/archive/refs/tags/v1.8.0.zip"
	wazero_cmd_zip := exec.Command("curl", "-L", wazero_zip_url, "-o", "package.zip")
	err = wazero_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wazero_bin_url := "https://github.com/tetratelabs/wazero/archive/refs/tags/v1.8.0.bin"
	wazero_cmd_bin := exec.Command("curl", "-L", wazero_bin_url, "-o", "binary.bin")
	err = wazero_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wazero_src_url := "https://github.com/tetratelabs/wazero/archive/refs/tags/v1.8.0.src.tar.gz"
	wazero_cmd_src := exec.Command("curl", "-L", wazero_src_url, "-o", "source.tar.gz")
	err = wazero_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wazero_cmd_direct := exec.Command("./binary")
	err = wazero_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: wabt")
	exec.Command("latte", "install", "wabt").Run()
}
