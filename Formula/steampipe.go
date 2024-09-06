package main

// Steampipe - Use SQL to instantly query your cloud services
// Homepage: https://steampipe.io/

import (
	"fmt"
	
	"os/exec"
)

func installSteampipe() {
	// Método 1: Descargar y extraer .tar.gz
	steampipe_tar_url := "https://github.com/turbot/steampipe/archive/refs/tags/v0.24.0.tar.gz"
	steampipe_cmd_tar := exec.Command("curl", "-L", steampipe_tar_url, "-o", "package.tar.gz")
	err := steampipe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	steampipe_zip_url := "https://github.com/turbot/steampipe/archive/refs/tags/v0.24.0.zip"
	steampipe_cmd_zip := exec.Command("curl", "-L", steampipe_zip_url, "-o", "package.zip")
	err = steampipe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	steampipe_bin_url := "https://github.com/turbot/steampipe/archive/refs/tags/v0.24.0.bin"
	steampipe_cmd_bin := exec.Command("curl", "-L", steampipe_bin_url, "-o", "binary.bin")
	err = steampipe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	steampipe_src_url := "https://github.com/turbot/steampipe/archive/refs/tags/v0.24.0.src.tar.gz"
	steampipe_cmd_src := exec.Command("curl", "-L", steampipe_src_url, "-o", "source.tar.gz")
	err = steampipe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	steampipe_cmd_direct := exec.Command("./binary")
	err = steampipe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
