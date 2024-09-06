package main

// Raven - Risk Analysis and Vulnerability Enumeration for CI/CD
// Homepage: https://github.com/CycodeLabs/raven

import (
	"fmt"
	
	"os/exec"
)

func installRaven() {
	// Método 1: Descargar y extraer .tar.gz
	raven_tar_url := "https://files.pythonhosted.org/packages/d8/b6/4bc5aecae28382720fca4e9492a623e3d821d96e8f4d06e4335c77779ebd/raven_cycode-1.0.9.tar.gz"
	raven_cmd_tar := exec.Command("curl", "-L", raven_tar_url, "-o", "package.tar.gz")
	err := raven_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	raven_zip_url := "https://files.pythonhosted.org/packages/d8/b6/4bc5aecae28382720fca4e9492a623e3d821d96e8f4d06e4335c77779ebd/raven_cycode-1.0.9.zip"
	raven_cmd_zip := exec.Command("curl", "-L", raven_zip_url, "-o", "package.zip")
	err = raven_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	raven_bin_url := "https://files.pythonhosted.org/packages/d8/b6/4bc5aecae28382720fca4e9492a623e3d821d96e8f4d06e4335c77779ebd/raven_cycode-1.0.9.bin"
	raven_cmd_bin := exec.Command("curl", "-L", raven_bin_url, "-o", "binary.bin")
	err = raven_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	raven_src_url := "https://files.pythonhosted.org/packages/d8/b6/4bc5aecae28382720fca4e9492a623e3d821d96e8f4d06e4335c77779ebd/raven_cycode-1.0.9.src.tar.gz"
	raven_cmd_src := exec.Command("curl", "-L", raven_src_url, "-o", "source.tar.gz")
	err = raven_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	raven_cmd_direct := exec.Command("./binary")
	err = raven_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
