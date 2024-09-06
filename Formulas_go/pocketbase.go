package main

// Pocketbase - Open source backend for your next project in 1 file
// Homepage: https://pocketbase.io/

import (
	"fmt"
	
	"os/exec"
)

func installPocketbase() {
	// Método 1: Descargar y extraer .tar.gz
	pocketbase_tar_url := "https://github.com/pocketbase/pocketbase/archive/refs/tags/v0.22.20.tar.gz"
	pocketbase_cmd_tar := exec.Command("curl", "-L", pocketbase_tar_url, "-o", "package.tar.gz")
	err := pocketbase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pocketbase_zip_url := "https://github.com/pocketbase/pocketbase/archive/refs/tags/v0.22.20.zip"
	pocketbase_cmd_zip := exec.Command("curl", "-L", pocketbase_zip_url, "-o", "package.zip")
	err = pocketbase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pocketbase_bin_url := "https://github.com/pocketbase/pocketbase/archive/refs/tags/v0.22.20.bin"
	pocketbase_cmd_bin := exec.Command("curl", "-L", pocketbase_bin_url, "-o", "binary.bin")
	err = pocketbase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pocketbase_src_url := "https://github.com/pocketbase/pocketbase/archive/refs/tags/v0.22.20.src.tar.gz"
	pocketbase_cmd_src := exec.Command("curl", "-L", pocketbase_src_url, "-o", "source.tar.gz")
	err = pocketbase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pocketbase_cmd_direct := exec.Command("./binary")
	err = pocketbase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
