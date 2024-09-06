package main

// Sampler - Tool for shell commands execution, visualization and alerting
// Homepage: https://sampler.dev

import (
	"fmt"
	
	"os/exec"
)

func installSampler() {
	// Método 1: Descargar y extraer .tar.gz
	sampler_tar_url := "https://github.com/sqshq/sampler/archive/refs/tags/v1.1.0.tar.gz"
	sampler_cmd_tar := exec.Command("curl", "-L", sampler_tar_url, "-o", "package.tar.gz")
	err := sampler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sampler_zip_url := "https://github.com/sqshq/sampler/archive/refs/tags/v1.1.0.zip"
	sampler_cmd_zip := exec.Command("curl", "-L", sampler_zip_url, "-o", "package.zip")
	err = sampler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sampler_bin_url := "https://github.com/sqshq/sampler/archive/refs/tags/v1.1.0.bin"
	sampler_cmd_bin := exec.Command("curl", "-L", sampler_bin_url, "-o", "binary.bin")
	err = sampler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sampler_src_url := "https://github.com/sqshq/sampler/archive/refs/tags/v1.1.0.src.tar.gz"
	sampler_cmd_src := exec.Command("curl", "-L", sampler_src_url, "-o", "source.tar.gz")
	err = sampler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sampler_cmd_direct := exec.Command("./binary")
	err = sampler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
