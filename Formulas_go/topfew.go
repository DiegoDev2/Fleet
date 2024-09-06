package main

// Topfew - Finds the field values which appear most often in a stream of records
// Homepage: https://github.com/timbray/topfew

import (
	"fmt"
	
	"os/exec"
)

func installTopfew() {
	// Método 1: Descargar y extraer .tar.gz
	topfew_tar_url := "https://github.com/timbray/topfew/archive/refs/tags/v2.0.0.tar.gz"
	topfew_cmd_tar := exec.Command("curl", "-L", topfew_tar_url, "-o", "package.tar.gz")
	err := topfew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	topfew_zip_url := "https://github.com/timbray/topfew/archive/refs/tags/v2.0.0.zip"
	topfew_cmd_zip := exec.Command("curl", "-L", topfew_zip_url, "-o", "package.zip")
	err = topfew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	topfew_bin_url := "https://github.com/timbray/topfew/archive/refs/tags/v2.0.0.bin"
	topfew_cmd_bin := exec.Command("curl", "-L", topfew_bin_url, "-o", "binary.bin")
	err = topfew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	topfew_src_url := "https://github.com/timbray/topfew/archive/refs/tags/v2.0.0.src.tar.gz"
	topfew_cmd_src := exec.Command("curl", "-L", topfew_src_url, "-o", "source.tar.gz")
	err = topfew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	topfew_cmd_direct := exec.Command("./binary")
	err = topfew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
