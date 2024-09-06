package main

// Thanos - Highly available Prometheus setup with long term storage capabilities
// Homepage: https://thanos.io

import (
	"fmt"
	
	"os/exec"
)

func installThanos() {
	// Método 1: Descargar y extraer .tar.gz
	thanos_tar_url := "https://github.com/thanos-io/thanos/archive/refs/tags/v0.36.1.tar.gz"
	thanos_cmd_tar := exec.Command("curl", "-L", thanos_tar_url, "-o", "package.tar.gz")
	err := thanos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thanos_zip_url := "https://github.com/thanos-io/thanos/archive/refs/tags/v0.36.1.zip"
	thanos_cmd_zip := exec.Command("curl", "-L", thanos_zip_url, "-o", "package.zip")
	err = thanos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thanos_bin_url := "https://github.com/thanos-io/thanos/archive/refs/tags/v0.36.1.bin"
	thanos_cmd_bin := exec.Command("curl", "-L", thanos_bin_url, "-o", "binary.bin")
	err = thanos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thanos_src_url := "https://github.com/thanos-io/thanos/archive/refs/tags/v0.36.1.src.tar.gz"
	thanos_cmd_src := exec.Command("curl", "-L", thanos_src_url, "-o", "source.tar.gz")
	err = thanos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thanos_cmd_direct := exec.Command("./binary")
	err = thanos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
