package main

// PythonRequests - Python HTTP for Humans
// Homepage: https://requests.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installPythonRequests() {
	// Método 1: Descargar y extraer .tar.gz
	pythonrequests_tar_url := "https://files.pythonhosted.org/packages/9d/be/10918a2eac4ae9f02f6cfe6414b7a155ccd8f7f9d4380d62fd5b955065c3/requests-2.31.0.tar.gz"
	pythonrequests_cmd_tar := exec.Command("curl", "-L", pythonrequests_tar_url, "-o", "package.tar.gz")
	err := pythonrequests_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonrequests_zip_url := "https://files.pythonhosted.org/packages/9d/be/10918a2eac4ae9f02f6cfe6414b7a155ccd8f7f9d4380d62fd5b955065c3/requests-2.31.0.zip"
	pythonrequests_cmd_zip := exec.Command("curl", "-L", pythonrequests_zip_url, "-o", "package.zip")
	err = pythonrequests_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonrequests_bin_url := "https://files.pythonhosted.org/packages/9d/be/10918a2eac4ae9f02f6cfe6414b7a155ccd8f7f9d4380d62fd5b955065c3/requests-2.31.0.bin"
	pythonrequests_cmd_bin := exec.Command("curl", "-L", pythonrequests_bin_url, "-o", "binary.bin")
	err = pythonrequests_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonrequests_src_url := "https://files.pythonhosted.org/packages/9d/be/10918a2eac4ae9f02f6cfe6414b7a155ccd8f7f9d4380d62fd5b955065c3/requests-2.31.0.src.tar.gz"
	pythonrequests_cmd_src := exec.Command("curl", "-L", pythonrequests_src_url, "-o", "source.tar.gz")
	err = pythonrequests_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonrequests_cmd_direct := exec.Command("./binary")
	err = pythonrequests_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python-charset-normalizer")
exec.Command("latte", "install", "python-charset-normalizer")
	fmt.Println("Instalando dependencia: python-idna")
exec.Command("latte", "install", "python-idna")
	fmt.Println("Instalando dependencia: python-urllib3")
exec.Command("latte", "install", "python-urllib3")
}
