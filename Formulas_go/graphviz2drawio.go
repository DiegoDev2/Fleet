package main

// Graphviz2drawio - Convert graphviz (dot) files into draw.io / lucid (mxGraph) format
// Homepage: https://github.com/hbmartin/graphviz2drawio/

import (
	"fmt"
	
	"os/exec"
)

func installGraphviz2drawio() {
	// Método 1: Descargar y extraer .tar.gz
	graphviz2drawio_tar_url := "https://files.pythonhosted.org/packages/2d/c5/bb43966bc97368fc7eed9d8a79f0bc7eba8484cf6066f687720b616e957a/graphviz2drawio-1.0.0.tar.gz"
	graphviz2drawio_cmd_tar := exec.Command("curl", "-L", graphviz2drawio_tar_url, "-o", "package.tar.gz")
	err := graphviz2drawio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphviz2drawio_zip_url := "https://files.pythonhosted.org/packages/2d/c5/bb43966bc97368fc7eed9d8a79f0bc7eba8484cf6066f687720b616e957a/graphviz2drawio-1.0.0.zip"
	graphviz2drawio_cmd_zip := exec.Command("curl", "-L", graphviz2drawio_zip_url, "-o", "package.zip")
	err = graphviz2drawio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphviz2drawio_bin_url := "https://files.pythonhosted.org/packages/2d/c5/bb43966bc97368fc7eed9d8a79f0bc7eba8484cf6066f687720b616e957a/graphviz2drawio-1.0.0.bin"
	graphviz2drawio_cmd_bin := exec.Command("curl", "-L", graphviz2drawio_bin_url, "-o", "binary.bin")
	err = graphviz2drawio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphviz2drawio_src_url := "https://files.pythonhosted.org/packages/2d/c5/bb43966bc97368fc7eed9d8a79f0bc7eba8484cf6066f687720b616e957a/graphviz2drawio-1.0.0.src.tar.gz"
	graphviz2drawio_cmd_src := exec.Command("curl", "-L", graphviz2drawio_src_url, "-o", "source.tar.gz")
	err = graphviz2drawio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphviz2drawio_cmd_direct := exec.Command("./binary")
	err = graphviz2drawio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
