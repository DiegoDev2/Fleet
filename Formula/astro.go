package main

// Astro - To build and run Airflow DAGs locally and interact with the Astronomer API
// Homepage: https://www.astronomer.io/

import (
	"fmt"
	
	"os/exec"
)

func installAstro() {
	// Método 1: Descargar y extraer .tar.gz
	astro_tar_url := "https://github.com/astronomer/astro-cli/archive/refs/tags/v1.28.1.tar.gz"
	astro_cmd_tar := exec.Command("curl", "-L", astro_tar_url, "-o", "package.tar.gz")
	err := astro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	astro_zip_url := "https://github.com/astronomer/astro-cli/archive/refs/tags/v1.28.1.zip"
	astro_cmd_zip := exec.Command("curl", "-L", astro_zip_url, "-o", "package.zip")
	err = astro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	astro_bin_url := "https://github.com/astronomer/astro-cli/archive/refs/tags/v1.28.1.bin"
	astro_cmd_bin := exec.Command("curl", "-L", astro_bin_url, "-o", "binary.bin")
	err = astro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	astro_src_url := "https://github.com/astronomer/astro-cli/archive/refs/tags/v1.28.1.src.tar.gz"
	astro_cmd_src := exec.Command("curl", "-L", astro_src_url, "-o", "source.tar.gz")
	err = astro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	astro_cmd_direct := exec.Command("./binary")
	err = astro_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
