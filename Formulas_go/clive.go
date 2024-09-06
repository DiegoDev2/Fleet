package main

// Clive - Automates terminal operations
// Homepage: https://github.com/koki-develop/clive

import (
	"fmt"
	
	"os/exec"
)

func installClive() {
	// Método 1: Descargar y extraer .tar.gz
	clive_tar_url := "https://github.com/koki-develop/clive/archive/refs/tags/v0.12.9.tar.gz"
	clive_cmd_tar := exec.Command("curl", "-L", clive_tar_url, "-o", "package.tar.gz")
	err := clive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clive_zip_url := "https://github.com/koki-develop/clive/archive/refs/tags/v0.12.9.zip"
	clive_cmd_zip := exec.Command("curl", "-L", clive_zip_url, "-o", "package.zip")
	err = clive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clive_bin_url := "https://github.com/koki-develop/clive/archive/refs/tags/v0.12.9.bin"
	clive_cmd_bin := exec.Command("curl", "-L", clive_bin_url, "-o", "binary.bin")
	err = clive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clive_src_url := "https://github.com/koki-develop/clive/archive/refs/tags/v0.12.9.src.tar.gz"
	clive_cmd_src := exec.Command("curl", "-L", clive_src_url, "-o", "source.tar.gz")
	err = clive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clive_cmd_direct := exec.Command("./binary")
	err = clive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: ttyd")
exec.Command("latte", "install", "ttyd")
}
