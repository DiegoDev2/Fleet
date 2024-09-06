package main

// Minigraph - Proof-of-concept seq-to-graph mapper and graph generator
// Homepage: https://lh3.github.io/minigraph

import (
	"fmt"
	
	"os/exec"
)

func installMinigraph() {
	// Método 1: Descargar y extraer .tar.gz
	minigraph_tar_url := "https://github.com/lh3/minigraph/archive/refs/tags/v0.21.tar.gz"
	minigraph_cmd_tar := exec.Command("curl", "-L", minigraph_tar_url, "-o", "package.tar.gz")
	err := minigraph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minigraph_zip_url := "https://github.com/lh3/minigraph/archive/refs/tags/v0.21.zip"
	minigraph_cmd_zip := exec.Command("curl", "-L", minigraph_zip_url, "-o", "package.zip")
	err = minigraph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minigraph_bin_url := "https://github.com/lh3/minigraph/archive/refs/tags/v0.21.bin"
	minigraph_cmd_bin := exec.Command("curl", "-L", minigraph_bin_url, "-o", "binary.bin")
	err = minigraph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minigraph_src_url := "https://github.com/lh3/minigraph/archive/refs/tags/v0.21.src.tar.gz"
	minigraph_cmd_src := exec.Command("curl", "-L", minigraph_src_url, "-o", "source.tar.gz")
	err = minigraph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minigraph_cmd_direct := exec.Command("./binary")
	err = minigraph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
