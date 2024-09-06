package main

// PythonLspServer - Python Language Server for the Language Server Protocol
// Homepage: https://github.com/python-lsp/python-lsp-server

import (
	"fmt"
	
	"os/exec"
)

func installPythonLspServer() {
	// Método 1: Descargar y extraer .tar.gz
	pythonlspserver_tar_url := "https://files.pythonhosted.org/packages/2b/15/b7e1577b9ca358e008b06910bf23cfa0a8be130ee9f319a262a3c610ee8d/python_lsp_server-1.12.0.tar.gz"
	pythonlspserver_cmd_tar := exec.Command("curl", "-L", pythonlspserver_tar_url, "-o", "package.tar.gz")
	err := pythonlspserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonlspserver_zip_url := "https://files.pythonhosted.org/packages/2b/15/b7e1577b9ca358e008b06910bf23cfa0a8be130ee9f319a262a3c610ee8d/python_lsp_server-1.12.0.zip"
	pythonlspserver_cmd_zip := exec.Command("curl", "-L", pythonlspserver_zip_url, "-o", "package.zip")
	err = pythonlspserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonlspserver_bin_url := "https://files.pythonhosted.org/packages/2b/15/b7e1577b9ca358e008b06910bf23cfa0a8be130ee9f319a262a3c610ee8d/python_lsp_server-1.12.0.bin"
	pythonlspserver_cmd_bin := exec.Command("curl", "-L", pythonlspserver_bin_url, "-o", "binary.bin")
	err = pythonlspserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonlspserver_src_url := "https://files.pythonhosted.org/packages/2b/15/b7e1577b9ca358e008b06910bf23cfa0a8be130ee9f319a262a3c610ee8d/python_lsp_server-1.12.0.src.tar.gz"
	pythonlspserver_cmd_src := exec.Command("curl", "-L", pythonlspserver_src_url, "-o", "source.tar.gz")
	err = pythonlspserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonlspserver_cmd_direct := exec.Command("./binary")
	err = pythonlspserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
