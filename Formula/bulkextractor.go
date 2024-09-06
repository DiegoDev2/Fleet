package main

// BulkExtractor - Stream-based forensics tool
// Homepage: https://github.com/simsong/bulk_extractor/wiki

import (
	"fmt"
	
	"os/exec"
)

func installBulkExtractor() {
	// Método 1: Descargar y extraer .tar.gz
	bulkextractor_tar_url := "https://github.com/simsong/bulk_extractor/releases/download/v2.1.1/bulk_extractor-2.1.1.tar.gz"
	bulkextractor_cmd_tar := exec.Command("curl", "-L", bulkextractor_tar_url, "-o", "package.tar.gz")
	err := bulkextractor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bulkextractor_zip_url := "https://github.com/simsong/bulk_extractor/releases/download/v2.1.1/bulk_extractor-2.1.1.zip"
	bulkextractor_cmd_zip := exec.Command("curl", "-L", bulkextractor_zip_url, "-o", "package.zip")
	err = bulkextractor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bulkextractor_bin_url := "https://github.com/simsong/bulk_extractor/releases/download/v2.1.1/bulk_extractor-2.1.1.bin"
	bulkextractor_cmd_bin := exec.Command("curl", "-L", bulkextractor_bin_url, "-o", "binary.bin")
	err = bulkextractor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bulkextractor_src_url := "https://github.com/simsong/bulk_extractor/releases/download/v2.1.1/bulk_extractor-2.1.1.src.tar.gz"
	bulkextractor_cmd_src := exec.Command("curl", "-L", bulkextractor_src_url, "-o", "source.tar.gz")
	err = bulkextractor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bulkextractor_cmd_direct := exec.Command("./binary")
	err = bulkextractor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
}
