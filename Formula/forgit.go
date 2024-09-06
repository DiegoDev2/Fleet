package main

// Forgit - Interactive git commands in the terminal
// Homepage: https://github.com/wfxr/forgit

import (
	"fmt"
	
	"os/exec"
)

func installForgit() {
	// Método 1: Descargar y extraer .tar.gz
	forgit_tar_url := "https://github.com/wfxr/forgit/releases/download/24.09.0/forgit-24.09.0.tar.gz"
	forgit_cmd_tar := exec.Command("curl", "-L", forgit_tar_url, "-o", "package.tar.gz")
	err := forgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	forgit_zip_url := "https://github.com/wfxr/forgit/releases/download/24.09.0/forgit-24.09.0.zip"
	forgit_cmd_zip := exec.Command("curl", "-L", forgit_zip_url, "-o", "package.zip")
	err = forgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	forgit_bin_url := "https://github.com/wfxr/forgit/releases/download/24.09.0/forgit-24.09.0.bin"
	forgit_cmd_bin := exec.Command("curl", "-L", forgit_bin_url, "-o", "binary.bin")
	err = forgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	forgit_src_url := "https://github.com/wfxr/forgit/releases/download/24.09.0/forgit-24.09.0.src.tar.gz"
	forgit_cmd_src := exec.Command("curl", "-L", forgit_src_url, "-o", "source.tar.gz")
	err = forgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	forgit_cmd_direct := exec.Command("./binary")
	err = forgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fzf")
	exec.Command("latte", "install", "fzf").Run()
}
