package main

// Atlas - Database toolkit
// Homepage: https://atlasgo.io/

import (
	"fmt"
	
	"os/exec"
)

func installAtlas() {
	// Método 1: Descargar y extraer .tar.gz
	atlas_tar_url := "https://github.com/ariga/atlas/archive/refs/tags/v0.27.0.tar.gz"
	atlas_cmd_tar := exec.Command("curl", "-L", atlas_tar_url, "-o", "package.tar.gz")
	err := atlas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atlas_zip_url := "https://github.com/ariga/atlas/archive/refs/tags/v0.27.0.zip"
	atlas_cmd_zip := exec.Command("curl", "-L", atlas_zip_url, "-o", "package.zip")
	err = atlas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atlas_bin_url := "https://github.com/ariga/atlas/archive/refs/tags/v0.27.0.bin"
	atlas_cmd_bin := exec.Command("curl", "-L", atlas_bin_url, "-o", "binary.bin")
	err = atlas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atlas_src_url := "https://github.com/ariga/atlas/archive/refs/tags/v0.27.0.src.tar.gz"
	atlas_cmd_src := exec.Command("curl", "-L", atlas_src_url, "-o", "source.tar.gz")
	err = atlas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atlas_cmd_direct := exec.Command("./binary")
	err = atlas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
