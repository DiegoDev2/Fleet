package main

// Skeema - Declarative pure-SQL schema management for MySQL and MariaDB
// Homepage: https://www.skeema.io/

import (
	"fmt"
	
	"os/exec"
)

func installSkeema() {
	// Método 1: Descargar y extraer .tar.gz
	skeema_tar_url := "https://github.com/skeema/skeema/archive/refs/tags/v1.12.0.tar.gz"
	skeema_cmd_tar := exec.Command("curl", "-L", skeema_tar_url, "-o", "package.tar.gz")
	err := skeema_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skeema_zip_url := "https://github.com/skeema/skeema/archive/refs/tags/v1.12.0.zip"
	skeema_cmd_zip := exec.Command("curl", "-L", skeema_zip_url, "-o", "package.zip")
	err = skeema_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skeema_bin_url := "https://github.com/skeema/skeema/archive/refs/tags/v1.12.0.bin"
	skeema_cmd_bin := exec.Command("curl", "-L", skeema_bin_url, "-o", "binary.bin")
	err = skeema_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skeema_src_url := "https://github.com/skeema/skeema/archive/refs/tags/v1.12.0.src.tar.gz"
	skeema_cmd_src := exec.Command("curl", "-L", skeema_src_url, "-o", "source.tar.gz")
	err = skeema_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skeema_cmd_direct := exec.Command("./binary")
	err = skeema_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
