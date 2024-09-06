package main

// Checkmake - Linter/analyzer for Makefiles
// Homepage: https://github.com/mrtazz/checkmake

import (
	"fmt"
	
	"os/exec"
)

func installCheckmake() {
	// Método 1: Descargar y extraer .tar.gz
	checkmake_tar_url := "https://github.com/mrtazz/checkmake/archive/refs/tags/0.2.2.tar.gz"
	checkmake_cmd_tar := exec.Command("curl", "-L", checkmake_tar_url, "-o", "package.tar.gz")
	err := checkmake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	checkmake_zip_url := "https://github.com/mrtazz/checkmake/archive/refs/tags/0.2.2.zip"
	checkmake_cmd_zip := exec.Command("curl", "-L", checkmake_zip_url, "-o", "package.zip")
	err = checkmake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	checkmake_bin_url := "https://github.com/mrtazz/checkmake/archive/refs/tags/0.2.2.bin"
	checkmake_cmd_bin := exec.Command("curl", "-L", checkmake_bin_url, "-o", "binary.bin")
	err = checkmake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	checkmake_src_url := "https://github.com/mrtazz/checkmake/archive/refs/tags/0.2.2.src.tar.gz"
	checkmake_cmd_src := exec.Command("curl", "-L", checkmake_src_url, "-o", "source.tar.gz")
	err = checkmake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	checkmake_cmd_direct := exec.Command("./binary")
	err = checkmake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
}
