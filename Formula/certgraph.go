package main

// Certgraph - Crawl the graph of certificate Alternate Names
// Homepage: https://lanrat.github.io/certgraph/

import (
	"fmt"
	
	"os/exec"
)

func installCertgraph() {
	// Método 1: Descargar y extraer .tar.gz
	certgraph_tar_url := "https://github.com/lanrat/certgraph/archive/refs/tags/20220513.tar.gz"
	certgraph_cmd_tar := exec.Command("curl", "-L", certgraph_tar_url, "-o", "package.tar.gz")
	err := certgraph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	certgraph_zip_url := "https://github.com/lanrat/certgraph/archive/refs/tags/20220513.zip"
	certgraph_cmd_zip := exec.Command("curl", "-L", certgraph_zip_url, "-o", "package.zip")
	err = certgraph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	certgraph_bin_url := "https://github.com/lanrat/certgraph/archive/refs/tags/20220513.bin"
	certgraph_cmd_bin := exec.Command("curl", "-L", certgraph_bin_url, "-o", "binary.bin")
	err = certgraph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	certgraph_src_url := "https://github.com/lanrat/certgraph/archive/refs/tags/20220513.src.tar.gz"
	certgraph_cmd_src := exec.Command("curl", "-L", certgraph_src_url, "-o", "source.tar.gz")
	err = certgraph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	certgraph_cmd_direct := exec.Command("./binary")
	err = certgraph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
