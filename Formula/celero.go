package main

// Celero - C++ Benchmark Authoring Library/Framework
// Homepage: https://github.com/DigitalInBlue/Celero

import (
	"fmt"
	
	"os/exec"
)

func installCelero() {
	// Método 1: Descargar y extraer .tar.gz
	celero_tar_url := "https://github.com/DigitalInBlue/Celero/archive/refs/tags/v2.9.0.tar.gz"
	celero_cmd_tar := exec.Command("curl", "-L", celero_tar_url, "-o", "package.tar.gz")
	err := celero_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	celero_zip_url := "https://github.com/DigitalInBlue/Celero/archive/refs/tags/v2.9.0.zip"
	celero_cmd_zip := exec.Command("curl", "-L", celero_zip_url, "-o", "package.zip")
	err = celero_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	celero_bin_url := "https://github.com/DigitalInBlue/Celero/archive/refs/tags/v2.9.0.bin"
	celero_cmd_bin := exec.Command("curl", "-L", celero_bin_url, "-o", "binary.bin")
	err = celero_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	celero_src_url := "https://github.com/DigitalInBlue/Celero/archive/refs/tags/v2.9.0.src.tar.gz"
	celero_cmd_src := exec.Command("curl", "-L", celero_src_url, "-o", "source.tar.gz")
	err = celero_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	celero_cmd_direct := exec.Command("./binary")
	err = celero_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
