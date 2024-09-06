package main

// Meilisearch - Ultra relevant, instant and typo-tolerant full-text search API
// Homepage: https://docs.meilisearch.com/

import (
	"fmt"
	
	"os/exec"
)

func installMeilisearch() {
	// Método 1: Descargar y extraer .tar.gz
	meilisearch_tar_url := "https://github.com/meilisearch/meilisearch/archive/refs/tags/v1.10.1.tar.gz"
	meilisearch_cmd_tar := exec.Command("curl", "-L", meilisearch_tar_url, "-o", "package.tar.gz")
	err := meilisearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	meilisearch_zip_url := "https://github.com/meilisearch/meilisearch/archive/refs/tags/v1.10.1.zip"
	meilisearch_cmd_zip := exec.Command("curl", "-L", meilisearch_zip_url, "-o", "package.zip")
	err = meilisearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	meilisearch_bin_url := "https://github.com/meilisearch/meilisearch/archive/refs/tags/v1.10.1.bin"
	meilisearch_cmd_bin := exec.Command("curl", "-L", meilisearch_bin_url, "-o", "binary.bin")
	err = meilisearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	meilisearch_src_url := "https://github.com/meilisearch/meilisearch/archive/refs/tags/v1.10.1.src.tar.gz"
	meilisearch_cmd_src := exec.Command("curl", "-L", meilisearch_src_url, "-o", "source.tar.gz")
	err = meilisearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	meilisearch_cmd_direct := exec.Command("./binary")
	err = meilisearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
