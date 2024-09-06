package main

// Ephemeralpg - Run tests on an isolated, temporary Postgres database
// Homepage: https://eradman.com/ephemeralpg/

import (
	"fmt"
	
	"os/exec"
)

func installEphemeralpg() {
	// Método 1: Descargar y extraer .tar.gz
	ephemeralpg_tar_url := "https://eradman.com/ephemeralpg/code/ephemeralpg-3.3.tar.gz"
	ephemeralpg_cmd_tar := exec.Command("curl", "-L", ephemeralpg_tar_url, "-o", "package.tar.gz")
	err := ephemeralpg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ephemeralpg_zip_url := "https://eradman.com/ephemeralpg/code/ephemeralpg-3.3.zip"
	ephemeralpg_cmd_zip := exec.Command("curl", "-L", ephemeralpg_zip_url, "-o", "package.zip")
	err = ephemeralpg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ephemeralpg_bin_url := "https://eradman.com/ephemeralpg/code/ephemeralpg-3.3.bin"
	ephemeralpg_cmd_bin := exec.Command("curl", "-L", ephemeralpg_bin_url, "-o", "binary.bin")
	err = ephemeralpg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ephemeralpg_src_url := "https://eradman.com/ephemeralpg/code/ephemeralpg-3.3.src.tar.gz"
	ephemeralpg_cmd_src := exec.Command("curl", "-L", ephemeralpg_src_url, "-o", "source.tar.gz")
	err = ephemeralpg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ephemeralpg_cmd_direct := exec.Command("./binary")
	err = ephemeralpg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
