package main

// Grokmirror - Framework to smartly mirror git repositories
// Homepage: https://github.com/mricon/grokmirror

import (
	"fmt"
	
	"os/exec"
)

func installGrokmirror() {
	// Método 1: Descargar y extraer .tar.gz
	grokmirror_tar_url := "https://files.pythonhosted.org/packages/26/91/af8831185ef4e5bef5d210039ab67abdc8c27a09a585d3963a10cf774789/grokmirror-2.0.12.tar.gz"
	grokmirror_cmd_tar := exec.Command("curl", "-L", grokmirror_tar_url, "-o", "package.tar.gz")
	err := grokmirror_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grokmirror_zip_url := "https://files.pythonhosted.org/packages/26/91/af8831185ef4e5bef5d210039ab67abdc8c27a09a585d3963a10cf774789/grokmirror-2.0.12.zip"
	grokmirror_cmd_zip := exec.Command("curl", "-L", grokmirror_zip_url, "-o", "package.zip")
	err = grokmirror_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grokmirror_bin_url := "https://files.pythonhosted.org/packages/26/91/af8831185ef4e5bef5d210039ab67abdc8c27a09a585d3963a10cf774789/grokmirror-2.0.12.bin"
	grokmirror_cmd_bin := exec.Command("curl", "-L", grokmirror_bin_url, "-o", "binary.bin")
	err = grokmirror_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grokmirror_src_url := "https://files.pythonhosted.org/packages/26/91/af8831185ef4e5bef5d210039ab67abdc8c27a09a585d3963a10cf774789/grokmirror-2.0.12.src.tar.gz"
	grokmirror_cmd_src := exec.Command("curl", "-L", grokmirror_src_url, "-o", "source.tar.gz")
	err = grokmirror_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grokmirror_cmd_direct := exec.Command("./binary")
	err = grokmirror_cmd_direct.Run()
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
