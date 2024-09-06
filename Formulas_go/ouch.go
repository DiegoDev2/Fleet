package main

// Ouch - Painless compression and decompression for your terminal
// Homepage: https://github.com/ouch-org/ouch

import (
	"fmt"
	
	"os/exec"
)

func installOuch() {
	// Método 1: Descargar y extraer .tar.gz
	ouch_tar_url := "https://github.com/ouch-org/ouch/archive/refs/tags/0.5.1.tar.gz"
	ouch_cmd_tar := exec.Command("curl", "-L", ouch_tar_url, "-o", "package.tar.gz")
	err := ouch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ouch_zip_url := "https://github.com/ouch-org/ouch/archive/refs/tags/0.5.1.zip"
	ouch_cmd_zip := exec.Command("curl", "-L", ouch_zip_url, "-o", "package.zip")
	err = ouch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ouch_bin_url := "https://github.com/ouch-org/ouch/archive/refs/tags/0.5.1.bin"
	ouch_cmd_bin := exec.Command("curl", "-L", ouch_bin_url, "-o", "binary.bin")
	err = ouch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ouch_src_url := "https://github.com/ouch-org/ouch/archive/refs/tags/0.5.1.src.tar.gz"
	ouch_cmd_src := exec.Command("curl", "-L", ouch_src_url, "-o", "source.tar.gz")
	err = ouch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ouch_cmd_direct := exec.Command("./binary")
	err = ouch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
