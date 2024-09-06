package main

// Taplo - TOML toolkit written in Rust
// Homepage: https://taplo.tamasfe.dev

import (
	"fmt"
	
	"os/exec"
)

func installTaplo() {
	// Método 1: Descargar y extraer .tar.gz
	taplo_tar_url := "https://github.com/tamasfe/taplo/archive/refs/tags/0.9.3.tar.gz"
	taplo_cmd_tar := exec.Command("curl", "-L", taplo_tar_url, "-o", "package.tar.gz")
	err := taplo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taplo_zip_url := "https://github.com/tamasfe/taplo/archive/refs/tags/0.9.3.zip"
	taplo_cmd_zip := exec.Command("curl", "-L", taplo_zip_url, "-o", "package.zip")
	err = taplo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taplo_bin_url := "https://github.com/tamasfe/taplo/archive/refs/tags/0.9.3.bin"
	taplo_cmd_bin := exec.Command("curl", "-L", taplo_bin_url, "-o", "binary.bin")
	err = taplo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taplo_src_url := "https://github.com/tamasfe/taplo/archive/refs/tags/0.9.3.src.tar.gz"
	taplo_cmd_src := exec.Command("curl", "-L", taplo_src_url, "-o", "source.tar.gz")
	err = taplo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taplo_cmd_direct := exec.Command("./binary")
	err = taplo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
