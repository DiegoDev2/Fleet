package main

// Zrepl - One-stop ZFS backup & replication solution
// Homepage: https://zrepl.github.io

import (
	"fmt"
	
	"os/exec"
)

func installZrepl() {
	// Método 1: Descargar y extraer .tar.gz
	zrepl_tar_url := "https://github.com/zrepl/zrepl/archive/refs/tags/v0.6.1.tar.gz"
	zrepl_cmd_tar := exec.Command("curl", "-L", zrepl_tar_url, "-o", "package.tar.gz")
	err := zrepl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zrepl_zip_url := "https://github.com/zrepl/zrepl/archive/refs/tags/v0.6.1.zip"
	zrepl_cmd_zip := exec.Command("curl", "-L", zrepl_zip_url, "-o", "package.zip")
	err = zrepl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zrepl_bin_url := "https://github.com/zrepl/zrepl/archive/refs/tags/v0.6.1.bin"
	zrepl_cmd_bin := exec.Command("curl", "-L", zrepl_bin_url, "-o", "binary.bin")
	err = zrepl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zrepl_src_url := "https://github.com/zrepl/zrepl/archive/refs/tags/v0.6.1.src.tar.gz"
	zrepl_cmd_src := exec.Command("curl", "-L", zrepl_src_url, "-o", "source.tar.gz")
	err = zrepl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zrepl_cmd_direct := exec.Command("./binary")
	err = zrepl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
