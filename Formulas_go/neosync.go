package main

// Neosync - CLI for interfacing with Neosync
// Homepage: https://www.neosync.dev/

import (
	"fmt"
	
	"os/exec"
)

func installNeosync() {
	// Método 1: Descargar y extraer .tar.gz
	neosync_tar_url := "https://github.com/nucleuscloud/neosync/archive/refs/tags/v0.4.62.tar.gz"
	neosync_cmd_tar := exec.Command("curl", "-L", neosync_tar_url, "-o", "package.tar.gz")
	err := neosync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neosync_zip_url := "https://github.com/nucleuscloud/neosync/archive/refs/tags/v0.4.62.zip"
	neosync_cmd_zip := exec.Command("curl", "-L", neosync_zip_url, "-o", "package.zip")
	err = neosync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neosync_bin_url := "https://github.com/nucleuscloud/neosync/archive/refs/tags/v0.4.62.bin"
	neosync_cmd_bin := exec.Command("curl", "-L", neosync_bin_url, "-o", "binary.bin")
	err = neosync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neosync_src_url := "https://github.com/nucleuscloud/neosync/archive/refs/tags/v0.4.62.src.tar.gz"
	neosync_cmd_src := exec.Command("curl", "-L", neosync_src_url, "-o", "source.tar.gz")
	err = neosync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neosync_cmd_direct := exec.Command("./binary")
	err = neosync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
