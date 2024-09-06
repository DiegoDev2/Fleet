package main

// Flamegraph - Stack trace visualizer
// Homepage: https://github.com/brendangregg/FlameGraph

import (
	"fmt"
	
	"os/exec"
)

func installFlamegraph() {
	// Método 1: Descargar y extraer .tar.gz
	flamegraph_tar_url := "https://github.com/brendangregg/FlameGraph/archive/refs/tags/v1.0.tar.gz"
	flamegraph_cmd_tar := exec.Command("curl", "-L", flamegraph_tar_url, "-o", "package.tar.gz")
	err := flamegraph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flamegraph_zip_url := "https://github.com/brendangregg/FlameGraph/archive/refs/tags/v1.0.zip"
	flamegraph_cmd_zip := exec.Command("curl", "-L", flamegraph_zip_url, "-o", "package.zip")
	err = flamegraph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flamegraph_bin_url := "https://github.com/brendangregg/FlameGraph/archive/refs/tags/v1.0.bin"
	flamegraph_cmd_bin := exec.Command("curl", "-L", flamegraph_bin_url, "-o", "binary.bin")
	err = flamegraph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flamegraph_src_url := "https://github.com/brendangregg/FlameGraph/archive/refs/tags/v1.0.src.tar.gz"
	flamegraph_cmd_src := exec.Command("curl", "-L", flamegraph_src_url, "-o", "source.tar.gz")
	err = flamegraph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flamegraph_cmd_direct := exec.Command("./binary")
	err = flamegraph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
