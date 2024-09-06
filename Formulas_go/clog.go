package main

// Clog - Colorized pattern-matching log tail utility
// Homepage: https://taskwarrior.org/docs/clog/

import (
	"fmt"
	
	"os/exec"
)

func installClog() {
	// Método 1: Descargar y extraer .tar.gz
	clog_tar_url := "https://github.com/GothenburgBitFactory/clog/releases/download/v1.3.0/clog-1.3.0.tar.gz"
	clog_cmd_tar := exec.Command("curl", "-L", clog_tar_url, "-o", "package.tar.gz")
	err := clog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clog_zip_url := "https://github.com/GothenburgBitFactory/clog/releases/download/v1.3.0/clog-1.3.0.zip"
	clog_cmd_zip := exec.Command("curl", "-L", clog_zip_url, "-o", "package.zip")
	err = clog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clog_bin_url := "https://github.com/GothenburgBitFactory/clog/releases/download/v1.3.0/clog-1.3.0.bin"
	clog_cmd_bin := exec.Command("curl", "-L", clog_bin_url, "-o", "binary.bin")
	err = clog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clog_src_url := "https://github.com/GothenburgBitFactory/clog/releases/download/v1.3.0/clog-1.3.0.src.tar.gz"
	clog_cmd_src := exec.Command("curl", "-L", clog_src_url, "-o", "source.tar.gz")
	err = clog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clog_cmd_direct := exec.Command("./binary")
	err = clog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
