package main

// Pomsky - Regular expression language
// Homepage: https://pomsky-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installPomsky() {
	// Método 1: Descargar y extraer .tar.gz
	pomsky_tar_url := "https://github.com/rulex-rs/pomsky/archive/refs/tags/v0.11.tar.gz"
	pomsky_cmd_tar := exec.Command("curl", "-L", pomsky_tar_url, "-o", "package.tar.gz")
	err := pomsky_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pomsky_zip_url := "https://github.com/rulex-rs/pomsky/archive/refs/tags/v0.11.zip"
	pomsky_cmd_zip := exec.Command("curl", "-L", pomsky_zip_url, "-o", "package.zip")
	err = pomsky_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pomsky_bin_url := "https://github.com/rulex-rs/pomsky/archive/refs/tags/v0.11.bin"
	pomsky_cmd_bin := exec.Command("curl", "-L", pomsky_bin_url, "-o", "binary.bin")
	err = pomsky_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pomsky_src_url := "https://github.com/rulex-rs/pomsky/archive/refs/tags/v0.11.src.tar.gz"
	pomsky_cmd_src := exec.Command("curl", "-L", pomsky_src_url, "-o", "source.tar.gz")
	err = pomsky_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pomsky_cmd_direct := exec.Command("./binary")
	err = pomsky_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
