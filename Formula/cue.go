package main

// Cue - Validate and define text-based and dynamic configuration
// Homepage: https://cuelang.org/

import (
	"fmt"
	
	"os/exec"
)

func installCue() {
	// Método 1: Descargar y extraer .tar.gz
	cue_tar_url := "https://github.com/cue-lang/cue/archive/refs/tags/v0.10.0.tar.gz"
	cue_cmd_tar := exec.Command("curl", "-L", cue_tar_url, "-o", "package.tar.gz")
	err := cue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cue_zip_url := "https://github.com/cue-lang/cue/archive/refs/tags/v0.10.0.zip"
	cue_cmd_zip := exec.Command("curl", "-L", cue_zip_url, "-o", "package.zip")
	err = cue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cue_bin_url := "https://github.com/cue-lang/cue/archive/refs/tags/v0.10.0.bin"
	cue_cmd_bin := exec.Command("curl", "-L", cue_bin_url, "-o", "binary.bin")
	err = cue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cue_src_url := "https://github.com/cue-lang/cue/archive/refs/tags/v0.10.0.src.tar.gz"
	cue_cmd_src := exec.Command("curl", "-L", cue_src_url, "-o", "source.tar.gz")
	err = cue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cue_cmd_direct := exec.Command("./binary")
	err = cue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
