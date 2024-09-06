package main

// Opensearch - Open source distributed and RESTful search engine
// Homepage: https://github.com/opensearch-project/OpenSearch

import (
	"fmt"
	
	"os/exec"
)

func installOpensearch() {
	// Método 1: Descargar y extraer .tar.gz
	opensearch_tar_url := "https://github.com/opensearch-project/OpenSearch/archive/refs/tags/2.16.0.tar.gz"
	opensearch_cmd_tar := exec.Command("curl", "-L", opensearch_tar_url, "-o", "package.tar.gz")
	err := opensearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensearch_zip_url := "https://github.com/opensearch-project/OpenSearch/archive/refs/tags/2.16.0.zip"
	opensearch_cmd_zip := exec.Command("curl", "-L", opensearch_zip_url, "-o", "package.zip")
	err = opensearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensearch_bin_url := "https://github.com/opensearch-project/OpenSearch/archive/refs/tags/2.16.0.bin"
	opensearch_cmd_bin := exec.Command("curl", "-L", opensearch_bin_url, "-o", "binary.bin")
	err = opensearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensearch_src_url := "https://github.com/opensearch-project/OpenSearch/archive/refs/tags/2.16.0.src.tar.gz"
	opensearch_cmd_src := exec.Command("curl", "-L", opensearch_src_url, "-o", "source.tar.gz")
	err = opensearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensearch_cmd_direct := exec.Command("./binary")
	err = opensearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
	exec.Command("latte", "install", "gradle").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
