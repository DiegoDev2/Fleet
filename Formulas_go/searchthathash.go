package main

// SearchThatHash - Searches Hash APIs to crack your hash quickly
// Homepage: https://github.com/HashPals/Search-That-Hash

import (
	"fmt"
	
	"os/exec"
)

func installSearchThatHash() {
	// Método 1: Descargar y extraer .tar.gz
	searchthathash_tar_url := "https://files.pythonhosted.org/packages/5e/b9/a304a92ba77a9e18b3023b66634e71cded5285cef7e3b56d3c1874e9d84e/search-that-hash-0.2.8.tar.gz"
	searchthathash_cmd_tar := exec.Command("curl", "-L", searchthathash_tar_url, "-o", "package.tar.gz")
	err := searchthathash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	searchthathash_zip_url := "https://files.pythonhosted.org/packages/5e/b9/a304a92ba77a9e18b3023b66634e71cded5285cef7e3b56d3c1874e9d84e/search-that-hash-0.2.8.zip"
	searchthathash_cmd_zip := exec.Command("curl", "-L", searchthathash_zip_url, "-o", "package.zip")
	err = searchthathash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	searchthathash_bin_url := "https://files.pythonhosted.org/packages/5e/b9/a304a92ba77a9e18b3023b66634e71cded5285cef7e3b56d3c1874e9d84e/search-that-hash-0.2.8.bin"
	searchthathash_cmd_bin := exec.Command("curl", "-L", searchthathash_bin_url, "-o", "binary.bin")
	err = searchthathash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	searchthathash_src_url := "https://files.pythonhosted.org/packages/5e/b9/a304a92ba77a9e18b3023b66634e71cded5285cef7e3b56d3c1874e9d84e/search-that-hash-0.2.8.src.tar.gz"
	searchthathash_cmd_src := exec.Command("curl", "-L", searchthathash_src_url, "-o", "source.tar.gz")
	err = searchthathash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	searchthathash_cmd_direct := exec.Command("./binary")
	err = searchthathash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
