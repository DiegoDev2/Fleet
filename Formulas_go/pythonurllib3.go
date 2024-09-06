package main

// PythonUrllib3 - HTTP library with thread-safe connection pooling, file post, and more
// Homepage: https://urllib3.readthedocs.io/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installPythonUrllib3() {
	// Método 1: Descargar y extraer .tar.gz
	pythonurllib3_tar_url := "https://files.pythonhosted.org/packages/7a/50/7fd50a27caa0652cd4caf224aa87741ea41d3265ad13f010886167cfcc79/urllib3-2.2.1.tar.gz"
	pythonurllib3_cmd_tar := exec.Command("curl", "-L", pythonurllib3_tar_url, "-o", "package.tar.gz")
	err := pythonurllib3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonurllib3_zip_url := "https://files.pythonhosted.org/packages/7a/50/7fd50a27caa0652cd4caf224aa87741ea41d3265ad13f010886167cfcc79/urllib3-2.2.1.zip"
	pythonurllib3_cmd_zip := exec.Command("curl", "-L", pythonurllib3_zip_url, "-o", "package.zip")
	err = pythonurllib3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonurllib3_bin_url := "https://files.pythonhosted.org/packages/7a/50/7fd50a27caa0652cd4caf224aa87741ea41d3265ad13f010886167cfcc79/urllib3-2.2.1.bin"
	pythonurllib3_cmd_bin := exec.Command("curl", "-L", pythonurllib3_bin_url, "-o", "binary.bin")
	err = pythonurllib3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonurllib3_src_url := "https://files.pythonhosted.org/packages/7a/50/7fd50a27caa0652cd4caf224aa87741ea41d3265ad13f010886167cfcc79/urllib3-2.2.1.src.tar.gz"
	pythonurllib3_cmd_src := exec.Command("curl", "-L", pythonurllib3_src_url, "-o", "source.tar.gz")
	err = pythonurllib3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonurllib3_cmd_direct := exec.Command("./binary")
	err = pythonurllib3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
