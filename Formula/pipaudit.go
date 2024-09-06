package main

// PipAudit - Audits Python environments and dependency trees for known vulnerabilities
// Homepage: https://pypi.org/project/pip-audit/

import (
	"fmt"
	
	"os/exec"
)

func installPipAudit() {
	// Método 1: Descargar y extraer .tar.gz
	pipaudit_tar_url := "https://files.pythonhosted.org/packages/46/2f/d030d0d3a50b776f910dd87dc1d57dd4a27bfad176b85882f463632e4747/pip_audit-2.7.3.tar.gz"
	pipaudit_cmd_tar := exec.Command("curl", "-L", pipaudit_tar_url, "-o", "package.tar.gz")
	err := pipaudit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipaudit_zip_url := "https://files.pythonhosted.org/packages/46/2f/d030d0d3a50b776f910dd87dc1d57dd4a27bfad176b85882f463632e4747/pip_audit-2.7.3.zip"
	pipaudit_cmd_zip := exec.Command("curl", "-L", pipaudit_zip_url, "-o", "package.zip")
	err = pipaudit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipaudit_bin_url := "https://files.pythonhosted.org/packages/46/2f/d030d0d3a50b776f910dd87dc1d57dd4a27bfad176b85882f463632e4747/pip_audit-2.7.3.bin"
	pipaudit_cmd_bin := exec.Command("curl", "-L", pipaudit_bin_url, "-o", "binary.bin")
	err = pipaudit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipaudit_src_url := "https://files.pythonhosted.org/packages/46/2f/d030d0d3a50b776f910dd87dc1d57dd4a27bfad176b85882f463632e4747/pip_audit-2.7.3.src.tar.gz"
	pipaudit_cmd_src := exec.Command("curl", "-L", pipaudit_src_url, "-o", "source.tar.gz")
	err = pipaudit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipaudit_cmd_direct := exec.Command("./binary")
	err = pipaudit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
