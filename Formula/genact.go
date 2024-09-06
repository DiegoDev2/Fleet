package main

// Genact - Nonsense activity generator
// Homepage: https://github.com/svenstaro/genact

import (
	"fmt"
	
	"os/exec"
)

func installGenact() {
	// Método 1: Descargar y extraer .tar.gz
	genact_tar_url := "https://github.com/svenstaro/genact/archive/refs/tags/v1.4.2.tar.gz"
	genact_cmd_tar := exec.Command("curl", "-L", genact_tar_url, "-o", "package.tar.gz")
	err := genact_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	genact_zip_url := "https://github.com/svenstaro/genact/archive/refs/tags/v1.4.2.zip"
	genact_cmd_zip := exec.Command("curl", "-L", genact_zip_url, "-o", "package.zip")
	err = genact_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	genact_bin_url := "https://github.com/svenstaro/genact/archive/refs/tags/v1.4.2.bin"
	genact_cmd_bin := exec.Command("curl", "-L", genact_bin_url, "-o", "binary.bin")
	err = genact_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	genact_src_url := "https://github.com/svenstaro/genact/archive/refs/tags/v1.4.2.src.tar.gz"
	genact_cmd_src := exec.Command("curl", "-L", genact_src_url, "-o", "source.tar.gz")
	err = genact_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	genact_cmd_direct := exec.Command("./binary")
	err = genact_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
