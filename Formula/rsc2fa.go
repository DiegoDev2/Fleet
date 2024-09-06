package main

// Rsc2fa - Two-factor authentication on the command-line
// Homepage: https://pkg.go.dev/rsc.io/2fa

import (
	"fmt"
	
	"os/exec"
)

func installRsc2fa() {
	// Método 1: Descargar y extraer .tar.gz
	rsc2fa_tar_url := "https://github.com/rsc/2fa/archive/refs/tags/v1.2.0.tar.gz"
	rsc2fa_cmd_tar := exec.Command("curl", "-L", rsc2fa_tar_url, "-o", "package.tar.gz")
	err := rsc2fa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rsc2fa_zip_url := "https://github.com/rsc/2fa/archive/refs/tags/v1.2.0.zip"
	rsc2fa_cmd_zip := exec.Command("curl", "-L", rsc2fa_zip_url, "-o", "package.zip")
	err = rsc2fa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rsc2fa_bin_url := "https://github.com/rsc/2fa/archive/refs/tags/v1.2.0.bin"
	rsc2fa_cmd_bin := exec.Command("curl", "-L", rsc2fa_bin_url, "-o", "binary.bin")
	err = rsc2fa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rsc2fa_src_url := "https://github.com/rsc/2fa/archive/refs/tags/v1.2.0.src.tar.gz"
	rsc2fa_cmd_src := exec.Command("curl", "-L", rsc2fa_src_url, "-o", "source.tar.gz")
	err = rsc2fa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rsc2fa_cmd_direct := exec.Command("./binary")
	err = rsc2fa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
