package main

// Funcoeszz - Dozens of command-line mini-applications (Portuguese)
// Homepage: https://funcoeszz.net/

import (
	"fmt"
	
	"os/exec"
)

func installFuncoeszz() {
	// Método 1: Descargar y extraer .tar.gz
	funcoeszz_tar_url := "https://funcoeszz.net/download/funcoeszz-21.1.sh"
	funcoeszz_cmd_tar := exec.Command("curl", "-L", funcoeszz_tar_url, "-o", "package.tar.gz")
	err := funcoeszz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	funcoeszz_zip_url := "https://funcoeszz.net/download/funcoeszz-21.1.sh"
	funcoeszz_cmd_zip := exec.Command("curl", "-L", funcoeszz_zip_url, "-o", "package.zip")
	err = funcoeszz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	funcoeszz_bin_url := "https://funcoeszz.net/download/funcoeszz-21.1.sh"
	funcoeszz_cmd_bin := exec.Command("curl", "-L", funcoeszz_bin_url, "-o", "binary.bin")
	err = funcoeszz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	funcoeszz_src_url := "https://funcoeszz.net/download/funcoeszz-21.1.sh"
	funcoeszz_cmd_src := exec.Command("curl", "-L", funcoeszz_src_url, "-o", "source.tar.gz")
	err = funcoeszz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	funcoeszz_cmd_direct := exec.Command("./binary")
	err = funcoeszz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
}
