package main

// Grizzly - Command-line tool for managing and automating Grafana dashboards
// Homepage: https://grafana.github.io/grizzly/

import (
	"fmt"
	
	"os/exec"
)

func installGrizzly() {
	// Método 1: Descargar y extraer .tar.gz
	grizzly_tar_url := "https://github.com/grafana/grizzly/archive/refs/tags/v0.4.7.tar.gz"
	grizzly_cmd_tar := exec.Command("curl", "-L", grizzly_tar_url, "-o", "package.tar.gz")
	err := grizzly_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grizzly_zip_url := "https://github.com/grafana/grizzly/archive/refs/tags/v0.4.7.zip"
	grizzly_cmd_zip := exec.Command("curl", "-L", grizzly_zip_url, "-o", "package.zip")
	err = grizzly_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grizzly_bin_url := "https://github.com/grafana/grizzly/archive/refs/tags/v0.4.7.bin"
	grizzly_cmd_bin := exec.Command("curl", "-L", grizzly_bin_url, "-o", "binary.bin")
	err = grizzly_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grizzly_src_url := "https://github.com/grafana/grizzly/archive/refs/tags/v0.4.7.src.tar.gz"
	grizzly_cmd_src := exec.Command("curl", "-L", grizzly_src_url, "-o", "source.tar.gz")
	err = grizzly_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grizzly_cmd_direct := exec.Command("./binary")
	err = grizzly_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
