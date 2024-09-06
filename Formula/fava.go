package main

// Fava - Web interface for the double-entry bookkeeping software Beancount
// Homepage: https://beancount.github.io/fava/

import (
	"fmt"
	
	"os/exec"
)

func installFava() {
	// Método 1: Descargar y extraer .tar.gz
	fava_tar_url := "https://files.pythonhosted.org/packages/58/6a/2fbb724e81f5b62e335c9f01a96ae6a0b20130c3eb8a67aa889722863833/fava-1.28.tar.gz"
	fava_cmd_tar := exec.Command("curl", "-L", fava_tar_url, "-o", "package.tar.gz")
	err := fava_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fava_zip_url := "https://files.pythonhosted.org/packages/58/6a/2fbb724e81f5b62e335c9f01a96ae6a0b20130c3eb8a67aa889722863833/fava-1.28.zip"
	fava_cmd_zip := exec.Command("curl", "-L", fava_zip_url, "-o", "package.zip")
	err = fava_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fava_bin_url := "https://files.pythonhosted.org/packages/58/6a/2fbb724e81f5b62e335c9f01a96ae6a0b20130c3eb8a67aa889722863833/fava-1.28.bin"
	fava_cmd_bin := exec.Command("curl", "-L", fava_bin_url, "-o", "binary.bin")
	err = fava_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fava_src_url := "https://files.pythonhosted.org/packages/58/6a/2fbb724e81f5b62e335c9f01a96ae6a0b20130c3eb8a67aa889722863833/fava-1.28.src.tar.gz"
	fava_cmd_src := exec.Command("curl", "-L", fava_src_url, "-o", "source.tar.gz")
	err = fava_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fava_cmd_direct := exec.Command("./binary")
	err = fava_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
