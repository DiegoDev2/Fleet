package main

// Dexter - Automatic indexer for Postgres
// Homepage: https://github.com/ankane/dexter

import (
	"fmt"
	
	"os/exec"
)

func installDexter() {
	// Método 1: Descargar y extraer .tar.gz
	dexter_tar_url := "https://github.com/ankane/dexter/archive/refs/tags/v0.5.5.tar.gz"
	dexter_cmd_tar := exec.Command("curl", "-L", dexter_tar_url, "-o", "package.tar.gz")
	err := dexter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dexter_zip_url := "https://github.com/ankane/dexter/archive/refs/tags/v0.5.5.zip"
	dexter_cmd_zip := exec.Command("curl", "-L", dexter_zip_url, "-o", "package.zip")
	err = dexter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dexter_bin_url := "https://github.com/ankane/dexter/archive/refs/tags/v0.5.5.bin"
	dexter_cmd_bin := exec.Command("curl", "-L", dexter_bin_url, "-o", "binary.bin")
	err = dexter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dexter_src_url := "https://github.com/ankane/dexter/archive/refs/tags/v0.5.5.src.tar.gz"
	dexter_cmd_src := exec.Command("curl", "-L", dexter_src_url, "-o", "source.tar.gz")
	err = dexter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dexter_cmd_direct := exec.Command("./binary")
	err = dexter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@16")
exec.Command("latte", "install", "postgresql@16")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
