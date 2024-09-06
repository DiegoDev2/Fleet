package main

// Joshuto - Ranger-like terminal file manager written in Rust
// Homepage: https://github.com/kamiyaa/joshuto

import (
	"fmt"
	
	"os/exec"
)

func installJoshuto() {
	// Método 1: Descargar y extraer .tar.gz
	joshuto_tar_url := "https://github.com/kamiyaa/joshuto/archive/refs/tags/v0.9.8.tar.gz"
	joshuto_cmd_tar := exec.Command("curl", "-L", joshuto_tar_url, "-o", "package.tar.gz")
	err := joshuto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	joshuto_zip_url := "https://github.com/kamiyaa/joshuto/archive/refs/tags/v0.9.8.zip"
	joshuto_cmd_zip := exec.Command("curl", "-L", joshuto_zip_url, "-o", "package.zip")
	err = joshuto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	joshuto_bin_url := "https://github.com/kamiyaa/joshuto/archive/refs/tags/v0.9.8.bin"
	joshuto_cmd_bin := exec.Command("curl", "-L", joshuto_bin_url, "-o", "binary.bin")
	err = joshuto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	joshuto_src_url := "https://github.com/kamiyaa/joshuto/archive/refs/tags/v0.9.8.src.tar.gz"
	joshuto_cmd_src := exec.Command("curl", "-L", joshuto_src_url, "-o", "source.tar.gz")
	err = joshuto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	joshuto_cmd_direct := exec.Command("./binary")
	err = joshuto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
