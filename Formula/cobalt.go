package main

// Cobalt - Static site generator written in Rust
// Homepage: https://cobalt-org.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCobalt() {
	// Método 1: Descargar y extraer .tar.gz
	cobalt_tar_url := "https://github.com/cobalt-org/cobalt.rs/archive/refs/tags/v0.19.6.tar.gz"
	cobalt_cmd_tar := exec.Command("curl", "-L", cobalt_tar_url, "-o", "package.tar.gz")
	err := cobalt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cobalt_zip_url := "https://github.com/cobalt-org/cobalt.rs/archive/refs/tags/v0.19.6.zip"
	cobalt_cmd_zip := exec.Command("curl", "-L", cobalt_zip_url, "-o", "package.zip")
	err = cobalt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cobalt_bin_url := "https://github.com/cobalt-org/cobalt.rs/archive/refs/tags/v0.19.6.bin"
	cobalt_cmd_bin := exec.Command("curl", "-L", cobalt_bin_url, "-o", "binary.bin")
	err = cobalt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cobalt_src_url := "https://github.com/cobalt-org/cobalt.rs/archive/refs/tags/v0.19.6.src.tar.gz"
	cobalt_cmd_src := exec.Command("curl", "-L", cobalt_src_url, "-o", "source.tar.gz")
	err = cobalt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cobalt_cmd_direct := exec.Command("./binary")
	err = cobalt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
