package main

// Zf - Command-line fuzzy finder that prioritizes matches on filenames
// Homepage: https://github.com/natecraddock/zf

import (
	"fmt"
	
	"os/exec"
)

func installZf() {
	// Método 1: Descargar y extraer .tar.gz
	zf_tar_url := "https://github.com/natecraddock/zf/archive/refs/tags/0.9.2.tar.gz"
	zf_cmd_tar := exec.Command("curl", "-L", zf_tar_url, "-o", "package.tar.gz")
	err := zf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zf_zip_url := "https://github.com/natecraddock/zf/archive/refs/tags/0.9.2.zip"
	zf_cmd_zip := exec.Command("curl", "-L", zf_zip_url, "-o", "package.zip")
	err = zf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zf_bin_url := "https://github.com/natecraddock/zf/archive/refs/tags/0.9.2.bin"
	zf_cmd_bin := exec.Command("curl", "-L", zf_bin_url, "-o", "binary.bin")
	err = zf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zf_src_url := "https://github.com/natecraddock/zf/archive/refs/tags/0.9.2.src.tar.gz"
	zf_cmd_src := exec.Command("curl", "-L", zf_src_url, "-o", "source.tar.gz")
	err = zf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zf_cmd_direct := exec.Command("./binary")
	err = zf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zig")
	exec.Command("latte", "install", "zig").Run()
}
