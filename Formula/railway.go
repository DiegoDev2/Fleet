package main

// Railway - Develop and deploy code with zero configuration
// Homepage: https://railway.app/

import (
	"fmt"
	
	"os/exec"
)

func installRailway() {
	// Método 1: Descargar y extraer .tar.gz
	railway_tar_url := "https://github.com/railwayapp/cli/archive/refs/tags/v3.13.0.tar.gz"
	railway_cmd_tar := exec.Command("curl", "-L", railway_tar_url, "-o", "package.tar.gz")
	err := railway_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	railway_zip_url := "https://github.com/railwayapp/cli/archive/refs/tags/v3.13.0.zip"
	railway_cmd_zip := exec.Command("curl", "-L", railway_zip_url, "-o", "package.zip")
	err = railway_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	railway_bin_url := "https://github.com/railwayapp/cli/archive/refs/tags/v3.13.0.bin"
	railway_cmd_bin := exec.Command("curl", "-L", railway_bin_url, "-o", "binary.bin")
	err = railway_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	railway_src_url := "https://github.com/railwayapp/cli/archive/refs/tags/v3.13.0.src.tar.gz"
	railway_cmd_src := exec.Command("curl", "-L", railway_src_url, "-o", "source.tar.gz")
	err = railway_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	railway_cmd_direct := exec.Command("./binary")
	err = railway_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
