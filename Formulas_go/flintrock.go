package main

// Flintrock - Tool for launching Apache Spark clusters
// Homepage: https://github.com/nchammas/flintrock

import (
	"fmt"
	
	"os/exec"
)

func installFlintrock() {
	// Método 1: Descargar y extraer .tar.gz
	flintrock_tar_url := "https://files.pythonhosted.org/packages/e9/3b/810c7757f6abb0a73a50c2da6da2dacb5af85a04b056aef81323b2b6a082/Flintrock-2.1.0.tar.gz"
	flintrock_cmd_tar := exec.Command("curl", "-L", flintrock_tar_url, "-o", "package.tar.gz")
	err := flintrock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flintrock_zip_url := "https://files.pythonhosted.org/packages/e9/3b/810c7757f6abb0a73a50c2da6da2dacb5af85a04b056aef81323b2b6a082/Flintrock-2.1.0.zip"
	flintrock_cmd_zip := exec.Command("curl", "-L", flintrock_zip_url, "-o", "package.zip")
	err = flintrock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flintrock_bin_url := "https://files.pythonhosted.org/packages/e9/3b/810c7757f6abb0a73a50c2da6da2dacb5af85a04b056aef81323b2b6a082/Flintrock-2.1.0.bin"
	flintrock_cmd_bin := exec.Command("curl", "-L", flintrock_bin_url, "-o", "binary.bin")
	err = flintrock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flintrock_src_url := "https://files.pythonhosted.org/packages/e9/3b/810c7757f6abb0a73a50c2da6da2dacb5af85a04b056aef81323b2b6a082/Flintrock-2.1.0.src.tar.gz"
	flintrock_cmd_src := exec.Command("curl", "-L", flintrock_src_url, "-o", "source.tar.gz")
	err = flintrock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flintrock_cmd_direct := exec.Command("./binary")
	err = flintrock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
