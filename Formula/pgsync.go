package main

// Pgsync - Sync Postgres data between databases
// Homepage: https://github.com/ankane/pgsync

import (
	"fmt"
	
	"os/exec"
)

func installPgsync() {
	// Método 1: Descargar y extraer .tar.gz
	pgsync_tar_url := "https://github.com/ankane/pgsync/archive/refs/tags/v0.8.0.tar.gz"
	pgsync_cmd_tar := exec.Command("curl", "-L", pgsync_tar_url, "-o", "package.tar.gz")
	err := pgsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgsync_zip_url := "https://github.com/ankane/pgsync/archive/refs/tags/v0.8.0.zip"
	pgsync_cmd_zip := exec.Command("curl", "-L", pgsync_zip_url, "-o", "package.zip")
	err = pgsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgsync_bin_url := "https://github.com/ankane/pgsync/archive/refs/tags/v0.8.0.bin"
	pgsync_cmd_bin := exec.Command("curl", "-L", pgsync_bin_url, "-o", "binary.bin")
	err = pgsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgsync_src_url := "https://github.com/ankane/pgsync/archive/refs/tags/v0.8.0.src.tar.gz"
	pgsync_cmd_src := exec.Command("curl", "-L", pgsync_src_url, "-o", "source.tar.gz")
	err = pgsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgsync_cmd_direct := exec.Command("./binary")
	err = pgsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
