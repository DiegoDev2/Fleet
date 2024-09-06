package main

// Tofuenv - OpenTofu version manager inspired by tfenv
// Homepage: https://github.com/tofuutils/tofuenv

import (
	"fmt"
	
	"os/exec"
)

func installTofuenv() {
	// Método 1: Descargar y extraer .tar.gz
	tofuenv_tar_url := "https://github.com/tofuutils/tofuenv/archive/refs/tags/v1.0.6.tar.gz"
	tofuenv_cmd_tar := exec.Command("curl", "-L", tofuenv_tar_url, "-o", "package.tar.gz")
	err := tofuenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tofuenv_zip_url := "https://github.com/tofuutils/tofuenv/archive/refs/tags/v1.0.6.zip"
	tofuenv_cmd_zip := exec.Command("curl", "-L", tofuenv_zip_url, "-o", "package.zip")
	err = tofuenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tofuenv_bin_url := "https://github.com/tofuutils/tofuenv/archive/refs/tags/v1.0.6.bin"
	tofuenv_cmd_bin := exec.Command("curl", "-L", tofuenv_bin_url, "-o", "binary.bin")
	err = tofuenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tofuenv_src_url := "https://github.com/tofuutils/tofuenv/archive/refs/tags/v1.0.6.src.tar.gz"
	tofuenv_cmd_src := exec.Command("curl", "-L", tofuenv_src_url, "-o", "source.tar.gz")
	err = tofuenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tofuenv_cmd_direct := exec.Command("./binary")
	err = tofuenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: grep")
exec.Command("latte", "install", "grep")
	fmt.Println("Instalando dependencia: jq")
exec.Command("latte", "install", "jq")
}
