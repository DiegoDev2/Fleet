package main

// CargoDepgraph - Creates dependency graphs for cargo projects
// Homepage: https://sr.ht/~jplatte/cargo-depgraph/

import (
	"fmt"
	
	"os/exec"
)

func installCargoDepgraph() {
	// Método 1: Descargar y extraer .tar.gz
	cargodepgraph_tar_url := "https://git.sr.ht/~jplatte/cargo-depgraph/archive/v1.6.0.tar.gz"
	cargodepgraph_cmd_tar := exec.Command("curl", "-L", cargodepgraph_tar_url, "-o", "package.tar.gz")
	err := cargodepgraph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargodepgraph_zip_url := "https://git.sr.ht/~jplatte/cargo-depgraph/archive/v1.6.0.zip"
	cargodepgraph_cmd_zip := exec.Command("curl", "-L", cargodepgraph_zip_url, "-o", "package.zip")
	err = cargodepgraph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargodepgraph_bin_url := "https://git.sr.ht/~jplatte/cargo-depgraph/archive/v1.6.0.bin"
	cargodepgraph_cmd_bin := exec.Command("curl", "-L", cargodepgraph_bin_url, "-o", "binary.bin")
	err = cargodepgraph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargodepgraph_src_url := "https://git.sr.ht/~jplatte/cargo-depgraph/archive/v1.6.0.src.tar.gz"
	cargodepgraph_cmd_src := exec.Command("curl", "-L", cargodepgraph_src_url, "-o", "source.tar.gz")
	err = cargodepgraph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargodepgraph_cmd_direct := exec.Command("./binary")
	err = cargodepgraph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
}
