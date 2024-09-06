package main

// Shub - Scrapinghub command-line client
// Homepage: https://shub.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installShub() {
	// Método 1: Descargar y extraer .tar.gz
	shub_tar_url := "https://files.pythonhosted.org/packages/70/ad/b4fa99366cd3c8db8812438fb1e8b6f8a10b2935b0ee28ac238ade864a8f/shub-2.15.4.tar.gz"
	shub_cmd_tar := exec.Command("curl", "-L", shub_tar_url, "-o", "package.tar.gz")
	err := shub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shub_zip_url := "https://files.pythonhosted.org/packages/70/ad/b4fa99366cd3c8db8812438fb1e8b6f8a10b2935b0ee28ac238ade864a8f/shub-2.15.4.zip"
	shub_cmd_zip := exec.Command("curl", "-L", shub_zip_url, "-o", "package.zip")
	err = shub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shub_bin_url := "https://files.pythonhosted.org/packages/70/ad/b4fa99366cd3c8db8812438fb1e8b6f8a10b2935b0ee28ac238ade864a8f/shub-2.15.4.bin"
	shub_cmd_bin := exec.Command("curl", "-L", shub_bin_url, "-o", "binary.bin")
	err = shub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shub_src_url := "https://files.pythonhosted.org/packages/70/ad/b4fa99366cd3c8db8812438fb1e8b6f8a10b2935b0ee28ac238ade864a8f/shub-2.15.4.src.tar.gz"
	shub_cmd_src := exec.Command("curl", "-L", shub_src_url, "-o", "source.tar.gz")
	err = shub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shub_cmd_direct := exec.Command("./binary")
	err = shub_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
