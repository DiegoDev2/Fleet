package main

// Igraph - Network analysis package
// Homepage: https://igraph.org/

import (
	"fmt"
	
	"os/exec"
)

func installIgraph() {
	// Método 1: Descargar y extraer .tar.gz
	igraph_tar_url := "https://github.com/igraph/igraph/releases/download/0.10.13/igraph-0.10.13.tar.gz"
	igraph_cmd_tar := exec.Command("curl", "-L", igraph_tar_url, "-o", "package.tar.gz")
	err := igraph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	igraph_zip_url := "https://github.com/igraph/igraph/releases/download/0.10.13/igraph-0.10.13.zip"
	igraph_cmd_zip := exec.Command("curl", "-L", igraph_zip_url, "-o", "package.zip")
	err = igraph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	igraph_bin_url := "https://github.com/igraph/igraph/releases/download/0.10.13/igraph-0.10.13.bin"
	igraph_cmd_bin := exec.Command("curl", "-L", igraph_bin_url, "-o", "binary.bin")
	err = igraph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	igraph_src_url := "https://github.com/igraph/igraph/releases/download/0.10.13/igraph-0.10.13.src.tar.gz"
	igraph_cmd_src := exec.Command("curl", "-L", igraph_src_url, "-o", "source.tar.gz")
	err = igraph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	igraph_cmd_direct := exec.Command("./binary")
	err = igraph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: arpack")
	exec.Command("latte", "install", "arpack").Run()
	fmt.Println("Instalando dependencia: glpk")
	exec.Command("latte", "install", "glpk").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
