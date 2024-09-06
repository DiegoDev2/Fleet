package main

// Scrapy - Web crawling & scraping framework
// Homepage: https://scrapy.org

import (
	"fmt"
	
	"os/exec"
)

func installScrapy() {
	// Método 1: Descargar y extraer .tar.gz
	scrapy_tar_url := "https://files.pythonhosted.org/packages/f2/1f/5524416a64c030fbe18caeba079e7176836b281bf9eb50b79efdf8015063/scrapy-2.11.2.tar.gz"
	scrapy_cmd_tar := exec.Command("curl", "-L", scrapy_tar_url, "-o", "package.tar.gz")
	err := scrapy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scrapy_zip_url := "https://files.pythonhosted.org/packages/f2/1f/5524416a64c030fbe18caeba079e7176836b281bf9eb50b79efdf8015063/scrapy-2.11.2.zip"
	scrapy_cmd_zip := exec.Command("curl", "-L", scrapy_zip_url, "-o", "package.zip")
	err = scrapy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scrapy_bin_url := "https://files.pythonhosted.org/packages/f2/1f/5524416a64c030fbe18caeba079e7176836b281bf9eb50b79efdf8015063/scrapy-2.11.2.bin"
	scrapy_cmd_bin := exec.Command("curl", "-L", scrapy_bin_url, "-o", "binary.bin")
	err = scrapy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scrapy_src_url := "https://files.pythonhosted.org/packages/f2/1f/5524416a64c030fbe18caeba079e7176836b281bf9eb50b79efdf8015063/scrapy-2.11.2.src.tar.gz"
	scrapy_cmd_src := exec.Command("curl", "-L", scrapy_src_url, "-o", "source.tar.gz")
	err = scrapy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scrapy_cmd_direct := exec.Command("./binary")
	err = scrapy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
