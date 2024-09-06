package main

// Fastmod - Fast partial replacement for the codemod tool
// Homepage: https://github.com/facebookincubator/fastmod

import (
	"fmt"
	
	"os/exec"
)

func installFastmod() {
	// Método 1: Descargar y extraer .tar.gz
	fastmod_tar_url := "https://github.com/facebookincubator/fastmod/archive/refs/tags/v0.4.4.tar.gz"
	fastmod_cmd_tar := exec.Command("curl", "-L", fastmod_tar_url, "-o", "package.tar.gz")
	err := fastmod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastmod_zip_url := "https://github.com/facebookincubator/fastmod/archive/refs/tags/v0.4.4.zip"
	fastmod_cmd_zip := exec.Command("curl", "-L", fastmod_zip_url, "-o", "package.zip")
	err = fastmod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastmod_bin_url := "https://github.com/facebookincubator/fastmod/archive/refs/tags/v0.4.4.bin"
	fastmod_cmd_bin := exec.Command("curl", "-L", fastmod_bin_url, "-o", "binary.bin")
	err = fastmod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastmod_src_url := "https://github.com/facebookincubator/fastmod/archive/refs/tags/v0.4.4.src.tar.gz"
	fastmod_cmd_src := exec.Command("curl", "-L", fastmod_src_url, "-o", "source.tar.gz")
	err = fastmod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastmod_cmd_direct := exec.Command("./binary")
	err = fastmod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
