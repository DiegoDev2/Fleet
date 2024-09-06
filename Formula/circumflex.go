package main

// Circumflex - Hacker News in your terminal
// Homepage: https://github.com/bensadeh/circumflex

import (
	"fmt"
	
	"os/exec"
)

func installCircumflex() {
	// Método 1: Descargar y extraer .tar.gz
	circumflex_tar_url := "https://github.com/bensadeh/circumflex/archive/refs/tags/3.7.tar.gz"
	circumflex_cmd_tar := exec.Command("curl", "-L", circumflex_tar_url, "-o", "package.tar.gz")
	err := circumflex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	circumflex_zip_url := "https://github.com/bensadeh/circumflex/archive/refs/tags/3.7.zip"
	circumflex_cmd_zip := exec.Command("curl", "-L", circumflex_zip_url, "-o", "package.zip")
	err = circumflex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	circumflex_bin_url := "https://github.com/bensadeh/circumflex/archive/refs/tags/3.7.bin"
	circumflex_cmd_bin := exec.Command("curl", "-L", circumflex_bin_url, "-o", "binary.bin")
	err = circumflex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	circumflex_src_url := "https://github.com/bensadeh/circumflex/archive/refs/tags/3.7.src.tar.gz"
	circumflex_cmd_src := exec.Command("curl", "-L", circumflex_src_url, "-o", "source.tar.gz")
	err = circumflex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	circumflex_cmd_direct := exec.Command("./binary")
	err = circumflex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: less")
	exec.Command("latte", "install", "less").Run()
}
