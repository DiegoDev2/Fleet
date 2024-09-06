package main

// Hyperfine - Command-line benchmarking tool
// Homepage: https://github.com/sharkdp/hyperfine

import (
	"fmt"
	
	"os/exec"
)

func installHyperfine() {
	// Método 1: Descargar y extraer .tar.gz
	hyperfine_tar_url := "https://github.com/sharkdp/hyperfine/archive/refs/tags/v1.18.0.tar.gz"
	hyperfine_cmd_tar := exec.Command("curl", "-L", hyperfine_tar_url, "-o", "package.tar.gz")
	err := hyperfine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hyperfine_zip_url := "https://github.com/sharkdp/hyperfine/archive/refs/tags/v1.18.0.zip"
	hyperfine_cmd_zip := exec.Command("curl", "-L", hyperfine_zip_url, "-o", "package.zip")
	err = hyperfine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hyperfine_bin_url := "https://github.com/sharkdp/hyperfine/archive/refs/tags/v1.18.0.bin"
	hyperfine_cmd_bin := exec.Command("curl", "-L", hyperfine_bin_url, "-o", "binary.bin")
	err = hyperfine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hyperfine_src_url := "https://github.com/sharkdp/hyperfine/archive/refs/tags/v1.18.0.src.tar.gz"
	hyperfine_cmd_src := exec.Command("curl", "-L", hyperfine_src_url, "-o", "source.tar.gz")
	err = hyperfine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hyperfine_cmd_direct := exec.Command("./binary")
	err = hyperfine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
