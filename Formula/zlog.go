package main

// Zlog - High-performance C logging library
// Homepage: https://github.com/HardySimpson/zlog

import (
	"fmt"
	
	"os/exec"
)

func installZlog() {
	// Método 1: Descargar y extraer .tar.gz
	zlog_tar_url := "https://github.com/HardySimpson/zlog/archive/refs/tags/1.2.18.tar.gz"
	zlog_cmd_tar := exec.Command("curl", "-L", zlog_tar_url, "-o", "package.tar.gz")
	err := zlog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zlog_zip_url := "https://github.com/HardySimpson/zlog/archive/refs/tags/1.2.18.zip"
	zlog_cmd_zip := exec.Command("curl", "-L", zlog_zip_url, "-o", "package.zip")
	err = zlog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zlog_bin_url := "https://github.com/HardySimpson/zlog/archive/refs/tags/1.2.18.bin"
	zlog_cmd_bin := exec.Command("curl", "-L", zlog_bin_url, "-o", "binary.bin")
	err = zlog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zlog_src_url := "https://github.com/HardySimpson/zlog/archive/refs/tags/1.2.18.src.tar.gz"
	zlog_cmd_src := exec.Command("curl", "-L", zlog_src_url, "-o", "source.tar.gz")
	err = zlog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zlog_cmd_direct := exec.Command("./binary")
	err = zlog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
