package main

// SafeRm - Wraps rm to prevent dangerous deletion of files
// Homepage: https://launchpad.net/safe-rm

import (
	"fmt"
	
	"os/exec"
)

func installSafeRm() {
	// Método 1: Descargar y extraer .tar.gz
	saferm_tar_url := "https://launchpad.net/safe-rm/trunk/1.1.0/+download/safe-rm-1.1.0.tar.gz"
	saferm_cmd_tar := exec.Command("curl", "-L", saferm_tar_url, "-o", "package.tar.gz")
	err := saferm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	saferm_zip_url := "https://launchpad.net/safe-rm/trunk/1.1.0/+download/safe-rm-1.1.0.zip"
	saferm_cmd_zip := exec.Command("curl", "-L", saferm_zip_url, "-o", "package.zip")
	err = saferm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	saferm_bin_url := "https://launchpad.net/safe-rm/trunk/1.1.0/+download/safe-rm-1.1.0.bin"
	saferm_cmd_bin := exec.Command("curl", "-L", saferm_bin_url, "-o", "binary.bin")
	err = saferm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	saferm_src_url := "https://launchpad.net/safe-rm/trunk/1.1.0/+download/safe-rm-1.1.0.src.tar.gz"
	saferm_cmd_src := exec.Command("curl", "-L", saferm_src_url, "-o", "source.tar.gz")
	err = saferm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	saferm_cmd_direct := exec.Command("./binary")
	err = saferm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
