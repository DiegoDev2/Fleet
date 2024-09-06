package main

// Otree - Command-line tool to view objects (JSON/YAML/TOML) in TUI tree widget
// Homepage: https://github.com/fioncat/otree

import (
	"fmt"
	
	"os/exec"
)

func installOtree() {
	// Método 1: Descargar y extraer .tar.gz
	otree_tar_url := "https://github.com/fioncat/otree/archive/refs/tags/v0.2.0.tar.gz"
	otree_cmd_tar := exec.Command("curl", "-L", otree_tar_url, "-o", "package.tar.gz")
	err := otree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	otree_zip_url := "https://github.com/fioncat/otree/archive/refs/tags/v0.2.0.zip"
	otree_cmd_zip := exec.Command("curl", "-L", otree_zip_url, "-o", "package.zip")
	err = otree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	otree_bin_url := "https://github.com/fioncat/otree/archive/refs/tags/v0.2.0.bin"
	otree_cmd_bin := exec.Command("curl", "-L", otree_bin_url, "-o", "binary.bin")
	err = otree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	otree_src_url := "https://github.com/fioncat/otree/archive/refs/tags/v0.2.0.src.tar.gz"
	otree_cmd_src := exec.Command("curl", "-L", otree_src_url, "-o", "source.tar.gz")
	err = otree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	otree_cmd_direct := exec.Command("./binary")
	err = otree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
