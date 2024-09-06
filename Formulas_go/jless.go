package main

// Jless - Command-line pager for JSON data
// Homepage: https://jless.io/

import (
	"fmt"
	
	"os/exec"
)

func installJless() {
	// Método 1: Descargar y extraer .tar.gz
	jless_tar_url := "https://github.com/PaulJuliusMartinez/jless/archive/refs/tags/v0.9.0.tar.gz"
	jless_cmd_tar := exec.Command("curl", "-L", jless_tar_url, "-o", "package.tar.gz")
	err := jless_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jless_zip_url := "https://github.com/PaulJuliusMartinez/jless/archive/refs/tags/v0.9.0.zip"
	jless_cmd_zip := exec.Command("curl", "-L", jless_zip_url, "-o", "package.zip")
	err = jless_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jless_bin_url := "https://github.com/PaulJuliusMartinez/jless/archive/refs/tags/v0.9.0.bin"
	jless_cmd_bin := exec.Command("curl", "-L", jless_bin_url, "-o", "binary.bin")
	err = jless_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jless_src_url := "https://github.com/PaulJuliusMartinez/jless/archive/refs/tags/v0.9.0.src.tar.gz"
	jless_cmd_src := exec.Command("curl", "-L", jless_src_url, "-o", "source.tar.gz")
	err = jless_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jless_cmd_direct := exec.Command("./binary")
	err = jless_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
}
